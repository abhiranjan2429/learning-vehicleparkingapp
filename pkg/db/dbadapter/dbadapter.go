package dbadapter

import (
	"github.com/abhiranjan2429/vehicleparkingapp/pkg/model"
)

type (
	DatabaseAdapter interface {
		ReadData(string, model.DbSchema, interface{}) error
		SaveData(interface{}, model.DbSchema) error
	}
)
