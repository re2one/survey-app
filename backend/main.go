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

	qr := repository.NewQuestionRepository(db)
	qc := controller.NewQuestionController(qr, sr)

	cr := repository.NewChoiceRepository(db)
	cc := controller.NewChoiceController(cr, qr)

	ar := repository.NewAnsweredRepository(db)
	fc := controller.NewFullQuestionsController(qr, ar, sr, ur)

	mr := repository.NewMultipleChoiceRepository(db)
	mc := controller.NewMultipleChoiceController(mr, qr)

	router := mux.NewRouter()
	router.HandleFunc("/api/signup", uc.Signup).Methods(http.MethodPost)
	router.HandleFunc("/api/login", uc.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/pubkey", pc.Get).Methods(http.MethodGet)
	router.HandleFunc("/api/refresh", authorizer.IsAuthorized("user", uc.RefreshToken)).Methods(http.MethodGet)

	router.HandleFunc("/api/surveys", authorizer.IsAuthorized("user", sc.GetAll)).Methods(http.MethodGet)
	router.HandleFunc("/api/surveys/{id}", authorizer.IsAuthorized("user", sc.Get)).Methods(http.MethodGet)
	router.HandleFunc("/api/surveys", authorizer.IsAuthorized("admin", sc.Post)).Methods(http.MethodPost)
	router.HandleFunc("/api/surveys", authorizer.IsAuthorized("admin", sc.Put)).Methods(http.MethodPut)
	router.HandleFunc("/api/surveys/{id}", authorizer.IsAuthorized("admin", sc.Delete)).Methods(http.MethodDelete)

	router.HandleFunc("/api/questions/{surveyId}", authorizer.IsAuthorized("user", qc.GetAll)).Methods(http.MethodGet)
	router.HandleFunc("/api/questions/single/{id}", authorizer.IsAuthorized("user", qc.Get)).Methods(http.MethodGet)
	router.HandleFunc("/api/questions/{surveyId}", authorizer.IsAuthorized("admin", qc.Post)).Methods(http.MethodPost)
	router.HandleFunc("/api/questions", authorizer.IsAuthorized("admin", qc.Put)).Methods(http.MethodPut)
	router.HandleFunc("/api/questions/{id}", authorizer.IsAuthorized("admin", qc.Delete)).Methods(http.MethodDelete)

	router.HandleFunc("/api/choices/{questionId}", authorizer.IsAuthorized("user", cc.GetAll)).Methods(http.MethodGet)
	router.HandleFunc("/api/choices/single/{id}", authorizer.IsAuthorized("user", cc.Get)).Methods(http.MethodGet)
	router.HandleFunc("/api/choices/{questionId}", authorizer.IsAuthorized("admin", cc.Post)).Methods(http.MethodPost)
	router.HandleFunc("/api/choices", authorizer.IsAuthorized("admin", cc.Put)).Methods(http.MethodPut)
	router.HandleFunc("/api/choices/{id}", authorizer.IsAuthorized("admin", cc.Delete)).Methods(http.MethodDelete)

	router.HandleFunc("/api/fullquestions/{surveyid}/{email}", authorizer.IsAuthorized("user", fc.GetAll)).Methods(http.MethodGet)
	router.HandleFunc("/api/fullquestions/{email}", authorizer.IsAuthorized("user", fc.Post)).Methods(http.MethodPost)

	router.HandleFunc("/api/answer/multiplechoice", authorizer.IsAuthorized("user", mc.Post)).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8081", router))
}
