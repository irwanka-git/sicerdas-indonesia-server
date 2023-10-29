package controller

import (
	"encoding/json"
	"fmt"
	"irwanka/sicerdas/helper"
	"irwanka/sicerdas/repository"
	"irwanka/sicerdas/service"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

var (
	uploadGambarRepository repository.UploadFirebaseRepository = repository.NewUploadFirebaseRepository()
	uploadGambarService    service.UploadFirebaseService       = service.NewUploadFirebaseRepository(uploadGambarRepository)
)

type UploadController interface {
	UploadGambarToFirebase(w http.ResponseWriter, r *http.Request)
	SinkronGambarQuizTemplateToFirebase(w http.ResponseWriter, r *http.Request)
	SinkronGambarInfoCerdasToFirebase(w http.ResponseWriter, r *http.Request)
}

func NewUploadController() UploadController {
	return &controller{}
}

func (*controller) SinkronGambarQuizTemplateToFirebase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	_, claims, _ := jwtauth.FromContext(r.Context())

	sub := fmt.Sprintf("%v", claims["sub"])
	user, err := userService.GetlUserByUuid(sub)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}
	if user.Username != "admin" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "akses tidak valid"})
		return
	}
	quiz_template, _ := quizService.GetAllQuizTemplate()
	for i := 0; i < len(quiz_template); i++ {
		if !strings.Contains(quiz_template[i].Gambar, "https://storage.googleapis.com") {
			path := fmt.Sprintf("/var/www/html/public/gambar/%v", quiz_template[i].Gambar)
			directory := "gambar"

			if helper.FileExists((path)) {
				url, errUplaod := uploadGambarService.UploadGambarToFirebase(path, directory)
				if errUplaod == nil {
					quizService.UpdateGambarQuizTemplate(quiz_template[i].IDQuizTemplate, url)
				}
			} else {
				fmt.Println(path)
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.ResponseMessage{Status: true, Message: "Berhasil Sinkrob Gambar Quiz ke Firebase"})
}

func (*controller) UploadGambarToFirebase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	_, claims, _ := jwtauth.FromContext(r.Context())

	sub := fmt.Sprintf("%v", claims["sub"])
	user, err := userService.GetlUserByUuid(sub)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}
	if user.Username != "admin" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "akses tidak valid"})
		return
	}

	filename := chi.URLParam(r, "filename")
	if filename == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "Request Tidak Valid"})
		return
	}
	path := fmt.Sprintf("/var/www/html/public/gambar/%v", filename)

	directory := "gambar"
	url, err := uploadGambarService.UploadGambarToFirebase(path, directory)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.ResponseData{Status: true, Message: "Berhasil Upload Gambar ke Firebase", Data: url})
}

func (*controller) SinkronGambarInfoCerdasToFirebase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	_, claims, _ := jwtauth.FromContext(r.Context())

	sub := fmt.Sprintf("%v", claims["sub"])
	user, err := userService.GetlUserByUuid(sub)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}
	if user.Username != "admin" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "akses tidak valid"})
		return
	}

	info_cerdas, _ := userService.GetAllInfoCerdas()
	for i := 0; i < len(info_cerdas); i++ {
		if !strings.Contains(info_cerdas[i].Gambar, "https://storage.googleapis.com") {
			path := fmt.Sprintf("/var/www/html/public/gambar/%v", info_cerdas[i].Gambar)
			directory := "gambar"
			if helper.FileExists((path)) {
				url, errUplaod := uploadGambarService.UploadGambarToFirebase(path, directory)
				if errUplaod == nil {
					userService.UpdateGambarInfo(info_cerdas[i].IDInfo, url)
				}
			} else {
				fmt.Println(path)
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.ResponseMessage{Status: true, Message: "Berhasil Sinkron Gambar Info Cerdas ke Firebase"})
}
