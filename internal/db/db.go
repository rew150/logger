package db

import (
	"fmt"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Create() (*sqlx.DB, error) {
	host := strings.TrimSpace(os.Getenv("POSTGRES_HOST"))
	port := "5432"
	user := strings.TrimSpace(os.Getenv("POSTGRES_USER"))
	password := strings.TrimSpace(os.Getenv("POSTGRES_PASSWORD"))
	db := strings.TrimSpace(os.Getenv("POSTGRES_DB"))

	info := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		db,
	)
	return sqlx.Connect("postgres", info)
}
