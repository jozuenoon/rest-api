package paymentservice_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPaymentservice(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Paymentservice Suite")
}
