package producer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/abhiranjan2429/vehicleparkingapp/pkg/model"
	constant "github.com/abhiranjan2429/vehicleparkingapp/pkg/pconstant"
	kafka "github.com/segmentio/kafka-go"
)

const (
	broker1Address = constant.Broker1Address
	broker2Address = constant.Broker2Address
	broker3Address = constant.Broker3Address
)

func Produce(ctx context.Context, vechicleNumber string, parkingID string, topic string) {
	v := &model.Vehicle{
		VehicleNumber: vechicleNumber,
		ParkingID:     parkingID,
	}
	vehicleJSON, _ := json.Marshal(v)

	fmt.Println(string(vehicleJSON))
	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address}, //, broker2Address, broker3Address},
		Topic:   topic,
	})
	defer w.Close()

	for {

		err := w.WriteMessages(ctx, kafka.Message{
			// create an arbitrary message payload for the value
			Value: (vehicleJSON),
			Key:   []byte(v.ParkingID),
		})
		if err == nil {
			fmt.Println("published", topic)
			//panic("could not write message  " + err.Error())
			break
		}
		fmt.Println(err)
	}

	// log a confirmation once the message is written
	fmt.Println("Vehicle with", vechicleNumber)

}
