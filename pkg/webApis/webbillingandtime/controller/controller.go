package controller

var (
	billingController billing
	prktimeController prktime
	homeController    home
)

func Startup() {
	billingController.registerRoutes()
	prktimeController.registerRoutes()
	homeController.registerRoutes()

}
