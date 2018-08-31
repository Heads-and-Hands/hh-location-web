package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"beacon/hh-location/middleware"
	"beacon/hh-location/handlers"
	"beacon/hh-location/provider"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	log.Println("Hello!")

	SetupCloseHandler()

	r := mux.NewRouter()

	r.Handle("/", middleware.WebClientCommonHandler(handlers.IndexHandler)).Methods("GET")
	r.Handle("/position", middleware.WebClientCommonHandler(handlers.PositionGetHandler)).Methods("GET")
	r.Handle("/device", middleware.WebClientCommonHandler(handlers.DeviceHandler)).Methods("GET")

	r.Handle("/beacon", middleware.MobileClientCommonHandler(handlers.BeaconGetHandler)).Methods("GET")
	r.Handle("/position", middleware.MobileClientCommonHandler(handlers.PositionPostHandler)).Methods("POST")

	http.ListenAndServe(":8877", r)

}

func onClose() {
	provider.GetProvider().Close()
	log.Println("Buy!")
}

func SetupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		onClose()
		os.Exit(0)
	}()
}