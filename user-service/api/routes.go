package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/social-media/user-service/controller"
)

func BuildRouter(host string) {
	v1Router := mux.NewRouter()
	routeAPIs(v1Router)
	srv := &http.Server{
		Handler: v1Router,
		Addr:    host,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func routeAPIs(router *mux.Router) {
	router.HandleFunc("/user", controller.GetUser).Methods("GET")
	router.HandleFunc("/user", controller.CreateUser).Methods("POST")
}
