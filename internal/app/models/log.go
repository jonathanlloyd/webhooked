package models

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type LogID string

type Log struct {
	ID LogID
}

func NewLog() Log {
	uid := uuid.NewV4()
	logID := fmt.Sprintf("log-%s", uid)
	return Log{
		ID: LogID(logID),
	}
}
