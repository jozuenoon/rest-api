package restapi_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/jozuenoon/rest-api/client"
	. "github.com/jozuenoon/rest-api/cmd/payment"
	"github.com/nsf/jsondiff"
)

var _ = Describe("RestApi", func() {
	var (
		tmpDir string
		err    error
		cancel func()
		ctx    context.Context
		srv    Server
	)

	client := NewAPIClient(&Configuration{
		BasePath: "http://127.0.0.1:8080",
	})

	api := client.PaymentServiceApi

	tmpDir, err = ioutil.TempDir("", "payment_database")
	ctx, cancel = context.WithCancel(context.Background())

	BeforeSuite(func() {
		srv = Server{
			DatabasePath: tmpDir,
		}
		go srv.Run(ctx)
	})

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

	It("should create object", func() {
		date := "2017-01-02"
		resp, _, err := api.CreatePayment(context.Background(), PaymentapiPaymentAttributes{
			ProcessingDate: date,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.Data).Should(HaveLen(1))
		Expect(resp.Data[0].Id).ShouldNot(BeEmpty())
		Expect(resp.Data[0].Attributes.ProcessingDate).Should(Equal(date))
	})

	It("should return created object", func() {
		date := "2017-01-03"
		cresp, _, err := api.CreatePayment(context.Background(), PaymentapiPaymentAttributes{
			ProcessingDate: date,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(cresp.Data).Should(HaveLen(1))
		resp, _, err := api.GetPayment(context.Background(), cresp.Data[0].Id)
		Expect(err).NotTo(HaveOccurred())
		Expect(cresp).Should(Equal(resp))
	})

	It("should change values on object update", func() {
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
		amount := "332234.6"
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

	It("should delete object", func() {
		date := "2017-01-02"
		amount := "3345.4"
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

		_, _, err = api.GetPayment(context.Background(), cresp.Data[0].Id)
		Expect(err).Should(HaveOccurred())
	})

	It("should create multiple resource objects", func() {
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

		amount := "43434.5"
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

	It("should create and retrieve complete resource data from json", func() {
		body := PaymentapiPaymentAttributes{}
		err := json.Unmarshal(exampleData, &body)
		Expect(err).NotTo(HaveOccurred())

		resp, _, err := api.CreatePayment(context.Background(), body)
		Expect(err).NotTo(HaveOccurred())

		gresp, _, err := api.GetPayment(context.Background(), resp.Data[0].Id)
		Expect(err).NotTo(HaveOccurred())
		Expect(gresp.Data).Should(HaveLen(1))

		payload1, err := json.Marshal(gresp.Data[0].Attributes)
		Expect(err).NotTo(HaveOccurred())

		itsEqual, err := AreEqualJSON(payload1, exampleData)
		Expect(err).NotTo(HaveOccurred())
		Expect(itsEqual).Should(BeTrue())
	})

	It("should reject non ULID id", func() {
		gresp, hresp, err := api.GetPayment(context.Background(), "some_invalid_id")
		Expect(err).To(HaveOccurred())
		Expect(hresp.StatusCode).To(Equal(400))
		Expect(gresp.Data).To(BeNil())
	})

	It("should error on non existing object", func() {
		gresp, hresp, err := api.GetPayment(context.Background(), "01D9CK9NF4JQSFZB9P1K3S1QC7")
		Expect(err).To(HaveOccurred())
		Expect(hresp.StatusCode).To(Equal(404))
		Expect(gresp.Data).To(BeNil())
	})
})

var exampleData = []byte(`
{
	"amount": "100.21",
	"numeric_reference": "1002001",
	"payment_id": "123456789012345678",
	"payment_purpose": "Paying for goods/services",
	"payment_scheme": "FPS",
	"payment_type": "Credit",
	"processing_date": "2017-01-18",
	"reference": "Payment for Em's piano lessons",
	"scheme_payment_sub_type": "InternetBanking",
	"scheme_payment_type": "ImmediatePayment",
	"currency": "GBP",
	"end_to_end_reference": "Wil piano Jan",
	"beneficiary_party": {
		"account_name": "W Owens",
		"account_number": "31926819",
		"account_number_code": "BBAN",
		"account_type": 1,
		"address": "1 The Beneficiary Localtown SE2",
		"bank_id": "403000",
		"bank_id_code": "GBDSC",
		"name": "Wilfred Jeremiah Owens"
	},
	"charges_information": {
		"bearer_code": "SHAR",
		"sender_charges": [
		{
			"amount": "5.00",
			"currency": "GBP"
		},
		{
			"amount": "10.00",
			"currency": "USD"
		}
		],
		"receiver_charges_amount": "1.00",
		"receiver_charges_currency": "USD"
	},
	"debtor_party": {
		"account_name": "EJ Brown Black",
		"account_number": "GB29XABC10161234567801",
		"account_number_code": "IBAN",
		"address": "10 Debtor Crescent Sourcetown NE1",
		"bank_id": "203301",
		"bank_id_code": "GBDSC",
		"name": "Emelia Jane Brown"
	},
	"fx": {
		"contract_reference": "FX123",
		"exchange_rate": "2.00000",
		"original_amount": "200.42",
		"original_currency": "USD"
	},
	"sponsor_party": {
		"account_number": "56781234",
		"bank_id": "123123",
		"bank_id_code": "GBDSC"
	}
}
`)

// Inspired by https://gist.github.com/turtlemonvh/e4f7404e28387fadb8ad275a99596f67
func AreEqualJSON(payload1, payload2 []byte) (bool, error) {
	var data1 interface{}
	var data2 interface{}

	err := json.Unmarshal(payload1, &data1)
	if err != nil {
		return false, fmt.Errorf("failed to mashal payload 1 :: %s", err.Error())
	}
	err = json.Unmarshal(payload2, &data2)
	if err != nil {
		return false, fmt.Errorf("failed to mashal payload 2 :: %s", err.Error())
	}

	if !reflect.DeepEqual(data1, data2) {
		opts := jsondiff.DefaultConsoleOptions()
		diff, info := jsondiff.Compare(payload2, payload1, &opts)
		return false, fmt.Errorf("%s\n====\n%s", info, diff.String())
	}

	return true, nil
}
