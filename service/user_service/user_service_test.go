package user_service_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/derUbermenk/go-user_auth_api/service/user_service"
)

var _ = Describe("UserService", func() {
	us := user_service.NewUserService(&user_service.UserRepositoryDouble{})

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
			It("returns a nil user, success false, and err nil", func() {
				new_user_request := user_service.NewUserRequest{
					Email: "valid_email@email.com",
					Name:  "valid name",
					// missing password
				}

				user, success, err := us.CreateUser(new_user_request)

				Expect(user).To((BeNil()))
				Expect(success).To(BeFalse())
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("FetchUserByEmail", func() {
		It("returns a user when it exists", func() {
			email := "existing_user_email@email.com"
			user, err := us.FetchUserByEmail(email)

			Expect(user).To(Not(BeNil()))
			Expect(err).To(BeNil())
		})

		It("returns a nil when it does not exist", func() {
			email := "non_existing_user_email@email.com"
			user, err := us.FetchUserByEmail(email)

			Expect(user).To((BeNil()))
			Expect(err).To(BeNil())
		})
	})

	Describe("FetchUser", func() {
		It("returns the user when it exists", func() {
			existing_id := 1
			user, err := us.FetchUser(existing_id)

			Expect(user).To(Not(BeNil()))
			Expect(err).To(BeNil())
		})

		It("returns a nil user when it does not exist", func() {
			existing_id := 2
			user, err := us.FetchUser(existing_id)

			Expect(user).To((BeNil()))
			Expect(err).To(BeNil())
		})
	})

	Describe("FetchUserSelf", func() {
		It("returns the user when it exists", func() {
			existing_id := 1
			user, err := us.FetchUserSelf(existing_id)

			Expect(user).To(Not(BeNil()))
			Expect(err).To(BeNil())
		})

		It("returns a nil user when it does not exist", func() {
			existing_id := 2
			user, err := us.FetchUserSelf(existing_id)

			Expect(user).To((BeNil()))
			Expect(err).To(BeNil())
		})
	})
})
