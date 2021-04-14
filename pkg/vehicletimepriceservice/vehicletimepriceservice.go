package vehicletimepriceservice

import (
	"fmt"
	"time"

	"github.com/abhiranjan2429/vehicleparkingapp/pkg/db/dbadapter"
	"github.com/abhiranjan2429/vehicleparkingapp/pkg/db/mdb"
	"github.com/abhiranjan2429/vehicleparkingapp/pkg/model"
	constant "github.com/abhiranjan2429/vehicleparkingapp/pkg/pconstant"
)

func TimeandBillingService(ch chan model.Vehicle) {

	dbConnection := model.Connectionstring{
		Connectionstring: constant.DbConnectionSring,
	}
	dbSchemaExit := model.DbSchema{
		Database:   constant.DbSchemaDatabse,
		Collection: constant.DbSchemaCollectionExit,
		UniqueKey:  constant.UniqueKey,
	}
	dbSchemaEntry := model.DbSchema{
		Database:   constant.DbSchemaDatabse,
		Collection: constant.DbSchemaCollectionEntry,
		UniqueKey:  constant.UniqueKey,
	}
	dbSchemaBilling := model.DbSchema{
		Database:   constant.DbSchemaDatabse,
		Collection: constant.DbSchemaCollectionBilling,
	}
	//vehicle.Time = time.Now().Format("01-02-2006 15:04:05")
	for {

		vehicle := <-ch
		fmt.Println("exit serice", vehicle)
		var db dbadapter.DatabaseAdapter
		db = mdb.New(dbConnection)
		var entryVehicle, exitVehicle model.Vehicle
		err := db.ReadData(vehicle.ParkingID, dbSchemaEntry, &entryVehicle)
		err = db.ReadData(vehicle.ParkingID, dbSchemaExit, &exitVehicle)
		if err != nil {
			fmt.Println(err)
		}
		timedifference, ptime := calculateTime(entryVehicle.Time, exitVehicle.Time)
		//formattedTimedifference := time.Parse("", timedifference)
		// Here 0.99 is considered to 1 and 1.00 is considerd to be 2
		price := float64((int(timedifference) + 1) * int(constant.PricePerHour))
		// TODO take only parking time to calculate billing
		result := &model.BillingandTime{
			ParkingID:   vehicle.ParkingID,
			ParkingTime: ptime,
			Amount:      price,
		}
		db.SaveData(result, dbSchemaBilling)
	}

}
func calculateTime(entry string, exit string) (float64, string) {

	entryTime, err := time.Parse(constant.TimeFormat, entry)
	if err != nil {
		fmt.Println(err)
	}
	exitTime, err := time.Parse(constant.TimeFormat, exit)
	if err != nil {
		fmt.Println(err)
	}

	c := time.Duration((exitTime.Sub(entryTime).Minutes()))

	e := exitTime.Sub(entryTime)
	out := time.Time{}.Add(e)
	ptime := out.Format(constant.PrkTimeFormat)

	//t, _ := time.Parse("15:05", c.String())
	diff := exitTime.Sub(entryTime).Hours()
	fmt.Println(c)
	return diff, ptime
}
