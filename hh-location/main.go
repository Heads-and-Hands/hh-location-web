package main

import (
	"github.com/gorilla/mux"
	"hh-location-web/hh-location/handlers"
	"hh-location-web/hh-location/middleware"
	"hh-location-web/hh-location/myAdminConfig"
	"hh-location-web/hh-location/provider"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	onStart()

	r := mux.NewRouter()

	r.Handle("/", middleware.WebCommonHandler(handlers.IndexHandler)).Methods("GET")

	r.Handle("/device", middleware.WebCommonHandler(handlers.DeviceGetHandler)).Methods("GET")
	r.Handle("/device", middleware.MobileCommonHandler(handlers.DevicePostHandler)).Methods("POST")
	r.Handle("/owner", middleware.MobileCommonHandler(handlers.OwnerPostHandler)).Methods("POST")

	r.Handle("/beacon", middleware.MobileCommonHandler(handlers.BeaconGetHandler)).Methods("GET")

	r.Handle("/position", middleware.WebCommonHandler(handlers.PositionGetHandler)).Methods("GET")
	r.Handle("/position", middleware.MobileCommonHandler(handlers.PositionPostHandler)).Methods("POST")

	log.Println("Hello admin!")
	m := myAdminConfig.Init()
	r.PathPrefix("/admin").Handler(m)

	http.ListenAndServe(":8877", r)
}

func onStart() {
	log.Println("Hello!")

	SetupCloseHandler()
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
