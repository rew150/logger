package logger

type Log struct {
	Timestamp string `json:"timestamp" db:"ts"`
	Message   string `json:"message" db:"msg"`
}
