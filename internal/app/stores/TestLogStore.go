package stores

import (
	"errors"

	"github.com/jonathanlloyd/webhooked/internal/app/models"
)

type TestLogStore struct {
	logs map[models.LogID]models.Log
}

func NewTestLogStore() *TestLogStore {
	return &TestLogStore{
		logs: map[models.LogID]models.Log{},
	}
}

func (t *TestLogStore) InsertLog(log models.Log) error {
	if _, ok := t.logs[log.ID]; ok {
		return errors.New("A log with that ID already exists")
	}

	t.logs[log.ID] = log

	return nil
}

func (t *TestLogStore) Logs() []models.Log {
	var allLogs []models.Log
	for _, log := range t.logs {
		allLogs = append(allLogs, log)
	}
	return allLogs
}
