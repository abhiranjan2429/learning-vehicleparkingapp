package controller

import (
	"encoding/json"
	"net/http"

	"github.com/abhiranjan2429/vehicleparkingapp/pkg/db/dbadapter"
	"github.com/abhiranjan2429/vehicleparkingapp/pkg/db/mdb"
	"github.com/abhiranjan2429/vehicleparkingapp/pkg/model"
	constant "github.com/abhiranjan2429/vehicleparkingapp/pkg/pconstant"
)

type (
	prktime struct {
	}
)

func (e prktime) registerRoutes() {
	http.HandleFunc("/prktime", e.handleEntry)
}
func (e prktime) handleEntry(w http.ResponseWriter, r *http.Request) {
	dbConnection := model.Connectionstring{
		Connectionstring: constant.DbConnectionSring,
	}
	dbSchemaBilling := model.DbSchema{
		Database:   constant.DbSchemaDatabse,
		Collection: constant.DbSchemaCollectionBilling,
		UniqueKey:  constant.UniqueKey,
	}
	vehicle := &model.Vehicle{}
	err := json.NewDecoder(r.Body).Decode(vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var db dbadapter.DatabaseAdapter
	db = mdb.New(dbConnection)
	var b model.BillingandTime
	err = db.ReadData(vehicle.ParkingID, dbSchemaBilling, &b)
	result := model.ParkignTime{
		PrkTime: b.ParkingTime,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)

}
