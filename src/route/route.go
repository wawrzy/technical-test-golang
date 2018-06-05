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
	arrayRoutes = append(arrayRoutes, Route{Path: "/signin", Callback: controller.Signin})
	arrayRoutes = append(arrayRoutes, Route{Path: "/user", Callback: controller.User})
	arrayRoutes = append(arrayRoutes, Route{Path: "/ticket", Callback: controller.Ticket})
	arrayRoutes = append(arrayRoutes, Route{Path: "/ticket/close", Callback: controller.TicketClose})
	arrayRoutes = append(arrayRoutes, Route{Path: "/ticket/archive", Callback: controller.TicketArchive})
	arrayRoutes = append(arrayRoutes, Route{Path: "/message", Callback: controller.Message})

	return arrayRoutes
}