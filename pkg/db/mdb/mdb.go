package mdb

import (
	"context"
	"fmt"
	"sync"

	"github.com/abhiranjan2429/vehicleparkingapp/pkg/db/dbadapter"
	"github.com/abhiranjan2429/vehicleparkingapp/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodb struct {
	connectionstring model.Connectionstring
}

var clientInstance *mongo.Client

//Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error

//Used to execute client creation procedure only once.
var mongoOnce sync.Once

func New(db model.Connectionstring) dbadapter.DatabaseAdapter {

	m := &Mongodb{
		connectionstring: db,
	}
	//getMongoClient(m.database)
	return m
}

func (m_db *Mongodb) ReadData(key string, db model.DbSchema, model interface{}) error {

	filter := bson.D{primitive.E{Key: db.UniqueKey, Value: key}}
	client, err := getMongoClient(m_db.connectionstring)
	if err != nil {
		return err
	}
	collection := client.Database(db.Database).Collection(db.Collection)
	err = collection.FindOne(context.TODO(), filter).Decode(model)
	return nil
}
func (m_db *Mongodb) SaveData(model interface{}, db model.DbSchema) error {

	v, _ := bson.Marshal(model)
	client, err := getMongoClient(m_db.connectionstring)
	if err != nil {
		return err
	}
	collection := client.Database(db.Database).Collection(db.Collection)
	result, err := collection.InsertOne(context.TODO(), v)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil

}

//GetMongoClient - Return mongodb connection to work with
func getMongoClient(m_db model.Connectionstring) (*mongo.Client, error) {

	//Perform connection creation operation only once.
	mongoOnce.Do(func() {

		clientOptions := options.Client().ApplyURI(m_db.Connectionstring)

		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}
