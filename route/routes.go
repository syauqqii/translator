package route

import (
	"net/http"
	"translator/controller"
)

// SetupRoutes => menetapkan rute HTTP
func SetupRoutes() {
	http.HandleFunc("/translate", controller.TranslateHandler)
}