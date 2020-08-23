package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"backend/adapter/presenter"
	"backend/adapter/repository"
	"backend/common"
	"backend/controller"
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

	authenticator := common.NewAuthenticator(string(privateKey))
	authorizer := common.NewAuthorizer(string(privateKey))

	ur := repository.NewUserRepository(db)
	rr := repository.NewRoleRepository(db)
	up := presenter.NewUserPresenter(&authenticator)
	ui := interactor.NewUserInteractor(ur, rr, up, authorizer, authenticator)
	uc := controller.NewUserController(ui)

	sr := repository.NewSurveyRepository(db)
	sc := controller.NewSurveyController(sr)

	pc := controller.NewPubkeyController()

	router := mux.NewRouter()
	router.HandleFunc("/api/signup", uc.Signup).Methods(http.MethodPost)
	router.HandleFunc("/api/login", uc.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/pubkey", pc.Get).Methods(http.MethodGet)
	router.HandleFunc("/api/surveys", authorizer.IsAuthorized("user", sc.GetAll)).Methods(http.MethodGet)
	router.HandleFunc("/api/surveys/:title", authorizer.IsAuthorized("user", sc.Test)).Methods(http.MethodGet)
	router.HandleFunc("/api/surveys", authorizer.IsAuthorized("admin", sc.Test)).Methods(http.MethodPost)
	router.HandleFunc("/api/surveys", authorizer.IsAuthorized("admin", sc.Test)).Methods(http.MethodPut)
	router.HandleFunc("/api/surveys", authorizer.IsAuthorized("admin", sc.Test)).Methods(http.MethodDelete)
	router.HandleFunc("/api/refresh", authorizer.IsAuthorized("user", uc.RefreshToken)).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8081", router))
}
