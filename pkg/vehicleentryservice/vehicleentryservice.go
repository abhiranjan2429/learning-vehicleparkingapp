package vehicleentryservice

import (
	"context"
	"fmt"
	"time"

	"github.com/abhiranjan2429/vehicleparkingapp/pkg/consumer"
	"github.com/abhiranjan2429/vehicleparkingapp/pkg/db/dbadapter"
	"github.com/abhiranjan2429/vehicleparkingapp/pkg/db/mdb"
	"github.com/abhiranjan2429/vehicleparkingapp/pkg/model"
	constant "github.com/abhiranjan2429/vehicleparkingapp/pkg/pconstant"
)

func EntryService() {

	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	var db dbadapter.DatabaseAdapter
	dbConnection := model.Connectionstring{
		Connectionstring: constant.DbConnectionSring,
	}
	dbSchema := model.DbSchema{
		Database:   constant.DbSchemaDatabse,
		Collection: constant.DbSchemaCollectionEntry,
	}
	for {
		vehicle := consumer.Consume(ctx, constant.Entrytopic)
		vehicle.Time = time.Now().Format(constant.TimeFormat)
		fmt.Println("Saving the entry to db", vehicle)
		db = mdb.New(dbConnection)
		db.SaveData(*vehicle, dbSchema)

	}

	// vehicle := &model.Vehicle{
	// 	"KA-03-JJ-4455",
	// 	"1234567",
	// 	"",
	// }

}
