package users_handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/derUbermenk/go-user_auth_api/handler/users_handler"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// "github.com/derUbermenk/go-user_auth_api/handler/users_handler"
)

var _ = Describe("UsersHandler", func() {
	var r *gin.Engine

	Context(
		"Non Authentication protected routes",
		func() {
			Describe("Create", func() {

				BeforeEach(func() {
					r = gin.Default()
					r.POST("/user/create", users_handler.Create(FakeUserService{}))
				})

				It("responds with status 200: OK when user has been created", func() {
					valid_input := map[string]interface{}{
						"email":    "valid_email@email.com",
						"name":     "valid_name",
						"password": "valid_password",
					}

					body, _ := json.Marshal(valid_input)

					rec := httptest.NewRecorder()
					req, _ := http.NewRequest("POST", "/user/create", bytes.NewReader(body))

					r.ServeHTTP(rec, req)

					Expect(rec.Result().StatusCode).To((Equal(200)))
				})

				It("responds with status 400: bad request when there are missing params or invalid inputs", func() {
					valid_input := map[string]interface{}{
						"email":    "invalid_email@email.com",
						"name":     "invalid_name",
						"password": "invalid_password",
					}

					body, _ := json.Marshal(valid_input)

					rec := httptest.NewRecorder()
					req, _ := http.NewRequest("POST", "/user/create", bytes.NewReader(body))

					r.ServeHTTP(rec, req)
					Expect(rec.Result().StatusCode).To((Equal(400)))
				})
			})

			Describe("Show_by_email", func() {
				It("responds with status 404: Not Found when user with email does not exist", func() {

				})

				It("responds with status 200: Ok when user with email exists", func() {

				})
			})
		})
})
