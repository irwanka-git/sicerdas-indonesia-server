package controller

import (
	"encoding/json"
	"html/template"
	"irwanka/sicerdas/helper"
	"irwanka/sicerdas/repository"
	"irwanka/sicerdas/service"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
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
}

func NewReportController() ReportController {
	return &controller{}
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
	listTemplate = append(listTemplate, "templates/views/preview.html")
	listTemplateKomponen := helper.GetArrFilename("templates/views/komponen/")
	listTemplate = append(listTemplate, listTemplateKomponen...)
	t := template.Must(template.New("preview.html").Funcs(funcMap).ParseFiles(listTemplate...))
	quizUser, errQuizUser := reportRepository.GetQuizUserDummyFromTemplate(report.IDQuizTemplate)
	if errQuizUser != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "Quiz Dummy Belum ada"})
		return
	}

	skoring, errSkoring := reportService.GetDataSkoringFromReportTabel(report.TabelReferensi, int(quizUser.IDQuiz), int(quizUser.IDUser))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errSkoring.Error()})
		return
	}

	data := map[string]interface{}{
		"nama_komponen": report.NamaReport,
		"blade":         report.Blade,
		"skoring":       skoring,
	}
	err = t.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}

}
