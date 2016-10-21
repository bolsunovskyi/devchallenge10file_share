package main

import (
	"file_share/config"
	"fmt"
	"net/http"
	"file_share/router"
	"time"
)

func main() {
	if !config.Read() {
		fmt.Printf("Unable to load config file: %s\n", config.File)
	}

	appRouter := router.GetRouter()

	http.Handle("/", appRouter)

	server := http.Server{
		Addr: 		fmt.Sprintf(":%d", config.Config.Port),
		Handler: 	appRouter,
		ReadTimeout: 	time.Hour,
		WriteTimeout:	time.Hour,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}