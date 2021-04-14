package vehicleexitservice

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

func ExitService(ch chan model.Vehicle) {

	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking

	dbConnection := model.Connectionstring{
		Connectionstring: constant.DbConnectionSring,
	}
	dbSchema := model.DbSchema{
		Database:   constant.DbSchemaDatabse,
		Collection: constant.DbSchemaCollectionExit,
	}
	// vehicle := &model.Vehicle{
	// 	"KA-03-JJ-4455",
	// 	"1234567",
	// 	"",
	// }
	for {
		vehicle := consumer.Consume(ctx, constant.Exittopic)
		fmt.Println("exit serice", vehicle)
		var db dbadapter.DatabaseAdapter
		vehicle.Time = time.Now().Format(constant.TimeFormat)
		fmt.Println("Saving the entry to db", vehicle)
		db = mdb.New(dbConnection)
		db.SaveData(*vehicle, dbSchema)
		ch <- *vehicle
	}

}
