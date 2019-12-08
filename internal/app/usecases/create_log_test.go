package usecases_test

import (
	"strings"
	"testing"

	"github.com/jonathanlloyd/webhooked/internal/app/stores"
	"github.com/jonathanlloyd/webhooked/internal/app/usecases"
	"github.com/stretchr/testify/assert"
)

func TestCreateLogSuccess(t *testing.T) {
	logStore := stores.NewTestLogStore()

	err := usecases.CreateLog(logStore)
	assert.Nil(t, err)

	logs := logStore.Logs()
	assert.Equal(t, len(logs), 1)

	logID := string(logs[0].ID)
	assert.NotEqual(t, len(logID), 0)
	assert.True(t, strings.HasPrefix(logID, "log-"))
}
