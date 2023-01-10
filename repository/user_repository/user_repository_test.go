package user_repository_test

import (
	"log"
	"os"
	"os/exec"

	"github.com/derUbermenk/go-user_auth_api/repository"
	"github.com/derUbermenk/go-user_auth_api/repository/user_repository"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func CreateDB() {
	cmd := exec.Command("createdb", "user_authed_api_test")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		log.Printf("%v\n", err.Error())
		AbortSuite(err.Error())
	}
}

func DropDB() {
	// cmd := exec.Command("dropdb", "user_authed_api_test")
	migrations_path := "file:/home/chester/Documents/code_projects/go-projects/user_auth_api/repository/migrations"
	db_connstring := "postgres://chester:baba_yetu@localhost:5432/user_authed_api_test?sslmode=disable"
	cmd := exec.Command("migrate", "-source", migrations_path, "-database", db_connstring, "drop", "-f")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		AbortSuite(err.Error())
	}
}

func RunMigrations() {

	/*
		m, err := migrate.New(
			"file:/home/chester/Documents/code_projects/go-projects/user_auth_api/repository/migrations",
			"postgres://chester:baba_yetu@localhost:5432/user_authed_api_test?sslmode=disable",
		)

		if err != nil {
			AbortSuite(err.Error())
		}

		if err := m.Up(); err != nil {
			AbortSuite(err.Error())
		}
	*/

	// hacky way, above does not allow drop database

	migrations_path := "file:/home/chester/Documents/code_projects/go-projects/user_auth_api/repository/migrations"
	db_connstring := "postgres://chester:baba_yetu@localhost:5432/user_authed_api_test?sslmode=disable"

	cmd := exec.Command("migrate", "-source", migrations_path, "-database", db_connstring, "up")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		AbortSuite(err.Error())
	}
}

func ConnectTestDB() *sqlx.DB {
	db_connstring := "postgres://chester:baba_yetu@localhost:5432/user_authed_api_test?sslmode=disable"
	db, err := sqlx.Connect("postgres", db_connstring)

	if err != nil {
		AbortSuite(err.Error())
	}

	if err := db.Ping(); err != nil {
		AbortSuite(err.Error())
	}

	return db
}

var _ = Describe("UserRepository", func() {

	Describe("Create", func() {
		BeforeEach(func() {
			DropDB()
			RunMigrations()
		})

		AfterEach(func() {
			DropDB()
		})

		It("returns a user with the id of the created record", func() {
			db := ConnectTestDB()
			ur := user_repository.NewUserRepository(db)

			input := map[string]interface{}{
				"email":    "email@email.com",
				"password": "user_password",
			}

			user, err := ur.Create(input)

			Expect(user).To(Equal(repository.User{ID: 1}))
			Expect(err).To(BeNil())
		})

		It("returns a user value repository.User when the entered user email is already in the db", func() {
			db := ConnectTestDB()
			ur := user_repository.NewUserRepository(db)

			input := map[string]interface{}{
				"email":    "email@email.com",
				"password": "user_password",
			}

			ur.Create(input)
			user, err := ur.Create(input)

			Expect(user).To(Equal(repository.User{}))
			Expect(err).ToNot(BeNil())
		})
	})

	Describe("FindByEmail", func() {
		BeforeEach(func() {
			DropDB()
			RunMigrations()
		})

		AfterEach(func() {
			DropDB()
		})

		It("returns the expected user record", func() {
			db := ConnectTestDB()
			ur := user_repository.NewUserRepository(db)

			input := map[string]interface{}{
				"email":    "email@email.com",
				"password": "user_password",
			}

			_, err := ur.Create(input)

			Expect(err).To(BeNil())

			user, err := ur.FindByEmail("email@email.com")
			Expect(user).To(Equal(repository.User{
				ID:    1,
				Email: "email@email.com",
			}))

			Expect(err).To(BeNil())
		})

		It("does not include the password of the user", func() {
			db := ConnectTestDB()
			ur := user_repository.NewUserRepository(db)

			input := map[string]interface{}{
				"email":    "email@email.com",
				"password": "user_password",
			}

			_, err := ur.Create(input)

			Expect(err).To(BeNil())

			user, err := ur.FindByEmail("email@email.com")
			Expect(user.Password).To(BeEmpty())

			Expect(err).To(BeNil())
		})

		Context("when the user with email does not exist", func() {
			It("returns a zero valued user", func() {
				db := ConnectTestDB()
				ur := user_repository.NewUserRepository(db)

				input := map[string]interface{}{
					"email":    "email@email.com",
					"password": "user_password",
				}

				_, err := ur.Create(input)

				Expect(err).To(BeNil())

				user, err := ur.FindByEmail("non_existing_email@email.com")
				Expect(user).To(BeZero())
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("FetchUser", func() {
		BeforeEach(func() {
			DropDB()
			RunMigrations()
		})

		AfterEach(func() {
			DropDB()
		})

		It("returns the specific user without exposing the password", func() {
			db := ConnectTestDB()
			ur := user_repository.NewUserRepository(db)

			input := map[string]interface{}{
				"email":    "email@email.com",
				"password": "user_password",
			}

			_, err := ur.Create(input)

			Expect(err).To(BeNil())

			user, err := ur.FetchUser(1)
			Expect(user).To(Equal(repository.User{
				ID:    1,
				Email: "email@email.com",
			}))

			Expect(err).To(BeNil())
		})

		It("returns returns a zero value user when the user with given user id does not exist", func() {
			db := ConnectTestDB()
			ur := user_repository.NewUserRepository(db)

			input := map[string]interface{}{
				"email":    "email@email.com",
				"password": "user_password",
			}

			_, err := ur.Create(input)

			Expect(err).To(BeNil())

			user, err := ur.FetchUser(77)
			Expect(user).To(BeZero())

			Expect(err).To(BeNil())
		})
	})
})
