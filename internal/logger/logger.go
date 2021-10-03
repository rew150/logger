package logger

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/rew150/logger/internal/db"
)

type LoggerService interface {
	GetAll(context.Context) ([]Log, error)
	Append(context.Context, Log) error
}

type service struct {
	db *sqlx.DB
}

var Service LoggerService

func init() {
	db, err := db.Create()
	if err != nil {
		log.Fatalf("Error init db: %s\n", err)
	}
	Service = service{
		db: db,
	}
}

func (s service) GetAll(ctx context.Context) (logs []Log, err error) {
	logs = []Log{}
	err = s.db.SelectContext(ctx, &logs, `
		SELECT * FROM logs
	`)
	return
}

func (s service) Append(ctx context.Context, log Log) (err error) {
	_, err = s.db.NamedExecContext(ctx, `
		INSERT INTO logs (
			ts,
			msg
		) VALUES (
			:ts,
			:msg
		)
	`, &log)
	return
}
