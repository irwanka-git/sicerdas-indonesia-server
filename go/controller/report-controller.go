package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"irwanka/sicerdas/entity"
	"irwanka/sicerdas/helper"
	"irwanka/sicerdas/repository"
	"irwanka/sicerdas/service"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	strip "github.com/grokify/html-strip-tags-go"
)

const (
	ID_USER_DUMMY = 100
)

var (
	reportRepository repository.ReportRepository = repository.NewReportRepository()
	reportService    service.ReportService       = service.NewReporService(reportRepository)
)

type ReportController interface {
	PreviewKomponenReportDummy(w http.ResponseWriter, r *http.Request)
	PreviewLampiranReportDummy(w http.ResponseWriter, r *http.Request)
	RenderCoverReport(w http.ResponseWriter, r *http.Request)
	RenderReportUtama(w http.ResponseWriter, r *http.Request)
	RenderReportLampiran(w http.ResponseWriter, r *http.Request)
	ExportReportPDFToFirebase(w http.ResponseWriter, r *http.Request)
}

func NewReportController() ReportController {
	return &controller{}
}

func (*controller) RenderCoverReport(w http.ResponseWriter, r *http.Request) {
	id_quiz, _ := strconv.Atoi(chi.URLParam(r, "id_quiz"))
	id_user, _ := strconv.Atoi(chi.URLParam(r, "id_user"))
	quiz, _ := reportRepository.GetDetilQuizCetak(id_quiz)
	user, _ := userRepository.GetDataUserById(id_user)
	biro, _ := userRepository.GetDataUserById(int(quiz.IDUserBiro))
	var cover = "default.png"
	if biro.CoverBiroGambar != "" {
		cover = biro.CoverBiroGambar
	}
	var listTemplate = []string{}
	listTemplateLampiran := helper.GetArrFilename("templates/cover/")
	if len(listTemplateLampiran) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "Template tidak ditemukan"})
		return
	}
	listTemplate = append(listTemplate, listTemplateLampiran...)
	t := template.Must(template.New("report.html").ParseFiles(listTemplate...))

	data := map[string]interface{}{
		"user":    user,
		"quiz":    quiz,
		"tanggal": helper.StringTimeTglIndo(quiz.Tanggal),
		"cover":   cover,
	}
	err := t.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}
}

func (*controller) ExportReportPDFToFirebase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	_, claims, _ := jwtauth.FromContext(r.Context())

	sub := fmt.Sprintf("%v", claims["sub"])
	userToken, err := userService.GetlUserByUuid(sub)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}
	if userToken.Username != "admin" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "akses tidak valid"})
		return
	}

	id_quiz, _ := strconv.Atoi(chi.URLParam(r, "id_quiz"))
	id_user, _ := strconv.Atoi(chi.URLParam(r, "id_user"))
	id_model := chi.URLParam(r, "id_model")
	model, _ := reportRepository.GetModelReport(id_model)

	url_docker := os.Getenv("URL_HOST_DOCKER")
	quiz, _ := reportRepository.GetDetilQuizCetak(id_quiz)
	nomor_seri := reportRepository.GenerateQRCodeNomorSeriCetak()
	user, _ := userRepository.GetDataUserById(id_user)
	quizTemplate, _ := reportRepository.GetDetilQuizTemplate(int(quiz.IDQuizTemplate))

	//render cover
	path_export_cover := fmt.Sprintf("/app/export/pdf/%v.%v.cover.pdf", user.ID, quiz.Token)
	url_render_cover := fmt.Sprintf("%v/render-cover/%v/%v", url_docker, id_quiz, id_user)
	exCover := exec.Command("/app/export.sh", url_render_cover, path_export_cover, "Portrait", quizTemplate.Kertas, nomor_seri, "cover")

	_, err2 := exCover.Output()
	// fmt.Println(string(out))
	if err2 != nil {
		json.NewEncoder(w).Encode(helper.ResponseMessage{Status: false, Message: err2.Error()})
		return
	}
	//render repot utama
	url_render_utama := fmt.Sprintf("%v/render-report-utama/%v/%v/%v/%v", url_docker, id_quiz, id_user, id_model, nomor_seri)
	fmt.Println(url_render_utama)
	path_export_utama := fmt.Sprintf("/app/export/pdf/%v.%v.%v.pdf", user.ID, quiz.Token, model.Direktori)
	ex := exec.Command("/app/export.sh", url_render_utama, path_export_utama, "Portrait", quizTemplate.Kertas, nomor_seri, "utama")
	// fmt.Println(ex.Args)

	_, err3 := ex.Output()
	// fmt.Println(string(out))
	if err3 != nil {
		//fmt.Println(err3.Error())
		json.NewEncoder(w).Encode(helper.ResponseMessage{Status: false, Message: err3.Error()})
		return
	}

	//render lampiran
	listLampiran := reportRepository.GetListLampiranReport(int(quizTemplate.IDQuizTemplate))
	var listExportLampiran = []string{}
	for i := 0; i < len(listLampiran); i++ {
		url_render_lampiran := fmt.Sprintf("%v/render-report-lampiran/%v/%v/%v/%v", url_docker, id_quiz, id_user, listLampiran[i].IDReport, nomor_seri)
		path_export_lampiran := fmt.Sprintf("/app/export/pdf/%v.%v.lampiran.%v.pdf", user.ID, quiz.Token, listLampiran[i].IDReport)
		// fmt.Println(url_render_lampiran)
		listExportLampiran = append(listExportLampiran, path_export_lampiran)
		exLampiran := exec.Command("/app/export.sh", url_render_lampiran, path_export_lampiran, listLampiran[i].Orientasi, quizTemplate.Kertas, nomor_seri, "lampiran")
		// fmt.Println(ex.Args)
		_, err4 := exLampiran.Output()
		if err4 != nil {
			//fmt.Println(err3.Error())
			json.NewEncoder(w).Encode(helper.ResponseMessage{Status: false, Message: err4.Error()})
			return
		}
	}

	//buat list pdf untuk pasing ke pdfunite
	var listPDF = []string{}
	listPDF = append(listPDF, path_export_cover)
	listPDF = append(listPDF, path_export_utama)
	for i := 0; i < len(listExportLampiran); i++ {
		listPDF = append(listPDF, listExportLampiran[i])
	}

	//merge cover + utama + lampiran => final;
	path_export_report_final := fmt.Sprintf("/app/export/pdf/%v-%v.pdf", helper.CleanNamaFileOnly(user.NamaPengguna), nomor_seri)
	listPDF = append(listPDF, path_export_report_final)
	exMergePDF := exec.Command("pdfunite", listPDF...)
	_, err5 := exMergePDF.Output()
	if err5 != nil {
		//fmt.Println(err3.Error())
		json.NewEncoder(w).Encode(helper.ResponseMessage{Status: false, Message: err5.Error()})
		return
	}

	//upload firebase
	directory := fmt.Sprintf("report-individu/%v/%v", quiz.Token, model.Direktori)
	url_firebase_result, err := uploadFirebaseService.UploadReportPDFToFirebase(path_export_report_final, directory)
	if err != nil {
		fmt.Println(err3)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Status: false, Message: err.Error()})
		return
	}

	for i := 0; i < len(listPDF); i++ {
		os.Remove(listPDF[i])
	}
	os.Remove(fmt.Sprintf("templates/assets/qrcode/%v.png", nomor_seri))
	//update no_seri
	reportRepository.UpdateNomorSeriCetak(id_quiz, id_user, nomor_seri, url_firebase_result)
	json.NewEncoder(w).Encode(helper.ResponseData{Status: true, Message: "Berhasil", Data: url_firebase_result})
}

func (*controller) RenderReportUtama(w http.ResponseWriter, r *http.Request) {
	id_quiz, _ := strconv.Atoi(chi.URLParam(r, "id_quiz"))
	id_user, _ := strconv.Atoi(chi.URLParam(r, "id_user"))
	id_model := chi.URLParam(r, "id_model")
	nomor_seri := chi.URLParam(r, "nomor_seri")
	renderReportKomponenUtama(id_quiz, id_user, id_model, nomor_seri, w)
}

func (*controller) RenderReportLampiran(w http.ResponseWriter, r *http.Request) {
	id_quiz, _ := strconv.Atoi(chi.URLParam(r, "id_quiz"))
	id_user, _ := strconv.Atoi(chi.URLParam(r, "id_user"))
	id_report, _ := strconv.Atoi(chi.URLParam(r, "id_report"))
	nomor_seri := chi.URLParam(r, "nomor_seri")
	renderReportLampiran(id_quiz, id_user, id_report, nomor_seri, w)
}

func renderReportLampiran(id_quiz int, id_user int, id_report int, nomor_seri string, w http.ResponseWriter) {
	quiz, _ := reportRepository.GetDetilQuizCetak(id_quiz)
	report := reportRepository.GetDetailReportLampiran(id_report, int(quiz.IDQuizTemplate))
	userBiro, _ := userRepository.GetDataUserById(int(quiz.IDUserBiro))
	// fmt.Println(report)
	if report == nil {
		return
	}

	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"unscapeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
		"stripTags": func(original string) string {
			stripped := strip.StripTags(original)
			result := strings.ReplaceAll(stripped, "&nbsp;", " ")
			return result
		},
		"replaceStrip": func(original string) string {
			stripped := strings.ReplaceAll(original, "_", " ")
			return stripped
		},
	}

	var listTemplate = []string{}

	listTemplateLampiran := helper.GetArrFilename("templates/lampiran/")
	if len(listTemplateLampiran) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "Template tidak ditemukan"})
		return
	}
	listTemplate = append(listTemplate, listTemplateLampiran...)
	t := template.Must(template.New("report.html").Funcs(funcMap).ParseFiles(listTemplate...))

	var skoring any
	var errSkoring error
	skoring, errSkoring = reportService.GetDataSkoringLampiranFromReportTabel(report.TabelReferensi, id_quiz, id_user)
	if errSkoring != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errSkoring.Error()})
		return
	}

	data := map[string]interface{}{
		"nama_komponen": report.NamaReport,
		"blade":         report.Blade,
		"skoring":       skoring,
		"urutan":        report.Urutan,
		"userBiro":      userBiro,
	}
	err := t.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}
}

func renderReportKomponenUtama(id_quiz int, id_user int, id_model string, nomor_seri string, w http.ResponseWriter) {

	model, _ := reportRepository.GetModelReport(id_model)
	quiz, _ := reportRepository.GetDetilQuizCetak(id_quiz)
	user, _ := userRepository.GetDataUserById(id_user)
	userQuiz := reportRepository.GetDetilQuizSesiUserCetak(id_quiz, id_user)
	userBiro, _ := userRepository.GetDataUserById(int(quiz.IDUserBiro))
	listKomponenUtama, _ := reportRepository.GetListKomponenUtama(int(quiz.IDQuizTemplate), id_model)

	var komponenReport = []entity.QuizReportRender{}
	var urutan = 0
	for i := 0; i < len(listKomponenUtama); i++ {
		var skoring any
		var errSkoring error
		skoring, errSkoring = reportService.GetDataSkoringFromReportTabel(listKomponenUtama[i].TabelReferensi, int(id_quiz), int(id_user))
		if errSkoring != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errSkoring.Error()})
			return
		}

		if listKomponenUtama[i].TabelReferensi != "-" {
			urutan = urutan + 1
		}

		if listKomponenUtama[i].TabelReferensi == "saran" {
			skoring, _ = reportRepository.GetDetilQuizTemplate(int(quiz.IDQuizTemplate))
		}
		var ttd_asesor = "-"
		if quiz.TtdAsesor != "" {
			ttd_asesor = quiz.TtdAsesor
		}
		if listKomponenUtama[i].TabelReferensi == "ttd" {
			ttdInfo := map[string]interface{}{
				"quiz":      quiz,
				"skoringAt": helper.StringTimeTglIndo(userQuiz.SkoringAt),
				"ttdAsesor": ttd_asesor,
			}
			skoring = ttdInfo
			fmt.Println(skoring)
		}

		data := map[string]interface{}{
			"skoring": skoring,
			"urutan":  urutan,
		}
		var temp = entity.QuizReportRender{}
		temp.Blade = listKomponenUtama[i].Blade
		temp.Skoring = data
		komponenReport = append(komponenReport, temp)
	}

	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"unscapeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
		"stripTags": func(original string) string {
			stripped := strip.StripTags(original)
			result := strings.ReplaceAll(stripped, "&nbsp;", " ")
			return result
		},
		"replaceStrip": func(original string) string {
			stripped := strings.ReplaceAll(original, "_", " ")
			return stripped
		},
	}

	var listTemplate = []string{}
	listTemplate = append(listTemplate, fmt.Sprintf("templates/%v/report.html", model.Direktori))
	listTemplateKomponen := helper.GetArrFilename(fmt.Sprintf("templates/%v/komponen/", model.Direktori))
	if len(listTemplateKomponen) == 0 {
		fmt.Println("File template report tidak ditemukan")
		return
	}
	listTemplate = append(listTemplate, listTemplateKomponen...)
	t := template.Must(template.New("report.html").Funcs(funcMap).ParseFiles(listTemplate...))

	data := map[string]interface{}{
		"komponen":  komponenReport,
		"quiz":      quiz,
		"tanggal":   helper.StringTimeTglIndo(quiz.Tanggal),
		"user":      user,
		"nomorSeri": nomor_seri,
		"userBiro":  userBiro,
	}
	err := t.Execute(w, data)
	if err != nil {
		return
	}
}

func (*controller) PreviewLampiranReportDummy(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	if uuid == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "Request Tidak Valid"})
		return
	}

	report, err := reportRepository.GetLampiranReportTemplate(uuid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}

	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"unscapeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
		"stripTags": func(original string) string {
			stripped := strip.StripTags(original)
			result := strings.ReplaceAll(stripped, "&nbsp;", " ")
			return result
		},
		"replaceStrip": func(original string) string {
			stripped := strings.ReplaceAll(original, "_", " ")
			return stripped
		},
	}

	var listTemplate = []string{}

	listTemplateLampiran := helper.GetArrFilename("templates/lampiran/")
	if len(listTemplateLampiran) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "Template tidak ditemukan"})
		return
	}
	listTemplate = append(listTemplate, listTemplateLampiran...)
	t := template.Must(template.New("preview.html").Funcs(funcMap).ParseFiles(listTemplate...))
	quizUser, errQuizUser := reportRepository.GetQuizUserDummyFromTemplate(report.IDQuizTemplate)
	if errQuizUser != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "Quiz Dummy Belum ada"})
		return
	}
	var skoring any
	var errSkoring error
	skoring, errSkoring = reportService.GetDataSkoringLampiranFromReportTabel(report.TabelReferensi, int(quizUser.IDQuiz), int(quizUser.IDUser))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errSkoring.Error()})
		return
	}

	data := map[string]interface{}{
		"nama_komponen": report.NamaReport,
		"blade":         report.Blade,
		"skoring":       skoring,
		"urutan":        report.Urutan,
	}
	err = t.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}

}

func (*controller) PreviewKomponenReportDummy(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	if uuid == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "Request Tidak Valid"})
		return
	}

	report, err := reportRepository.GetKomponenReportTemplate(uuid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}
	model, _ := reportRepository.GetModelReport(report.Model)

	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"unscapeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
		"stripTags": func(original string) string {
			stripped := strip.StripTags(original)
			result := strings.ReplaceAll(stripped, "&nbsp;", " ")
			return result
		},
		"replaceStrip": func(original string) string {
			stripped := strings.ReplaceAll(original, "_", " ")
			return stripped
		},
	}

	var listTemplate = []string{}
	fmt.Println(model)

	listTemplate = append(listTemplate, fmt.Sprintf("templates/%v/preview.html", model.Direktori))
	listTemplateKomponen := helper.GetArrFilename(fmt.Sprintf("templates/%v/komponen/", model.Direktori))
	if len(listTemplateKomponen) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "Template tidak ditemukan"})
		return
	}
	listTemplate = append(listTemplate, listTemplateKomponen...)
	t := template.Must(template.New("preview.html").Funcs(funcMap).ParseFiles(listTemplate...))
	quizUser, errQuizUser := reportRepository.GetQuizUserDummyFromTemplate(report.IDQuizTemplate)
	if errQuizUser != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "Quiz Dummy Belum ada"})
		return
	}
	var skoring any
	var errSkoring error
	skoring, errSkoring = reportService.GetDataSkoringFromReportTabel(report.TabelReferensi, int(quizUser.IDQuiz), int(quizUser.IDUser))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errSkoring.Error()})
		return
	}

	if report.TabelReferensi == "saran" {
		skoring, _ = reportRepository.GetDetilQuizTemplate(int(report.IDQuizTemplate))
		// fmt.Println(skoring)
	}

	if report.TabelReferensi == "ttd" {
		userQuiz := reportRepository.GetDetilQuizSesiUserCetak(int(quizUser.IDQuiz), int(quizUser.IDUser))
		quiz, _ := reportRepository.GetDetilQuizCetak(int(quizUser.IDQuiz))
		ttdInfo := map[string]interface{}{
			"quiz":      quiz,
			"skoringAt": helper.StringTimeTglIndo(userQuiz.SkoringAt),
		}
		skoring = ttdInfo
	}

	data := map[string]interface{}{
		"nama_komponen": report.NamaReport,
		"blade":         report.Blade,
		"urutan":        "",
		"skoring":       skoring,
	}
	err = t.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}

}
