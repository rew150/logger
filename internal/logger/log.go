package logger

import "time"

type Log struct {
	Timestamp time.Time `json:"timestamp" db:"ts"`
	Message   string    `json:"message" db:"msg"`
}
