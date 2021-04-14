package main

import (
	"fmt"
	"net/http"

	constant "github.com/abhiranjan2429/vehicleparkingapp/pkg/pconstant"
	"github.com/abhiranjan2429/vehicleparkingapp/pkg/webApis/webbillingandtime/controller"
)

func main() {
	fmt.Println("webbillingandtime launched")
	controller.Startup()
	http.ListenAndServe(constant.BillingProducerPort, nil)
}
