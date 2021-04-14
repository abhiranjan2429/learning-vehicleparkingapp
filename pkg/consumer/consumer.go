package consumer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/abhiranjan2429/vehicleparkingapp/pkg/model"
	constant "github.com/abhiranjan2429/vehicleparkingapp/pkg/pconstant"
	"github.com/segmentio/kafka-go"
)

const (
	broker1Address = constant.Broker1Address
	broker2Address = constant.Broker2Address
	broker3Address = constant.Broker3Address
)

func Consume(ctx context.Context, topic string) *model.Vehicle {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address}, //, broker2Address, broker3Address},
		Topic:   topic,
		GroupID: "my-group",
	})
	defer r.Close()

	// the `ReadMessage` method blocks until we receive the next event
	msg, err := r.ReadMessage(ctx)
	if err != nil {
		fmt.Println(err)
		panic("could not read message " + err.Error())
	}
	v := &model.Vehicle{}
	err = json.Unmarshal(msg.Value, v)

	// after receiving the message, log its value
	fmt.Println("Vehicle Number is", v.VehicleNumber)
	fmt.Println("Vehicle's ParkingID is ", v.ParkingID)

	return v
}
