package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"backend/common"
	"backend/frameworks/persistence"
	"backend/frameworks/registry"
)

func main() {

	dbhost := flag.String("dbhost", "localhost:3306", "some bad help")
	flag.Parse()

	appKey := "saldkjhasdlkfjblwkjahwpdojabldpmfblaowidh"
	auth := common.NewAuth(appKey)
	db := persistence.NewDB(*dbhost)
	db.LogMode(true)
	defer db.Close()

	router := mux.NewRouter()
	reg := registry.NewRegistry(db, &auth)
	// apiRouter := router.PathPrefix("/api").Subrouter()

	userController := reg.NewUserHandler()
	router.HandleFunc("/api/signup", userController.Signup).Methods("POST")
	router.HandleFunc("/api/login", userController.Login).Methods("POST")

	/*corsRouter := handlers.CORS(
	handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
	handlers.AllowedOrigins([]string{"*"}),
	)(router)*/

	log.Fatal(http.ListenAndServe(":8081", router))
}
