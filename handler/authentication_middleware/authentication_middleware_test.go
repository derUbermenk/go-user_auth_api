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
	var r *gin.Engine

	Describe("Authenticate", func() {
		BeforeEach(func() {
			r = gin.Default()
			r.POST("/authenticate", authentication_middleware.Authenticate())
		})

		It("sets a context specific value user_id when the user is authenticated", func() {
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)

			req, _ := http.NewRequest("POST", "/authenticate", nil)
			req.AddCookie(
				&http.Cookie{
					Name:  "session",
					Value: "1",
				},
			)

			r.ServeHTTP(rec, req)

			_, user_id_exists := c.Keys["user_id"]

			Expect(user_id_exists).To(BeTrue())
		})

		It("responds with status 401: unauthorized when the user is not authenticated", func() {
			rec := httptest.NewRecorder()

			req, _ := http.NewRequest("POST", "/authenticate", nil)
			r.ServeHTTP(rec, req)

			Expect(rec.Result().StatusCode).To(Equal(401))
		})
	})
})
