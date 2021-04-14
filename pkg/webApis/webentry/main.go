package main

import (
	"fmt"
	"net/http"

	constant "github.com/abhiranjan2429/vehicleparkingapp/pkg/pconstant"
	"github.com/abhiranjan2429/vehicleparkingapp/pkg/webApis/webentry/controller"
)

func main() {
	fmt.Println("webentry launched")

	controller.Startup()
	http.ListenAndServe(constant.EntryProducerPort, nil)
}
