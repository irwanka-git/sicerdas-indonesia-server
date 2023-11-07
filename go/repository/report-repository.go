package repository

import (
	"errors"
	"fmt"
	"irwanka/sicerdas/entity"
	"reflect"
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

	GetReferensiMinatMAN() ([]*entity.RefPilihanMinatMan, error)
	GetSkoringPeminatanMAN(id_quiz int, id_user int) (*entity.SkorPeminatanMan, error)

	GetReferensiMinatTMI() ([]*entity.RefPilihanMinatTMI, error)
	GetSkoringPeminatanTMI(id_quiz int, id_user int) (*entity.SkorMinatIndonesium, error)

	GetInterprestasiTipologiJung(kode string) (entity.InterprestasiTipologiJung, error)
	GetSkoringMBTI(id_quiz int, id_user int) (*entity.SkorMbti, error)

	GetKomponenKarakteristikPribadi() ([]*entity.RefKomponenKarakteristikPribadi, error)
	GetSkoringKarakteristikPribadi(id_quiz int, id_user int) (*entity.SkorKarakteristikPribadi, error)

	GetResultGayaPekerjaan(id_quiz int, id_user int) ([]*entity.ResultGayaPekerjaan, error)
	GetSkoringGayaPekerjaan(id_quiz int, id_user int) (*entity.SkorGayaPekerjaan, error)

	GetResultGayaBelajar(id_quiz int, id_user int) ([]*entity.ResultGayaBelajar, error)

	GetResultPeminatanSMK(id_quiz int, id_user int) ([]*entity.ResultPeminatanSMK, error)
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

func (*repo) GetReferensiMinatMAN() ([]*entity.RefPilihanMinatMan, error) {
	var data []*entity.RefPilihanMinatMan
	db.Table("ref_pilihan_minat_sma").Scan(&data)
	return data, nil
}
func (*repo) GetSkoringPeminatanMAN(id_quiz int, id_user int) (*entity.SkorPeminatanMan, error) {
	var data *entity.SkorPeminatanMan
	db.Table("skor_peminatan_man").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&data)
	return data, nil
}

func (*repo) GetReferensiMinatTMI() ([]*entity.RefPilihanMinatTMI, error) {
	var data []*entity.RefPilihanMinatTMI
	db.Table("soal_tmi").Scan(&data)
	return data, nil
}
func (*repo) GetSkoringPeminatanTMI(id_quiz int, id_user int) (*entity.SkorMinatIndonesium, error) {
	var data *entity.SkorMinatIndonesium
	db.Table("skor_minat_indonesia").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&data)
	return data, nil
}

func (*repo) GetInterprestasiTipologiJung(kode string) (entity.InterprestasiTipologiJung, error) {
	var result = entity.InterprestasiTipologiJung{}
	db.Table("interprestasi_tipologi_jung").Where("kode = ?", kode).First(&result)
	return result, nil
}

func (*repo) GetSkoringMBTI(id_quiz int, id_user int) (*entity.SkorMbti, error) {
	var data *entity.SkorMbti
	db.Table("skor_mbti").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&data)
	return data, nil
}

func (*repo) GetKomponenKarakteristikPribadi() ([]*entity.RefKomponenKarakteristikPribadi, error) {
	var result []*entity.RefKomponenKarakteristikPribadi
	db.Table("ref_komponen_karakteristik_pribadi").Order("id_komponen asc").Scan(&result)
	return result, nil
}
func (*repo) GetSkoringKarakteristikPribadi(id_quiz int, id_user int) (*entity.SkorKarakteristikPribadi, error) {
	var data *entity.SkorKarakteristikPribadi
	db.Table("skor_karakteristik_pribadi").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&data)
	return data, nil
}

func (*repo) GetResultGayaPekerjaan(id_quiz int, id_user int) ([]*entity.ResultGayaPekerjaan, error) {
	var data []*entity.ResultGayaPekerjaan
	db.Raw(`SELECT A.id,
				A.rangking,
				C.nama_komponen,
				C.kode,
				C.cetak_komponen,
				A.skor,
				A.klasifikasi ,
				C.deskripsi,
      			C.pekerjaan
			FROM
				ref_skoring_gaya_pekerjaan AS A,
				ref_komponen_gaya_pekerjaan AS C 
			WHERE
				C.kode = A.kode 
				AND A.id_quiz = ? 
				AND A.id_user = ? 
			ORDER BY
				A.rangking`, id_quiz, id_user).Scan(&data)
	return data, nil
}

func (*repo) GetSkoringGayaPekerjaan(id_quiz int, id_user int) (*entity.SkorGayaPekerjaan, error) {
	var data *entity.SkorGayaPekerjaan
	db.Table("skor_gaya_pekerjaan").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&data)
	return data, nil
}

func (*repo) GetResultGayaBelajar(id_quiz int, id_user int) ([]*entity.ResultGayaBelajar, error) {
	var ref []*entity.ResultGayaBelajar
	db.Raw(`select a.*, 
	concat('Gaya', INITCAP(a.nama)) as field_name, 
	concat('Klasifikasi', INITCAP(a.nama)) as klasifikasi_name
	from ref_gaya_belajar as a order by a.kode`).Scan(&ref)

	var skoring *entity.SkorGayaBelajar
	db.Table("skor_gaya_belajar").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&skoring)

	var result = []*entity.ResultGayaBelajar{}
	rv := reflect.ValueOf(skoring)
	for i := 0; i < len(ref); i++ {
		klasifikasiName := ref[i].KlasifikasiName
		// fmt.Println(klasifikasiName)
		klasifikasi := reflect.Indirect(rv).FieldByName(klasifikasiName).String()
		fieldname := ref[i].FieldName
		// fmt.Println(fieldname)
		skor := reflect.Indirect(rv).FieldByName(fieldname).Int()
		var temp = entity.ResultGayaBelajar{}
		temp.Deskripsi = ref[i].Deskripsi
		temp.FieldName = ref[i].FieldName
		temp.Klasifikasi = klasifikasi
		temp.KlasifikasiName = ref[i].Klasifikasi

		temp.Kode = ref[i].Kode
		temp.Gambar = ref[i].Gambar
		temp.Skor = int(skor)
		temp.Nama = ref[i].Nama
		result = append(result, &temp)
	}
	return result, nil
}

func (*repo) GetResultPeminatanSMK(id_quiz int, id_user int) ([]*entity.ResultPeminatanSMK, error) {
	var ref []*entity.ResultPeminatanSMK
	db.Raw(`select b.nomor , b.gambar , b.keterangan  , b.deskripsi  
	from quiz_sesi_mapping_smk as a, soal_peminatan_smk  as b 
		where a.id_kegiatan  = b.id_kegiatan and a.id_quiz  = ?`, id_quiz).Scan(&ref)

	var result []*entity.ResultPeminatanSMK
	var skoring *entity.SkorPeminatanSmk
	db.Table("skor_peminatan_smk").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&skoring)
	rv := reflect.ValueOf(skoring)
	for p := 1; p <= len(ref); p++ {
		fieldminat := fmt.Sprintf("Minat%v", p)
		fmt.Println(fieldminat)
		nomor := reflect.Indirect(rv).FieldByName(fieldminat).String()
		fmt.Println(nomor)
		for i := 0; i < len(ref); i++ {
			if nomor == ref[i].Nomor {
				var temp = entity.ResultPeminatanSMK{}
				temp.Deskripsi = ref[i].Deskripsi
				temp.Gambar = ref[i].Gambar
				temp.Keterangan = ref[i].Keterangan
				temp.Nomor = ref[i].Nomor
				temp.Urutan = p
				result = append(result, &temp)
			}
		}
	}
	return result, nil
}
