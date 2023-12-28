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

type JawabanQuiz struct {
	Kategori string   `json:"kategori"`
	Jawaban  []string `json:"jawaban"`
}

type JawabanString struct {
	Kategori string `json:"kategori"`
	Jawaban  string `json:"jawaban"`
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

type SkorHitungFieldKlasifikasi struct {
	IDUser       int32  `json:"id_user"`
	IDQuiz       int32  `json:"id_quiz"`
	Skor         int16  `json:"skor"`
	FieldSkoring string `json:"field_skoring"`
	Klasifikasi  string `json:"klasifikasi"`
}

type SkorRekomendasi struct {
	Skor        int16  `json:"skor"`
	Rekomendasi string `json:"rekomendasi"`
}

type RefBidangKognitif struct {
	NamaBidang   string `gorm:"column:nama_bidang;primaryKey" json:"nama_bidang"`
	FieldSkoring string `gorm:"column:field_skoring" json:"field_skoring"`
	Deskripsi    string `gorm:"column:deskripsi" json:"deskripsi"`
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

type SkorSikapPelajaranMk struct {
	IDUser         int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz         int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	SikapAgm       int32  `gorm:"column:sikap_agm" json:"sikap_agm"`
	SikapPkn       int32  `gorm:"column:sikap_pkn" json:"sikap_pkn"`
	SikapInd       int32  `gorm:"column:sikap_ind" json:"sikap_ind"`
	SikapEng       int32  `gorm:"column:sikap_eng" json:"sikap_eng"`
	SikapMat       int32  `gorm:"column:sikap_mat" json:"sikap_mat"`
	SikapFis       int32  `gorm:"column:sikap_fis" json:"sikap_fis"`
	SikapBio       int32  `gorm:"column:sikap_bio" json:"sikap_bio"`
	SikapKim       int32  `gorm:"column:sikap_kim" json:"sikap_kim"`
	SikapEko       int32  `gorm:"column:sikap_eko" json:"sikap_eko"`
	SikapSej       int32  `gorm:"column:sikap_sej" json:"sikap_sej"`
	SikapSos       int32  `gorm:"column:sikap_sos" json:"sikap_sos"`
	SikapGeo       int32  `gorm:"column:sikap_geo" json:"sikap_geo"`
	SikapSbd       int32  `gorm:"column:sikap_sbd" json:"sikap_sbd"`
	SikapOrg       int32  `gorm:"column:sikap_org" json:"sikap_org"`
	SikapMlk       int32  `gorm:"column:sikap_mlk" json:"sikap_mlk"`
	SikapTik       int32  `gorm:"column:sikap_tik" json:"sikap_tik"`
	KlasifikasiAgm string `gorm:"column:klasifikasi_agm" json:"klasifikasi_agm"`
	KlasifikasiPkn string `gorm:"column:klasifikasi_pkn" json:"klasifikasi_pkn"`
	KlasifikasiInd string `gorm:"column:klasifikasi_ind" json:"klasifikasi_ind"`
	KlasifikasiEng string `gorm:"column:klasifikasi_eng" json:"klasifikasi_eng"`
	KlasifikasiMat string `gorm:"column:klasifikasi_mat" json:"klasifikasi_mat"`
	KlasifikasiFis string `gorm:"column:klasifikasi_fis" json:"klasifikasi_fis"`
	KlasifikasiBio string `gorm:"column:klasifikasi_bio" json:"klasifikasi_bio"`
	KlasifikasiKim string `gorm:"column:klasifikasi_kim" json:"klasifikasi_kim"`
	KlasifikasiEko string `gorm:"column:klasifikasi_eko" json:"klasifikasi_eko"`
	KlasifikasiSej string `gorm:"column:klasifikasi_sej" json:"klasifikasi_sej"`
	KlasifikasiSos string `gorm:"column:klasifikasi_sos" json:"klasifikasi_sos"`
	KlasifikasiGeo string `gorm:"column:klasifikasi_geo" json:"klasifikasi_geo"`
	KlasifikasiSbd string `gorm:"column:klasifikasi_sbd" json:"klasifikasi_sbd"`
	KlasifikasiOrg string `gorm:"column:klasifikasi_org" json:"klasifikasi_org"`
	KlasifikasiMlk string `gorm:"column:klasifikasi_mlk" json:"klasifikasi_mlk"`
	KlasifikasiTik string `gorm:"column:klasifikasi_tik" json:"klasifikasi_tik"`
}

type SkorPeminatanSmk struct {
	IDUser  int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz  int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	Minat1  string `gorm:"column:minat_1;default:null;" json:"minat_1"`
	Minat2  string `gorm:"column:minat_2;default:null;" json:"minat_2"`
	Minat3  string `gorm:"column:minat_3;default:null;" json:"minat_3"`
	Minat4  string `gorm:"column:minat_4;default:null;" json:"minat_4"`
	Minat5  string `gorm:"column:minat_5;default:null;" json:"minat_5"`
	Minat6  string `gorm:"column:minat_6;default:null;" json:"minat_6"`
	Minat7  string `gorm:"column:minat_7;default:null;" json:"minat_7"`
	Minat8  string `gorm:"column:minat_8;default:null;" json:"minat_8"`
	Minat9  string `gorm:"column:minat_9;default:null;" json:"minat_9"`
	Minat10 string `gorm:"column:minat_10;default:null;" json:"minat_10"`
	Minat11 string `gorm:"column:minat_11;default:null;" json:"minat_11"`
	Minat12 string `gorm:"column:minat_12;default:null;" json:"minat_12"`
	Minat13 string `gorm:"column:minat_13;default:null;" json:"minat_13"`
	Minat14 string `gorm:"column:minat_14;default:null;" json:"minat_14"`
	Minat15 string `gorm:"column:minat_15;default:null;" json:"minat_15"`
}

type RefKlasifikasiPeminatanSMA struct {
	ID          int32  `json:"id"`
	Skor        int32  `json:"skor"`
	Klasifikasi string `json:"klasifikasi"`
}

type SkorPeminatanSma struct {
	IDUser                    int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz                    int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	MinatSains                int32  `gorm:"column:minat_sains" json:"minat_sains"`
	MinatHumaniora            int32  `gorm:"column:minat_humaniora" json:"minat_humaniora"`
	MinatBahasa               int32  `gorm:"column:minat_bahasa" json:"minat_bahasa"`
	MinatRentang              int32  `gorm:"column:minat_rentang" json:"minat_rentang"`
	RekomMinat                string `gorm:"column:rekom_minat" json:"rekom_minat"`
	KlasifikasiMinatSains     string `gorm:"column:klasifikasi_minat_sains" json:"klasifikasi_minat_sains"`
	KlasifikasiMinatHumaniora string `gorm:"column:klasifikasi_minat_humaniora" json:"klasifikasi_minat_humaniora"`
	KlasifikasiMinatBahasa    string `gorm:"column:klasifikasi_minat_bahasa" json:"klasifikasi_minat_bahasa"`
}

type RefKlasifikasiPeminatanMAN struct {
	ID          int32  `json:"id"`
	Skor        int32  `json:"skor"`
	Klasifikasi string `json:"klasifikasi"`
}

type SkorPeminatanMan struct {
	IDUser                    int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz                    int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	MinatSains                int32  `gorm:"column:minat_sains" json:"minat_sains"`
	MinatHumaniora            int32  `gorm:"column:minat_humaniora" json:"minat_humaniora"`
	MinatBahasa               int32  `gorm:"column:minat_bahasa" json:"minat_bahasa"`
	MinatRentang              int32  `gorm:"column:minat_rentang" json:"minat_rentang"`
	RekomMinat                string `gorm:"column:rekom_minat" json:"rekom_minat"`
	MinatAgama                int32  `gorm:"column:minat_agama" json:"minat_agama"`
	KlasifikasiMinatSains     string `gorm:"column:klasifikasi_minat_sains" json:"klasifikasi_minat_sains"`
	KlasifikasiMinatHumaniora string `gorm:"column:klasifikasi_minat_humaniora" json:"klasifikasi_minat_humaniora"`
	KlasifikasiMinatBahasa    string `gorm:"column:klasifikasi_minat_bahasa" json:"klasifikasi_minat_bahasa"`
	KlasifikasiMinatAgama     string `gorm:"column:klasifikasi_minat_agama" json:"klasifikasi_minat_agama"`
}

type RefSekolahDinas struct {
	No               string `gorm:"column:no;primaryKey" json:"no"`
	NamaSekolahDinas string `gorm:"column:nama_sekolah_dinas" json:"nama_sekolah_dinas"`
	Akronim          string `gorm:"column:akronim" json:"akronim"`
	IDKelompok       int16  `gorm:"column:id_kelompok;comment:SKORING PSIKOTES LENGKAP" json:"id_kelompok"` // SKORING PSIKOTES LENGKAP
	IDKelas          int16  `gorm:"column:id_kelas;comment:SKORING MINAT SMA V2" json:"id_kelas"`           // SKORING MINAT SMA V2
	Icon             string `gorm:"column:icon;default:'default.png'" json:"icon"`
}

type RefSkoringKuliahDinas struct {
	IDSkoringSekolahDinas int32 `gorm:"column:id_skoring_sekolah_dinas;primaryKey;autoIncrement:true" json:"id_skoring_sekolah_dinas"`
	IDQuiz                int32 `gorm:"column:id_quiz;comment:ID_QUIZ" json:"id_quiz"` // ID_QUIZ
	IDUser                int32 `gorm:"column:id_user;comment:ID_USER" json:"id_user"` // ID_USER

	No string `gorm:"column:no;comment:NO SEKOLAH DINAS (ref_sekolah_dinas)" json:"no"` // NO SEKOLAH DINAS (ref_sekolah_dinas)
	B1 int32  `gorm:"column:b1" json:"b1"`
	B2 int32  `gorm:"column:b2" json:"b2"`
	B3 int32  `gorm:"column:b3" json:"b3"`
	B4 int32  `gorm:"column:b4" json:"b4"`
	B5 int32  `gorm:"column:b5" json:"b5"`
	B6 int32  `gorm:"column:b6" json:"b6"`
	B7 int32  `gorm:"column:b7" json:"b7"`
	B8 int32  `gorm:"column:b8" json:"b8"`
	B9 int32  `gorm:"column:b9" json:"b9"`

	Total    int32 `gorm:"column:total" json:"total"`
	Rangking int32 `gorm:"column:rangking" json:"rangking"`
}

type SkorKuliahDinas struct {
	IDUser      int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz      int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	MinatDinas1 string `gorm:"column:minat_dinas1" json:"minat_dinas1"`
	MinatDinas2 string `gorm:"column:minat_dinas2" json:"minat_dinas2"`
	MinatDinas3 string `gorm:"column:minat_dinas3" json:"minat_dinas3"`
}

type SkorKuliahAlam struct {
	IDUser    int32 `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz    int32 `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	MinatIpa1 int32 `gorm:"column:minat_ipa1" json:"minat_ipa1"`
	MinatIpa2 int32 `gorm:"column:minat_ipa2" json:"minat_ipa2"`
	MinatIpa3 int32 `gorm:"column:minat_ipa3" json:"minat_ipa3"`
	MinatIpa4 int32 `gorm:"column:minat_ipa4" json:"minat_ipa4"`
	MinatIpa5 int32 `gorm:"column:minat_ipa5" json:"minat_ipa5"`
}

type SkorKuliahSosial struct {
	IDUser    int32 `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz    int32 `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	MinatIps1 int32 `gorm:"column:minat_ips1" json:"minat_ips1"`
	MinatIps2 int32 `gorm:"column:minat_ips2" json:"minat_ips2"`
	MinatIps3 int32 `gorm:"column:minat_ips3" json:"minat_ips3"`
	MinatIps4 int32 `gorm:"column:minat_ips4" json:"minat_ips4"`
	MinatIps5 int32 `gorm:"column:minat_ips5" json:"minat_ips5"`
}

type SkorKuliahAgama struct {
	IDUser    int32 `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz    int32 `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	MinatAgm1 int32 `gorm:"column:minat_agm1" json:"minat_agm1"`
	MinatAgm2 int32 `gorm:"column:minat_agm2" json:"minat_agm2"`
	MinatAgm3 int32 `gorm:"column:minat_agm3" json:"minat_agm3"`
	MinatAgm4 int32 `gorm:"column:minat_agm4" json:"minat_agm4"`
	MinatAgm5 int32 `gorm:"column:minat_agm5" json:"minat_agm5"`
}

type SkorHitungMBTI struct {
	IDUser        int32  `json:"id_user"`
	IDQuiz        int32  `json:"id_quiz"`
	SkorA         int32  `json:"skor_a"`
	SkorB         int32  `json:"skor_b"`
	Kolom         string `json:"kolom"`
	Kode          string `json:"kode"`
	Klasifikasi   string `json:"klasifikasi"`
	FieldSkoringA string `json:"field_skoring_a"`
	FieldSkoringB string `json:"field_skoring_b"`
}

type SkorMbti struct {
	IDUser       int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz       int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	TipojungE    int32  `gorm:"column:tipojung_e" json:"tipojung_e"`
	TipojungI    int32  `gorm:"column:tipojung_i" json:"tipojung_i"`
	TipojungS    int32  `gorm:"column:tipojung_s" json:"tipojung_s"`
	TipojungN    int32  `gorm:"column:tipojung_n" json:"tipojung_n"`
	TipojungT    int32  `gorm:"column:tipojung_t" json:"tipojung_t"`
	TipojungF    int32  `gorm:"column:tipojung_f" json:"tipojung_f"`
	TipojungJ    int32  `gorm:"column:tipojung_j" json:"tipojung_j"`
	TipojungP    int32  `gorm:"column:tipojung_p" json:"tipojung_p"`
	TipojungKode string `gorm:"column:tipojung_kode" json:"tipojung_kode"`
}

type SkorKarakteristikPribadi struct {
	IDUser                  int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz                  int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	PribadiMotivasi         int32  `gorm:"column:pribadi_motivasi" json:"pribadi_motivasi"`
	PribadiJuang            int32  `gorm:"column:pribadi_juang" json:"pribadi_juang"`
	PribadiYakin            int32  `gorm:"column:pribadi_yakin" json:"pribadi_yakin"`
	PribadiPercaya          int32  `gorm:"column:pribadi_percaya" json:"pribadi_percaya"`
	PribadiKonsep           int32  `gorm:"column:pribadi_konsep" json:"pribadi_konsep"`
	PribadiKreativitas      int32  `gorm:"column:pribadi_kreativitas" json:"pribadi_kreativitas"`
	PribadiMimpin           int32  `gorm:"column:pribadi_mimpin" json:"pribadi_mimpin"`
	PribadiEntrepreneur     int32  `gorm:"column:pribadi_entrepreneur" json:"pribadi_entrepreneur"`
	PribadiStress           int32  `gorm:"column:pribadi_stress" json:"pribadi_stress"`
	PribadiEmosi            int32  `gorm:"column:pribadi_emosi" json:"pribadi_emosi"`
	PribadiSosial           int32  `gorm:"column:pribadi_sosial" json:"pribadi_sosial"`
	PribadiEmpati           int32  `gorm:"column:pribadi_empati" json:"pribadi_empati"`
	KlasifikasiMotivasi     string `gorm:"column:klasifikasi_motivasi" json:"klasifikasi_motivasi"`
	KlasifikasiJuang        string `gorm:"column:klasifikasi_juang" json:"klasifikasi_juang"`
	KlasifikasiYakin        string `gorm:"column:klasifikasi_yakin" json:"klasifikasi_yakin"`
	KlasifikasiPercaya      string `gorm:"column:klasifikasi_percaya" json:"klasifikasi_percaya"`
	KlasifikasiKonsep       string `gorm:"column:klasifikasi_konsep" json:"klasifikasi_konsep"`
	KlasifikasiKreativitas  string `gorm:"column:klasifikasi_kreativitas" json:"klasifikasi_kreativitas"`
	KlasifikasiMimpin       string `gorm:"column:klasifikasi_mimpin" json:"klasifikasi_mimpin"`
	KlasifikasiEntrepreneur string `gorm:"column:klasifikasi_entrepreneur" json:"klasifikasi_entrepreneur"`
	KlasifikasiStress       string `gorm:"column:klasifikasi_stress" json:"klasifikasi_stress"`
	KlasifikasiEmosi        string `gorm:"column:klasifikasi_emosi" json:"klasifikasi_emosi"`
	KlasifikasiSosial       string `gorm:"column:klasifikasi_sosial" json:"klasifikasi_sosial"`
	KlasifikasiEmpati       string `gorm:"column:klasifikasi_empati" json:"klasifikasi_empati"`
}

type SkorMinatIndonesium struct {
	IDUser        int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz        int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	TmiIlmuAlam   int32  `gorm:"column:tmi_ilmu_alam" json:"tmi_ilmu_alam"`
	TmiIlmuSosial int32  `gorm:"column:tmi_ilmu_sosial" json:"tmi_ilmu_sosial"`
	TmiRentang    int32  `gorm:"column:tmi_rentang" json:"tmi_rentang"`
	RekomTmi      string `gorm:"column:rekom_tmi" json:"rekom_tmi"`
	Minat1        int32  `gorm:"column:minat1" json:"minat1"`
	Minat2        int32  `gorm:"column:minat2" json:"minat2"`
	Minat3        int32  `gorm:"column:minat3" json:"minat3"`
	Minat4        int32  `gorm:"column:minat4" json:"minat4"`
	Minat5        int32  `gorm:"column:minat5" json:"minat5"`
	Minat6        int32  `gorm:"column:minat6" json:"minat6"`
	Minat7        int32  `gorm:"column:minat7" json:"minat7"`
}

type RefKecerdasanMajemuk struct {
	No             string `gorm:"column:no;primaryKey" json:"no"`
	NamaKecerdasan string `gorm:"column:nama_kecerdasan" json:"nama_kecerdasan"`
	NamaKecil      string `gorm:"column:nama_kecil" json:"nama_kecil"`
	Icon           string `gorm:"column:icon" json:"icon"`
}

type RefSkoringKecerdasanMajemuk struct {
	IDSkoringKecerdasanMajemuk int32  `gorm:"column:id_skoring_kecerdasan_majemuk;primaryKey;autoIncrement:true" json:"id_skoring_kecerdasan_majemuk"`
	IDQuiz                     int32  `gorm:"column:id_quiz;comment:ID_QUIZ" json:"id_quiz"` // ID_QUIZ
	IDUser                     int32  `gorm:"column:id_user;comment:ID_USER" json:"id_user"` // ID_USER
	No                         string `gorm:"column:no;" json:"no"`
	B1                         int32  `gorm:"column:b1" json:"b1"`
	B2                         int32  `gorm:"column:b2" json:"b2"`
	B3                         int32  `gorm:"column:b3" json:"b3"`
	B4                         int32  `gorm:"column:b4" json:"b4"`
	B5                         int32  `gorm:"column:b5" json:"b5"`
	B6                         int32  `gorm:"column:b6" json:"b6"`
	B7                         int32  `gorm:"column:b7" json:"b7"`
	B8                         int32  `gorm:"column:b8" json:"b8"`
	B9                         int32  `gorm:"column:b9" json:"b9"`
	Total                      int32  `gorm:"column:total" json:"total"`
	Rangking                   int32  `gorm:"column:rangking" json:"rangking"`
}

type SkorKecerdasanMajemuk struct {
	IDUser int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	Km1    string `gorm:"column:km_1" json:"km_1"`
	Km2    string `gorm:"column:km_2" json:"km_2"`
	Km3    string `gorm:"column:km_3" json:"km_3"`
	Km4    string `gorm:"column:km_4" json:"km_4"`
	Km5    string `gorm:"column:km_5" json:"km_5"`
}

type SkorSuasanaKerja struct {
	IDUser        int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz        int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	SuasanaKerja1 string `gorm:"column:suasana_kerja1;default:0" json:"suasana_kerja1"`
	SuasanaKerja2 string `gorm:"column:suasana_kerja2;default:0" json:"suasana_kerja2"`
	SuasanaKerja3 string `gorm:"column:suasana_kerja3;default:0" json:"suasana_kerja3"`
}

type SkorGayaBelajar struct {
	IDUser                int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz                int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	GayaAuditoris         int16  `gorm:"column:gaya_auditoris" json:"gaya_auditoris"`
	GayaVisual            int16  `gorm:"column:gaya_visual" json:"gaya_visual"`
	GayaKinestetik        int16  `gorm:"column:gaya_kinestetik" json:"gaya_kinestetik"`
	KlasifikasiAuditoris  string `gorm:"column:klasifikasi_auditoris" json:"klasifikasi_auditoris"`
	KlasifikasiVisual     string `gorm:"column:klasifikasi_visual" json:"klasifikasi_visual"`
	KlasifikasiKinestetik string `gorm:"column:klasifikasi_kinestetik" json:"klasifikasi_kinestetik"`
}

type SkorKejiwaanDewasa struct {
	IDUser int32 `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz int32 `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`

	SkorDepresi     int32 `gorm:"column:skor_depresi" json:"skor_depresi"`
	SkorKecemasan   int32 `gorm:"column:skor_kecemasan" json:"skor_kecemasan"`
	SkorManipulatif int32 `gorm:"column:skor_manipulatif" json:"skor_manipulatif"`
	SkorParanoid    int32 `gorm:"column:skor_paranoid" json:"skor_paranoid"`
	SkorPsikopat    int32 `gorm:"column:skor_psikopat" json:"skor_psikopat"`
	SkorScizopernia int32 `gorm:"column:skor_scizopernia" json:"skor_scizopernia"`
	SkorHisteria    int32 `gorm:"column:skor_histeria" json:"skor_histeria"`
	SkorHipokondria int32 `gorm:"column:skor_hipokondria" json:"skor_hipokondria"`
	SkorHipomania   int32 `gorm:"column:skor_hipomania" json:"skor_hipomania"`
	SkorIntroversi  int32 `gorm:"column:skor_introversi" json:"skor_introversi"`

	NilaiDepresi     int32 `gorm:"column:nilai_depresi" json:"nilai_depresi"`
	NilaiKecemasan   int32 `gorm:"column:nilai_kecemasan" json:"nilai_kecemasan"`
	NilaiManipulatif int32 `gorm:"column:nilai_manipulatif" json:"nilai_manipulatif"`
	NilaiParanoid    int32 `gorm:"column:nilai_paranoid" json:"nilai_paranoid"`
	NilaiPsikopat    int32 `gorm:"column:nilai_psikopat" json:"nilai_psikopat"`
	NilaiScizopernia int32 `gorm:"column:nilai_scizopernia" json:"nilai_scizopernia"`
	NilaiHisteria    int32 `gorm:"column:nilai_histeria" json:"nilai_histeria"`
	NilaiHipokondria int32 `gorm:"column:nilai_hipokondria" json:"nilai_hipokondria"`
	NilaiHipomania   int32 `gorm:"column:nilai_hipomania" json:"nilai_hipomania"`
	NilaiIntroversi  int32 `gorm:"column:nilai_introversi" json:"nilai_introversi"`
}

type SkorHitungNilaiFieldSkoring struct {
	Skor         int16  `json:"skor"`
	Nilai        int16  `json:"nilai"`
	FieldSkoring string `json:"field_skoring"`
}

type SkorKesehatanMental struct {
	IDUser               int32 `gorm:"column:id_user;not null" json:"id_user"`
	IDQuiz               int32 `gorm:"column:id_quiz;not null" json:"id_quiz"`
	SkorStressKehidupan  int32 `gorm:"column:skor_stress_kehidupan" json:"skor_stress_kehidupan"`
	SkorOverThinking     int32 `gorm:"column:skor_over_thinking" json:"skor_over_thinking"`
	SkorAdiksiMedsos     int32 `gorm:"column:skor_adiksi_medsos" json:"skor_adiksi_medsos"`
	SkorImpulsiveBuying  int32 `gorm:"column:skor_impulsive_buying" json:"skor_impulsive_buying"`
	SkorGangguanMood     int32 `gorm:"column:skor_gangguan_mood" json:"skor_gangguan_mood"`
	SkorGangguanMakan    int32 `gorm:"column:skor_gangguan_makan" json:"skor_gangguan_makan"`
	SkorPenampilanTubuh  int32 `gorm:"column:skor_penampilan_tubuh" json:"skor_penampilan_tubuh"`
	SkorKecemasanBicara  int32 `gorm:"column:skor_kecemasan_bicara" json:"skor_kecemasan_bicara"`
	SkorPanicAttack      int32 `gorm:"column:skor_panic_attack" json:"skor_panic_attack"`
	SkorBipolarDisorder  int32 `gorm:"column:skor_bipolar_disorder" json:"skor_bipolar_disorder"`
	SkorAdiksiZat        int32 `gorm:"column:skor_adiksi_zat" json:"skor_adiksi_zat"`
	NilaiStressKehidupan int32 `gorm:"column:nilai_stress_kehidupan" json:"nilai_stress_kehidupan"`
	NilaiOverThinking    int32 `gorm:"column:nilai_over_thinking" json:"nilai_over_thinking"`
	NilaiAdiksiMedsos    int32 `gorm:"column:nilai_adiksi_medsos" json:"nilai_adiksi_medsos"`
	NilaiImpulsiveBuying int32 `gorm:"column:nilai_impulsive_buying" json:"nilai_impulsive_buying"`
	NilaiGangguanMood    int32 `gorm:"column:nilai_gangguan_mood" json:"nilai_gangguan_mood"`
	NilaiGangguanMakan   int32 `gorm:"column:nilai_gangguan_makan" json:"nilai_gangguan_makan"`
	NilaiPenampilanTubuh int32 `gorm:"column:nilai_penampilan_tubuh" json:"nilai_penampilan_tubuh"`
	NilaiKecemasanBicara int32 `gorm:"column:nilai_kecemasan_bicara" json:"nilai_kecemasan_bicara"`
	NilaiPanicAttack     int32 `gorm:"column:nilai_panic_attack" json:"nilai_panic_attack"`
	NilaiBipolarDisorder int32 `gorm:"column:nilai_bipolar_disorder" json:"nilai_bipolar_disorder"`
	NilaiAdiksiZat       int32 `gorm:"column:nilai_adiksi_zat" json:"nilai_adiksi_zat"`
}

type SkorModeBelajar struct {
	IDUser        int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz        int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	IDModeBelajar int32  `gorm:"column:id_mode_belajar" json:"id_mode_belajar"`
	Prioritas1    string `gorm:"column:prioritas_1" json:"prioritas_1"`
	Prioritas2    string `gorm:"column:prioritas_2" json:"prioritas_2"`
	Prioritas3    string `gorm:"column:prioritas_3" json:"prioritas_3"`
	Prioritas4    string `gorm:"column:prioritas_4" json:"prioritas_4"`
	Prioritas5    string `gorm:"column:prioritas_5" json:"prioritas_5"`
}

type SkorSsct struct {
	IDUser      int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz      int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	Urutan      int32  `gorm:"column:urutan;primaryKey" json:"urutan"`
	Skor        int32  `gorm:"column:skor" json:"skor"`
	Klasifikasi string `gorm:"column:klasifikasi" json:"klasifikasi"`
}

type SkorRekomKuliahA struct {
	IDQuiz            int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	IDUser            int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	RekomKuliahAlam   string `gorm:"column:rekom_kuliah_alam" json:"rekom_kuliah_alam"`
	RekomKuliahSosial string `gorm:"column:rekom_kuliah_sosial" json:"rekom_kuliah_sosial"`
	RekomKuliahDinas  string `gorm:"column:rekom_kuliah_dinas" json:"rekom_kuliah_dinas"`
}

type SkorRekomKuliahB struct {
	IDQuiz            int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	IDUser            int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	RekomKuliahAlam   string `gorm:"column:rekom_kuliah_alam" json:"rekom_kuliah_alam"`
	RekomKuliahSosial string `gorm:"column:rekom_kuliah_sosial" json:"rekom_kuliah_sosial"`
	RekomKuliahDinas  string `gorm:"column:rekom_kuliah_dinas" json:"rekom_kuliah_dinas"`
	RekomKuliahAgama  string `gorm:"column:rekom_kuliah_agama" json:"rekom_kuliah_agama"`
}

type SkorRekomPeminatanSma struct {
	IDUser              int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz              int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	RekomMinat          string `gorm:"column:rekom_minat" json:"rekom_minat"`
	RekomSikapPelajaran string `gorm:"column:rekom_sikap_pelajaran" json:"rekom_sikap_pelajaran"`
	RekomMapel          string `gorm:"column:rekom_mapel" json:"rekom_mapel"`
}

type RefRekomendasiAkhirPeminatanSma struct {
	ID                  int32  `gorm:"column:id;primaryKey" json:"id"`
	RekomMinat          string `gorm:"column:rekom_minat" json:"rekom_minat"`
	RekomSikapPelajaran string `gorm:"column:rekom_sikap_pelajaran" json:"rekom_sikap_pelajaran"`
	RekomAkhir          string `gorm:"column:rekom_akhir" json:"rekom_akhir"`
}

type SkorKategori struct {
	Kategori    string `json:"kategori"`
	Klasifikasi string `json:"klasifikasi"`
	Skor        int8   `json:"skor"`
}

type SkorDisc struct {
	IDUser       int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz       int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	SkorD        int32  `gorm:"column:skor_d" json:"skor_d"`
	SkorI        int32  `gorm:"column:skor_i" json:"skor_i"`
	SkorS        int32  `gorm:"column:skor_s" json:"skor_s"`
	SkorC        int32  `gorm:"column:skor_c" json:"skor_c"`
	KlasifikasiD string `gorm:"column:klasifikasi_d" json:"klasifikasi_d"`
	KlasifikasiS string `gorm:"column:klasifikasi_s" json:"klasifikasi_s"`
	KlasifikasiI string `gorm:"column:klasifikasi_i" json:"klasifikasi_i"`
	KlasifikasiC string `gorm:"column:klasifikasi_c" json:"klasifikasi_c"`
}

type SkorModeKerja struct {
	IDUser      int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz      int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	IDModeKerja int32  `gorm:"column:id_mode_kerja;primaryKey" json:"id_mode_kerja"`
	Prioritas1  string `gorm:"column:prioritas_1" json:"prioritas_1"`
	Prioritas2  string `gorm:"column:prioritas_2" json:"prioritas_2"`
	Prioritas3  string `gorm:"column:prioritas_3" json:"prioritas_3"`
	Prioritas4  string `gorm:"column:prioritas_4" json:"prioritas_4"`
	Prioritas5  string `gorm:"column:prioritas_5" json:"prioritas_5"`
}

type SkorKepribadianManajerial struct {
	IDUser                          int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz                          int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	Visioner                        int32  `gorm:"column:visioner" json:"visioner"`
	KestabilanEmosi                 int32  `gorm:"column:kestabilan_emosi" json:"kestabilan_emosi"`
	Memotivasi                      int32  `gorm:"column:memotivasi" json:"memotivasi"`
	ManajemenResiko                 int32  `gorm:"column:manajemen_resiko" json:"manajemen_resiko"`
	PengambilanKeputusan            int32  `gorm:"column:pengambilan_keputusan" json:"pengambilan_keputusan"`
	PenyesuaianDiri                 int32  `gorm:"column:penyesuaian_diri" json:"penyesuaian_diri"`
	ManajemenWaktu                  int32  `gorm:"column:manajemen_waktu" json:"manajemen_waktu"`
	MotivasiPrestasi                int32  `gorm:"column:motivasi_prestasi" json:"motivasi_prestasi"`
	Integritas                      int32  `gorm:"column:integritas" json:"integritas"`
	PelayananPublik                 int32  `gorm:"column:pelayanan_publik" json:"pelayanan_publik"`
	KlasifikasiVisioner             string `gorm:"column:klasifikasi_visioner" json:"klasifikasi_visioner"`
	KlasifikasiKestabilanEmosi      string `gorm:"column:klasifikasi_kestabilan_emosi" json:"klasifikasi_kestabilan_emosi"`
	KlasifikasiMemotivasi           string `gorm:"column:klasifikasi_memotivasi" json:"klasifikasi_memotivasi"`
	KlasifikasiManajemenResiko      string `gorm:"column:klasifikasi_manajemen_resiko" json:"klasifikasi_manajemen_resiko"`
	KlasifikasiPengambilanKeputusan string `gorm:"column:klasifikasi_pengambilan_keputusan" json:"klasifikasi_pengambilan_keputusan"`
	KlasifikasiPenyesuaianDiri      string `gorm:"column:klasifikasi_penyesuaian_diri" json:"klasifikasi_penyesuaian_diri"`
	KlasifikasiManajemenWaktu       string `gorm:"column:klasifikasi_manajemen_waktu" json:"klasifikasi_manajemen_waktu"`
	KlasifikasiMotivasiPrestasi     string `gorm:"column:klasifikasi_motivasi_prestasi" json:"klasifikasi_motivasi_prestasi"`
	KlasifikasiIntegritas           string `gorm:"column:klasifikasi_integritas" json:"klasifikasi_integritas"`
	KlasifikasiPelayananPublik      string `gorm:"column:klasifikasi_pelayanan_publik" json:"klasifikasi_pelayanan_publik"`
}

type SkorWlb struct {
	IDUser                int32 `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz                int32 `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	KedamaianHati         int32 `gorm:"column:kedamaian_hati" json:"kedamaian_hati"`
	PengembanganDiri      int32 `gorm:"column:pengembangan_diri" json:"pengembangan_diri"`
	Ibadah                int32 `gorm:"column:ibadah" json:"ibadah"`
	Pendapatan            int32 `gorm:"column:pendapatan" json:"pendapatan"`
	HubunganSosial        int32 `gorm:"column:hubungan_sosial" json:"hubungan_sosial"`
	Kesehatan             int32 `gorm:"column:kesehatan" json:"kesehatan"`
	RekanKerja            int32 `gorm:"column:rekan_kerja" json:"rekan_kerja"`
	NilaiKedamaianHati    int32 `gorm:"column:nilai_kedamaian_hati" json:"nilai_kedamaian_hati"`
	NilaiPengembanganDiri int32 `gorm:"column:nilai_pengembangan_diri" json:"nilai_pengembangan_diri"`
	NilaiIbadah           int32 `gorm:"column:nilai_ibadah" json:"nilai_ibadah"`
	NilaiPendapatan       int32 `gorm:"column:nilai_pendapatan" json:"nilai_pendapatan"`
	NilaiHubunganSosial   int32 `gorm:"column:nilai_hubungan_sosial" json:"nilai_hubungan_sosial"`
	NilaiKesehatan        int32 `gorm:"column:nilai_kesehatan" json:"nilai_kesehatan"`
	NilaiRekanKerja       int32 `gorm:"column:nilai_rekan_kerja" json:"nilai_rekan_kerja"`
}
