package stores

import (
	"errors"

	"github.com/jonathanlloyd/webhooked/internal/app/models"
)

type TestLogStoreOption func(tls TestLogStore) TestLogStore

type ErrFunc func() bool

func WithLogStoreErrCallback(shouldErr ErrFunc) TestLogStoreOption {
	return func(tls TestLogStore) TestLogStore {
		tls.shouldErr = shouldErr
		return tls
	}
}

type TestLogStore struct {
	shouldErr ErrFunc
	logs      map[models.LogID]models.Log
}

func NewTestLogStore(options ...TestLogStoreOption) *TestLogStore {
	tls := TestLogStore{
		shouldErr: func() bool { return false },
		logs:      map[models.LogID]models.Log{},
	}
	for _, option := range options {
		tls = option(tls)
	}
	return &tls
}

func (t *TestLogStore) InsertLog(log models.Log) error {
	if t.shouldErr() {
		return errors.New("Something has gone horribly wrong :(")
	}

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
