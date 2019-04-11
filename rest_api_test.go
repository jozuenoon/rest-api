package rest_api_test

import (
	"context"
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/jozuenoon/rest-api/client"
	"github.com/jozuenoon/rest-api/cmd/payment"
)

var _ = Describe("RestApi", func() {
	var (
		tmpDir string
		err    error
		cancel func()
		ctx    context.Context
		srv    *payment.Server
	)

	client := NewAPIClient(&Configuration{
		BasePath: "http://127.0.0.1:8080",
	})

	api := client.PaymentServiceApi

	tmpDir, err = ioutil.TempDir("", "payment_database")
	ctx, cancel = context.WithCancel(context.Background())
	srv = &payment.Server{
		DatabasePath: tmpDir,
	}
	go srv.Run(ctx)

	AfterSuite(func() {
		cancel()
		os.RemoveAll(tmpDir)
	})

	It("should not error", func() {
		Expect(err).NotTo(HaveOccurred())
	})

	It("should be empty if nothing is loaded", func() {
		resp, _, err := api.GetPayment2(context.Background(), nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.Data).Should(HaveLen(0))
	})

	It("should create content", func() {
		date := "2017-01-01"
		resp, _, err := api.CreatePayment(context.Background(), PaymentapiPaymentAttributes{
			ProcessingDate: date,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.Data).Should(HaveLen(1))
		Expect(resp.Data[0].Id).ShouldNot(BeEmpty())
		Expect(resp.Data[0].Attributes.ProcessingDate).Should(Equal(date))
	})

	It("should return created content", func() {
		date := "2017-01-01"
		cresp, _, err := api.CreatePayment(context.Background(), PaymentapiPaymentAttributes{
			ProcessingDate: date,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(cresp.Data).Should(HaveLen(1))
		resp, _, err := api.GetPayment(context.Background(), cresp.Data[0].Id)
		Expect(err).NotTo(HaveOccurred())
		Expect(cresp).Should(Equal(resp))
	})

	It("should change values with update", func() {
		date := "2017-01-02"
		cresp, _, err := api.CreatePayment(context.Background(), PaymentapiPaymentAttributes{
			ProcessingDate: date,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(cresp.Data).Should(HaveLen(1))

		resp, _, err := api.GetPayment(context.Background(), cresp.Data[0].Id)
		Expect(err).NotTo(HaveOccurred())
		Expect(cresp).Should(Equal(resp))
		Expect(cresp.Data[0].Attributes.ProcessingDate).Should(Equal(date))
		Expect(cresp.Data[0].Attributes.Amount).Should(Equal(""))

		date2 := "2018-01-02"
		amount := "332234"
		uresp, _, err := api.UpdatePayment(context.Background(), cresp.Data[0].Id, PaymentapiPaymentAttributes{
			ProcessingDate: date2,
			Amount:         amount,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(uresp.Data).Should(HaveLen(1))
		Expect(uresp.Data[0].Attributes.ProcessingDate).Should(Equal(date2))
		Expect(uresp.Data[0].Attributes.Amount).Should(Equal(amount))

		resp, _, err = api.GetPayment(context.Background(), cresp.Data[0].Id)
		Expect(err).NotTo(HaveOccurred())
		Expect(uresp).Should(Equal(resp))
		Expect(uresp.Data[0].Attributes.ProcessingDate).Should(Equal(date2))
		Expect(uresp.Data[0].Attributes.Amount).Should(Equal(amount))
	})

	It("should delete payment", func() {
		date := "2017-01-02"
		amount := "3345"
		cresp, _, err := api.CreatePayment(context.Background(), PaymentapiPaymentAttributes{
			ProcessingDate: date,
			Amount:         amount,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(cresp.Data).Should(HaveLen(1))

		resp, _, err := api.GetPayment(context.Background(), cresp.Data[0].Id)
		Expect(err).NotTo(HaveOccurred())
		Expect(cresp).Should(Equal(resp))
		Expect(cresp.Data[0].Attributes.ProcessingDate).Should(Equal(date))
		Expect(cresp.Data[0].Attributes.Amount).Should(Equal(amount))

		_, _, err = api.DeletePayment(context.Background(), cresp.Data[0].Id)
		Expect(err).NotTo(HaveOccurred())

		resp, _, err = api.GetPayment(context.Background(), cresp.Data[0].Id)
		Expect(err).Should(HaveOccurred())
	})

	It("should create multiple contents", func() {
		resp, _, err := api.GetPayment2(context.Background(), nil)
		Expect(err).NotTo(HaveOccurred())
		existingEnteries := len(resp.Data)

		date := "2017-01-01"
		cresp, _, err := api.CreatePayment(context.Background(), PaymentapiPaymentAttributes{
			ProcessingDate: date,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(cresp.Data).Should(HaveLen(1))
		Expect(cresp.Data[0].Id).ShouldNot(BeEmpty())
		Expect(cresp.Data[0].Attributes.ProcessingDate).Should(Equal(date))

		amount := "43434"
		resp, _, err = api.CreatePayment(context.Background(), PaymentapiPaymentAttributes{
			Amount: amount,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.Data).Should(HaveLen(1))
		Expect(resp.Data[0].Id).ShouldNot(BeEmpty())
		Expect(resp.Data[0].Attributes.Amount).Should(Equal(amount))

		resp, _, err = api.GetPayment2(context.Background(), nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.Data).Should(HaveLen(existingEnteries + 2))
	})
})
