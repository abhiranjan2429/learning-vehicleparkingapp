package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/abhiranjan2429/vehicleparkingapp/pkg/model"
	constant "github.com/abhiranjan2429/vehicleparkingapp/pkg/pconstant"
	"github.com/abhiranjan2429/vehicleparkingapp/pkg/producer"
	"github.com/google/uuid"
)

type (
	entry struct {
	}
)

func (e entry) registerRoutes() {
	http.HandleFunc("/entry", e.handleEntry)
}
func (e entry) handleEntry(w http.ResponseWriter, r *http.Request) {
	vehicle := &model.Vehicle{}
	err := json.NewDecoder(r.Body).Decode(vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	parkingID := uuid.NewString()
	vehicle.ParkingID = parkingID
	fmt.Println("Vehicle entering no. is  ", vehicle.VehicleNumber, " and the parkingID is ", vehicle.ParkingID)

	// create a new context
	ctx := context.Background()
	fmt.Println("ENTRY publishing")
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	producer.Produce(ctx, vehicle.VehicleNumber, vehicle.ParkingID, constant.Entrytopic)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehicle)
}
