package model

type (
	Vehicle struct {
		VehicleNumber string `json :"vehiclenumber,omitempty" bson:"vehiclenumber,omitempty"`
		ParkingID     string `json :"parkingid,omitempty" bson:"parkingid,omitempty"`
		Time          string `json : "time,omitempty" bson : "time,omitempty"`
	}
	Connectionstring struct {
		Connectionstring string
	}
	DbSchema struct {
		Database   string
		Collection string
		UniqueKey  string
	}
	BillingandTime struct {
		ParkingID   string  `json :"parkingid,omitempty" bson:"parkingid,omitempty"`
		ParkingTime string  `json :"parkingtime,omitempty" bson:"parkingtime,omitempty"`
		Amount      float64 `json :"amount,omitempty" bson:"amount,omitempty"`
	}
	Bill struct {
		Bill string `json :bill,omitempty`
	}
	ParkignTime struct {
		PrkTime string `json: prktime,omitempty`
	}
)
