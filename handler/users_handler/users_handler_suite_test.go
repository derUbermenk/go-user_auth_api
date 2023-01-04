package users_handler_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUsersHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UsersHandler Suite")
}
