package route

import (
	"net/http"
	"translator/controller"
	"translator/helper"
)

func SetupRoutes() {
	http.HandleFunc("/translate", helper.RateLimitedHandler(controller.TranslateHandler))
}
