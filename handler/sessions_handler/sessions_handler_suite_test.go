package sessions_handler_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSessionsHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SessionsHandler Suite")
}
