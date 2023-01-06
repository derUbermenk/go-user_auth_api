package authentication_middleware_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAuthenticationMiddleware(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AuthenticationMiddleware Suite")
}
