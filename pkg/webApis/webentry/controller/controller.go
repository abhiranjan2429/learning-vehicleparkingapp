package controller

var (
	entryController entry
	homeController  home
)

func Startup() {
	entryController.registerRoutes()
	homeController.registerRoutes()
}
