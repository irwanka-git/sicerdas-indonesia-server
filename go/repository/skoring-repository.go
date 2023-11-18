package repository

import (
	"encoding/json"
	"fmt"
	"irwanka/sicerdas/entity"
	"irwanka/sicerdas/helper"
	"reflect"
	"strconv"
	"strings"
)

type SkoringRepository interface {
	GetStatusRunningSkoring() (*entity.StatusRunningSkoring, error)
	StartRunningSkoring(mulai string) error
	StopRunningSkoring(selesai string, user int) error
	GetUserSesiBelumSkoring() ([]*entity.QuizSesiUserSkoring, error)
	ClearTabelTemporaryJawabanUser(id_quiz int32, id_user int32) error
	GenerateTabelTemporaryJawabanUser(id_quiz int32, id_user int32) error
	GetKategoriTabelSkoring(id_quiz int32) ([]*entity.KategoriTabel, error)
	GetTabelSkoring(id_quiz int32) ([]*entity.KategoriTabel, error)

	GetListSkoringGabungan() ([]*entity.QuizSesiReport, error)

	//SKORING SESI
	SkoringKognitif(id_quiz int32, id_user int32) error
	SkoringKognitifPMK(id_quiz int32, id_user int32) error
	SkoringGayaPekerjaan(id_quiz int32, id_user int32) error
	SkoringSikapPelajaran(id_quiz int32, id_user int32) error
	SkoringSikapPelajaranMK(id_quiz int32, id_user int32) error
	SkoringPeminatanSMK(id_quiz int32, id_user int32) error
	SkoringPeminatanSMA(id_quiz int32, id_user int32) error
	SkoringPeminatanMAN(id_quiz int32, id_user int32) error
	SkoringMinatKuliahKedinasan(id_quiz int32, id_user int32) error
	SkoringKuliahAlam(id_quiz int32, id_user int32) error
	SkoringKuliahSosial(id_quiz int32, id_user int32) error
	SkoringKuliahAgama(id_quiz int32, id_user int32) error
	SkoringMBTI(id_quiz int32, id_user int32) error
	SkoringKarakteristikPribadi(id_quiz int32, id_user int32) error
	SkoringTesMinatIndonesia(id_quiz int32, id_user int32) error
	SkoringKecerdasanMajemuk(id_quiz int32, id_user int32) error
	SkoringSuasanaKerja(id_quiz int32, id_user int32) error
	SkoringGayaBelajar(id_quiz int32, id_user int32) error
	SkoringKejiwaanDewasa(id_quiz int32, id_user int32) error
	SkoringKesehatanMental(id_quiz int32, id_user int32) error
	SkoringModeBelajar(id_quiz int32, id_user int32) error
	SkoringSSCT(id_quiz int32, id_user int32) error
	SkoringDISC(id_quiz int32, id_user int32) error
	SkoringModeKerja(id_quiz int32, id_user int32) error
	SkoringKepribadianManajerial(id_quiz int32, id_user int32) error
	SkoringWLB(id_quiz int32, id_user int32) error

	//SKORING GABUNGAN
	SkoringRekomKuliahA(id_quiz int32, id_user int32) error
	SkoringRekomKuliahB(id_quiz int32, id_user int32) error
	SkoringRekomPeminatanSMA(id_quiz int32, id_user int32) error

	FinishSkoring(id_quiz int32, id_user int32, waktu string) error
}

func NewSkoringRepository() SkoringRepository {
	return &repo{}
}

func (*repo) GetListSkoringGabungan() ([]*entity.QuizSesiReport, error) {
	var result []*entity.QuizSesiReport
	db.Raw(`select * from quiz_sesi_report where jenis = 1 and tabel_referensi != tabel_terkait`).Scan(&result)
	return result, nil
}

func (*repo) GetStatusRunningSkoring() (*entity.StatusRunningSkoring, error) {
	var cek *entity.StatusRunningSkoring
	db.Table("status_skoring").First(&cek)
	return cek, nil
}
func (*repo) StartRunningSkoring(mulai string) error {
	db.Table("status_skoring").Where("id = ?", 1).UpdateColumn("status", 1)
	db.Table("status_skoring").Where("id = ?", 1).UpdateColumn("mulai", mulai)
	return nil
}

func (*repo) StopRunningSkoring(selesai string, jumlah int) error {
	db.Table("status_skoring").Where("id = ?", 1).UpdateColumn("status", 0)
	db.Table("status_skoring").Where("id = ?", 1).UpdateColumn("selesai", selesai)
	db.Table("status_skoring").Where("id = ?", 1).UpdateColumn("jumlah", jumlah)
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
		and ( c.skoring_tabel = '' or c.skoring_tabel is null)
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
			db.Exec("delete from quiz_sesi_user_jawaban where id_quiz = ? and id_user = ? and kategori = ? and urutan = ? ", id_quiz, id_user, kategori, urutan)
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

	var refKlasifikasi []*entity.RefKlasifikasiPeminatanSMA
	db.Table("ref_klasifikasi_minat_sma").Scan(&refKlasifikasi)
	var klasifikasi_minat_sains = ""
	var klasifikasi_minat_bahasa = ""
	var klasifikasi_minat_humaniora = ""
	for i := 0; i < len(refKlasifikasi); i++ {
		if refKlasifikasi[i].Skor == skoring.MinatSains {
			klasifikasi_minat_sains = refKlasifikasi[i].Klasifikasi
		}
		if refKlasifikasi[i].Skor == skoring.MinatHumaniora {
			klasifikasi_minat_humaniora = refKlasifikasi[i].Klasifikasi
		}
		if refKlasifikasi[i].Skor == skoring.MinatBahasa {
			klasifikasi_minat_bahasa = refKlasifikasi[i].Klasifikasi
		}
	}

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
			"minat_rentang":               skorRekomendasi.Skor,
			"klasifikasi_minat_sains":     klasifikasi_minat_sains,
			"klasifikasi_minat_humaniora": klasifikasi_minat_humaniora,
			"klasifikasi_minat_bahasa":    klasifikasi_minat_bahasa,
			"rekom_minat":                 skorRekomendasi.Rekomendasi})

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

	var refKlasifikasi []*entity.RefKlasifikasiPeminatanMAN
	db.Table("ref_klasifikasi_minat_man").Scan(&refKlasifikasi)
	var klasifikasi_minat_sains = ""
	var klasifikasi_minat_bahasa = ""
	var klasifikasi_minat_humaniora = ""
	var klasifikasi_minat_agama = ""

	for i := 0; i < len(refKlasifikasi); i++ {
		if refKlasifikasi[i].Skor == skoring.MinatSains {
			klasifikasi_minat_sains = refKlasifikasi[i].Klasifikasi
		}
		if refKlasifikasi[i].Skor == skoring.MinatHumaniora {
			klasifikasi_minat_humaniora = refKlasifikasi[i].Klasifikasi
		}
		if refKlasifikasi[i].Skor == skoring.MinatBahasa {
			klasifikasi_minat_bahasa = refKlasifikasi[i].Klasifikasi
		}
		if refKlasifikasi[i].Skor == skoring.MinatAgama {
			klasifikasi_minat_agama = refKlasifikasi[i].Klasifikasi
		}
	}

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
			"minat_rentang":               skorRekomendasi.Skor,
			"klasifikasi_minat_sains":     klasifikasi_minat_sains,
			"klasifikasi_minat_humaniora": klasifikasi_minat_humaniora,
			"klasifikasi_minat_bahasa":    klasifikasi_minat_bahasa,
			"klasifikasi_minat_agama":     klasifikasi_minat_agama,
			"rekom_minat":                 skorRekomendasi.Rekomendasi})

	return nil
}

func (*repo) SkoringMinatKuliahKedinasan(id_quiz int32, id_user int32) error {
	kategori := "SKALA_PMK_SEKOLAH_KEDINASAN"
	tabel := "skor_kuliah_dinas"
	var tabelSkoring entity.SkorKuliahDinas
	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Delete(tabelSkoring)
	var tabelRefSkoring entity.SkorKuliahDinas
	db.Table("ref_skoring_kuliah_dinas").Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Delete(tabelRefSkoring)

	var refSekolahDinas []*entity.RefSekolahDinas
	db.Table("ref_sekolah_dinas").Order("no asc").Scan(&refSekolahDinas)

	var quizSesiUserJawaban []*entity.QuizSesiUserJawaban
	db.Table("quiz_sesi_user_jawaban").
		Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).
		Where("kategori =? ", kategori).
		Order("urutan asc").Scan(&quizSesiUserJawaban)
	var jawabanUser = []string{}
	for i := 0; i < len(quizSesiUserJawaban); i++ {
		jawabanUser = append(jawabanUser, quizSesiUserJawaban[i].Jawaban)
	}

	for i := 0; i < len(refSekolahDinas); i++ {
		var refSkoring entity.RefSkoringKuliahDinas
		refSkoring.B1 = int32(strings.Index(jawabanUser[0], refSekolahDinas[i].No)) + 1
		refSkoring.B2 = int32(strings.Index(jawabanUser[1], refSekolahDinas[i].No)) + 1
		refSkoring.B3 = int32(strings.Index(jawabanUser[2], refSekolahDinas[i].No)) + 1
		refSkoring.B4 = int32(strings.Index(jawabanUser[3], refSekolahDinas[i].No)) + 1
		refSkoring.B5 = int32(strings.Index(jawabanUser[4], refSekolahDinas[i].No)) + 1
		refSkoring.B6 = int32(strings.Index(jawabanUser[5], refSekolahDinas[i].No)) + 1
		refSkoring.B7 = int32(strings.Index(jawabanUser[6], refSekolahDinas[i].No)) + 1
		refSkoring.B8 = int32(strings.Index(jawabanUser[7], refSekolahDinas[i].No)) + 1
		refSkoring.B9 = int32(strings.Index(jawabanUser[8], refSekolahDinas[i].No)) + 1
		refSkoring.IDQuiz = id_quiz
		refSkoring.IDUser = id_user
		refSkoring.No = refSekolahDinas[i].No
		refSkoring.Total = refSkoring.B1 + refSkoring.B2 + refSkoring.B3 + refSkoring.B4 + refSkoring.B5 + refSkoring.B6 + refSkoring.B7 + refSkoring.B8 + refSkoring.B9
		refSkoring.Rangking = 0
		db.Table("ref_skoring_kuliah_dinas").Create(&refSkoring)
	}

	var currentRangkingRefSkoring []*entity.RefSkoringKuliahDinas
	db.Table("ref_skoring_kuliah_dinas").Where("id_quiz = ?", id_quiz).
		Where("id_user = ? ", id_user).
		Order("total asc").
		Order("no asc").
		Scan(&currentRangkingRefSkoring)
	var rangking = 1
	var minat_dinas1 = ""
	var minat_dinas2 = ""
	var minat_dinas3 = ""
	for i := 0; i < len(currentRangkingRefSkoring); i++ {
		db.Table("ref_skoring_kuliah_dinas").
			Where("id_skoring_sekolah_dinas = ? ", currentRangkingRefSkoring[i].IDSkoringSekolahDinas).
			UpdateColumn("rangking", rangking)
		if rangking == 1 {
			minat_dinas1 = currentRangkingRefSkoring[i].No
		}
		if rangking == 2 {
			minat_dinas2 = currentRangkingRefSkoring[i].No
		}
		if rangking == 3 {
			minat_dinas3 = currentRangkingRefSkoring[i].No
		}
		rangking++
	}

	var skoring entity.SkorKuliahDinas
	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user
	skoring.MinatDinas1 = minat_dinas1
	skoring.MinatDinas2 = minat_dinas2
	skoring.MinatDinas3 = minat_dinas3
	db.Table(tabel).Create(&skoring)

	return nil
}

func (*repo) SkoringKuliahAlam(id_quiz int32, id_user int32) error {
	kategori := "SKALA_PMK_MINAT_ILMU_ALAM"
	tabel := "skor_kuliah_alam"
	var skoring entity.SkorKuliahAlam
	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Delete(skoring)

	var quizSesiUserJawaban []*entity.QuizSesiUserJawaban
	db.Table("quiz_sesi_user_jawaban").
		Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).
		Where("kategori =? ", kategori).
		Order("urutan asc").Scan(&quizSesiUserJawaban)

	var minat_ipa1 = 0
	var minat_ipa2 = 0
	var minat_ipa3 = 0
	var minat_ipa4 = 0
	var minat_ipa5 = 0

	for i := 0; i < len(quizSesiUserJawaban); i++ {
		jawaban, _ := strconv.Atoi(quizSesiUserJawaban[i].Jawaban)
		if quizSesiUserJawaban[i].Urutan == 1 {
			minat_ipa1 = jawaban
		}
		if quizSesiUserJawaban[i].Urutan == 2 {
			minat_ipa2 = jawaban
		}
		if quizSesiUserJawaban[i].Urutan == 3 {
			minat_ipa3 = jawaban
		}
		if quizSesiUserJawaban[i].Urutan == 4 {
			minat_ipa4 = jawaban
		}
		if quizSesiUserJawaban[i].Urutan == 5 {
			minat_ipa5 = jawaban
		}
	}
	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user
	skoring.MinatIpa1 = int32(minat_ipa1)
	skoring.MinatIpa2 = int32(minat_ipa2)
	skoring.MinatIpa3 = int32(minat_ipa3)
	skoring.MinatIpa4 = int32(minat_ipa4)
	skoring.MinatIpa5 = int32(minat_ipa5)

	db.Table(tabel).Create(&skoring)

	return nil
}

func (*repo) SkoringKuliahSosial(id_quiz int32, id_user int32) error {
	kategori := "SKALA_PMK_MINAT_ILMU_SOSIAL"
	tabel := "skor_kuliah_sosial"
	var skoring entity.SkorKuliahSosial
	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Delete(skoring)

	var quizSesiUserJawaban []*entity.QuizSesiUserJawaban
	db.Table("quiz_sesi_user_jawaban").
		Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).
		Where("kategori =? ", kategori).
		Order("urutan asc").Scan(&quizSesiUserJawaban)

	var minat_ips1 = 0
	var minat_ips2 = 0
	var minat_ips3 = 0
	var minat_ips4 = 0
	var minat_ips5 = 0

	for i := 0; i < len(quizSesiUserJawaban); i++ {
		jawaban, _ := strconv.Atoi(quizSesiUserJawaban[i].Jawaban)
		if quizSesiUserJawaban[i].Urutan == 1 {
			minat_ips1 = jawaban
		}
		if quizSesiUserJawaban[i].Urutan == 2 {
			minat_ips2 = jawaban
		}
		if quizSesiUserJawaban[i].Urutan == 3 {
			minat_ips3 = jawaban
		}
		if quizSesiUserJawaban[i].Urutan == 4 {
			minat_ips4 = jawaban
		}
		if quizSesiUserJawaban[i].Urutan == 5 {
			minat_ips5 = jawaban
		}
	}
	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user
	skoring.MinatIps1 = int32(minat_ips1)
	skoring.MinatIps2 = int32(minat_ips2)
	skoring.MinatIps3 = int32(minat_ips3)
	skoring.MinatIps4 = int32(minat_ips4)
	skoring.MinatIps5 = int32(minat_ips5)

	db.Table(tabel).Create(&skoring)

	return nil
}

func (*repo) SkoringKuliahAgama(id_quiz int32, id_user int32) error {
	kategori := "SKALA_PMK_ILMU_AGAMA"
	tabel := "skor_kuliah_agama"
	var skoring entity.SkorKuliahAgama
	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Delete(skoring)

	var quizSesiUserJawaban []*entity.QuizSesiUserJawaban
	db.Table("quiz_sesi_user_jawaban").
		Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).
		Where("kategori =? ", kategori).
		Order("urutan asc").Scan(&quizSesiUserJawaban)

	var minat_agm1 = 0
	var minat_agm2 = 0
	var minat_agm3 = 0
	var minat_agm4 = 0
	var minat_agm5 = 0

	for i := 0; i < len(quizSesiUserJawaban); i++ {
		jawaban, _ := strconv.Atoi(quizSesiUserJawaban[i].Jawaban)
		if quizSesiUserJawaban[i].Urutan == 1 {
			minat_agm1 = jawaban
		}
		if quizSesiUserJawaban[i].Urutan == 2 {
			minat_agm2 = jawaban
		}
		if quizSesiUserJawaban[i].Urutan == 3 {
			minat_agm3 = jawaban
		}
		if quizSesiUserJawaban[i].Urutan == 4 {
			minat_agm4 = jawaban
		}
		if quizSesiUserJawaban[i].Urutan == 5 {
			minat_agm5 = jawaban
		}
	}
	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user
	skoring.MinatAgm1 = int32(minat_agm1)
	skoring.MinatAgm2 = int32(minat_agm2)
	skoring.MinatAgm3 = int32(minat_agm3)
	skoring.MinatAgm4 = int32(minat_agm4)
	skoring.MinatAgm5 = int32(minat_agm5)

	db.Table(tabel).Create(&skoring)

	return nil
}

func (*repo) SkoringMBTI(id_quiz int32, id_user int32) error {
	kategori := "SKALA_TES_TIPOLOGI_JUNG"
	tabel := "skor_mbti"
	var skoring entity.SkorMbti
	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Delete(skoring)

	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user

	var skorHitungMBTI []*entity.SkorHitungMBTI
	db.Raw(`select x.id_user, 
				x.id_quiz, 
				y.kolom, 
				z.kode, 
				x.skor_a as skor_a, 
				x.skor_b as skor_b, 
				z.field_skoring_a, 
				z.field_skoring_b
				from (SELECT
						a.kolom,
						b.id_quiz, 
						b.id_user,
						sum(case when b.jawaban = 'A' then 1 else 0 end) as skor_a,
						sum(case when b.jawaban = 'B' then 1 else 0 end) as skor_b
					FROM
						soal_tipologi_jung AS a,
						quiz_sesi_user_jawaban AS b 
					WHERE
						a.urutan = b.urutan 
						AND b.id_quiz =  ?
						AND b.id_user =  ?
						AND b.kategori = ? 
					GROUP BY a.kolom, b.id_quiz, b.id_user) as x, 
					ref_skoring_tipologi_jung as y,
					ref_klasifikasi_tipologi_jung as z 
				where 
					x.kolom::VARCHAR = y.kolom::VARCHAR
					and x.skor_a = y.skor_a
					and y.kode_klasifikasi = z.kode
				order by 
					x.id_user, y.kolom`, id_quiz, id_user, kategori).Scan(&skorHitungMBTI)

	r := reflect.ValueOf(&skoring).Elem()
	rt := r.Type()
	var tipejung_kode = ""
	for i := 0; i < len(skorHitungMBTI); i++ {
		skorA := skorHitungMBTI[i].SkorA
		skorB := skorHitungMBTI[i].SkorB
		fied_skoringA := skorHitungMBTI[i].FieldSkoringA
		fied_skoringB := skorHitungMBTI[i].FieldSkoringB
		for n := 0; n < rt.NumField(); n++ {
			fieldnameJson := rt.Field(n).Tag.Get("json")
			fieldname := rt.Field(n).Name
			// fmt.Println(fieldname + ":" + fieldnameJson)
			if fieldnameJson == fied_skoringA {
				// fmt.Printf(" set fieldnameJson : %v ", skorA)
				reflect.ValueOf(&skoring).Elem().FieldByName(fieldname).SetInt(int64(skorA))
			}
			if fieldnameJson == fied_skoringB {
				// fmt.Printf(" set fieldnameJson : %v ", skorB)
				reflect.ValueOf(&skoring).Elem().FieldByName(fieldname).SetInt(int64(skorB))
			}
			// fmt.Println()
		}
		tipejung_kode = fmt.Sprintf("%v%v", tipejung_kode, skorHitungMBTI[i].Kode)
	}
	skoring.TipojungKode = tipejung_kode

	db.Table(tabel).Create(&skoring)

	return nil
}

func (*repo) SkoringKarakteristikPribadi(id_quiz int32, id_user int32) error {
	kategori := "SKALA_TES_KARAKTERISTIK_PRIBADI"
	tabel := "skor_karakteristik_pribadi"
	var skoring entity.SkorKarakteristikPribadi
	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Delete(skoring)
	//update skor
	db.Exec(`update quiz_sesi_user_jawaban
			set skor = get_skor_jawaban_karakteristik_pribadi(jawaban)
			where id_quiz = ? and id_user = ? and kategori = ?`, id_quiz, id_user, kategori)
	var skorHitung []*entity.SkorHitungFieldKlasifikasi
	db.Raw(`select x.id_user, x.id_quiz, x.total as skor , x.field_skoring, y.klasifikasi from (SELECT 
				a.id_quiz, a.id_user,
				c.field_skoring,
				sum(a.skor) as total
			from quiz_sesi_user_jawaban as a  , 
				soal_karakteristik_pribadi as b , 
				ref_komponen_karakteristik_pribadi as c
			where 
				a.urutan = b.urutan 
				and c.id_komponen = b.id_komponen
				and  a.id_quiz = ? 
				and a.id_user = ?
				and a.kategori = ?
			group by 
				a.id_user, a.id_quiz, c.field_skoring) as x, ref_karakter_pribadi as y 
				where x.total = y.skor`, id_quiz, id_user, kategori).Scan(&skorHitung)

	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user

	r := reflect.ValueOf(&skoring).Elem()
	rt := r.Type()
	for i := 0; i < len(skorHitung); i++ {
		skor := skorHitung[i].Skor
		klasifikasi := skorHitung[i].Klasifikasi
		for n := 0; n < rt.NumField(); n++ {
			fieldnameJson := rt.Field(n).Tag.Get("json")
			fieldname := rt.Field(n).Name
			if fieldnameJson == skorHitung[i].FieldSkoring {
				komponen := helper.Capitalize(strings.ReplaceAll(fieldnameJson, "pribadi_", ""))
				reflect.ValueOf(&skoring).Elem().FieldByName(fieldname).SetInt(int64(skor))
				klasifikasiName := fmt.Sprintf("Klasifikasi%v", komponen)
				reflect.ValueOf(&skoring).Elem().FieldByName(klasifikasiName).SetString(klasifikasi)
			}
		}
	}
	db.Table(tabel).Create(&skoring)

	return nil
}

func (*repo) SkoringTesMinatIndonesia(id_quiz int32, id_user int32) error {
	kategori := "SKALA_TES_MINAT_INDONESIA"
	tabel := "skor_minat_indonesia"
	skor_maksimal := 7

	var skoring entity.SkorMinatIndonesium
	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Delete(skoring)

	//update skor
	db.Exec(`update quiz_sesi_user_jawaban
			set skor = ? - urutan
			where id_quiz = ? and id_user = ? and kategori = ?`, skor_maksimal+1, id_quiz, id_user, kategori)

	var skorHitung []*entity.SkorHitungFieldKlasifikasi

	db.Raw(`select
				a.id_quiz, 
				a.id_user,  
				c.field_skoring,
				sum(a.skor) as skor 
			from quiz_sesi_user_jawaban as a, soal_tmi as b, ref_kelompok_tmi as c 
			where  
				a.id_quiz = ? 
				and a.id_user = ?  
				and a.kategori = ?
				and a.jawaban = lpad(b.urutan::TEXT ,2,'0'::TEXT)
				and c.kelompok::VARCHAR = b.kelompok::VARCHAR 
			GROUP BY 
				a.id_quiz, 
				a.id_user,  
				c.field_skoring`, id_quiz, id_user, kategori).Scan(&skorHitung)

	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user

	r := reflect.ValueOf(&skoring).Elem()
	rt := r.Type()
	for i := 0; i < len(skorHitung); i++ {
		skor := skorHitung[i].Skor
		for n := 0; n < rt.NumField(); n++ {
			if rt.Field(n).Tag.Get("json") == skorHitung[i].FieldSkoring {
				reflect.ValueOf(&skoring).Elem().FieldByName(rt.Field(n).Name).SetInt(int64(skor))
			}
		}
	}
	var urutanMinat []*entity.QuizSesiUserJawaban
	db.Table("quiz_sesi_user_jawaban").
		Where("id_quiz = ?", id_quiz).
		Where("id_user = ?", id_user).Where("kategori = ?", kategori).Order("skor desc").
		Scan(&urutanMinat)
	var urutan = 1
	for i := 0; i < len(urutanMinat); i++ {
		jawaban, _ := strconv.Atoi(urutanMinat[i].Jawaban)
		fieldMinat := fmt.Sprintf("Minat%v", urutan)
		reflect.ValueOf(&skoring).Elem().FieldByName(fieldMinat).SetInt(int64(jawaban))
		urutan++
	}

	db.Table(tabel).Create(&skoring)

	var skorRekomendasi []*entity.SkorRekomendasi
	db.Raw(`select 
				a.tmi_ilmu_alam -  a.tmi_ilmu_sosial as skor,
				b.rekomendasi 
			from skor_minat_indonesia as a,
				ref_rekomendasi_minat_tmi as b  
			where a.id_quiz = ? and a.id_user = ? 
				and b.perbedaan = (a.tmi_ilmu_alam -  a.tmi_ilmu_sosial)`, id_quiz, id_user).Scan(&skorRekomendasi)

	for i := 0; i < len(skorRekomendasi); i++ {
		skor := skorRekomendasi[i].Skor
		rekomendasi := skorRekomendasi[i].Rekomendasi
		db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).
			UpdateColumns(map[string]interface{}{
				"tmi_rentang": skor,
				"rekom_tmi":   rekomendasi})
	}

	return nil
}

func (*repo) SkoringKecerdasanMajemuk(id_quiz int32, id_user int32) error {
	kategori := "SKALA_KECERDASAN_MAJEMUK"
	tabel := "skor_kecerdasan_majemuk"
	var skoring entity.SkorKecerdasanMajemuk
	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Delete(skoring)

	var refKecerdasanMajemuk []*entity.RefKecerdasanMajemuk
	db.Table("ref_kecerdasan_majemuk").Order("no asc").Scan(&refKecerdasanMajemuk)

	var quizSesiUserJawaban []*entity.QuizSesiUserJawaban
	db.Table("quiz_sesi_user_jawaban").Where("id_quiz = ?", id_quiz).
		Where("id_user = ? ", id_user).
		Where("kategori = ?", kategori).
		Order("urutan asc").
		Scan(&quizSesiUserJawaban)
	var jawabanUser = []string{}
	for i := 0; i < len(quizSesiUserJawaban); i++ {
		jawabanUser = append(jawabanUser, quizSesiUserJawaban[i].Jawaban)
	}

	var clearSkoringKecerdasanMajemuk entity.RefSkoringKecerdasanMajemuk
	db.Table("ref_skoring_kecerdasan_majemuk").Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Delete(clearSkoringKecerdasanMajemuk)

	for i := 0; i < len(refKecerdasanMajemuk); i++ {
		var refSkoringKecerdasanMajemuk entity.RefSkoringKecerdasanMajemuk
		refSkoringKecerdasanMajemuk.IDQuiz = id_quiz
		refSkoringKecerdasanMajemuk.IDUser = id_user
		no := refKecerdasanMajemuk[i].No
		refSkoringKecerdasanMajemuk.B1 = int32(strings.Index(jawabanUser[0], no)) + 1
		refSkoringKecerdasanMajemuk.B2 = int32(strings.Index(jawabanUser[1], no)) + 1
		refSkoringKecerdasanMajemuk.B3 = int32(strings.Index(jawabanUser[2], no)) + 1
		refSkoringKecerdasanMajemuk.B4 = int32(strings.Index(jawabanUser[3], no)) + 1
		refSkoringKecerdasanMajemuk.B5 = int32(strings.Index(jawabanUser[4], no)) + 1
		refSkoringKecerdasanMajemuk.B6 = int32(strings.Index(jawabanUser[5], no)) + 1
		refSkoringKecerdasanMajemuk.B7 = int32(strings.Index(jawabanUser[6], no)) + 1
		refSkoringKecerdasanMajemuk.B8 = int32(strings.Index(jawabanUser[7], no)) + 1
		refSkoringKecerdasanMajemuk.B9 = int32(strings.Index(jawabanUser[8], no)) + 1
		refSkoringKecerdasanMajemuk.No = no
		refSkoringKecerdasanMajemuk.Total = refSkoringKecerdasanMajemuk.B1 + +refSkoringKecerdasanMajemuk.B2 + refSkoringKecerdasanMajemuk.B3 + refSkoringKecerdasanMajemuk.B4 + refSkoringKecerdasanMajemuk.B5 + refSkoringKecerdasanMajemuk.B6 + refSkoringKecerdasanMajemuk.B7 + refSkoringKecerdasanMajemuk.B8 + refSkoringKecerdasanMajemuk.B9
		refSkoringKecerdasanMajemuk.Rangking = 0
		db.Table("ref_skoring_kecerdasan_majemuk").Create(&refSkoringKecerdasanMajemuk)
	}
	var currentSkoringKecerdasanMajemuk []*entity.RefSkoringKecerdasanMajemuk
	db.Table("ref_skoring_kecerdasan_majemuk").Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Order("total asc, no asc").Scan(&currentSkoringKecerdasanMajemuk)

	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user
	var nomor = 1
	var rangking = 1
	for i := 0; i < len(currentSkoringKecerdasanMajemuk); i++ {
		id_skoring_kecerdasan_majemuk := currentSkoringKecerdasanMajemuk[i].IDSkoringKecerdasanMajemuk
		db.Table("ref_skoring_kecerdasan_majemuk").
			Where("id_skoring_kecerdasan_majemuk = ?", id_skoring_kecerdasan_majemuk).
			UpdateColumn("rangking", rangking)

		if nomor == 1 {
			skoring.Km1 = currentSkoringKecerdasanMajemuk[i].No
		}
		if nomor == 2 {
			skoring.Km2 = currentSkoringKecerdasanMajemuk[i].No
		}
		if nomor == 3 {
			skoring.Km3 = currentSkoringKecerdasanMajemuk[i].No
		}
		if nomor == 4 {
			skoring.Km4 = currentSkoringKecerdasanMajemuk[i].No
		}
		if nomor == 5 {
			skoring.Km5 = currentSkoringKecerdasanMajemuk[i].No
		}
		nomor++
		rangking++
	}

	db.Table(tabel).Create(&skoring)
	return nil
}

func (*repo) SkoringSuasanaKerja(id_quiz int32, id_user int32) error {
	kategori := "SKALA_PMK_SUASANA_KERJA"
	tabel := "skor_suasana_kerja"

	var skoring entity.SkorSuasanaKerja
	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Delete(skoring)
	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user

	var quizSesiUserJawaban []*entity.QuizSesiUserJawaban
	db.Table("quiz_sesi_user_jawaban").
		Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).
		Where("kategori =? ", kategori).
		Order("urutan asc").Scan(&quizSesiUserJawaban)

	for i := 0; i < len(quizSesiUserJawaban); i++ {
		jawaban := quizSesiUserJawaban[i].Jawaban
		if quizSesiUserJawaban[i].Urutan == 1 {
			skoring.SuasanaKerja1 = jawaban
		}
		if quizSesiUserJawaban[i].Urutan == 2 {
			skoring.SuasanaKerja2 = jawaban
		}
		if quizSesiUserJawaban[i].Urutan == 3 {
			skoring.SuasanaKerja3 = jawaban
		}
	}
	db.Table(tabel).Create(&skoring)
	return nil
}

func (*repo) SkoringGayaBelajar(id_quiz int32, id_user int32) error {
	kategori := "SKALA_GAYA_BELAJAR"
	tabel := "skor_gaya_belajar"

	var skoring entity.SkorGayaBelajar
	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Delete(skoring)
	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user
	var skorHitung []*entity.SkorHitungFieldKlasifikasi

	db.Raw(`select x.*, y.klasifikasi from (SELECT
		jawaban as field_skoring,
		COUNT ( * ) AS skor 
	FROM
		quiz_sesi_user_jawaban as a 
	WHERE
		id_quiz = ? 
		AND id_user = ?
		AND kategori = ?
	GROUP BY
		jawaban 
	ORDER BY 
		jawaban) as x, ref_skor_gaya_belajar as y where x.skor = y.skor`, id_quiz, id_user, kategori).Scan(&skorHitung)

	for i := 0; i < len(skorHitung); i++ {
		if skorHitung[i].FieldSkoring == "A" {
			skoring.GayaAuditoris = skorHitung[i].Skor
			skoring.KlasifikasiAuditoris = skorHitung[i].Klasifikasi
		}
		if skorHitung[i].FieldSkoring == "B" {
			skoring.GayaVisual = skorHitung[i].Skor
			skoring.KlasifikasiVisual = skorHitung[i].Klasifikasi
		}
		if skorHitung[i].FieldSkoring == "C" {
			skoring.GayaKinestetik = skorHitung[i].Skor
			skoring.KlasifikasiKinestetik = skorHitung[i].Klasifikasi
		}
	}
	db.Table(tabel).Create(&skoring)

	return nil
}

func (*repo) SkoringKejiwaanDewasa(id_quiz int32, id_user int32) error {
	kategori := "TES_KEJIWAAN_DEWASA_ID"
	tabel := "skor_kejiwaan_dewasa"

	var skoring entity.SkorKejiwaanDewasa
	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Delete(skoring)
	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user

	var skorHitung []*entity.SkorHitungNilaiFieldSkoring
	db.Raw(`select x.*, y.nilai from (select 
		substring(jawaban from 1 for 1)::int  + 
		substring(jawaban from 2 for 1)::int +
		substring(jawaban from 3 for 1)::int +
		substring(jawaban from 4 for 1)::int +
		substring(jawaban from 5 for 1)::int +
		substring(jawaban from 6 for 1)::int +
		substring(jawaban from 7 for 1)::int +
		substring(jawaban from 8 for 1)::int +
		substring(jawaban from 9 for 1)::int +
		substring(jawaban from 10 for 1)::int as skor,
		b.field_skoring 
	from quiz_sesi_user_jawaban as a , ref_model_kejiwaan_dewasa as b  
		where a.id_quiz  = ?
			 and a.id_user = ? 
			 and a.kategori  = ?
			 and a.urutan  = b.id ) as x, ref_skoring_kejiwaan_dewasa as y
			 where x.skor = y.skor`, id_quiz, id_user, kategori).Scan(&skorHitung)
	r := reflect.ValueOf(&skoring).Elem()
	rt := r.Type()
	for i := 0; i < len(skorHitung); i++ {
		skor := skorHitung[i].Skor
		nilai := skorHitung[i].Nilai
		field_skoring := skorHitung[i].FieldSkoring
		komponen := helper.Capitalize(strings.ReplaceAll(field_skoring, "skor_", ""))
		for n := 0; n < rt.NumField(); n++ {
			fieldnameJson := rt.Field(n).Tag.Get("json")
			fieldname := rt.Field(n).Name
			if fieldnameJson == field_skoring {
				reflect.ValueOf(&skoring).Elem().FieldByName(fieldname).SetInt(int64(skor))
				klasifikasiName := fmt.Sprintf("Nilai%v", komponen)
				reflect.ValueOf(&skoring).Elem().FieldByName(klasifikasiName).SetInt(int64(nilai))
			}
		}
	}
	db.Table(tabel).Create(&skoring)

	return nil
}

func (*repo) SkoringKesehatanMental(id_quiz int32, id_user int32) error {
	kategori := "TES_KESEHATAN_MENTAL_ID"
	tabel := "skor_kesehatan_mental"

	var skoring entity.SkorKesehatanMental
	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Delete(skoring)
	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user

	var skorHitung []*entity.SkorHitungNilaiFieldSkoring
	db.Raw(`select x.*, y.nilai from (select 
		substring(jawaban from 1 for 1)::int  + 
		substring(jawaban from 2 for 1)::int +
		substring(jawaban from 3 for 1)::int +
		substring(jawaban from 4 for 1)::int +
		substring(jawaban from 5 for 1)::int +
		substring(jawaban from 6 for 1)::int +
		substring(jawaban from 7 for 1)::int +
		substring(jawaban from 8 for 1)::int +
		substring(jawaban from 9 for 1)::int +
		substring(jawaban from 10 for 1)::int as skor,
		b.field_skoring 
	from quiz_sesi_user_jawaban as a , ref_model_kesehatan_mental as b  
		where a.id_quiz  = ?
			 and a.id_user = ? 
			 and a.kategori  = ?
			 and a.urutan  = b.id ) as x, ref_skoring_kesehatan_mental as y
			 where x.skor = y.skor`, id_quiz, id_user, kategori).Scan(&skorHitung)
	r := reflect.ValueOf(&skoring).Elem()
	rt := r.Type()
	for i := 0; i < len(skorHitung); i++ {
		skor := skorHitung[i].Skor
		nilai := skorHitung[i].Nilai
		field_skoring := skorHitung[i].FieldSkoring
		for n := 0; n < rt.NumField(); n++ {
			fieldnameJson := rt.Field(n).Tag.Get("json")
			fieldname := rt.Field(n).Name
			if fieldnameJson == field_skoring {
				reflect.ValueOf(&skoring).Elem().FieldByName(fieldname).SetInt(int64(skor))
				klasifikasiName := strings.ReplaceAll(fieldname, "Skor", "Nilai")
				reflect.ValueOf(&skoring).Elem().FieldByName(klasifikasiName).SetInt(int64(nilai))
			}
		}
	}
	db.Table(tabel).Create(&skoring)

	return nil
}

func (*repo) SkoringModeBelajar(id_quiz int32, id_user int32) error {
	kategori := "TES_MODE_BELAJAR"
	tabel := "skor_mode_belajar"

	var tabelSkoring entity.SkorModeBelajar
	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Delete(tabelSkoring)

	var skorHitung []*entity.SkorModeBelajar
	db.Raw(`select 
			urutan as id_mode_belajar, 
			substring(jawaban from 1 for 1) as prioritas_1, 
			substring(jawaban from 2 for 1) as prioritas_2, 
			substring(jawaban from 3 for 1) as prioritas_3, 
			substring(jawaban from 4 for 1) as prioritas_4,
			substring(jawaban from 5 for 1) as prioritas_5 
			from quiz_sesi_user_jawaban as a
			where a.id_quiz  = ? and a.id_user = ?
			and a.kategori  = ? `, id_quiz, id_user, kategori).Scan(&skorHitung)

	for i := 0; i < len(skorHitung); i++ {
		var skoring entity.SkorModeBelajar
		skoring.IDQuiz = id_quiz
		skoring.IDUser = id_user
		skoring.IDModeBelajar = skorHitung[i].IDModeBelajar
		skoring.Prioritas1 = skorHitung[i].Prioritas1
		skoring.Prioritas2 = skorHitung[i].Prioritas2
		skoring.Prioritas3 = skorHitung[i].Prioritas3
		skoring.Prioritas4 = skorHitung[i].Prioritas4
		skoring.Prioritas5 = skorHitung[i].Prioritas5
		db.Table(tabel).Create(&skoring)
	}

	return nil
}

func (*repo) SkoringSSCT(id_quiz int32, id_user int32) error {
	kategori := "SSCT_REMAJA"
	tabel := "skor_ssct"

	var tabelSkoring entity.SkorSsct
	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Delete(tabelSkoring)

	var skorHitung []*entity.SkorSsct
	db.Raw(`select x.urutan, x.skor , y.klasifikasi from (select 
		urutan, 
		substring(jawaban from 1 for 1)::int +  
		substring(jawaban from 2 for 1)::int +  
		substring(jawaban from 3 for 1)::int as skor
		from quiz_sesi_user_jawaban as a
		where a.id_quiz  = ? and a.id_user = ?
		and a.kategori  = ? ) as x , ref_skala_skor_ssct as y 
		where x.skor = y.skor `, id_quiz, id_user, kategori).Scan(&skorHitung)

	for i := 0; i < len(skorHitung); i++ {
		var skoring entity.SkorSsct
		skoring.IDQuiz = id_quiz
		skoring.IDUser = id_user
		skoring.Urutan = skorHitung[i].Urutan
		skoring.Skor = skorHitung[i].Skor
		skoring.Klasifikasi = skorHitung[i].Klasifikasi
		db.Table(tabel).Create(&skoring)
	}

	return nil
}

func (*repo) SkoringDISC(id_quiz int32, id_user int32) error {
	kategori := "SKALA_DISC"
	tabel := "skor_disc"
	var tabelSkoring entity.SkorDisc
	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Delete(tabelSkoring)

	var skorHitung []*entity.SkorKategori

	db.Raw(`select  jawaban as kategori, count(*) as skor , '' as klasifikasi
					from quiz_sesi_user_jawaban where id_quiz  = ? and id_user = ? and kategori = ?
					group  by jawaban `, id_quiz, id_user, kategori).Scan(&skorHitung)
	var skoring entity.SkorDisc
	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user
	for i := 0; i < len(skorHitung); i++ {

		if skorHitung[i].Kategori == "D" {
			skoring.SkorD = int32(skorHitung[i].Skor)
			skoring.KlasifikasiD = skorHitung[i].Klasifikasi
		}
		if skorHitung[i].Kategori == "S" {
			skoring.SkorS = int32(skorHitung[i].Skor)
			skoring.KlasifikasiS = skorHitung[i].Klasifikasi
		}
		if skorHitung[i].Kategori == "I" {
			skoring.SkorI = int32(skorHitung[i].Skor)
			skoring.KlasifikasiI = skorHitung[i].Klasifikasi
		}
		if skorHitung[i].Kategori == "C" {
			skoring.SkorC = int32(skorHitung[i].Skor)
			skoring.KlasifikasiC = skorHitung[i].Klasifikasi
		}
	}
	db.Table(tabel).Create(&skoring)
	return nil
}

func (*repo) SkoringModeKerja(id_quiz int32, id_user int32) error {

	kategori := "TES_MODE_KERJA"
	tabel := "skor_mode_kerja"

	var tabelSkoring entity.SkorModeKerja
	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Delete(tabelSkoring)

	var skorHitung []*entity.SkorModeKerja
	db.Raw(`select 
			urutan as id_mode_kerja, 
			substring(jawaban from 1 for 1) as prioritas_1, 
			substring(jawaban from 2 for 1) as prioritas_2, 
			substring(jawaban from 3 for 1) as prioritas_3, 
			substring(jawaban from 4 for 1) as prioritas_4,
			substring(jawaban from 5 for 1) as prioritas_5 
			from quiz_sesi_user_jawaban as a
			where a.id_quiz  = ? and a.id_user = ?
			and a.kategori  = ? `, id_quiz, id_user, kategori).Scan(&skorHitung)

	for i := 0; i < len(skorHitung); i++ {
		var skoring entity.SkorModeKerja
		skoring.IDQuiz = id_quiz
		skoring.IDUser = id_user
		skoring.IDModeKerja = skorHitung[i].IDModeKerja
		skoring.Prioritas1 = skorHitung[i].Prioritas1
		skoring.Prioritas2 = skorHitung[i].Prioritas2
		skoring.Prioritas3 = skorHitung[i].Prioritas3
		skoring.Prioritas4 = skorHitung[i].Prioritas4
		skoring.Prioritas5 = skorHitung[i].Prioritas5
		db.Table(tabel).Create(&skoring)
	}

	return nil
}

func (*repo) SkoringKepribadianManajerial(id_quiz int32, id_user int32) error {
	kategori := "SKALA_KEPRIBADIAN_MANAJERIAL"
	tabel := "skor_kepribadian_manajerial"
	var skoring entity.SkorKepribadianManajerial
	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Delete(skoring)
	//update skor
	db.Exec(`update quiz_sesi_user_jawaban
			set skor = cast(jawaban as integer)
			where id_quiz = ? and id_user = ? and kategori = ?`, id_quiz, id_user, kategori)
	var skorHitung []*entity.SkorHitungFieldKlasifikasi

	db.Raw(`select x.id_user, x.id_quiz, x.total as skor , x.field_skoring, y.klasifikasi from (SELECT 
				a.id_quiz, a.id_user,
				c.field_skoring,
				sum(a.skor) as total
			from quiz_sesi_user_jawaban as a  , 
				soal_kepribadian_manajerial as b , 
				ref_komponen_kepribadian_manajerial as c
			where 
				a.urutan = b.urutan 
				and c.id = b.id_komponen
				and  a.id_quiz = ? 
				and a.id_user = ?
				and a.kategori = ?
			group by 
				a.id_user, a.id_quiz, c.field_skoring) as x, ref_klasifikasi_pribadi_manajerial as y 
				where x.total = y.skor`, id_quiz, id_user, kategori).Scan(&skorHitung)

	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user

	r := reflect.ValueOf(&skoring).Elem()
	rt := r.Type()
	for i := 0; i < len(skorHitung); i++ {
		skor := skorHitung[i].Skor
		klasifikasi := skorHitung[i].Klasifikasi
		for n := 0; n < rt.NumField(); n++ {
			fieldnameJson := rt.Field(n).Tag.Get("json")
			fieldname := rt.Field(n).Name
			if fieldnameJson == skorHitung[i].FieldSkoring {
				// komponen := helper.Capitalize(skorHitung[i].FieldSkoring)
				reflect.ValueOf(&skoring).Elem().FieldByName(fieldname).SetInt(int64(skor))
				klasifikasiName := fmt.Sprintf("Klasifikasi%v", fieldname)
				reflect.ValueOf(&skoring).Elem().FieldByName(klasifikasiName).SetString(klasifikasi)
			}
		}
	}
	db.Table(tabel).Create(&skoring)
	return nil
}

func (*repo) SkoringWLB(id_quiz int32, id_user int32) error {

	kategori := "TES_WORK_LIFE_BALANCED"
	tabel := "skor_wlb"

	var skoring entity.SkorWlb
	db.Table(tabel).Where("id_quiz = ?", id_quiz).Where("id_user = ? ", id_user).Delete(skoring)
	skoring.IDQuiz = id_quiz
	skoring.IDUser = id_user

	var jawabanUser []*entity.QuizSesiUserJawaban
	db.Table("quiz_sesi_user_jawaban").Where("id_quiz = ? ", id_quiz).Where("id_user = ? ", id_user).Where("kategori = ? ", kategori).Order("urutan asc").Scan(&jawabanUser)
	for i := 0; i < len(jawabanUser); i++ {
		urutan := jawabanUser[i].Urutan
		jawaban := jawabanUser[i].Jawaban

		jawabanArr := strings.Split(jawaban, "")
		var skor = 0
		for n := 0; n < len(jawabanArr); n++ {
			var cekSoal *entity.SoalWlb
			db.Table("soal_wlb").Where("id_model = ? ", urutan).Where("urutan = ? ", n).First(&cekSoal)
			if cekSoal.Kategori == "P" {
				if jawabanArr[n] == "Y" {
					skor = skor + 3
				}
				if jawabanArr[n] == "K" {
					skor = skor + 2
				}
				if jawabanArr[n] == "T" {
					skor = skor + 1
				}
			} else {
				if jawabanArr[n] == "Y" {
					skor = skor + 1
				}
				if jawabanArr[n] == "K" {
					skor = skor + 2
				}
				if jawabanArr[n] == "T" {
					skor = skor + 3
				}
			}
		}
		db.Exec(`update quiz_sesi_user_jawaban
			set skor = ?
			where id_quiz = ? and id_user = ? and kategori = ? and urutan = ? `, skor, id_quiz, id_user, kategori, urutan)
	}

	var skorHitung []*entity.SkorHitungNilaiFieldSkoring
	db.Raw(`select x.*, y.nilai from (select 
		a.skor as skor,
		b.field_skoring 
	from quiz_sesi_user_jawaban as a , ref_model_wlb as b  
		where a.id_quiz  = ?
			 and a.id_user = ? 
			 and a.kategori  = ?
			 and a.urutan  = b.id ) as x, ref_skoring_wlb as y
			 where x.skor = y.skor`, id_quiz, id_user, kategori).Scan(&skorHitung)
	r := reflect.ValueOf(&skoring).Elem()
	rt := r.Type()
	for i := 0; i < len(skorHitung); i++ {
		skor := skorHitung[i].Skor
		nilai := skorHitung[i].Nilai
		field_skoring := skorHitung[i].FieldSkoring
		for n := 0; n < rt.NumField(); n++ {
			fieldnameJson := rt.Field(n).Tag.Get("json")
			fieldname := rt.Field(n).Name
			if fieldnameJson == field_skoring {
				reflect.ValueOf(&skoring).Elem().FieldByName(fieldname).SetInt(int64(skor))
				nilaiName := fmt.Sprintf("Nilai%v", fieldname)
				reflect.ValueOf(&skoring).Elem().FieldByName(nilaiName).SetInt(int64(nilai))
			}
		}
	}
	db.Table(tabel).Create(&skoring)
	return nil
}

// SKORING GABUNGAN
func (*repo) SkoringRekomKuliahA(id_quiz int32, id_user int32) error {
	//tabel referensi: skor_rekom_kuliah_a
	//tabel terkait: skor_kuliah_alam,skor_kuliah_sosial,skor_kuliah_dinas
	var skoring_rekom = entity.SkorRekomKuliahA{}
	db.Table("skor_rekom_kuliah_a").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).Delete(&skoring_rekom)

	var refKuliahAlam []*entity.RefKuliahAlam
	db.Table("soal_minat_kuliah_eksakta").Scan(&refKuliahAlam)
	var skoringAlam *entity.SkorKuliahAlam
	db.Table("skor_kuliah_alam").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&skoringAlam)
	var listMinatAlam = []string{"", "", "", "", ""}
	for i := 0; i < len(refKuliahAlam); i++ {
		if refKuliahAlam[i].Urutan == skoringAlam.MinatIpa1 {
			listMinatAlam[0] = fmt.Sprintf("%v. %v", 1, refKuliahAlam[i].Jurusan)
		}
		if refKuliahAlam[i].Urutan == skoringAlam.MinatIpa2 {
			listMinatAlam[1] = fmt.Sprintf("%v. %v", 2, refKuliahAlam[i].Jurusan)
		}
		if refKuliahAlam[i].Urutan == skoringAlam.MinatIpa3 {
			listMinatAlam[2] = fmt.Sprintf("%v. %v", 3, refKuliahAlam[i].Jurusan)
		}
		if refKuliahAlam[i].Urutan == skoringAlam.MinatIpa4 {
			listMinatAlam[3] = fmt.Sprintf("%v. %v", 4, refKuliahAlam[i].Jurusan)
		}
		if refKuliahAlam[i].Urutan == skoringAlam.MinatIpa5 {
			listMinatAlam[4] = fmt.Sprintf("%v. %v", 5, refKuliahAlam[i].Jurusan)
		}

	}

	var refKuliahSosial []*entity.RefKuliahSosial
	db.Table("soal_minat_kuliah_sosial").Scan(&refKuliahSosial)
	var skoringSosial *entity.SkorKuliahSosial
	db.Table("skor_kuliah_sosial").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&skoringSosial)
	var listMinatSosial = []string{"", "", "", "", ""}
	for i := 0; i < len(refKuliahSosial); i++ {
		if refKuliahSosial[i].Urutan == skoringSosial.MinatIps1 {
			listMinatSosial[0] = fmt.Sprintf("%v. %v", 1, refKuliahSosial[i].Jurusan)
		}
		if refKuliahSosial[i].Urutan == skoringSosial.MinatIps2 {
			listMinatSosial[1] = fmt.Sprintf("%v. %v", 2, refKuliahSosial[i].Jurusan)
		}
		if refKuliahSosial[i].Urutan == skoringSosial.MinatIps3 {
			listMinatSosial[2] = fmt.Sprintf("%v. %v", 3, refKuliahSosial[i].Jurusan)
		}
		if refKuliahSosial[i].Urutan == skoringSosial.MinatIps4 {
			listMinatSosial[3] = fmt.Sprintf("%v. %v", 4, refKuliahSosial[i].Jurusan)
		}
		if refKuliahSosial[i].Urutan == skoringSosial.MinatIps5 {
			listMinatSosial[4] = fmt.Sprintf("%v. %v", 5, refKuliahSosial[i].Jurusan)
		}

	}

	var refKuliahDinas []*entity.RefSekolahDinas
	db.Table("ref_sekolah_dinas").Scan(&refKuliahDinas)
	var skoringDinas *entity.SkorKuliahDinas
	db.Table("skor_kuliah_dinas").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&skoringDinas)
	var listMinatDinas = []string{"", "", ""}
	for i := 0; i < len(refKuliahDinas); i++ {
		if refKuliahDinas[i].No == skoringDinas.MinatDinas1 {
			listMinatDinas[0] = fmt.Sprintf("%v. %v - %v", 1, refKuliahDinas[i].Akronim, refKuliahDinas[i].NamaSekolahDinas)
		}
		if refKuliahDinas[i].No == skoringDinas.MinatDinas2 {
			listMinatDinas[1] = fmt.Sprintf("%v. %v - %v", 2, refKuliahDinas[i].Akronim, refKuliahDinas[i].NamaSekolahDinas)
		}
		if refKuliahDinas[i].No == skoringDinas.MinatDinas3 {
			listMinatDinas[2] = fmt.Sprintf("%v. %v - %v", 3, refKuliahDinas[i].Akronim, refKuliahDinas[i].NamaSekolahDinas)
		}
	}

	skoring_rekom.IDQuiz = id_quiz
	skoring_rekom.IDUser = id_user
	skoring_rekom.RekomKuliahAlam = strings.Join(listMinatAlam, ";")
	skoring_rekom.RekomKuliahSosial = strings.Join(listMinatSosial, ";")
	skoring_rekom.RekomKuliahDinas = strings.Join(listMinatDinas, ";")

	db.Table("skor_rekom_kuliah_a").Create(&skoring_rekom)

	return nil
}

func (*repo) SkoringRekomKuliahB(id_quiz int32, id_user int32) error {
	//tabel referensi: skor_rekom_kuliah_b
	//tabel terkait: skor_kuliah_alam,skor_kuliah_sosial,skor_kuliah_dinas,skor_kuliah_agama

	var skoring_rekom = entity.SkorRekomKuliahB{}
	db.Table("skor_rekom_kuliah_b").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).Delete(&skoring_rekom)

	var refKuliahAlam []*entity.RefKuliahAlam
	db.Table("soal_minat_kuliah_eksakta").Scan(&refKuliahAlam)
	var skoringAlam *entity.SkorKuliahAlam
	db.Table("skor_kuliah_alam").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&skoringAlam)
	var listMinatAlam = []string{"", "", "", "", ""}
	for i := 0; i < len(refKuliahAlam); i++ {
		if refKuliahAlam[i].Urutan == skoringAlam.MinatIpa1 {
			listMinatAlam[0] = fmt.Sprintf("%v. %v", 1, refKuliahAlam[i].Jurusan)
		}
		if refKuliahAlam[i].Urutan == skoringAlam.MinatIpa2 {
			listMinatAlam[1] = fmt.Sprintf("%v. %v", 2, refKuliahAlam[i].Jurusan)
		}
		if refKuliahAlam[i].Urutan == skoringAlam.MinatIpa3 {
			listMinatAlam[2] = fmt.Sprintf("%v. %v", 3, refKuliahAlam[i].Jurusan)
		}
		if refKuliahAlam[i].Urutan == skoringAlam.MinatIpa4 {
			listMinatAlam[3] = fmt.Sprintf("%v. %v", 4, refKuliahAlam[i].Jurusan)
		}
		if refKuliahAlam[i].Urutan == skoringAlam.MinatIpa5 {
			listMinatAlam[4] = fmt.Sprintf("%v. %v", 5, refKuliahAlam[i].Jurusan)
		}

	}

	var refKuliahSosial []*entity.RefKuliahSosial
	db.Table("soal_minat_kuliah_sosial").Scan(&refKuliahSosial)
	var skoringSosial *entity.SkorKuliahSosial
	db.Table("skor_kuliah_sosial").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&skoringSosial)
	var listMinatSosial = []string{"", "", "", "", ""}
	for i := 0; i < len(refKuliahSosial); i++ {
		if refKuliahSosial[i].Urutan == skoringSosial.MinatIps1 {
			listMinatSosial[0] = fmt.Sprintf("%v. %v", 1, refKuliahSosial[i].Jurusan)
		}
		if refKuliahSosial[i].Urutan == skoringSosial.MinatIps2 {
			listMinatSosial[1] = fmt.Sprintf("%v. %v", 2, refKuliahSosial[i].Jurusan)
		}
		if refKuliahSosial[i].Urutan == skoringSosial.MinatIps3 {
			listMinatSosial[2] = fmt.Sprintf("%v. %v", 3, refKuliahSosial[i].Jurusan)
		}
		if refKuliahSosial[i].Urutan == skoringSosial.MinatIps4 {
			listMinatSosial[3] = fmt.Sprintf("%v. %v", 4, refKuliahSosial[i].Jurusan)
		}
		if refKuliahSosial[i].Urutan == skoringSosial.MinatIps5 {
			listMinatSosial[4] = fmt.Sprintf("%v. %v", 5, refKuliahSosial[i].Jurusan)
		}

	}

	var refKuliahDinas []*entity.RefSekolahDinas
	db.Table("ref_sekolah_dinas").Scan(&refKuliahDinas)
	var skoringDinas *entity.SkorKuliahDinas
	db.Table("skor_kuliah_dinas").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&skoringDinas)
	var listMinatDinas = []string{"", "", ""}
	for i := 0; i < len(refKuliahDinas); i++ {
		if refKuliahDinas[i].No == skoringDinas.MinatDinas1 {
			listMinatDinas[0] = fmt.Sprintf("%v. %v - %v", 1, refKuliahDinas[i].Akronim, refKuliahDinas[i].NamaSekolahDinas)
		}
		if refKuliahDinas[i].No == skoringDinas.MinatDinas2 {
			listMinatDinas[1] = fmt.Sprintf("%v. %v - %v", 2, refKuliahDinas[i].Akronim, refKuliahDinas[i].NamaSekolahDinas)
		}
		if refKuliahDinas[i].No == skoringDinas.MinatDinas3 {
			listMinatDinas[2] = fmt.Sprintf("%v. %v - %v", 3, refKuliahDinas[i].Akronim, refKuliahDinas[i].NamaSekolahDinas)
		}
	}

	var refKuliahAgama []*entity.RefKuliahAgama
	db.Table("soal_minat_kuliah_agama").Scan(&refKuliahAgama)
	var skoringAgama *entity.SkorKuliahAgama
	db.Table("skor_kuliah_agama").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&skoringAgama)
	var listMinatAgama = []string{"", "", "", "", ""}
	for i := 0; i < len(refKuliahAgama); i++ {
		if refKuliahAgama[i].Urutan == skoringAgama.MinatAgm1 {
			listMinatAgama[0] = fmt.Sprintf("%v. %v", 1, refKuliahAgama[i].Jurusan)
		}
		if refKuliahAgama[i].Urutan == skoringAgama.MinatAgm2 {
			listMinatAgama[1] = fmt.Sprintf("%v. %v", 2, refKuliahAgama[i].Jurusan)
		}
		if refKuliahAgama[i].Urutan == skoringAgama.MinatAgm3 {
			listMinatAgama[2] = fmt.Sprintf("%v. %v", 3, refKuliahAgama[i].Jurusan)
		}
		if refKuliahAgama[i].Urutan == skoringAgama.MinatAgm4 {
			listMinatAgama[3] = fmt.Sprintf("%v. %v", 4, refKuliahAgama[i].Jurusan)
		}
		if refKuliahAgama[i].Urutan == skoringAgama.MinatAgm5 {
			listMinatAgama[4] = fmt.Sprintf("%v. %v", 5, refKuliahAgama[i].Jurusan)
		}
	}

	skoring_rekom.IDQuiz = id_quiz
	skoring_rekom.IDUser = id_user
	skoring_rekom.RekomKuliahAlam = strings.Join(listMinatAlam, ";")
	skoring_rekom.RekomKuliahSosial = strings.Join(listMinatSosial, ";")
	skoring_rekom.RekomKuliahDinas = strings.Join(listMinatDinas, ";")
	skoring_rekom.RekomKuliahAgama = strings.Join(listMinatAgama, ";")
	db.Table("skor_rekom_kuliah_b").Create(&skoring_rekom)

	return nil
}

func (*repo) SkoringRekomPeminatanSMA(id_quiz int32, id_user int32) error {
	//tabel referensi: skor_rekom_peminatan_sma
	//tabel terkait: skor_peminatan_sma,skor_sikap_pelajaran
	var skoring_minat = entity.SkorRekomPeminatanSma{}
	db.Table("skor_rekom_peminatan_sma").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).Delete(&skoring_minat)
	var skoringMinatSMA *entity.SkorPeminatanSma
	db.Table("skor_peminatan_sma").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&skoringMinatSMA)

	var skoringSikapPelajaran *entity.SkorSikapPelajaran
	db.Table("skor_sikap_pelajaran").Where("id_quiz = ?", id_quiz).Where("id_user = ?", id_user).First(&skoringSikapPelajaran)

	var refRekomendasiAkhirPeminatanSma *entity.RefRekomendasiAkhirPeminatanSma
	db.Table("ref_rekomendasi_akhir_peminatan_sma").Where("rekom_minat = ?", skoringMinatSMA.RekomMinat).
		Where("rekom_sikap_pelajaran = ?", skoringSikapPelajaran.RekomSikapPelajaran).
		First(&refRekomendasiAkhirPeminatanSma)

	skoring_minat.IDQuiz = id_quiz
	skoring_minat.IDUser = id_user
	skoring_minat.RekomMinat = skoringMinatSMA.RekomMinat
	skoring_minat.RekomSikapPelajaran = skoringSikapPelajaran.RekomSikapPelajaran
	skoring_minat.RekomMapel = refRekomendasiAkhirPeminatanSma.RekomAkhir

	db.Table("skor_rekom_peminatan_sma").Create(&skoring_minat)

	return nil
}

func (*repo) FinishSkoring(id_quiz int32, id_user int32, waktu string) error {
	//update status skoring
	db.Table("quiz_sesi_user").
		Where("id_quiz = ?", id_quiz).
		Where("id_user = ?", id_user).
		UpdateColumns(map[string]interface{}{
			"skoring":    1,
			"skoring_at": waktu})
	return nil
}
