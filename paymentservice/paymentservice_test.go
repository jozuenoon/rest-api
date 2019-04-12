package paymentservice_test

import (
	"context"
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jozuenoon/rest-api/paymentapi"
	. "github.com/jozuenoon/rest-api/paymentservice"
)

var _ = Describe("Paymentservice", func() {
	var (
		svc    *PaymentService
		tmpDir string
		err    error
	)

	BeforeEach(func() {
		tmpDir, err = ioutil.TempDir("", "payment_database")
	})

	JustBeforeEach(func() {
		svc, err = NewPaymentService(tmpDir)
	})

	AfterEach(func() {
		err = os.RemoveAll(tmpDir)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should not error", func() {
		Expect(err).NotTo(HaveOccurred())
	})

	It("should be empty if nothing is loaded", func() {
		resp, err := svc.GetPayment(context.Background(), &paymentapi.PaymentRequest{})
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.Data).Should(HaveLen(0))
	})

	It("should create content", func() {
		date := "2017-01-01"
		resp, err := svc.CreatePayment(context.Background(), &paymentapi.PaymentAttributes{
			ProcessingDate: date,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.Data).Should(HaveLen(1))
		Expect(resp.Data[0].Id).ShouldNot(BeEmpty())
		Expect(resp.Data[0].Attributes.ProcessingDate).Should(Equal(date))
	})

	It("should return created content", func() {
		date := "2017-01-01"
		cresp, err := svc.CreatePayment(context.Background(), &paymentapi.PaymentAttributes{
			ProcessingDate: date,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(cresp.Data).Should(HaveLen(1))
		resp, err := svc.GetPayment(context.Background(), &paymentapi.PaymentRequest{
			Id: cresp.Data[0].Id,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(cresp).Should(Equal(resp))
	})

	It("should change values with update", func() {
		date := "2017-01-02"
		cresp, err := svc.CreatePayment(context.Background(), &paymentapi.PaymentAttributes{
			ProcessingDate: date,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(cresp.Data).Should(HaveLen(1))

		resp, err := svc.GetPayment(context.Background(), &paymentapi.PaymentRequest{
			Id: cresp.Data[0].Id,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(cresp).Should(Equal(resp))
		Expect(cresp.Data[0].Attributes.ProcessingDate).Should(Equal(date))
		Expect(cresp.Data[0].Attributes.Amount).Should(Equal(int64(0)))

		date2 := "2018-01-02"
		amount := int64(332234)
		uresp, err := svc.UpdatePayment(context.Background(), &paymentapi.PaymentUpdate{
			Id: cresp.Data[0].Id,
			PaymentAttributes: &paymentapi.PaymentAttributes{
				ProcessingDate: date2,
				Amount:         amount,
			},
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(uresp.Data).Should(HaveLen(1))
		Expect(uresp.Data[0].Attributes.ProcessingDate).Should(Equal(date2))
		Expect(uresp.Data[0].Attributes.Amount).Should(Equal(amount))

		resp, err = svc.GetPayment(context.Background(), &paymentapi.PaymentRequest{
			Id: cresp.Data[0].Id,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(uresp).Should(Equal(resp))
		Expect(uresp.Data[0].Attributes.ProcessingDate).Should(Equal(date2))
		Expect(uresp.Data[0].Attributes.Amount).Should(Equal(amount))
	})

	It("should delete payment", func() {
		date := "2017-01-02"
		amount := int64(3345)
		cresp, err := svc.CreatePayment(context.Background(), &paymentapi.PaymentAttributes{
			ProcessingDate: date,
			Amount:         amount,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(cresp.Data).Should(HaveLen(1))

		resp, err := svc.GetPayment(context.Background(), &paymentapi.PaymentRequest{
			Id: cresp.Data[0].Id,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(cresp).Should(Equal(resp))
		Expect(cresp.Data[0].Attributes.ProcessingDate).Should(Equal(date))
		Expect(cresp.Data[0].Attributes.Amount).Should(Equal(amount))

		_, err = svc.DeletePayment(context.Background(), &paymentapi.PaymentRequest{
			Id: cresp.Data[0].Id,
		})
		Expect(err).NotTo(HaveOccurred())

		resp, err = svc.GetPayment(context.Background(), &paymentapi.PaymentRequest{
			Id: cresp.Data[0].Id,
		})
		Expect(err).Should(HaveOccurred())
	})

	It("should create multiple contents", func() {
		date := "2017-01-01"
		cresp, err := svc.CreatePayment(context.Background(), &paymentapi.PaymentAttributes{
			ProcessingDate: date,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(cresp.Data).Should(HaveLen(1))
		Expect(cresp.Data[0].Id).ShouldNot(BeEmpty())
		Expect(cresp.Data[0].Attributes.ProcessingDate).Should(Equal(date))

		amount := int64(43434)
		resp, err := svc.CreatePayment(context.Background(), &paymentapi.PaymentAttributes{
			Amount: amount,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.Data).Should(HaveLen(1))
		Expect(resp.Data[0].Id).ShouldNot(BeEmpty())
		Expect(resp.Data[0].Attributes.Amount).Should(Equal(amount))

		resp, err = svc.GetPayment(context.Background(), &paymentapi.PaymentRequest{})
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.Data).Should(HaveLen(2))
	})
})
