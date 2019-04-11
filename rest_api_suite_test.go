package rest_api_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRestApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RestApi Suite")
}
