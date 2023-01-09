package user_repository_test

import (
	"log"
	"os"
	"os/exec"

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
			RunMigrations()
		})

		AfterEach(func() {
			DropDB()
		})

		It("creates a new user record given valid values", func() {
			db := ConnectTestDB()
			_ = user_repository.NewUserRepository(db)

			Expect(1).To(Equal(1))
		})

		It("creates a new user record given valid values", func() {

		})
	})
})
