package route

import (
	"net/http"
	"../controller"
)

type Route struct {
	Path		string
	Callback	func(w http.ResponseWriter, r *http.Request)
}

func Routes() []Route {
	var arrayRoutes []Route

	arrayRoutes = append(arrayRoutes, Route{Path: "/signup", Callback: controller.Signup})
	arrayRoutes = append(arrayRoutes, Route{Path: "/signup/:old_email", Callback: controller.Signup})
	arrayRoutes = append(arrayRoutes, Route{Path: "/signin", Callback: controller.Signin})

	return arrayRoutes
}