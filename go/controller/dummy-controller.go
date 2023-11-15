package controller

import (
	"encoding/json"
	"fmt"
	"irwanka/sicerdas/entity"
	"irwanka/sicerdas/helper"
	"irwanka/sicerdas/repository"
	"irwanka/sicerdas/service"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

type DummyController interface {
	CekExistDummyForTemplate(w http.ResponseWriter, r *http.Request)
	CreateDummyForTemplate(w http.ResponseWriter, r *http.Request)
}

func NewDummyController() DummyController {
	return &controller{}
}

func (*controller) CekExistDummyForTemplate(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	id_template, _ := strconv.Atoi(id)
	dummyRepo := repository.NewDummyQuizUserRepository()
	userQuiz, err := dummyRepo.CekDummyQuizUser(id_template)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "Belum ada data dummy untuk template"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.ResponseData{Status: true, Data: userQuiz})
}

func (*controller) CreateDummyForTemplate(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	id_template, _ := strconv.Atoi(id)

	dummyRepo := repository.NewDummyQuizUserRepository()
	skoringRepo := repository.NewSkoringRepository()
	skoringService := service.NewSkoringService(skoringRepo)
	oldquiz, errorCek := dummyRepo.CekDummyQuizUser(id_template)
	if errorCek != nil {
		fmt.Println(errorCek.Error())
	} else {
		if oldquiz.IDQuiz > 0 {
			errDelete := dummyRepo.DeleteOldDummy(int(oldquiz.IDQuiz))
			if errDelete != nil {
				return
			}
		}
	}

	quiz := dummyRepo.CreateDummySesiTemplate(id_template)
	if quiz.IDQuiz == 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Status: false, Message: "Gagal buat sesi dummy"})
		fmt.Println("Error gagal buat sesi dummy untuk template ini..")
		return
	}
	// quiz, _ := dummyRepo.CekDummyQuizUser(id_template)
	id_quiz := quiz.IDQuiz
	//fmt.Println(id_quiz)
	randomJawaban, errJawaban := dummyRepo.GenerateJawaban(int(id_quiz))
	if errJawaban != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Status: false, Message: "Jawaban Random dummy gagal dibuat"})
		fmt.Println("Error gagal buat jawaban sesi dummy untuk template ini..")
		return
	}
	var jawabanString = []entity.JawabanString{}
	for i := 0; i < len(randomJawaban); i++ {
		//fmt.Println(randomJawaban[i].Jawaban)
		//fmt.Println(randomJawaban[i].Kategori)
		tmp := entity.JawabanString{}
		tmp.Kategori = randomJawaban[i].Kategori
		jws, _ := json.Marshal(randomJawaban[i].Jawaban)
		tmp.Jawaban = string(jws)
		jawabanString = append(jawabanString, tmp)
	}
	j, _ := json.Marshal(jawabanString)
	jsonString := string(j)
	// log.Println(string(jsonString))

	dummyRepo.SubmitJawabanUser(int(id_quiz), jsonString)

	kategori_tabel, _ := skoringRepo.GetTabelSkoring(id_quiz)
	skoringRepo.ClearTabelTemporaryJawabanUser(id_quiz, 100)
	skoringRepo.GenerateTabelTemporaryJawabanUser(id_quiz, 100)
	skoringService.SkoringAllKategori(kategori_tabel, id_quiz, 100)
	skoringRepo.ClearTabelTemporaryJawabanUser(id_quiz, 100)
	now := helper.StringTimeYMDHIS(time.Now())
	skoringRepo.StartRunningSkoring(now)
	skoringRepo.FinishSkoring(id_quiz, 100, now)
	skoringRepo.StopRunningSkoring(now, 1)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.ResponseMessage{Status: true, Message: "Sesi dummy berhasil dibuat"})
}
