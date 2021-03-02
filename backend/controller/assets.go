package controller

import (
	"encoding/json"
	"fmt"
	"image/jpeg"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"backend/usecase/repository"
)

type assetsController struct {
	assetsRepository repository.AssetsRepository
}

type SingleAssetResponse struct {
	SurveyId   string `json:"surveyId"`
	QuestionId string `json:"questionId"`
}

type FilenameResponse struct {
	Filenames []string `json:"filenames"`
}

type AssetsController interface {
	Post(writer http.ResponseWriter, request *http.Request)
	Upload(writer http.ResponseWriter, request *http.Request)
	UploadPDF(writer http.ResponseWriter, request *http.Request)
	GetPDF(writer http.ResponseWriter, request *http.Request)
	GetAll(writer http.ResponseWriter, request *http.Request)
	Get(writer http.ResponseWriter, request *http.Request)
}

func NewAssetsController(a repository.AssetsRepository) AssetsController {
	return &assetsController{assetsRepository: a}
}

func (a *assetsController) Post(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)
	err := a.assetsRepository.Post(v["surveyId"], v["questionId"])
	if err != nil {
		log.Error().Err(err).Msg("unable to create asset folder")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	r := SingleAssetResponse{SurveyId: v["surveyId"], QuestionId: v["questionId"]}
	json.NewEncoder(writer).Encode(r)
	return
}

func (a *assetsController) Upload(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)

	request.ParseMultipartForm(32 << 20) // limit your max input length!

	// in your case file would be fileupload
	file, header, err := request.FormFile("fileKey")
	if err != nil {
		log.Error().Err(err).Caller().Msg("unable to unwrap file")
		writer.WriteHeader(http.StatusInternalServerError)
	}
	defer file.Close()

	name := strings.Split(header.Filename, ".")
	log.Info().Str("Filename", name[0]).Msg("Name of transferred file")

	image, err := jpeg.Decode(file)

	// save file to disk
	err = a.assetsRepository.SaveFile(v["surveyId"], v["questionId"], image, header.Filename)
	if err != nil {
		log.Error().Err(err).Caller().Msg("unable to write file to disc")
		writer.WriteHeader(http.StatusInternalServerError)

	}

	return
}

func (a *assetsController) UploadPDF(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)

	request.ParseMultipartForm(32 << 20) // limit your max input length!

	// in your case file would be fileupload
	file, header, err := request.FormFile("fileKey")
	if err != nil {
		log.Error().Err(err).Caller().Msg("unable to unwrap file")
		writer.WriteHeader(http.StatusInternalServerError)
	}
	defer file.Close()

	name := strings.Split(header.Filename, ".")
	log.Info().Str("Filename", name[0]).Msg("Name of transferred file")

	// image, err := jpeg.Decode(file)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error().Err(err).Caller().Msg("unable to decode file into bytes")
		writer.WriteHeader(http.StatusInternalServerError)
	}

	// save file to disk
	err = a.assetsRepository.SavePDF(v["type"], v["surveyId"], fileBytes)
	if err != nil {
		log.Error().Err(err).Caller().Msg("unable to write file to disc")
		writer.WriteHeader(http.StatusInternalServerError)

	}

	return
}

func (a *assetsController) GetAll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)
	filenames, err := a.assetsRepository.GetFilenames(v["surveyId"], v["questionId"])
	if err != nil {
		log.Error().Err(err).Caller().Msg("unable to retrieve filenames from disc")
		writer.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(writer).Encode(FilenameResponse{Filenames: filenames})
	return
}

func (a *assetsController) Get(writer http.ResponseWriter, request *http.Request) {
	v := mux.Vars(request)

	// img, err := a.assetsRepository.Get(v["surveyId"], v["questionId"], v["imageName"])
	path := fmt.Sprintf("assets/survey_%v/question_%v/%v", v["surveyId"], v["questionId"], v["imageName"])
	img, err := os.Open(path)
	defer img.Close()
	if err != nil {
		log.Error().Err(err).Caller().Msg("unable to retrieve image from disc")
		writer.WriteHeader(http.StatusInternalServerError)
	}
	writer.Header().Set("Content-Type", "image/jpeg") // <-- set the content-type header
	io.Copy(writer, img)
	return
}

func (a *assetsController) GetPDF(writer http.ResponseWriter, request *http.Request) {
	v := mux.Vars(request)

	var path string
	switch v["type"] {
	case "termsandconditions":
		path = "assets/termsandconditions.pdf"
	case "impressum":
		path = "assets/impressum.pdf"
	case "datenschutz":
		path = "assets/datenschutz.pdf"
	default:
		path = fmt.Sprintf("assets/survey_%v/%v/%v.pdf", v["surveyId"], v["type"], v["type"])
	}
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Error().Err(err).Caller().Msg("unable to retrieve image from disc")
		writer.WriteHeader(http.StatusInternalServerError)
	}
	writer.Header().Set("Content-Type", "application/pdf") // <-- set the content-type header
	io.Copy(writer, file)
	return
}
