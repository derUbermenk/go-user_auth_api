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
			req, _ := http.NewRequest("HEAD", "/authenticate", nil)
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
			req, _ := http.NewRequest("HEAD", "/authenticate", nil)

			c, _ := gin.CreateTestContext(rec)
			c.Request = req

			authentication_middleware.Authenticate()(c)

			Expect(c.Writer.Status()).To(Equal(401))
		})
	})

	Describe("Authorize Owner", func() {
		It("responds with status 403: Forbidden when the client is not the owner of resource", func() {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("HEAD", "/validate", nil)

			c, _ := gin.CreateTestContext(rec)
			c.Request = req

			c.Set("user_id", "2")
			c.AddParam("id", "1")

			authentication_middleware.AuthorizeOwner()(c)

			Expect(c.Writer.Status()).To(Equal(403))
		})

		It("does not respond with status 403: Forbidden when the client is the owner of resource", func() {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("HEAD", "/validate", nil)

			c, _ := gin.CreateTestContext(rec)
			c.Request = req

			c.Set("user_id", "1")
			c.AddParam("id", "1")

			authentication_middleware.AuthorizeOwner()(c)

			Expect(c.Writer.Status()).To(Not(Equal(403)))
		})
	})
})
