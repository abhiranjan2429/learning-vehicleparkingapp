package main

import (
	"fmt"
	"net/http"

	constant "github.com/abhiranjan2429/vehicleparkingapp/pkg/pconstant"
	"github.com/abhiranjan2429/vehicleparkingapp/pkg/webApis/webexit/controller"
)

func main() {
	fmt.Println("webexit launched")

	controller.Startup()
	http.ListenAndServe(constant.ExitProducerPort, nil)
}
