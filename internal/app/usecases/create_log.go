package usecases

import (
	"github.com/jonathanlloyd/webhooked/internal/app/models"
	"github.com/jonathanlloyd/webhooked/internal/app/stores"
)

func CreateLog(logStore stores.LogStore) error {
	log := models.NewLog()
	err := logStore.InsertLog(log)
	if err != nil {
		return err
	}
	return nil
}
