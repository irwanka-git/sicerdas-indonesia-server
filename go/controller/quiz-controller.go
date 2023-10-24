package controller

import (
	"encoding/json"
	"fmt"
	"irwanka/sicerdas/entity"
	"irwanka/sicerdas/helper"
	"irwanka/sicerdas/repository"
	"irwanka/sicerdas/service"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

var (
	quizRepository           repository.QuizRepository           = repository.NewQuizRepository()
	uploadFirebaseRepository repository.UploadFirebaseRepository = repository.NewUploadFirebaseRepository()
	quizService              service.QuisService                 = service.NewQuizService(quizRepository)
	uploadFirebaseService    service.UploadFirebaseService       = service.NewUploadFirebaseRepository(uploadFirebaseRepository)
)

type QuizController interface {
	GetListQuizSessionInfo(w http.ResponseWriter, r *http.Request)
	UploadQuizJsonToFirebase(w http.ResponseWriter, r *http.Request)
	GetQuizDetil(w http.ResponseWriter, r *http.Request)
}

func NewQuizController() QuizController {
	return &controller{}
}

func (*controller) GetListQuizSessionInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	token := chi.URLParam(r, "token")
	if token == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "Request Tidak Valid"})
		return
	}
	data, err := quizService.GetListInfoSessionQuiz(token)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.ResponseData{Status: true, Data: data})
}

func (*controller) UploadQuizJsonToFirebase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	// currentJWTToken := jwtauth.TokenFromHeader(r)
	// print(currentJWTToken)
	token := chi.URLParam(r, "token")
	if token == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "Request Tidak Valid"})
		return
	}
	quiz, err := quizService.GetlDetilQuizByToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}

	listInfoSesi, err := quizService.GetListInfoSessionQuiz(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}
	listSoalSesi, err := quizService.GetAllSoalSessionQuiz(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}

	var result = entity.QuizFirebaseStorage{}
	result.Session = listInfoSesi
	result.Soal = listSoalSesi
	jsonString, _ := json.Marshal(result)

	path := fmt.Sprintf("%v/%v.json", os.Getenv("PATH_JSON_SOAL"), token)
	os.WriteFile(path, jsonString, 0644)
	directory := fmt.Sprintf("soal/%v/%v", quiz.SkoringTabel, quiz.Tanggal.Format("2006-01-02"))
	url, err := uploadFirebaseService.UploadFileToFirebase(path, directory)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}
	err = quizService.UpdateURLFirebaseSoalQuiz(token, url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "Gagal update firebase soal"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.ResponseData{Status: true, Message: "Berhasil Upload Sesi ke Firebase", Data: url})
}

func (*controller) GetQuizDetil(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var token string
	if r.URL.Query().Get("token") != "" {
		token = helper.CleanString(r.URL.Query().Get("token"))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "token tidak valid"})
		return
	}
	data, err := quizService.GetlDetilQuizByToken(token)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.ResponseData{Status: true, Data: data})
}
