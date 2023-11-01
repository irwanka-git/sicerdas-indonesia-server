package repository

import (
	"encoding/json"
	"fmt"
	"irwanka/sicerdas/entity"
	"irwanka/sicerdas/helper"
	"reflect"
	"strings"
)

type SkoringRepository interface {
	GetStatusRunningSkoring() (*entity.StatusRunningSkoring, error)
	StartRunningSkoring() error
	StopRunningSkoring() error
	GetUserSesiBelumSkoring() ([]*entity.QuizSesiUserSkoring, error)
	ClearTabelTemporaryJawabanUser(id_quiz int32, id_user int32) error
	GenerateTabelTemporaryJawabanUser(id_quiz int32, id_user int32) error
	GetKategoriTabelSkoring(id_quiz int32) ([]*entity.KategoriTabel, error)
	GetTabelSkoring(id_quiz int32) ([]*entity.KategoriTabel, error)

	//SKORING KOGNITIF
	SkoringKognitif(id_quiz int32, id_user int32) error
	SkoringKognitifPMK(id_quiz int32, id_user int32) error
	SkoringGayaPekerjaan(id_quiz int32, id_user int32) error
	SkoringSikapPelajaran(id_quiz int32, id_user int32) error
	SkoringSikapPelajaranMK(id_quiz int32, id_user int32) error
	SkoringPeminatanSMK(id_quiz int32, id_user int32) error
	SkoringPeminatanSMA(id_quiz int32, id_user int32) error
	SkoringPeminatanMAN(id_quiz int32, id_user int32) error
	SkoringMinatKuliahKedinasan(id_quiz int32, id_user int32) error
}

func NewSkoringRepository() SkoringRepository {
	return &repo{}
}
func (*repo) GetStatusRunningSkoring() (*entity.StatusRunningSkoring, error) {
	var cek *entity.StatusRunningSkoring
	db.Table("status_skoring").First(&cek)
	return cek, nil
}
func (*repo) StartRunningSkoring() error {
	db.Table("status_skoring").Where("id = ?", 1).UpdateColumn("status", 1)
	return nil
}

func (*repo) StopRunningSkoring() error {
	db.Table("status_skoring").Where("id = ?", 1).UpdateColumn("status", 0)
	return nil
}

func (*repo) GetUserSesiBelumSkoring() ([]*entity.QuizSesiUserSkoring, error) {
	var list []*entity.QuizSesiUserSkoring
	db.Raw(`select 
		a.id_quiz, 
		a.id_user, 
		a.skoring, 
		a.submit , 
		c.token, 
		b.nama_sesi, 
		c.id_quiz_template,
		c.kota
	from  
		quiz_sesi_user  as a,  
		quiz_sesi_template as b , 
		quiz_sesi as c 
		where a.submit = 1 and a.skoring = 0
		and b.id_quiz_template = c.id_quiz_template
		and c.id_quiz = a.id_quiz
		and c.jenis = 'quiz' 
		and a.id_user = 243 and a.id_quiz = 166
		order by a.id_quiz_user asc
	`).Scan(&list)
	return list, nil
}

func (*repo) GetKategoriTabelSkoring(id_quiz int32) ([]*entity.KategoriTabel, error) {
	var list []*entity.KategoriTabel
	db.Raw(`select 
		b.kategori, b.tabel 
	from quiz_sesi_detil as a, quiz_sesi_master as b 
		where a.id_sesi_master = b.id_sesi_master
		and a.id_quiz = ? and b.tabel is not null 
		and trim(b.tabel) !='-' and trim(b.tabel) !=''
		order by tabel asc
		`, id_quiz).Scan(&list)
	return list, nil
}
func (*repo) GetTabelSkoring(id_quiz int32) ([]*entity.KategoriTabel, error) {
	var list []*entity.KategoriTabel
	db.Raw(`select 
		b.tabel , count(*) 
	from quiz_sesi_detil as a, quiz_sesi_master as b 
		where a.id_sesi_master = b.id_sesi_master
		and a.id_quiz = ? and b.tabel is not null 
		and trim(b.tabel) !='-' and trim(b.tabel) !=''
		group by tabel
		order by tabel asc
		`, id_quiz).Scan(&list)
	return list, nil
}

func (*repo) ClearTabelTemporaryJawabanUser(id_quiz int32, id_user int32) error {
	var delete = entity.QuizSesiUserJawaban{}
	db.Table("quiz_sesi_user_jawaban").
		Where("id_quiz = ? ", id_quiz).
		Where("id_user = ? ", id_user).
		Delete(delete)
	return nil
}

func (*repo) GenerateTabelTemporaryJawabanUser(id_quiz int32, id_user int32) error {
	var jawaban *entity.JawabanSubmit
	db.Table("quiz_sesi_user").
		Where("id_quiz = ?", id_quiz).
		Where("id_user = ?", id_user).
		First(&jawaban)
	// fmt.Println(jawaban.Jawaban)
	var data = []entity.JawabanMapping{}
	json.Unmarshal([]byte(jawaban.Jawaban), &data)
	for i := 0; i < len(data); i++ {
		json.Unmarshal([]byte(data[i].Jawaban), &data[i].JawabanArray)
	}
	for i := 0; i < len(data); i++ {
		kategori := data[i].Kategori
		for urutan := 1; urutan < len(data[i].JawabanArray); urutan++ {
			// db.Exec("delete * from quiz_sesi_user_jawaban where id_quiz = ? and id_user = ? and kategori = ? and urutan = ? ", id_quiz, id_user, kategori, urutan)
			var current *entity.JawabanSubmit
			db.Table("quiz_sesi_user_jawaban").
				Where("id_quiz = ? ", id_quiz).
				Where("id_user = ? ", id_user).
				Where("kategori = ? ", kategori).
				Where("urutan = ? ", urutan).
				Delete(&current)

			var insertJawaban = entity.QuizSesiUserJawaban{}
			insertJawaban.IDQuiz = id_quiz
			insertJawaban.IDUser = id_user
			insertJawaban.Kategori = kategori
			insertJawaban.Urutan = int32(urutan)
			insertJawaban.Jawaban = data[i].JawabanArray[urutan]
			insertJawaban.Skor = 0
			db.Table("quiz_sesi_user_jawaban").Create(&insertJawaban)
		}
	}
	return nil
}

func (*repo) SkoringKognitif(id_quiz int32, id_user int32) error {
	//regenereate  delete current if exist
	var skoring = entity.SkorKognitif{}
	db.Table("skor_kognitif").Where("id_user = ?", id_user).Where("id_quiz = ? ", id_quiz).Delete(&skoring)

	//koreksi KOGNITIF_ (KOGNITIF KELAS 1 SMA) prefix= KOGNITIF_
	// a.paket = 'NON' and   a.bidang::TEXT = replace(b.kategori, 'KOGNITIF_','')
	db.Exec(`update quiz_sesi_user_jawaban  as t set skor = 1 
				from ( 
				select b.id_quiz, 
				case when a.pilihan_jawaban = b.jawaban then 1 else 0 end as skor,
					b.id_jawaban_peserta 
				from  soal_kognitif as a , quiz_sesi_user_jawaban as b  
				where a.paket = 'NON' and   a.bidang::TEXT = replace(b.kategori, 'KOGNITIF_','')
				and a.urutan = b.urutan  and b.id_quiz = ?  and left(b.kategori,9) = 'KOGNITIF_'
				and b.id_user = ?) as x 
				where x.skor = 1 
				and t.id_jawaban_peserta = x.id_jawaban_peserta`,
		id_quiz,
		id_user)

	//update skoring dan klasifikasi
	type groupSkoring struct {
		Kategori    string `json:"kategori"`
		Skor        int    `json:"skor"`
		Klasifikasi string `json:"klasifikasi"`
	}
	var skors []*groupSkoring
	db.Raw(`select s.kategori, s.skor, 
				replace(case when kd15.klasifikasi is not null then  kd15.klasifikasi::text 
					when kd20.klasifikasi is not null then  kd20.klasifikasi::text  
					when s.skor = 0 then 'SANGAT_RENDAH'
					end,'_',' ') as klasifikasi
			from  
				(select kategori , sum(skor) as skor, count(skor) as soal 
				from quiz_sesi_user_jawaban where 
				id_quiz = ? and id_user  = ?  and  left(kategori,9) = 'KOGNITIF_'
				group by  kategori ) as s 
			left join ref_skala_kd_15 as kd15 on s.soal = 15 and s.skor = kd15.skor 
			left join ref_skala_kd_20 as kd20 on s.soal = 20 and s.skor = kd20.skor`,
		id_quiz, id_user).Scan(&skors)
	var sum_skor = 0
	for _, s := range skors {
		if s.Kategori == "KOGNITIF_INFORMASI_UMUM" {
			skoring.TpaIu = int32(s.Skor)
			skoring.KlasifikasiIu = s.Klasifikasi
		}
		if s.Kategori == "KOGNITIF_PENALARAN_ABSTRAK" {
			skoring.TpaPa = int32(s.Skor)
			skoring.KlasifikasiPa = s.Klasifikasi
		}
		if s.Kategori == "KOGNITIF_PENGERTIAN_MEKANIK" {
			skoring.TpaPm = int32(s.Skor)
			skoring.KlasifikasiPm = s.Klasifikasi
		}
		if s.Kategori == "KOGNITIF_PENALARAN_SPASIAL" {
			skoring.TpaPs = int32(s.Skor)
			skoring.KlasifikasiPs = s.Klasifikasi
		}
		if s.Kategori == "KOGNITIF_PENALARAN_KUANTITATIF" {
			skoring.TpaPk = int32(s.Skor)
			skoring.KlasifikasiPk = s.Klasifikasi
		}
		if s.Kategori == "KOGNITIF_PENALARAN_VERBAL" {
			skoring.TpaPv = int32(s.Skor)
			skoring.KlasifikasiPv = s.Klasifikasi
		}
		if s.Kategori == "KOGNITIF_CEPAT_TELITI" {
			skoring.TpaKt = int32(s.Skor)
			skoring.KlasifikasiKt = s.Klasifikasi
		}
		sum_skor = sum_skor + s.Skor
	}
	//konversi IQ 90 referensi tabel : ref_konversi_iq_90
	var klasIQ *entity.RefKonversiIq90
	db.Table("ref_konversi_iq_90").Where("skor_x = ? ", sum_skor).First(&klasIQ)
	skoring.SkorIq = klasIQ.TotIq
	skoring.KlasifikasiIq = klasIQ.Klasifikasi
	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user
	skoring.TpaIq = int32(sum_skor)

	db.Table("skor_kognitif").Create(&skoring)
	return nil
}

func (*repo) SkoringKognitifPMK(id_quiz int32, id_user int32) error {
	//regenereate  delete current if exist
	var skoring = entity.SkorKognitif{}
	db.Table("skor_kognitif_pmk").Where("id_user = ?", id_user).Where("id_quiz = ? ", id_quiz).Delete(&skoring)

	//koreksi KOGNITIF_ (KOGNITIF KELAS 3 SMA) prefix= KOGNITIF_PMK_
	// a.paket = 'PMK' and   a.bidang::TEXT = replace(b.kategori, 'KOGNITIF_PMK_','')
	db.Exec(`update quiz_sesi_user_jawaban  as t set skor = 1 
				from ( 
				select b.id_quiz, 
				case when a.pilihan_jawaban = b.jawaban then 1 else 0 end as skor,
					b.id_jawaban_peserta 
				from  soal_kognitif as a , quiz_sesi_user_jawaban as b  
				where a.paket = 'PMK' and   a.bidang::TEXT = replace(b.kategori, 'KOGNITIF_PMK_','')
				and a.urutan = b.urutan  and b.id_quiz = ?  and left(b.kategori,13) = 'KOGNITIF_PMK_'
				and b.id_user = ?) as x 
				where x.skor = 1 
				and t.id_jawaban_peserta = x.id_jawaban_peserta`,
		id_quiz,
		id_user)

	//update skoring dan klasifikasi
	type groupSkoring struct {
		Kategori    string `json:"kategori"`
		Skor        int    `json:"skor"`
		Klasifikasi string `json:"klasifikasi"`
	}
	var skors []*groupSkoring
	db.Raw(`select s.kategori, s.skor, 
				replace(case when kd15.klasifikasi is not null then  kd15.klasifikasi::text 
					when kd20.klasifikasi is not null then  kd20.klasifikasi::text  
					when s.skor = 0 then 'SANGAT_RENDAH'
					end,'_',' ') as klasifikasi
			from  
				(select kategori , sum(skor) as skor, count(skor) as soal 
				from quiz_sesi_user_jawaban where 
				id_quiz = ? and id_user  = ?  and  left(kategori,13) = 'KOGNITIF_PMK_'
				group by  kategori ) as s 
			left join ref_skala_kd_15 as kd15 on s.soal = 15 and s.skor = kd15.skor 
			left join ref_skala_kd_20 as kd20 on s.soal = 20 and s.skor = kd20.skor`,
		id_quiz, id_user).Scan(&skors)
	var sum_skor = 0
	for _, s := range skors {
		if s.Kategori == "KOGNITIF_PMK_INFORMASI_UMUM" {
			skoring.TpaIu = int32(s.Skor)
			skoring.KlasifikasiIu = s.Klasifikasi
		}
		if s.Kategori == "KOGNITIF_PMK_PENALARAN_ABSTRAK" {
			skoring.TpaPa = int32(s.Skor)
			skoring.KlasifikasiPa = s.Klasifikasi
		}
		if s.Kategori == "KOGNITIF_PMK_PENGERTIAN_MEKANIK" {
			skoring.TpaPm = int32(s.Skor)
			skoring.KlasifikasiPm = s.Klasifikasi
		}
		if s.Kategori == "KOGNITIF_PMK_PENALARAN_SPASIAL" {
			skoring.TpaPs = int32(s.Skor)
			skoring.KlasifikasiPs = s.Klasifikasi
		}
		if s.Kategori == "KOGNITIF_PMK_PENALARAN_KUANTITATIF" {
			skoring.TpaPk = int32(s.Skor)
			skoring.KlasifikasiPk = s.Klasifikasi
		}
		if s.Kategori == "KOGNITIF_PMK_PENALARAN_VERBAL" {
			skoring.TpaPv = int32(s.Skor)
			skoring.KlasifikasiPv = s.Klasifikasi
		}
		if s.Kategori == "KOGNITIF_PMK_CEPAT_TELITI" {
			skoring.TpaKt = int32(s.Skor)
			skoring.KlasifikasiKt = s.Klasifikasi
		}
		sum_skor = sum_skor + s.Skor
	}
	//konversi IQ 90 referensi tabel : ref_konversi_iq_90
	var klasIQ *entity.RefKonversiIq105
	db.Table("ref_konversi_iq_105").Where("skor_x = ? ", sum_skor).First(&klasIQ)
	skoring.SkorIq = klasIQ.TotIq
	skoring.KlasifikasiIq = klasIQ.Klasifikasi
	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user
	skoring.TpaIq = int32(sum_skor)

	db.Table("skor_kognitif_pmk").Create(&skoring)
	return nil
}

func (*repo) SkoringGayaPekerjaan(id_quiz int32, id_user int32) error {
	var refKomponen []*entity.RefKomponenGayaPekerjaan
	db.Table("ref_komponen_gaya_pekerjaan").Scan(&refKomponen)
	var kodeKomponen = make(map[string]string, len(refKomponen)+1)
	for i := 0; i < len(refKomponen); i++ {
		nomor := refKomponen[i].No
		kodeKomponen[nomor] = refKomponen[i].Kode
	}

	var refSkoring entity.RefSkoringGayaPekerjaan
	db.Table("ref_skoring_gaya_pekerjaan").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).Delete(&refSkoring)
	var skoringGayaPekerjaaan entity.SkorGayaPekerjaan
	db.Table("skor_gaya_pekerjaan").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).Delete(skoringGayaPekerjaaan)

	var dataSkor []*entity.DataSkorGayaPekerjaan
	db.Raw(`select 
				c.jawaban, c.c, c.t, c.u, b.*
			from quiz_sesi_user_jawaban as a, 
				soal_gaya_pekerjaan as b , 
				ref_skor_gaya_pekerjaan as c 
			where a.urutan = b.nomor 
				and a.id_quiz = ? 
				and a.id_user = ?
				and a.kategori='SKALA_GAYA_PEKERJAAN'
				and c.jawaban = a.jawaban`, id_quiz, id_user).Scan(&dataSkor)

	var skorKomponen = map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
		"h": 0,
		"i": 0,
		"j": 0,
		"k": 0,
		"l": 0}

	for i := 0; i < len(dataSkor); i++ {
		skorU := dataSkor[i].U
		skorT := dataSkor[i].T
		skorC := dataSkor[i].C
		r := reflect.ValueOf(dataSkor[i]).Elem()
		rt := r.Type()
		rv := reflect.ValueOf(dataSkor[i])
		for k, v := range skorKomponen {
			var valueSkor = v
			for n := 0; n < rt.NumField(); n++ {
				field := rt.Field(n)
				if field.Name == "Komponen"+strings.ToUpper(k) {
					value := reflect.Indirect(rv).FieldByName(field.Name)
					if strings.Contains(value.String(), "U") {
						valueSkor += int(skorU)
					}
					if strings.Contains(value.String(), "T") {
						valueSkor += int(skorT)
					}
					if strings.Contains(value.String(), "C") {
						valueSkor += int(skorC)
					}
				}
			}
			skorKomponen[k] = valueSkor
		}
	}
	skorKomponenRangking := helper.SortingMapIntDesc(skorKomponen)
	rangking := 1
	batasRangking := 3
	var skoring entity.SkorGayaPekerjaan
	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user
	for i := 0; i < len(skorKomponenRangking); i++ {
		nomor := skorKomponenRangking[i].Key
		skor := skorKomponenRangking[i].Value
		if rangking <= batasRangking {
			if rangking == 1 {
				skoring.RangkingGp1 = kodeKomponen[nomor]
			}
			if rangking == 2 {
				skoring.RangkingGp2 = kodeKomponen[nomor]
			}
			if rangking == 3 {
				skoring.RangkingGp3 = kodeKomponen[nomor]
			}
		}
		var itemRefSkoring entity.RefSkoringGayaPekerjaan
		itemRefSkoring.IDQuiz = id_quiz
		itemRefSkoring.IDUser = id_user
		itemRefSkoring.Kode = kodeKomponen[nomor]
		itemRefSkoring.Skor = int16(skor)
		itemRefSkoring.Rangking = int16(rangking)
		db.Table("ref_skoring_gaya_pekerjaan").Create(&itemRefSkoring)
		rangking++
	}
	db.Table("skor_gaya_pekerjaan").Create(&skoring)

	var klasifikasi []*entity.RefSkoringGayaPekerjaan

	db.Raw(`select a.id, b.akronim as klasifikasi
			from ref_skoring_gaya_pekerjaan  as a, 
				ref_klasifikasi_gaya_kerja as b  
			where a.skor >= b.skor_min and a.skor <= b.skor_max  
				and a.id_quiz = ? and a.id_user = ?
			order by a.rangking`, id_quiz, id_user).Scan(&klasifikasi)

	for i := 0; i < len(klasifikasi); i++ {
		db.Table("ref_skoring_gaya_pekerjaan").
			Where("id = ? ", klasifikasi[i].ID).
			UpdateColumn("klasifikasi", klasifikasi[i].Klasifikasi)
	}
	return nil
}

func (*repo) SkoringSikapPelajaran(id_quiz int32, id_user int32) error {
	kategori := "SIKAP_TERHADAP_PELAJARAN"
	var tabelSkoring entity.SkorSikapPelajaran
	db.Table("skor_sikap_pelajaran").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).Delete(tabelSkoring)

	db.Exec(`UPDATE  quiz_sesi_user_jawaban 
				set skor = SUBSTR(jawaban,1,1)::INTEGER 
				+ SUBSTR(jawaban,2,1)::INTEGER 
				+ SUBSTR(jawaban,3,1)::INTEGER
			where id_quiz = ? 
				and kategori = ?`, id_quiz, kategori)

	//hitung per pelajaran
	var skorHitung []*entity.SkorHitungSikapPelajaran
	db.Raw(`SELECT 
				a.id_quiz, 
				a.id_user, 
				a.skor, 
				b.field_skoring, 
				c.klasifikasi
			FROM quiz_sesi_user_jawaban as a, 
				soal_sikap_pelajaran as b , 
				ref_skala_sikap_pelajaran as c
			where 
				a.urutan = b.urutan
				and a.skor = c.skor
				and a.id_quiz = ?
				and a.id_user = ?
				and a.kategori = ? `, id_quiz, id_user, kategori).Scan(&skorHitung)

	var skoring entity.SkorSikapPelajaran

	r := reflect.ValueOf(&skoring).Elem()
	rt := r.Type()

	for i := 0; i < len(skorHitung); i++ {
		for n := 0; n < rt.NumField(); n++ {
			if rt.Field(n).Tag.Get("json") == skorHitung[i].FieldSkoring {
				sikapName := rt.Field(n).Name
				pelajaran := strings.Replace(sikapName, "Sikap", "", -1)
				klasifikasiName := "Klasifikasi" + pelajaran
				reflect.ValueOf(&skoring).Elem().FieldByName(sikapName).SetInt(int64(skorHitung[i].Skor))
				reflect.ValueOf(&skoring).Elem().FieldByName(klasifikasiName).SetString(skorHitung[i].Klasifikasi)
			}
		}
	}

	//hitung kelompok
	var skorHitungKelompok []*entity.SkorHitungSikapPelajaran
	db.Raw(`SELECT 
				a.id_quiz, 
				a.id_user, 
				c.field_skoring,
				sum(a.skor) as skor
				FROM quiz_sesi_user_jawaban as a, 
				soal_sikap_pelajaran as b  , 
				ref_kelompok_sikap_pelajaran as c 
			where 
				a.urutan = b.urutan and a.id_quiz = ? and a.id_user = ?
				and c.kelompok::VARCHAR = b.kelompok::VARCHAR 
				and a.kategori = ?
			group by 
				id_quiz, id_user, c.field_skoring`, id_quiz, id_user, kategori).Scan(&skorHitungKelompok)

	for i := 0; i < len(skorHitungKelompok); i++ {
		for n := 0; n < rt.NumField(); n++ {
			if rt.Field(n).Tag.Get("json") == skorHitungKelompok[i].FieldSkoring {
				kelompokName := rt.Field(n).Name
				reflect.ValueOf(&skoring).Elem().FieldByName(kelompokName).SetInt(int64(skorHitungKelompok[i].Skor))
			}
		}
	}
	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user
	db.Table("skor_sikap_pelajaran").Create(&skoring)
	var rekomSkoring *entity.SkorSikapPelajaran

	//update recomendasi
	db.Raw(`select 
				a.id_user, 
				a.id_quiz,
				a.sikap_ilmu_alam -  a.sikap_ilmu_sosial as sikap_rentang,
				b.rekomendasi as rekom_sikap_pelajaran 
			from skor_sikap_pelajaran as a, ref_rekomendasi_sikap_pelajaran as b  
			where 
				a.id_quiz = ? and a.id_user = ?
				and b.perbedaan = (a.sikap_ilmu_alam -  a.sikap_ilmu_sosial)`, id_quiz, id_user).First(&rekomSkoring)

	db.Table("skor_sikap_pelajaran").
		Where("id_quiz = ?", id_quiz).
		Where("id_user = ?", id_user).
		UpdateColumns(map[string]interface{}{
			"sikap_rentang":         rekomSkoring.SikapRentang,
			"rekom_sikap_pelajaran": rekomSkoring.RekomSikapPelajaran})

	return nil
}

func (*repo) SkoringSikapPelajaranMK(id_quiz int32, id_user int32) error {
	kategori := "SKALA_PMK_SIKAP_PELAJARAN"
	var tabelSkoring entity.SkorSikapPelajaranMk
	db.Table("skor_sikap_pelajaran_mk").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).Delete(tabelSkoring)

	db.Exec(`UPDATE  quiz_sesi_user_jawaban 
				set skor = SUBSTR(jawaban,1,1)::INTEGER 
				+ SUBSTR(jawaban,2,1)::INTEGER 
				+ SUBSTR(jawaban,3,1)::INTEGER
			where id_quiz = ? 
				and kategori = ?`, id_quiz, kategori)

	//hitung per pelajaran
	var skorHitung []*entity.SkorHitungSikapPelajaran
	db.Raw(`SELECT 
				a.id_quiz, 
				a.id_user, 
				a.skor, 
				b.field_skoring, 
				c.klasifikasi
			FROM quiz_sesi_user_jawaban as a, 
				soal_sikap_pelajaran_kuliah as b , 
				ref_skala_sikap_pelajaran as c
			where 
				a.urutan = b.urutan
				and a.skor = c.skor
				and a.id_quiz = ?
				and a.id_user = ?
				and a.kategori = ? `, id_quiz, id_user, kategori).Scan(&skorHitung)

	var skoring entity.SkorSikapPelajaranMk

	r := reflect.ValueOf(&skoring).Elem()
	rt := r.Type()

	for i := 0; i < len(skorHitung); i++ {
		for n := 0; n < rt.NumField(); n++ {
			if rt.Field(n).Tag.Get("json") == skorHitung[i].FieldSkoring {
				sikapName := rt.Field(n).Name
				pelajaran := strings.Replace(sikapName, "Sikap", "", -1)
				klasifikasiName := "Klasifikasi" + pelajaran
				reflect.ValueOf(&skoring).Elem().FieldByName(sikapName).SetInt(int64(skorHitung[i].Skor))
				reflect.ValueOf(&skoring).Elem().FieldByName(klasifikasiName).SetString(skorHitung[i].Klasifikasi)
			}
		}
	}

	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user
	db.Table("skor_sikap_pelajaran_mk").Create(&skoring)
	return nil
}

func (*repo) SkoringPeminatanSMK(id_quiz int32, id_user int32) error {
	tabel := "skor_peminatan_smk"
	kategori := "SKALA_PEMINATAN_SMK"
	var tabelSkoring entity.SkorPeminatanSmk

	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).Delete(tabelSkoring)

	var skor_maksimal int64
	db.Table("quiz_sesi_mapping_smk").Where("id_quiz = ?", id_quiz).Count(&skor_maksimal)
	db.Exec(`update quiz_sesi_user_jawaban 
				set skor = ? - urutan
			where 
				id_quiz = ? and id_user = ? and kategori = ?`, skor_maksimal+1, id_quiz, id_user, kategori)
	var quizSesiUserJawaban []*entity.QuizSesiUserJawaban
	db.Table("quiz_sesi_user_jawaban").
		Where("id_quiz = ?", id_quiz).
		Where("id_user = ?", id_user).
		Where("kategori = ?", kategori).
		Order("urutan asc").Scan(&quizSesiUserJawaban)

	var skoring entity.SkorPeminatanSmk
	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user
	r := reflect.ValueOf(&skoring).Elem()
	rt := r.Type()

	for n := 0; n < rt.NumField(); n++ {
		if strings.Contains(rt.Field(n).Tag.Get("json"), "minat_") {
			reflect.ValueOf(&skoring).Elem().FieldByName(rt.Field(n).Name).SetString("")
		}
	}

	for i := 0; i < len(quizSesiUserJawaban); i++ {
		urutan := quizSesiUserJawaban[i].Urutan
		jawaban := quizSesiUserJawaban[i].Jawaban
		fieldname := fmt.Sprintf("minat_%v", urutan)
		for n := 0; n < rt.NumField(); n++ {
			if rt.Field(n).Tag.Get("json") == fieldname {
				reflect.ValueOf(&skoring).Elem().FieldByName(rt.Field(n).Name).SetString(jawaban)
			}
		}
	}
	db.Table(tabel).Create(&skoring)
	return nil
}

func (*repo) SkoringPeminatanSMA(id_quiz int32, id_user int32) error {
	tabel := "skor_peminatan_sma"
	kategori := "SKALA_PEMINATAN_SMA"
	var tabelSkoring entity.SkorPeminatanSma
	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).Delete(tabelSkoring)

	var skorHitung []*entity.SkorHitungFieldKlasifikasi
	db.Raw(`select 
				b.id_quiz, 
				b.id_user, 
				a.field_skoring,
				count(b.jawaban) as skor
				from 
				ref_pilihan_minat_sma as a 
				left join quiz_sesi_user_jawaban as b 
				on a.kd_pilihan = b.jawaban
				and b.id_quiz = ? and b.id_user = ? and b.kategori = ?
			GROUP BY 
				b.id_quiz, 
				b.id_user, 
				a.field_skoring`, id_quiz, id_user, kategori).Scan(&skorHitung)

	var skoring entity.SkorPeminatanSma
	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user
	r := reflect.ValueOf(&skoring).Elem()
	rt := r.Type()

	for i := 0; i < len(skorHitung); i++ {
		skor := skorHitung[i].Skor
		fied_skoring := skorHitung[i].FieldSkoring
		for n := 0; n < rt.NumField(); n++ {
			if rt.Field(n).Tag.Get("json") == fied_skoring {
				reflect.ValueOf(&skoring).Elem().FieldByName(rt.Field(n).Name).SetInt(int64(skor))
			}
		}
	}

	db.Table(tabel).Create(&skoring)

	var skorRekomendasi *entity.SkorRekomendasi
	db.Raw(`SELECT
			x.minat_rentang as skor,
			y.rekomendasi as rekomendasi
			FROM
			(
				SELECT
				id_user,
				id_quiz,
				minat_sains,
				minat_bahasa,
				minat_humaniora,
				COALESCE ( minat_sains,0)  - 
				COALESCE(minat_bahasa,0) - COALESCE(minat_humaniora, 0) 
				AS minat_rentang
				FROM
				skor_peminatan_sma AS a
				WHERE
				a.id_quiz = ?
				and a.id_user = ?
			) AS x
			LEFT JOIN ref_rekomendasi_minat_sma AS y ON x.minat_rentang = y.perbedaan`, id_quiz, id_user).First(&skorRekomendasi)

	db.Table(tabel).
		Where("id_quiz = ?", id_quiz).
		Where("id_user = ?", id_user).
		UpdateColumns(map[string]interface{}{
			"minat_rentang": skorRekomendasi.Skor,
			"rekom_minat":   skorRekomendasi.Rekomendasi})

	return nil
}

func (*repo) SkoringPeminatanMAN(id_quiz int32, id_user int32) error {
	tabel := "skor_peminatan_man"
	kategori := "SKALA_PEMINATAN_MAN"
	var tabelSkoring entity.SkorPeminatanMan
	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).Delete(tabelSkoring)

	var skorHitung []*entity.SkorHitungFieldKlasifikasi
	db.Raw(`select 
				b.id_quiz, 
				b.id_user, 
				a.field_skoring,
				count(b.jawaban) as skor
				from 
				ref_pilihan_minat_man as a 
				left join quiz_sesi_user_jawaban as b 
				on a.kd_pilihan = b.jawaban
				and b.id_quiz = ? and b.id_user = ? and b.kategori = ?
			GROUP BY 
				b.id_quiz, 
				b.id_user, 
				a.field_skoring`, id_quiz, id_user, kategori).Scan(&skorHitung)

	var skoring entity.SkorPeminatanMan
	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user
	r := reflect.ValueOf(&skoring).Elem()
	rt := r.Type()

	for i := 0; i < len(skorHitung); i++ {
		skor := skorHitung[i].Skor
		fied_skoring := skorHitung[i].FieldSkoring
		for n := 0; n < rt.NumField(); n++ {
			if rt.Field(n).Tag.Get("json") == fied_skoring {
				reflect.ValueOf(&skoring).Elem().FieldByName(rt.Field(n).Name).SetInt(int64(skor))
			}
		}
	}

	db.Table(tabel).Create(&skoring)

	var skorRekomendasi *entity.SkorRekomendasi
	db.Raw(`SELECT
			x.minat_rentang as skor,
			y.rekomendasi as rekomendasi
			FROM
			(
				SELECT
				id_user,
				id_quiz,
				minat_sains,
				minat_bahasa,
				minat_humaniora,
				COALESCE ( minat_sains,0)  - 
				COALESCE(minat_bahasa,0) - COALESCE(minat_humaniora, 0)  - COALESCE(minat_agama, 0) 
				AS minat_rentang
				FROM
				skor_peminatan_man AS a
				WHERE
				a.id_quiz = ?
				and a.id_user = ?
			) AS x
			LEFT JOIN ref_rekomendasi_minat_man AS y ON x.minat_rentang = y.perbedaan`, id_quiz, id_user).First(&skorRekomendasi)

	db.Table(tabel).
		Where("id_quiz = ?", id_quiz).
		Where("id_user = ?", id_user).
		UpdateColumns(map[string]interface{}{
			"minat_rentang": skorRekomendasi.Skor,
			"rekom_minat":   skorRekomendasi.Rekomendasi})

	return nil
}

func (*repo) SkoringMinatKuliahKedinasan(id_quiz int32, id_user int32) error {
	return nil
}
