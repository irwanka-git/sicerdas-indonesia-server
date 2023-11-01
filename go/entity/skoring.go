package entity

type StatusRunningSkoring struct {
	Status int `json:"status"`
}

type QuizSesiUserSkoring struct {
	IDQuiz         int32  `json:"id_quiz"`
	IDUser         int32  `json:"id_user"`
	Skoring        int    `json:"skoring"`
	Submit         int    `json:"submit"`
	Token          string `json:"token"`
	NamaSesi       string `json:"nama_sesi"`
	Kota           string `json:"kota"`
	IDQuizTemplate int32  `json:"id_quiz_template"`
}

type KategoriTabel struct {
	Kategori string `json:"kategori"`
	Tabel    string `json:"tabel"`
}

type JawabanSubmit struct {
	IDQuiz  int32  `json:"id_quiz"`
	IDUser  int32  `json:"id_user"`
	Jawaban string `json:"jawaban"`
}

type JawabanMapping struct {
	Kategori     string   `json:"kategori"`
	Jawaban      string   `json:"jawaban"`
	JawabanArray []string `json:"jawaban_array"`
}

type QuizSesiUserJawaban struct {
	IDJawabanPeserta int32  `gorm:"column:id_jawaban_peserta;primaryKey;autoIncrement:true" json:"id_jawaban_peserta"`
	IDQuiz           int32  `gorm:"column:id_quiz" json:"id_quiz"`
	IDUser           int32  `gorm:"column:id_user" json:"id_user"`
	Kategori         string `gorm:"column:kategori" json:"kategori"`
	Urutan           int32  `gorm:"column:urutan" json:"urutan"`
	Jawaban          string `gorm:"column:jawaban" json:"jawaban"`
	Skor             int32  `gorm:"column:skor" json:"skor"`
}

type SkorKognitif struct {
	ID            int64   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	IDUser        int32   `gorm:"column:id_user" json:"id_user"`
	IDQuiz        int32   `gorm:"column:id_quiz" json:"id_quiz"`
	TpaIu         int32   `gorm:"column:tpa_iu" json:"tpa_iu"`
	TpaPv         int32   `gorm:"column:tpa_pv" json:"tpa_pv"`
	TpaPk         int32   `gorm:"column:tpa_pk" json:"tpa_pk"`
	TpaPa         int32   `gorm:"column:tpa_pa" json:"tpa_pa"`
	TpaPs         int32   `gorm:"column:tpa_ps" json:"tpa_ps"`
	TpaPm         int32   `gorm:"column:tpa_pm" json:"tpa_pm"`
	TpaKt         int32   `gorm:"column:tpa_kt" json:"tpa_kt"`
	TpaIq         int32   `gorm:"column:tpa_iq" json:"tpa_iq"`
	SkorIq        float32 `gorm:"column:skor_iq" json:"skor_iq"`
	KlasifikasiIu string  `gorm:"column:klasifikasi_iu" json:"klasifikasi_iu"`
	KlasifikasiPv string  `gorm:"column:klasifikasi_pv" json:"klasifikasi_pv"`
	KlasifikasiPk string  `gorm:"column:klasifikasi_pk" json:"klasifikasi_pk"`
	KlasifikasiPa string  `gorm:"column:klasifikasi_pa" json:"klasifikasi_pa"`
	KlasifikasiPs string  `gorm:"column:klasifikasi_ps" json:"klasifikasi_ps"`
	KlasifikasiPm string  `gorm:"column:klasifikasi_pm" json:"klasifikasi_pm"`
	KlasifikasiKt string  `gorm:"column:klasifikasi_kt" json:"klasifikasi_kt"`
	KlasifikasiIq string  `gorm:"column:klasifikasi_iq" json:"klasifikasi_iq"`
}

type RefKonversiIq90 struct {
	ID          int32   `gorm:"column:id;primaryKey" json:"id"`
	SkorX       int32   `gorm:"column:skor_x" json:"skor_x"`
	TotIq       float32 `gorm:"column:tot_iq" json:"tot_iq"`
	Klasifikasi string  `gorm:"column:klasifikasi" json:"klasifikasi"`
}

type RefKonversiIq105 struct {
	ID          int32   `gorm:"column:id;primaryKey" json:"id"`
	SkorX       int32   `gorm:"column:skor_x" json:"skor_x"`
	TotIq       float32 `gorm:"column:tot_iq" json:"tot_iq"`
	Klasifikasi string  `gorm:"column:klasifikasi" json:"klasifikasi"`
}

type RefKomponenGayaPekerjaan struct {
	NamaKomponen  string `gorm:"column:nama_komponen" json:"nama_komponen"`
	Kode          string `gorm:"column:kode" json:"kode"`
	No            string `gorm:"column:no;primaryKey" json:"no"`
	CetakKomponen string `gorm:"column:cetak_komponen" json:"cetak_komponen"`
	Deskripsi     string `gorm:"column:deskripsi" json:"deskripsi"`
	Pekerjaan     string `gorm:"column:pekerjaan" json:"pekerjaan"`
	Gambar        string `gorm:"column:gambar" json:"gambar"`
}

type RefSkoringGayaPekerjaan struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	IDUser      int32  `gorm:"column:id_user" json:"id_user"`
	IDQuiz      int32  `gorm:"column:id_quiz" json:"id_quiz"`
	Kode        string `gorm:"column:kode" json:"kode"`
	Skor        int16  `gorm:"column:skor" json:"skor"`
	Rangking    int16  `gorm:"column:rangking" json:"rangking"`
	Klasifikasi string `gorm:"column:klasifikasi" json:"klasifikasi"`
}

type RefSkorGayaPekerjaan struct {
	ID      int16  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Respon  string `gorm:"column:respon" json:"respon"`
	U       int16  `gorm:"column:u" json:"u"`
	T       int16  `gorm:"column:t" json:"t"`
	C       int16  `gorm:"column:c" json:"c"`
	Jawaban string `gorm:"column:jawaban" json:"jawaban"`
}

type RefKlasifikasiGayaKerja struct {
	ID          int16  `gorm:"column:id;primaryKey" json:"id"`
	Klasifikasi string `gorm:"column:klasifikasi" json:"klasifikasi"`
	SkorMin     int16  `gorm:"column:skor_min" json:"skor_min"`
	SkorMax     int16  `gorm:"column:skor_max" json:"skor_max"`
	Akronim     string `gorm:"column:akronim" json:"akronim"`
}

type SoalGayaPekerjaan struct {
	Nomor     int16  `gorm:"column:nomor;primaryKey" json:"nomor"`
	Deskripsi string `gorm:"column:deskripsi" json:"deskripsi"`
	KomponenA string `gorm:"column:komponen_a" json:"komponen_a"`
	KomponenB string `gorm:"column:komponen_b" json:"komponen_b"`
	KomponenC string `gorm:"column:komponen_c" json:"komponen_c"`
	KomponenD string `gorm:"column:komponen_d" json:"komponen_d"`
	KomponenE string `gorm:"column:komponen_e" json:"komponen_e"`
	KomponenF string `gorm:"column:komponen_f" json:"komponen_f"`
	KomponenG string `gorm:"column:komponen_g" json:"komponen_g"`
	KomponenH string `gorm:"column:komponen_h" json:"komponen_h"`
	KomponenI string `gorm:"column:komponen_i" json:"komponen_i"`
	KomponenJ string `gorm:"column:komponen_j" json:"komponen_j"`
	KomponenK string `gorm:"column:komponen_k" json:"komponen_k"`
	KomponenL string `gorm:"column:komponen_l" json:"komponen_l"`
	UUID      string `gorm:"column:uuid" json:"uuid"`
}

type DataSkorGayaPekerjaan struct {
	Nomor     int16  `gorm:"column:nomor;primaryKey" json:"nomor"`
	Deskripsi string `gorm:"column:deskripsi" json:"deskripsi"`
	KomponenA string `gorm:"column:komponen_a" json:"komponen_a"`
	KomponenB string `gorm:"column:komponen_b" json:"komponen_b"`
	KomponenC string `gorm:"column:komponen_c" json:"komponen_c"`
	KomponenD string `gorm:"column:komponen_d" json:"komponen_d"`
	KomponenE string `gorm:"column:komponen_e" json:"komponen_e"`
	KomponenF string `gorm:"column:komponen_f" json:"komponen_f"`
	KomponenG string `gorm:"column:komponen_g" json:"komponen_g"`
	KomponenH string `gorm:"column:komponen_h" json:"komponen_h"`
	KomponenI string `gorm:"column:komponen_i" json:"komponen_i"`
	KomponenJ string `gorm:"column:komponen_j" json:"komponen_j"`
	KomponenK string `gorm:"column:komponen_k" json:"komponen_k"`
	KomponenL string `gorm:"column:komponen_l" json:"komponen_l"`
	U         int16  `gorm:"column:u" json:"u"`
	T         int16  `gorm:"column:t" json:"t"`
	C         int16  `gorm:"column:c" json:"c"`
	Jawaban   string `gorm:"column:jawaban" json:"jawaban"`
}

type SkorGayaPekerjaan struct {
	IDQuiz      int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	IDUser      int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	RangkingGp1 string `gorm:"column:rangking_gp1" json:"rangking_gp1"`
	RangkingGp2 string `gorm:"column:rangking_gp2" json:"rangking_gp2"`
	RangkingGp3 string `gorm:"column:rangking_gp3" json:"rangking_gp3"`
}

type SoalSikapPelajaran struct {
	IDSoal        int32  `gorm:"column:id_soal;primaryKey;autoIncrement:true" json:"id_soal"`
	Urutan        int32  `gorm:"column:urutan" json:"urutan"`
	Pelajaran     string `gorm:"column:pelajaran" json:"pelajaran"`
	Kode          string `gorm:"column:kode" json:"kode"`
	SikapNegatif1 string `gorm:"column:sikap_negatif1" json:"sikap_negatif1"`
	SikapPositif1 string `gorm:"column:sikap_positif1" json:"sikap_positif1"`
	SikapNegatif2 string `gorm:"column:sikap_negatif2" json:"sikap_negatif2"`
	SikapPositif2 string `gorm:"column:sikap_positif2" json:"sikap_positif2"`
	SikapNegatif3 string `gorm:"column:sikap_negatif3" json:"sikap_negatif3"`
	SikapPositif3 string `gorm:"column:sikap_positif3" json:"sikap_positif3"`
	Kelompok      string `gorm:"column:kelompok" json:"kelompok"`
	FieldSkoring  string `gorm:"column:field_skoring" json:"field_skoring"`
	UUID          string `gorm:"column:uuid" json:"uuid"`
}

type SkorHitungSikapPelajaran struct {
	IDUser       int32  `json:"id_user"`
	IDQuiz       int32  `json:"id_quiz"`
	Skor         int16  `json:"skor"`
	FieldSkoring string `json:"field_skoring"`
	Klasifikasi  string `json:"klasifikasi"`
}

type RefSkalaSikapPelajaran struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Skor        int32  `gorm:"column:skor" json:"skor"`
	Klasifikasi string `gorm:"column:klasifikasi" json:"klasifikasi"`
}

type RefKelompokSikapPelajaran struct {
	Kelompok     string `gorm:"column:kelompok" json:"kelompok"`
	FieldSkoring string `gorm:"column:field_skoring" json:"field_skoring"`
}

type RefRekomendasiSikapPelajaran struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Perbedaan   int32  `gorm:"column:perbedaan;not null;comment:PERBEDAAN = (MAT+FIS+BIO ) - (EKO+SEJ+GEO)" json:"perbedaan"` // PERBEDAAN = (MAT+FIS+BIO ) - (EKO+SEJ+GEO)
	Rekomendasi string `gorm:"column:rekomendasi;not null;default:0" json:"rekomendasi"`
}

type SkorSikapPelajaran struct {
	IDUser              int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz              int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	SikapAgm            int32  `gorm:"column:sikap_agm" json:"sikap_agm"`
	SikapPkn            int32  `gorm:"column:sikap_pkn" json:"sikap_pkn"`
	SikapInd            int32  `gorm:"column:sikap_ind" json:"sikap_ind"`
	SikapEng            int32  `gorm:"column:sikap_eng" json:"sikap_eng"`
	SikapMat            int32  `gorm:"column:sikap_mat" json:"sikap_mat"`
	SikapFis            int32  `gorm:"column:sikap_fis" json:"sikap_fis"`
	SikapBio            int32  `gorm:"column:sikap_bio" json:"sikap_bio"`
	SikapEko            int32  `gorm:"column:sikap_eko" json:"sikap_eko"`
	SikapSej            int32  `gorm:"column:sikap_sej" json:"sikap_sej"`
	SikapGeo            int32  `gorm:"column:sikap_geo" json:"sikap_geo"`
	KlasifikasiAgm      string `gorm:"column:klasifikasi_agm" json:"klasifikasi_agm"`
	KlasifikasiPkn      string `gorm:"column:klasifikasi_pkn" json:"klasifikasi_pkn"`
	KlasifikasiInd      string `gorm:"column:klasifikasi_ind" json:"klasifikasi_ind"`
	KlasifikasiEng      string `gorm:"column:klasifikasi_eng" json:"klasifikasi_eng"`
	KlasifikasiMat      string `gorm:"column:klasifikasi_mat" json:"klasifikasi_mat"`
	KlasifikasiFis      string `gorm:"column:klasifikasi_fis" json:"klasifikasi_fis"`
	KlasifikasiBio      string `gorm:"column:klasifikasi_bio" json:"klasifikasi_bio"`
	KlasifikasiEko      string `gorm:"column:klasifikasi_eko" json:"klasifikasi_eko"`
	KlasifikasiSej      string `gorm:"column:klasifikasi_sej" json:"klasifikasi_sej"`
	KlasifikasiGeo      string `gorm:"column:klasifikasi_geo" json:"klasifikasi_geo"`
	SikapIlmuAlam       int32  `gorm:"column:sikap_ilmu_alam" json:"sikap_ilmu_alam"`
	SikapIlmuSosial     int32  `gorm:"column:sikap_ilmu_sosial" json:"sikap_ilmu_sosial"`
	SikapRentang        int32  `gorm:"column:sikap_rentang" json:"sikap_rentang"`
	RekomSikapPelajaran string `gorm:"column:rekom_sikap_pelajaran" json:"rekom_sikap_pelajaran"`
}
