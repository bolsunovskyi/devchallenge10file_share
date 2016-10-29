package main

import (
	"file_share/config"
	"fmt"
	"net/http"
	"file_share/router"
	"github.com/gorilla/handlers"
	"time"
)

func main() {
	if !config.Read("") {
		fmt.Printf("Unable to load config file: %s\n", config.File)
	}

	appRouter := router.GetRouter()

	http.Handle("/", appRouter)

	server := http.Server{
		Addr: 		fmt.Sprintf(":%d", config.Config.Port),
		Handler: 	handlers.CORS(
			handlers.AllowedHeaders(
				[]string{"Access-Token", "File-Folder", "Content-Type"}),
			handlers.AllowedMethods([]string{"PATCH", "POST", "PUT", "DELETE", "GET", "OPTIONS"}))(appRouter),
		ReadTimeout: 	time.Hour,
		WriteTimeout:	time.Hour,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Server started on port: %d\n", config.Config.Port)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}