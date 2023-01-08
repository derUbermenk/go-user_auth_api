package user_service_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/derUbermenk/go-user_auth_api/service/user_service"
)

var _ = Describe("UserService", func() {
	us := userService.UserServ(UserRepositoryDouble{})


	Describe("CreateUser", func() {
		Context("when user creation succeeded", func() {
			It("returns a user, success true, and err nil", func() {
				new_user_request := user_service.NewUserRequest{
					Email:    "valid_email@email.com",
					Name:     "valid name",
					Password: "valid_password",
				}

				user, success, err := us.CreateUser(new_user_request)

				Expect(user).To(Not(BeNil()))
				Expect(success).To(BeTrue())
				Expect(err).To(BeNil())
			})
		})

		Context("when there is a missing required param", func() {
			It("returns a nil user, success false, and err nil") {
				new_user_request := user_service.NewUserRequest{
					Email:    "valid_email@email.com",
					Name:     "valid name",
					// missing password
				}

				user, success, err := us.Create(new_user_request)

				Expect(user).To(Not(BeNil()))
				Expect(success).To(BeTrue())
				Expect(err).To(BeNil())
			}
		})
	})
})
