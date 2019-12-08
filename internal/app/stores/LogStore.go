package stores

import "github.com/jonathanlloyd/webhooked/internal/app/models"

type LogStore interface {
	InsertLog(log models.Log) error
}
