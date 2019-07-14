package rest_api

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRestApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RestApi Suite")
}
