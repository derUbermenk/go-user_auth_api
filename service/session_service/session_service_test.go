package session_service_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/derUbermenk/go-user_auth_api/service/session_service"
)

var _ = Describe("SessionService", func() {
	ss := session_service.NewSessionService(session_service.SessionRepositoryDouble{})

	Describe("CreateSession", func() {
		It("returns the id of the User, validity to true, when the credentials are valid", func() {
			valid_credentials := session_service.NewSessionRequest{
				Email:    "existing_email@email.com",
				Password: "valid_password",
			}

			user_id, is_validvalid_credentials, err := ss.CreateSession(valid_credentials)

			Expect(user_id).To(Not(BeZero()))
			Expect(is_validvalid_credentials).To(BeTrue())
			Expect(err).To(BeNil())
		})

		It("does not return the id of the User, and valid_credentials to false when the credentials are invalid", func() {
			valid_credentials := session_service.NewSessionRequest{
				Email:    "existing_email@email.com",
				Password: "invalid_password",
			}

			user_id, is_validvalid_credentials, err := ss.CreateSession(valid_credentials)
			Expect(user_id).To(BeZero())
			Expect(is_validvalid_credentials).To(BeFalse())
			Expect(err).To(BeNil())
		})

		It("returns user_id to 0 when the given email does not exist", func() {
			invalid_credentials := session_service.NewSessionRequest{
				Email:    "non_existing_email@email.com",
				Password: "invalid_password",
			}

			user_id, is_validvalid_credentials, err := ss.CreateSession(invalid_credentials)
			Expect(user_id).To(BeZero())
			Expect(is_validvalid_credentials).To(BeFalse())
			Expect(err).To(BeNil())
		})
	})
})
