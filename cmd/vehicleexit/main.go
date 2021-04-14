package main

import (
	"context"
	"fmt"
	"os"

	constant "github.com/abhiranjan2429/vehicleparkingapp/pkg/pconstant"
	"github.com/abhiranjan2429/vehicleparkingapp/pkg/producer"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("No vehicle number or parkingID found")
		os.Exit(0)
	}
	vehicleNumber := os.Args[1]
	parkingID := os.Args[2]
	fmt.Println("Vehicle exiting ", vehicleNumber)

	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	fmt.Println("EXIT publishing")
	producer.Produce(ctx, vehicleNumber, parkingID, constant.Exittopic)

}
