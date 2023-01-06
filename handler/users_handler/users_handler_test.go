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
)

var _ = Describe("UsersHandler", func() {
	var r *gin.Engine

	Context("Non Authentication protected handlers",
		func() {
			Describe("Create", func() {

				BeforeEach(func() {
					r = gin.Default()
					r.POST("/user/create", users_handler.Create(&users_handler.UserServiceDouble{}))
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
						"email": "invalid_email@email.com",
						// with missing name
						"password": "invalid_password",
					}

					body, _ := json.Marshal(valid_input)

					rec := httptest.NewRecorder()
					req, _ := http.NewRequest("POST", "/user/create", bytes.NewReader(body))

					r.ServeHTTP(rec, req)
					Expect(rec.Result().StatusCode).To((Equal(400)))
				})
			})

			Describe("FetchByEmail", func() {
				BeforeEach(func() {
					r = gin.Default()
					r.HEAD("/user/:email", users_handler.FetchByEmail(&users_handler.UserServiceDouble{}))
				})

				It("responds with status 200: Ok when user with email exists", func() {
					rec := httptest.NewRecorder()
					req, _ := http.NewRequest("HEAD", "/user/existing_email@email.com", nil)

					r.ServeHTTP(rec, req)
					Expect(rec.Result().StatusCode).To((Equal(200)))

				})

				It("responds with status 404: Not Found when user with email does not exist", Pending, func() {
					rec := httptest.NewRecorder()
					req, _ := http.NewRequest("", "/user/non_existing_email@email.com", nil)

					r.ServeHTTP(rec, req)
					Expect(rec.Result().StatusCode).To((Equal(404)))
				})
			})
		})

	Context("Authentication protected handlers",
		func() {
			Describe("Fetch", func() {

				BeforeEach(func() {
					r = gin.Default()
					r.GET("/user/:id", users_handler.Fetch(&users_handler.UserServiceDouble{}))
				})

				It("responds with status 200: OK when user was found and retrieved", func() {
					rec := httptest.NewRecorder()
					req, _ := http.NewRequest("GET", "/user/1", nil)

					r.ServeHTTP(rec, req)

					Expect(rec.Result().StatusCode).To((Equal(200)))
				})

				It("responds with status 404: Not Found when user was not found", func() {
					rec := httptest.NewRecorder()
					req, _ := http.NewRequest("GET", "/user/2", nil)

					r.ServeHTTP(rec, req)

					Expect(rec.Result().StatusCode).To((Equal(404)))
				})
			})

		},
	)

	Context("Authentication and Authorization protected handlers", func() {
		Describe("FetchSelf", func() {
			BeforeEach(func() {
				r = gin.Default()
				r.GET("/user/:id/self", users_handler.FetchSelf(&users_handler.UserServiceDouble{}))
			})

			It("responds with status 200: OK when user was found and retrieved", func() {
				rec := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/user/1/self", nil)

				r.ServeHTTP(rec, req)

				Expect(rec.Result().StatusCode).To((Equal(200)))
			})

			It("responds with status 404: Not Found when user was not found", func() {
				rec := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/user/2/self", nil)

				r.ServeHTTP(rec, req)

				Expect(rec.Result().StatusCode).To((Equal(404)))
			})
		})

		Describe("Delete", func() {
			BeforeEach(func() {
				r = gin.Default()
				r.DELETE("/user/:id/destroy", users_handler.Delete(&users_handler.UserServiceDouble{}))
			})

			It("responds with status 200: OK when user was successfully deleted", func() {
				rec := httptest.NewRecorder()
				req, _ := http.NewRequest("DELETE", "/user/1/destroy", nil)

				r.ServeHTTP(rec, req)

				Expect(rec.Result().StatusCode).To((Equal(200)))
			})

			It("responds with status 404: Not Found when user to be destroyed does not exist", func() {
				rec := httptest.NewRecorder()
				req, _ := http.NewRequest("DELETE", "/user/2/destroy", nil)

				r.ServeHTTP(rec, req)

				Expect(rec.Result().StatusCode).To((Equal(404)))
			})
		})
	})
})
