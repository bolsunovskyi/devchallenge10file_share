package router

import "github.com/gorilla/mux"

func GetRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range appRoutes {
		router.
		Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
