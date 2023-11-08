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
	"slices"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

var (
	quizRepository           repository.QuizRepository           = repository.NewQuizRepository()
	uploadFirebaseRepository repository.UploadFirebaseRepository = repository.NewUploadFirebaseRepository()
	quizService              service.QuisService                 = service.NewQuizService(quizRepository)
	uploadFirebaseService    service.UploadFirebaseService       = service.NewUploadFirebaseRepository(uploadFirebaseRepository)
)

type QuizController interface {
	//peserta
	GetListQuizSessionInfo(w http.ResponseWriter, r *http.Request)
	GetListQuizUser(w http.ResponseWriter, r *http.Request)
	GetSalamPembuka(w http.ResponseWriter, r *http.Request)
	GetDetilQuizUser(w http.ResponseWriter, r *http.Request)
	SubmitJawabanQuiz(w http.ResponseWriter, r *http.Request)

	//admin
	UploadQuizJsonToFirebase(w http.ResponseWriter, r *http.Request)
	GetQuizDetil(w http.ResponseWriter, r *http.Request)
	SinkronGambarQuizWithTemplate(w http.ResponseWriter, r *http.Request)
}

func NewQuizController() QuizController {
	return &controller{}
}

func (*controller) SubmitJawabanQuiz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	_, claims, _ := jwtauth.FromContext(r.Context())

	sub := fmt.Sprintf("%v", claims["sub"])
	user, err := userService.GetlUserByUuid(sub)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}
	token := helper.CleanParameterIDOnly(chi.URLParam(r, "token"))
	if token == "" {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "request tidak valid"})
		return
	}

	var parameter struct {
		Jawaban string `json:"jawaban"`
	}

	err = json.NewDecoder(r.Body).Decode(&parameter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "jawaban tidak valid", Status: false})
		return
	}
	quiz, errQuiz := quizService.GetlDetilQuizByToken(token)
	if errQuiz != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errQuiz.Error(), Status: false})
		return
	}
	errSubmit := quizService.SubmitJawabanQuiz(parameter.Jawaban, quiz.IDQuiz, *user)
	if errSubmit != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errSubmit.Error(), Status: false})
		return
	}

	w.WriteHeader(http.StatusOK)
	var nama_panggilan = ""
	nama := strings.Split(user.NamaPengguna, " ")
	if len(nama) > 1 {
		nama_panggilan = nama[0]
		if len(nama_panggilan) <= 3 {
			nama_panggilan = nama[1]
		}
	} else {
		nama_panggilan = user.NamaPengguna
	}
	pesan := fmt.Sprintf("<p>Terima kasih <b>%s</b>. <br><b>Jawaban kamu berhasil</b> dikirim. Kami akan segera periksa dan hasilnya nanti dapat kamu <b>download</b> melalui aplikasi.</p>", nama_panggilan)
	json.NewEncoder(w).Encode(helper.ResponseMessage{Status: true, Message: pesan})
}

func (*controller) GetSalamPembuka(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	_, claims, _ := jwtauth.FromContext(r.Context())

	sub := fmt.Sprintf("%v", claims["sub"])
	_, err := userService.GetlUserByUuid(sub)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}
	token := helper.CleanParameterIDOnly(chi.URLParam(r, "token"))
	if token == "" {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "request tidak valid"})
		return
	}

	salam, err := quizService.GetSalamPembuka(token)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error(), Status: false})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.ResponseMessage{Message: salam, Status: true})

}

func (*controller) GetDetilQuizUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	_, claims, _ := jwtauth.FromContext(r.Context())

	sub := fmt.Sprintf("%v", claims["sub"])
	user, err := userService.GetlUserByUuid(sub)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}
	token := chi.URLParam(r, "token")
	if token == "" {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "request tidak valid"})
		return
	}

	infoQuiz, err := quizService.GetStatusQuizUser(int(user.ID), token)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}

	listSession, err := quizService.GetListInfoSessionQuiz(token)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.ResponseDataInfo{Status: true, Data: listSession, Information: infoQuiz})
}

func (*controller) GetListQuizUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	_, claims, _ := jwtauth.FromContext(r.Context())

	sub := fmt.Sprintf("%v", claims["sub"])
	user, err := userService.GetlUserByUuid(sub)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}

	listQuiz, _ := quizService.GetlListQuizByUser(int(user.ID))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.ResponseData{Status: true, Data: listQuiz})
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
	directory := fmt.Sprintf("soal/%v/%v", quiz.IDQuizTemplate, quiz.Tanggal.Format("2006-01-02"))
	url, err := uploadFirebaseService.UploadSoalQuizToFirebase(path, directory)
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

	token := chi.URLParam(r, "token")
	if token == "" {
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

func (*controller) SinkronGambarQuizWithTemplate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	quiz_template, _ := quizService.GetAllQuizTemplate()
	quiz, _ := quizService.GetAllQuizSesi()
	for i := 0; i < len(quiz); i++ {
		id_quiz_template := quiz[i].IDQuizTemplate
		id_find := slices.IndexFunc(quiz_template, func(c *entity.QuizSesiTemplate) bool { return c.IDQuizTemplate == id_quiz_template })
		if id_find >= 0 {
			gambar := quiz_template[id_find].Gambar
			quizService.UpdateGambarQuizSesi(quiz[i].IDQuiz, gambar)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.ResponseMessage{Status: true, Message: "Berhasin sinkron gambar quiz dengan template"})
}
