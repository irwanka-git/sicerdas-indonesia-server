package service

import (
	"irwanka/sicerdas/entity"
	"irwanka/sicerdas/helper"
	"irwanka/sicerdas/repository"
	"reflect"
	"strings"
)

type ReportService interface {
	GetDataSkoringFromReportTabel(tabel_referensi string, id_quiz int, id_user int) (any, error)
}

var (
	reportRepository repository.ReportRepository
)

func NewReporService(repo repository.ReportRepository) ReportService {
	reportRepository = repo
	return &service{}
}

func (*service) GetDataSkoringFromReportTabel(tabel_referensi string, id_quiz int, id_user int) (any, error) {
	var skoring any
	if tabel_referensi == "skor_kognitif" {
		skoringCek, errSkoring := reportRepository.GetSkoringKognitif(id_quiz, id_user)
		if errSkoring != nil {
			return nil, errSkoring
		}
		skoring = skoringCek
	}

	if tabel_referensi == "skor_kognitif_pmk" {
		skoringCek, errSkoring := reportRepository.GetSkoringKognitifPMK(id_quiz, id_user)
		if errSkoring != nil {
			return nil, errSkoring
		}
		skoring = skoringCek
	}

	if tabel_referensi == "skor_kuliah_dinas" {
		refSekolahDinas, _ := reportRepository.GetReferensiSekolahDinas()
		skoringCek, errSkoring := reportRepository.GetSkoringKuliahDinas(id_quiz, id_user)
		if errSkoring != nil {
			return nil, errSkoring
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

	if tabel_referensi == "skor_kuliah_alam" {
		refMinatKuliah, _ := reportRepository.GetReferensiKuliahIlmuAlam()
		skoringCek, errSkoring := reportRepository.GetSkoringKuliahAlam(id_quiz, id_user)
		if errSkoring != nil {
			return nil, errSkoring
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

	if tabel_referensi == "skor_kuliah_sosial" {
		refMinatKuliah, _ := reportRepository.GetReferensiKuliahIlmuSosial()
		skoringCek, errSkoring := reportRepository.GetSkoringKuliahSosial(id_quiz, id_user)
		if errSkoring != nil {
			return nil, errSkoring
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

	if tabel_referensi == "skor_kuliah_agama" {
		refMinatKuliah, _ := reportRepository.GetReferensiKuliahIlmuAgama()
		skoringCek, errSkoring := reportRepository.GetSkoringKuliahAgama(id_quiz, id_user)
		if errSkoring != nil {
			return nil, errSkoring
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

	if tabel_referensi == "skor_suasana_kerja" {
		refMinat, _ := reportRepository.GetReferensiSuasanaKerja()
		skoringCek, errSkoring := reportRepository.GetSkoringSuasanaKerja(id_quiz, id_user)
		if errSkoring != nil {
			return nil, errSkoring
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

	if tabel_referensi == "skor_sikap_pelajaran" {
		refSikap, _ := reportRepository.GetReferensiSikapPelajaran()
		skoringCek, errSkoring := reportRepository.GetSkoringSikapPelajaran(id_quiz, id_user)
		if errSkoring != nil {
			return nil, errSkoring
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

	if tabel_referensi == "skor_sikap_pelajaran_mk" {
		refSikap, _ := reportRepository.GetReferensiSikapPelajaranMK()
		skoringCek, errSkoring := reportRepository.GetSkoringSikapPelajaranMK(id_quiz, id_user)
		if errSkoring != nil {
			return nil, errSkoring
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

	if tabel_referensi == "skor_peminatan_sma" {
		refMinat, _ := reportRepository.GetReferensiMinatSMA()
		skoringCek, errSkoring := reportRepository.GetSkoringPeminatanSMA(id_quiz, id_user)
		if errSkoring != nil {
			return nil, errSkoring
		}
		// fmt.Println(skoringCek.KlasifikasiMinatSains)
		dataSkoring := map[string]interface{}{
			"data": skoringCek,
			"ref":  refMinat,
		}

		skoring = dataSkoring
	}

	if tabel_referensi == "skor_peminatan_man" {
		refMinat, _ := reportRepository.GetReferensiMinatSMA()
		skoringCek, errSkoring := reportRepository.GetSkoringPeminatanSMA(id_quiz, id_user)
		if errSkoring != nil {
			return nil, errSkoring
		}
		// fmt.Println(skoringCek.KlasifikasiMinatSains)
		dataSkoring := map[string]interface{}{
			"data": skoringCek,
			"ref":  refMinat,
		}

		skoring = dataSkoring
	}

	if tabel_referensi == "skor_peminatan_smk" {
		result, _ := reportRepository.GetResultPeminatanSMK(id_quiz, id_user)
		skoring = result
	}

	if tabel_referensi == "skor_minat_indonesia" {
		refMinat, _ := reportRepository.GetReferensiMinatTMI()
		skoringCek, errSkoring := reportRepository.GetSkoringPeminatanTMI(id_quiz, id_user)
		if errSkoring != nil {
			return nil, errSkoring
		}
		var minat = [7]entity.ResultMinatTMI{}

		for i := 0; i < len(refMinat); i++ {
			var temp = entity.ResultMinatTMI{}
			temp.Minat = refMinat[i].Minat
			temp.Keterangan = refMinat[i].Keterangan
			if refMinat[i].Urutan == skoringCek.Minat1 {
				minat[0] = temp
			}
			if refMinat[i].Urutan == skoringCek.Minat2 {
				minat[1] = temp
			}
			if refMinat[i].Urutan == skoringCek.Minat3 {
				minat[2] = temp
			}
			if refMinat[i].Urutan == skoringCek.Minat4 {
				minat[3] = temp
			}
			if refMinat[i].Urutan == skoringCek.Minat5 {
				minat[4] = temp
			}
			if refMinat[i].Urutan == skoringCek.Minat6 {
				minat[5] = temp
			}
			if refMinat[i].Urutan == skoringCek.Minat7 {
				minat[6] = temp
			}
		}
		// fmt.Println(skoringCek.KlasifikasiMinatSains)
		dataSkoring := map[string]interface{}{
			"data":  skoringCek,
			"minat": minat,
		}
		skoring = dataSkoring
	}

	if tabel_referensi == "skor_mbti" {

		skoringCek, errSkoring := reportRepository.GetSkoringMBTI(id_quiz, id_user)
		if errSkoring != nil {
			return nil, errSkoring
		}
		tipologi, _ := reportRepository.GetInterprestasiTipologiJung(skoringCek.TipojungKode)

		// fmt.Println(skoringCek.KlasifikasiMinatSains)
		dataSkoring := map[string]interface{}{
			"data":     skoringCek,
			"tipologi": tipologi,
		}
		skoring = dataSkoring
	}

	if tabel_referensi == "skor_karakteristik_pribadi" {
		komponen, _ := reportRepository.GetKomponenKarakteristikPribadi()
		skoringCek, errSkoring := reportRepository.GetSkoringKarakteristikPribadi(id_quiz, id_user)
		if errSkoring != nil {
			return nil, errSkoring
		}
		var result = []entity.ResultKarakteristikPribadi{}

		for i := 0; i < len(komponen); i++ {
			rv := reflect.ValueOf(skoringCek)
			var temp = entity.ResultKarakteristikPribadi{}
			temp.Urutan = int32(i) + 1
			temp.NamaKomponen = komponen[i].NamaKomponen
			temp.Keterangan = komponen[i].Keterangan
			kode := helper.Capitalize(strings.ReplaceAll(strings.ToLower(komponen[i].FieldSkoring), "pribadi_", ""))
			klasifikasiName := "Klasifikasi" + kode
			temp.Klasifikasi = reflect.Indirect(rv).FieldByName(klasifikasiName).String()
			result = append(result, temp)
		}
		skoring = result
	}

	if tabel_referensi == "skor_gaya_pekerjaan" {
		result, _ := reportRepository.GetResultGayaPekerjaan(id_quiz, id_user)
		skoring = result
	}

	if tabel_referensi == "skor_gaya_belajar" {
		result, _ := reportRepository.GetResultGayaBelajar(id_quiz, id_user)
		skoring = result
	}

	return skoring, nil
}
