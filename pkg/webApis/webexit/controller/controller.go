package controller

var (
	exitController exit
	homeController home
)

func Startup() {
	exitController.registerRoutes()
	homeController.registerRoutes()
}
