package constant

//This is momentarily here later this will be moved to config file
const (
	//KAFKA
	Exittopic      = "exit_topic"
	Entrytopic     = "entry_topic"
	Broker1Address = "broker:9092"
	//Broker1Address = "localhost:9093"
	Broker2Address = "localhost:9094"
	Broker3Address = "localhost:9095"

	//DB
	//DbConnectionSring = "mongodb://localhost:27017"
	DbConnectionSring = "mongodb://mongodbdc:27017"
	DbSchemaDatabse   = "plot"
	UniqueKey         = "parkingid"
	//entry
	DbSchemaCollectionEntry = "vehicleentry"

	//exit
	DbSchemaCollectionExit = "vehicleexit"
	//billing
	DbSchemaCollectionBilling = "billing"
	//MISC
	TimeFormat    = "2006-01-02T15:04:05.000Z"
	PricePerHour  = 10.0
	PrkTimeFormat = "15:04:05"
	// Api ports provide only port number as string with : as suffix example ":443" or "":8080"
	EntryProducerPort   = ":9011"
	ExitProducerPort    = ":9012"
	BillingProducerPort = ":9013"
)
