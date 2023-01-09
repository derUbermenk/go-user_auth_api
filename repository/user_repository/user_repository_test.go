package user_repository_test

import (
	"log"
	"os"
	"os/exec"

	"github.com/derUbermenk/go-user_auth_api/repository/user_repository"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
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
	cmd := exec.Command("dropdb", "user_authed_api_test")
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

var _ = Describe("UserRepository", func() {
	// var db *sqlx.DB

	Describe("Create", func() {
		BeforeEach(func() {
			CreateDB()
			RunMigrations()
		})

		AfterEach(func() {
			DropDB()
		})

		It("creates a new user record given valid values", func() {
			_ = user_repository.NewUserRepository(db)
		})

		It("creates a new user record given valid values", func() {

		})

		// after reset db
	})
})
