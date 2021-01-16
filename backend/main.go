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

	rnd := common.NewRandomString()
	uc := controller.NewUserController(ui, rnd)

	assr := repository.NewAssetsRepository(db)
	assc := controller.NewAssetsController(assr)

	sr := repository.NewSurveyRepository(db)
	sc := controller.NewSurveyController(sr, assr)

	pc := controller.NewPubkeyController()

	par := repository.NewPuzzleAnswerRepository(db)
	pac := controller.NewPuzzleAnswerController(par)

	ar := repository.NewAnsweredRepository(db)
	qr := repository.NewQuestionRepository(db)
	qc := controller.NewQuestionController(qr, sr, ar, ur, par)

	cr := repository.NewChoiceRepository(db)
	cc := controller.NewChoiceController(cr, qr)

	mr := repository.NewMultipleChoiceRepository(db)
	mc := controller.NewMultipleChoiceController(mr, qr)

	fc := controller.NewFullQuestionsController(qr, ar, sr, ur, mr, cr)

	puzzr := repository.NewPuzzleRepository(db)
	puzzc := controller.NewPuzzleController(puzzr, ar, ur, qr)

	rc := controller.NewResultsController(qr, mr, puzzr, par, ur, ar)

	br := repository.NewBracketRepository(db)
	bc := controller.NewBracketController(br, sr)

	router := mux.NewRouter()

	router.HandleFunc("/api/signup", uc.Signup).Methods(http.MethodPost)
	router.HandleFunc("/api/login", uc.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/pubkey", pc.Get).Methods(http.MethodGet)
	router.HandleFunc("/api/refresh", authorizer.IsUser(uc.RefreshToken)).Methods(http.MethodGet)

	router.HandleFunc("/api/surveys", authorizer.IsUser(sc.GetAll)).Methods(http.MethodGet)
	router.HandleFunc("/api/surveys/{id}", authorizer.IsUser(sc.Get)).Methods(http.MethodGet)
	router.HandleFunc("/api/surveys", authorizer.IsAdmin(sc.Post)).Methods(http.MethodPost)
	router.HandleFunc("/api/surveys", authorizer.IsAdmin(sc.Put)).Methods(http.MethodPut)
	router.HandleFunc("/api/surveys/{id}", authorizer.IsAdmin(sc.Delete)).Methods(http.MethodDelete)

	router.HandleFunc("/api/questions/{surveyId}", authorizer.IsUser(qc.GetAll)).Methods(http.MethodGet)
	router.HandleFunc("/api/questions/single/{id}", authorizer.IsUser(qc.Get)).Methods(http.MethodGet)
	router.HandleFunc("/api/questions/{surveyId}", authorizer.IsAdmin(qc.Post)).Methods(http.MethodPost)
	router.HandleFunc("/api/questions/answered/{email}/{surveyId}", authorizer.IsAdmin(qc.GetAnswered)).Methods(http.MethodGet)
	router.HandleFunc("/api/questions", authorizer.IsAdmin(qc.Put)).Methods(http.MethodPut)
	router.HandleFunc("/api/questions/{id}", authorizer.IsAdmin(qc.Delete)).Methods(http.MethodDelete)

	router.HandleFunc("/api/choices/{questionId}", authorizer.IsUser(cc.GetAll)).Methods(http.MethodGet)
	router.HandleFunc("/api/choices/single/{id}", authorizer.IsUser(cc.Get)).Methods(http.MethodGet)
	router.HandleFunc("/api/choices/{questionId}", authorizer.IsAdmin(cc.Post)).Methods(http.MethodPost)
	router.HandleFunc("/api/choices", authorizer.IsAdmin(cc.Put)).Methods(http.MethodPut)
	router.HandleFunc("/api/choices/{id}", authorizer.IsAdmin(cc.Delete)).Methods(http.MethodDelete)

	router.HandleFunc("/api/fullquestions/{surveyid}/{email}", authorizer.IsUser(fc.GetAll)).Methods(http.MethodGet)
	router.HandleFunc("/api/fullquestions/{email}", authorizer.IsUser(fc.Post)).Methods(http.MethodPost)
	router.HandleFunc("/api/fullquestions/viewed/{email}", authorizer.IsUser(fc.Viewed)).Methods(http.MethodPost)

	router.HandleFunc("/api/answer/multiplechoice", authorizer.IsUser(mc.Post)).Methods(http.MethodPost)

	router.HandleFunc("/api/results/{surveyId}", authorizer.IsAdmin(rc.Get)).Methods(http.MethodGet)
	router.HandleFunc("/api/results/{surveyId}/{email}/{questionId}", authorizer.IsAdmin(rc.GetSingle)).Methods(http.MethodGet)
	router.HandleFunc("/api/results/try/average/{email}/{surveyId}", authorizer.IsUser(authorizer.IsOwner(rc.GetUserAverage))).Methods(http.MethodGet)

	router.HandleFunc("/api/assets/directory/{surveyId}/{questionId}", authorizer.IsAdmin(assc.Post)).Methods(http.MethodPost)
	router.HandleFunc("/api/assets/upload/{surveyId}/{questionId}", authorizer.IsAdmin(assc.Upload)).Methods(http.MethodPost)
	router.HandleFunc("/api/assets/static/{type}/{surveyId}", assc.GetPDF).Methods(http.MethodGet)
	router.HandleFunc("/api/assets/static/{type}/{surveyId}", authorizer.IsAdmin(assc.UploadPDF)).Methods(http.MethodPost)
	router.HandleFunc("/api/assets/{surveyId}/{questionId}", authorizer.IsUser(assc.GetAll)).Methods(http.MethodGet)
	router.HandleFunc("/api/assets/{surveyId}/{questionId}/{imageName}", assc.Get).Methods(http.MethodGet)

	router.HandleFunc("/api/puzzle/{surveyId}/{questionId}", authorizer.IsAdmin(puzzc.Put)).Methods(http.MethodPut)
	router.HandleFunc("/api/puzzle/{questionId}", authorizer.IsUser(puzzc.GetAll)).Methods(http.MethodGet)
	router.HandleFunc("/api/puzzle/{questionId}/{email}", authorizer.IsUser(puzzc.GetAllForQuestionaire)).Methods(http.MethodGet)

	router.HandleFunc("/api/brackets/{surveyId}", authorizer.IsAdmin(bc.Get)).Methods(http.MethodGet)
	router.HandleFunc("/api/brackets/{surveyId}", authorizer.IsAdmin(bc.Post)).Methods(http.MethodPost)

	router.HandleFunc("/api/answer/puzzle", authorizer.IsUser(pac.Post)).Methods(http.MethodPost)

	router.HandleFunc("/api/users", authorizer.IsAdmin(uc.Get)).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8081", router))
}
