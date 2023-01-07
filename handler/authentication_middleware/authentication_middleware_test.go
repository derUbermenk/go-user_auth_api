package authentication_middleware_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/derUbermenk/go-user_auth_api/handler/authentication_middleware"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AuthenticationMiddleware", func() {

	Describe("Authenticate", func() {
		It("sets a context specific value user_id when the user is authenticated", func() {
			req, _ := http.NewRequest("POST", "/authenticate", nil)
			req.AddCookie(
				&http.Cookie{
					Name:  "session",
					Value: "1",
				},
			)

			rec := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(rec)
			c.Request = req

			authentication_middleware.Authenticate()(c)
			user_id, user_id_exists := c.Keys["user_id"]

			Expect(user_id_exists).To(BeTrue())
			Expect(user_id).To(Equal("1"))
		})

		It("responds with status 401: unauthorized when the user is not authenticated", func() {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/authenticate", nil)

			c, _ := gin.CreateTestContext(rec)
			c.Request = req

			authentication_middleware.Authenticate()(c)

			Expect(c.Writer.Status()).To(Equal(401))
		})
	})
})
