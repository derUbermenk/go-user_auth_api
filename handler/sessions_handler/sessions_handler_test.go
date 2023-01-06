package sessions_handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/derUbermenk/go-user_auth_api/handler/sessions_handler"
)

var _ = Describe("SessionsHandler", func() {
	var r *gin.Engine

	Describe("Create", func() {
		BeforeEach(func() {
			r = gin.Default()
			r.POST("/sessions/create", sessions_handler.Create(&sessions_handler.SessionServiceDouble{}))
		})

		It("responds with 200: OK when login succeeds", func() {
			valid_credentials := map[string]interface{}{
				"email":    "valid_email@email.com",
				"password": "valid_password@email.com",
			}

			body, _ := json.Marshal(valid_credentials)

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/sessions/create", bytes.NewReader(body))

			r.ServeHTTP(rec, req)
			Expect(rec.Result().StatusCode).To((Equal(200)))
		})

		It("responds with status 400: bad request when there are missing params or invalid credentials", func() {
			invalid_credentials := map[string]interface{}{
				"email":    "invalid_email@email.com",
				"password": "invalid_password@email.com",
			}

			body, _ := json.Marshal(invalid_credentials)

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/sessions/create", bytes.NewReader(body))

			r.ServeHTTP(rec, req)
			Expect(rec.Result().StatusCode).To((Equal(401)))
		})
	})
})
