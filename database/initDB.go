package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func migrateDB(db *sql.DB) {
	migrations := &migrate.FileMigrationSource{
		Dir: "database/migrations",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)

	if err != nil {
		panic(err)
	}

	fmt.Println("applied migration", n, "times")
}

func InitDB(user string, name string, password string, host string, port string) *sql.DB {
	connStr := fmt.Sprintf("user=%v dbname=%v password=%v host=%v port=%v sslmode=disable", user, name, password, host, port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	migrateDB(db)
	return db
}
