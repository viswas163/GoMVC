package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/viswas163/UpstartProject/routes"
)

func main() {

	router := &routes.Router{R: mux.NewRouter()}
	router.HandleRoutes()

	server := &http.Server{
		Addr:    ":3001",
		Handler: router.R,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")
}
