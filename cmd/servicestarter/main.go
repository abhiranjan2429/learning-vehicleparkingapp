package main

import (
	"fmt"

	"github.com/abhiranjan2429/vehicleparkingapp/pkg/model"
	"github.com/abhiranjan2429/vehicleparkingapp/pkg/vehicleentryservice"
	"github.com/abhiranjan2429/vehicleparkingapp/pkg/vehicleexitservice"
	"github.com/abhiranjan2429/vehicleparkingapp/pkg/vehicletimepriceservice"
)

func main() {
	fmt.Println("servicestarter launched.")
	wait := make(chan int)
	sync := make(chan model.Vehicle, 10)

	go vehicleentryservice.EntryService()
	go vehicleexitservice.ExitService(sync)
	go vehicletimepriceservice.TimeandBillingService(sync)
	<-wait
}
