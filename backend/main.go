package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// "github.com/rs/cors"

	"backend/common"
	"backend/frameworks/persistence"
	"backend/frameworks/registry"
)

func main() {

	appKey := "saldkjhasdlkfjblwkjahwpdojabldpmfblaowidh"
	auth := common.NewAuth(appKey)
	db := persistence.NewDB()
	db.LogMode(true)
	defer db.Close()

	router := mux.NewRouter()
	reg := registry.NewRegistry(db, &auth)
	// apiRouter := router.PathPrefix("/api").Subrouter()

	userController := reg.NewUserHandler()
	router.HandleFunc("/signup", userController.Signup).Methods("POST")
	router.HandleFunc("/login", userController.Login).Methods("POST")

	// corsRouter := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8081", router))
}
