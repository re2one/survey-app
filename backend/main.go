package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"backend/adapter/controller"
	"backend/adapter/presenter"
	"backend/adapter/repository"
	"backend/common"
	"backend/frameworks/persistence"
	"backend/usecase/interactor"
)

func main() {

	dbhost := flag.String("dbhost", "localhost:3306", "some bad help")
	flag.Parse()

	db := persistence.NewDB(*dbhost)
	db.LogMode(true)
	defer db.Close()

	privateKey, err := ioutil.ReadFile("common/keys/jwtRS256.key")
	if err != nil {
		return
	}

	/*publicKey, err := ioutil.ReadFile("common/keys/jwtRS256.key.pub")
	if err != nil {
		return
	}*/
	/*
		fmt.Println(string(publicKey))
		fmt.Println(string(privateKey))
	*/
	authenticator := common.NewAuthenticator(string(privateKey))
	authorizer := common.NewAuthorizer(string(privateKey))

	ur := repository.NewUserRepository(db)
	rr := repository.NewRoleRepository(db)
	up := presenter.NewUserPresenter(&authenticator)
	ui := interactor.NewUserInteractor(ur, rr, up, authorizer, authenticator)
	uc := controller.NewUserController(ui)

	sc := controller.NewSurveyController()

	router := mux.NewRouter()
	router.HandleFunc("/api/signup", uc.Signup).Methods(http.MethodPost)
	router.HandleFunc("/api/login", uc.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/surveys", authorizer.IsAuthorized("user", sc.Test)).Methods(http.MethodGet)
	router.HandleFunc("/api/refresh", authorizer.IsAuthorized("user", uc.RefreshToken)).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8081", router))
}
