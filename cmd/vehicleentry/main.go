package main

import (
	"context"
	"fmt"
	"os"

	constant "github.com/abhiranjan2429/vehicleparkingapp/pkg/pconstant"
	"github.com/abhiranjan2429/vehicleparkingapp/pkg/producer"
	"github.com/google/uuid"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No vehicle number found")
		os.Exit(0)
	}
	arg := os.Args[1]
	parkingID := uuid.NewString()
	fmt.Println("Vehicle entering no. is  ", arg, " and the parkingID is ", parkingID)

	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	fmt.Println("ENTRY publishing")

	producer.Produce(ctx, arg, parkingID, constant.Entrytopic)

}
