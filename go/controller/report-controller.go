package controller

import (
	"encoding/json"
	"html/template"
	"irwanka/sicerdas/entity"
	"irwanka/sicerdas/helper"
	"irwanka/sicerdas/repository"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-chi/chi"
	strip "github.com/grokify/html-strip-tags-go"
)

const (
	ID_USER_DUMMY = 100
)

var (
	reportRepository repository.ReportRepository = repository.NewReportRepository()
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
			return stripped
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

	var skoring any

	if report.TabelReferensi == "skor_kognitif" {
		skoringCek, errSkoring := reportRepository.GetSkoringKognitif(int(quizUser.IDQuiz), int(quizUser.IDUser))
		if errSkoring != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errSkoring.Error()})
			return
		}
		skoring = skoringCek
	}

	if report.TabelReferensi == "skor_kognitif_pmk" {
		skoringCek, errSkoring := reportRepository.GetSkoringKognitifPMK(int(quizUser.IDQuiz), int(quizUser.IDUser))
		if errSkoring != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errSkoring.Error()})
			return
		}
		skoring = skoringCek
	}

	if report.TabelReferensi == "skor_kuliah_dinas" {
		refSekolahDinas, _ := reportRepository.GetReferensiSekolahDinas()
		skoringCek, errSkoring := reportRepository.GetSkoringKuliahDinas(int(quizUser.IDQuiz), int(quizUser.IDUser))
		if errSkoring != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errSkoring.Error()})
			return
		}
		var minatSekolah = []entity.RefSekolahDinas{}

		for i := 0; i < len(refSekolahDinas); i++ {
			if refSekolahDinas[i].No == skoringCek.MinatDinas1 {
				minatSekolah = append(minatSekolah, *refSekolahDinas[i])
			}
		}

		for i := 0; i < len(refSekolahDinas); i++ {
			if refSekolahDinas[i].No == skoringCek.MinatDinas2 {
				minatSekolah = append(minatSekolah, *refSekolahDinas[i])
			}
		}

		for i := 0; i < len(refSekolahDinas); i++ {
			if refSekolahDinas[i].No == skoringCek.MinatDinas3 {
				minatSekolah = append(minatSekolah, *refSekolahDinas[i])
			}
		}
		skoring = minatSekolah
	}

	if report.TabelReferensi == "skor_kuliah_alam" {
		refMinatKuliah, _ := reportRepository.GetReferensiKuliahIlmuAlam()
		skoringCek, errSkoring := reportRepository.GetSkoringKuliahAlam(int(quizUser.IDQuiz), int(quizUser.IDUser))
		if errSkoring != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errSkoring.Error()})
			return
		}
		var minatKuliah = []entity.RefKuliahAlam{}

		for i := 0; i < len(refMinatKuliah); i++ {
			if refMinatKuliah[i].Urutan == skoringCek.MinatIpa1 {
				minatKuliah = append(minatKuliah, *refMinatKuliah[i])
			}
		}

		for i := 0; i < len(refMinatKuliah); i++ {
			if refMinatKuliah[i].Urutan == skoringCek.MinatIpa2 {
				minatKuliah = append(minatKuliah, *refMinatKuliah[i])
			}
		}

		for i := 0; i < len(refMinatKuliah); i++ {
			if refMinatKuliah[i].Urutan == skoringCek.MinatIpa3 {
				minatKuliah = append(minatKuliah, *refMinatKuliah[i])
			}
		}

		for i := 0; i < len(refMinatKuliah); i++ {
			if refMinatKuliah[i].Urutan == skoringCek.MinatIpa4 {
				minatKuliah = append(minatKuliah, *refMinatKuliah[i])
			}
		}

		for i := 0; i < len(refMinatKuliah); i++ {
			if refMinatKuliah[i].Urutan == skoringCek.MinatIpa5 {
				minatKuliah = append(minatKuliah, *refMinatKuliah[i])
			}
		}

		skoring = minatKuliah
	}

	if report.TabelReferensi == "skor_kuliah_sosial" {
		refMinatKuliah, _ := reportRepository.GetReferensiKuliahIlmuSosial()
		skoringCek, errSkoring := reportRepository.GetSkoringKuliahSosial(int(quizUser.IDQuiz), int(quizUser.IDUser))
		if errSkoring != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errSkoring.Error()})
			return
		}
		var minatKuliah = []entity.RefKuliahSosial{}

		for i := 0; i < len(refMinatKuliah); i++ {
			if refMinatKuliah[i].Urutan == skoringCek.MinatIps1 {
				minatKuliah = append(minatKuliah, *refMinatKuliah[i])
			}
		}

		for i := 0; i < len(refMinatKuliah); i++ {
			if refMinatKuliah[i].Urutan == skoringCek.MinatIps2 {
				minatKuliah = append(minatKuliah, *refMinatKuliah[i])
			}
		}

		for i := 0; i < len(refMinatKuliah); i++ {
			if refMinatKuliah[i].Urutan == skoringCek.MinatIps3 {
				minatKuliah = append(minatKuliah, *refMinatKuliah[i])
			}
		}

		for i := 0; i < len(refMinatKuliah); i++ {
			if refMinatKuliah[i].Urutan == skoringCek.MinatIps4 {
				minatKuliah = append(minatKuliah, *refMinatKuliah[i])
			}
		}

		for i := 0; i < len(refMinatKuliah); i++ {
			if refMinatKuliah[i].Urutan == skoringCek.MinatIps5 {
				minatKuliah = append(minatKuliah, *refMinatKuliah[i])
			}
		}

		skoring = minatKuliah
	}

	if report.TabelReferensi == "skor_kuliah_agama" {
		refMinatKuliah, _ := reportRepository.GetReferensiKuliahIlmuAgama()
		skoringCek, errSkoring := reportRepository.GetSkoringKuliahAgama(int(quizUser.IDQuiz), int(quizUser.IDUser))
		if errSkoring != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errSkoring.Error()})
			return
		}
		var minatKuliah = []entity.RefKuliahAgama{}

		for i := 0; i < len(refMinatKuliah); i++ {
			if refMinatKuliah[i].Urutan == skoringCek.MinatAgm1 {
				minatKuliah = append(minatKuliah, *refMinatKuliah[i])
			}
		}

		for i := 0; i < len(refMinatKuliah); i++ {
			if refMinatKuliah[i].Urutan == skoringCek.MinatAgm2 {
				minatKuliah = append(minatKuliah, *refMinatKuliah[i])
			}
		}

		for i := 0; i < len(refMinatKuliah); i++ {
			if refMinatKuliah[i].Urutan == skoringCek.MinatAgm3 {
				minatKuliah = append(minatKuliah, *refMinatKuliah[i])
			}
		}

		for i := 0; i < len(refMinatKuliah); i++ {
			if refMinatKuliah[i].Urutan == skoringCek.MinatAgm4 {
				minatKuliah = append(minatKuliah, *refMinatKuliah[i])
			}
		}

		for i := 0; i < len(refMinatKuliah); i++ {
			if refMinatKuliah[i].Urutan == skoringCek.MinatAgm5 {
				minatKuliah = append(minatKuliah, *refMinatKuliah[i])
			}
		}

		skoring = minatKuliah
	}

	if report.TabelReferensi == "skor_suasana_kerja" {
		refMinat, _ := reportRepository.GetReferensiSuasanaKerja()
		skoringCek, errSkoring := reportRepository.GetSkoringSuasanaKerja(int(quizUser.IDQuiz), int(quizUser.IDUser))
		if errSkoring != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errSkoring.Error()})
			return
		}
		var minat = []entity.RefSuasanaKerja{}

		for i := 0; i < len(refMinat); i++ {
			if refMinat[i].Nomor == skoringCek.SuasanaKerja1 {
				minat = append(minat, *refMinat[i])
			}
		}

		for i := 0; i < len(refMinat); i++ {
			if refMinat[i].Nomor == skoringCek.SuasanaKerja2 {
				minat = append(minat, *refMinat[i])
			}
		}

		for i := 0; i < len(refMinat); i++ {
			if refMinat[i].Nomor == skoringCek.SuasanaKerja3 {
				minat = append(minat, *refMinat[i])
			}
		}

		skoring = minat
	}

	if report.TabelReferensi == "skor_sikap_pelajaran" {
		refSikap, _ := reportRepository.GetReferensiSikapPelajaran()
		skoringCek, errSkoring := reportRepository.GetSkoringSikapPelajaran(int(quizUser.IDQuiz), int(quizUser.IDUser))
		if errSkoring != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errSkoring.Error()})
			return
		}
		var sikap = []entity.ResultSikapPelajaran{}
		var currentKelompok = ""
		var countKelompok = map[string]int{}
		for i := 0; i < len(refSikap); i++ {
			rv := reflect.ValueOf(skoringCek)
			var temp = entity.ResultSikapPelajaran{}
			temp.Urutan = int32(i) + 1
			temp.Kelompok = refSikap[i].Kelompok

			if currentKelompok != temp.Kelompok {
				currentKelompok = temp.Kelompok
				countKelompok[currentKelompok] = 1
			} else {
				countKelompok[currentKelompok] = countKelompok[currentKelompok] + 1
			}
			temp.Pelajaran = refSikap[i].Pelajaran
			kode := helper.Capitalize(strings.ToLower(refSikap[i].Kode))
			klasifikasiName := "Klasifikasi" + kode
			temp.Klasifikasi = reflect.Indirect(rv).FieldByName(klasifikasiName).String()
			sikap = append(sikap, temp)
		}
		skoring = map[string]interface{}{
			"data":  sikap,
			"group": countKelompok,
		}
	}

	if report.TabelReferensi == "skor_sikap_pelajaran_mk" {
		refSikap, _ := reportRepository.GetReferensiSikapPelajaranMK()
		skoringCek, errSkoring := reportRepository.GetSkoringSikapPelajaranMK(int(quizUser.IDQuiz), int(quizUser.IDUser))
		if errSkoring != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errSkoring.Error()})
			return
		}
		var sikap = []entity.ResultSikapPelajaranMK{}
		for i := 0; i < len(refSikap); i++ {
			rv := reflect.ValueOf(skoringCek)
			var temp = entity.ResultSikapPelajaranMK{}
			temp.Urutan = int32(i) + 1
			temp.Kelompok = refSikap[i].Kelompok
			temp.Pelajaran = refSikap[i].Pelajaran
			kode := helper.Capitalize(strings.ToLower(refSikap[i].Kode))
			klasifikasiName := "Klasifikasi" + kode
			temp.Klasifikasi = reflect.Indirect(rv).FieldByName(klasifikasiName).String()
			sikap = append(sikap, temp)
		}
		skoring = sikap
	}

	if report.TabelReferensi == "skor_peminatan_sma" {
		// refMinat, _ := reportRepository.GetReferensiMinatSMA()
		// skoringCek, errSkoring := reportRepository.GetSkoringPeminatanSMA(int(quizUser.IDQuiz), int(quizUser.IDUser))
		// if errSkoring != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errSkoring.Error()})
		// 	return
		// }
		// var minat = []entity.ResultPeminatanSMA{}
		// for i := 0; i < len(refSikap); i++ {
		// 	rv := reflect.ValueOf(skoringCek)
		// 	var temp = entity.ResultPeminatanSMA{}
		// 	temp.Urutan = int32(i) + 1
		// 	temp.Kelompok = refSikap[i].Kelompok
		// 	temp.Pelajaran = refSikap[i].Pelajaran
		// 	kode := helper.Capitalize(strings.ToLower(refSikap[i].Kode))
		// 	klasifikasiName := "Klasifikasi" + kode
		// 	temp.Klasifikasi = reflect.Indirect(rv).FieldByName(klasifikasiName).String()
		// 	sikap = append(sikap, temp)
		// }
		// skoring = sikap
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
