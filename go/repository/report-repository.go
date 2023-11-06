package repository

import (
	"errors"
	"irwanka/sicerdas/entity"
)

var (
	dummyDataRepo DummyQuizUserRepository = NewDummyQuizUserRepository()
)

type ReportRepository interface {
	GetKomponenReportTemplate(uuid string) (*entity.QuizSesiReportAndTemplate, error)
	GetQuizUserDummyFromTemplate(id_quiz_template int32) (*entity.QuizSesiUser, error)
	//data skoring
	GetSkoringKognitif(id_quiz int, id_user int) (*entity.SkorKognitif, error)
	GetSkoringKognitifPMK(id_quiz int, id_user int) (*entity.SkorKognitif, error)

	GetReferensiSekolahDinas() ([]*entity.RefSekolahDinas, error)
	GetSkoringKuliahDinas(id_quiz int, id_user int) (*entity.SkorKuliahDinas, error)

	GetReferensiKuliahIlmuAlam() ([]*entity.RefKuliahAlam, error)
	GetSkoringKuliahAlam(id_quiz int, id_user int) (*entity.SkorKuliahAlam, error)

	GetReferensiKuliahIlmuSosial() ([]*entity.RefKuliahSosial, error)
	GetSkoringKuliahSosial(id_quiz int, id_user int) (*entity.SkorKuliahSosial, error)

	GetReferensiKuliahIlmuAgama() ([]*entity.RefKuliahAgama, error)
	GetSkoringKuliahAgama(id_quiz int, id_user int) (*entity.SkorKuliahAgama, error)

	GetReferensiSuasanaKerja() ([]*entity.RefSuasanaKerja, error)
	GetSkoringSuasanaKerja(id_quiz int, id_user int) (*entity.SkorSuasanaKerja, error)

	GetReferensiSikapPelajaran() ([]*entity.RefSikapPelajaran, error)
	GetSkoringSikapPelajaran(id_quiz int, id_user int) (*entity.SkorSikapPelajaran, error)
	GetReferensiSikapPelajaranMK() ([]*entity.RefSikapPelajaranMK, error)
	GetSkoringSikapPelajaranMK(id_quiz int, id_user int) (*entity.SkorSikapPelajaranMk, error)

	GetReferensiMinatSMA() ([]*entity.RefPilihanMinatSma, error)
	GetSkoringPeminatanSMA(id_quiz int, id_user int) (*entity.SkorPeminatanSma, error)
}

func NewReportRepository() ReportRepository {
	return &repo{}
}

func (*repo) GetKomponenReportTemplate(uuid string) (*entity.QuizSesiReportAndTemplate, error) {
	var report *entity.QuizSesiReportAndTemplate
	result := db.Raw(`select b.*, a.id_quiz_template from quiz_sesi_template_report as a , quiz_sesi_report as b 
	where a.id_report  = b.id_report  and a.uuid  =? `, uuid).First(&report)
	if result.RowsAffected == 0 {
		return nil, errors.New("not found template report")
	}
	return report, nil
}

func (*repo) GetQuizUserDummyFromTemplate(id_quiz_template int32) (*entity.QuizSesiUser, error) {
	return dummyDataRepo.CekDummyQuizUser(int(id_quiz_template))
}

func (*repo) GetSkoringKognitif(id_quiz int, id_user int) (*entity.SkorKognitif, error) {
	var data *entity.SkorKognitif
	db.Table("skor_kognitif").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&data)

	return data, nil
}

func (*repo) GetSkoringKognitifPMK(id_quiz int, id_user int) (*entity.SkorKognitif, error) {
	var data *entity.SkorKognitif
	db.Table("skor_kognitif_pmk").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&data)
	return data, nil
}

func (*repo) GetSkoringKuliahDinas(id_quiz int, id_user int) (*entity.SkorKuliahDinas, error) {
	var data *entity.SkorKuliahDinas
	db.Table("skor_kuliah_dinas").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&data)
	return data, nil
}

func (*repo) GetReferensiSekolahDinas() ([]*entity.RefSekolahDinas, error) {
	var data []*entity.RefSekolahDinas
	db.Table("ref_sekolah_dinas").Scan(&data)
	return data, nil
}

func (*repo) GetReferensiKuliahIlmuAlam() ([]*entity.RefKuliahAlam, error) {
	var data []*entity.RefKuliahAlam
	db.Table("soal_minat_kuliah_eksakta").Scan(&data)
	return data, nil
}

func (*repo) GetSkoringKuliahAlam(id_quiz int, id_user int) (*entity.SkorKuliahAlam, error) {
	var data *entity.SkorKuliahAlam
	db.Table("skor_kuliah_alam").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&data)
	return data, nil
}

func (*repo) GetReferensiKuliahIlmuSosial() ([]*entity.RefKuliahSosial, error) {
	var data []*entity.RefKuliahSosial
	db.Table("soal_minat_kuliah_sosial").Scan(&data)
	return data, nil
}

func (*repo) GetSkoringKuliahSosial(id_quiz int, id_user int) (*entity.SkorKuliahSosial, error) {
	var data *entity.SkorKuliahSosial
	db.Table("skor_kuliah_sosial").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&data)
	return data, nil
}

func (*repo) GetReferensiKuliahIlmuAgama() ([]*entity.RefKuliahAgama, error) {
	var data []*entity.RefKuliahAgama
	db.Table("soal_minat_kuliah_agama").Scan(&data)
	return data, nil

}
func (*repo) GetSkoringKuliahAgama(id_quiz int, id_user int) (*entity.SkorKuliahAgama, error) {
	var data *entity.SkorKuliahAgama
	db.Table("skor_kuliah_agama").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&data)
	return data, nil
}

func (*repo) GetReferensiSuasanaKerja() ([]*entity.RefSuasanaKerja, error) {
	var data []*entity.RefSuasanaKerja
	db.Table("soal_minat_kuliah_suasana_kerja").Scan(&data)
	return data, nil
}
func (*repo) GetSkoringSuasanaKerja(id_quiz int, id_user int) (*entity.SkorSuasanaKerja, error) {
	var data *entity.SkorSuasanaKerja
	db.Table("skor_suasana_kerja").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&data)
	return data, nil
}

func (*repo) GetReferensiSikapPelajaranMK() ([]*entity.RefSikapPelajaranMK, error) {
	var data []*entity.RefSikapPelajaranMK
	db.Table("soal_sikap_pelajaran_kuliah").Scan(&data)
	return data, nil
}
func (*repo) GetSkoringSikapPelajaranMK(id_quiz int, id_user int) (*entity.SkorSikapPelajaranMk, error) {
	var data *entity.SkorSikapPelajaranMk
	db.Table("skor_sikap_pelajaran_mk").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&data)
	return data, nil
}

func (*repo) GetReferensiSikapPelajaran() ([]*entity.RefSikapPelajaran, error) {
	var data []*entity.RefSikapPelajaran
	db.Table("soal_sikap_pelajaran").Scan(&data)
	return data, nil
}
func (*repo) GetSkoringSikapPelajaran(id_quiz int, id_user int) (*entity.SkorSikapPelajaran, error) {
	var data *entity.SkorSikapPelajaran
	db.Table("skor_sikap_pelajaran").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&data)
	return data, nil
}

func (*repo) GetReferensiMinatSMA() ([]*entity.RefPilihanMinatSma, error) {
	var data []*entity.RefPilihanMinatSma
	db.Table("ref_pilihan_minat_sma").Scan(&data)
	return data, nil
}
func (*repo) GetSkoringPeminatanSMA(id_quiz int, id_user int) (*entity.SkorPeminatanSma, error) {
	var data *entity.SkorPeminatanSma
	db.Table("skor_peminatan_sma").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&data)
	return data, nil
}
