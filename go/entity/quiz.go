package entity

import (
	"time"
)

type Quiz struct {
	IDQuiz               int32     `gorm:"column:id_quiz;primaryKey;autoIncrement:true" json:"id_quiz"`
	Token                string    `gorm:"column:token" json:"token"`
	NamaSesi             string    `gorm:"column:nama_sesi" json:"nama_sesi"`
	Lokasi               string    `gorm:"column:lokasi" json:"lokasi"`
	Tanggal              time.Time `gorm:"column:tanggal" json:"tanggal"`
	Open                 int32     `gorm:"column:open;default:1" json:"open"`
	Gambar               string    `gorm:"column:gambar" json:"gambar"`
	SkoringTabel         string    `gorm:"column:skoring_tabel" json:"skoring_tabel"`
	IDQuizTemplate       int32     `gorm:"column:id_quiz_template" json:"id_quiz_template"`
	IDUserBiro           int32     `gorm:"column:id_user_biro;comment:Id User BIRO" json:"id_user_biro"` // Id User BIRO
	Kota                 string    `gorm:"column:kota" json:"kota"`
	NamaAsesor           string    `gorm:"column:nama_asesor" json:"nama_asesor"`
	NomorSipp            string    `gorm:"column:nomor_sipp" json:"nomor_sipp"`
	UUID                 string    `gorm:"column:uuid" json:"uuid"`
	TtdAsesor            string    `gorm:"column:ttd_asesor" json:"ttd_asesor"`
	IDLokasi             int16     `gorm:"column:id_lokasi" json:"id_lokasi"`
	Jenis                string    `gorm:"column:jenis" json:"jenis"`
	FilenameReportZip    string    `gorm:"column:filename_report_zip;default:NULL" json:"filename_report_zip"`
	JSONURL              string    `gorm:"column:json_url" json:"json_url"`
	JSONURLEncrypt       string    `gorm:"column:json_url_encrypt" json:"json_url_encrypt"`
	CoverTemplate        string    `gorm:"column:cover_template;default:'default.pdf'" json:"cover_template"`
	FilenameReportZipDoc string    `gorm:"column:filename_report_zip_doc" json:"filename_report_zip_doc"`
	ModelReport          string    `gorm:"column:model_report" json:"model_report"`
}

type QuizUserApi struct {
	Open        int    `json:"open"`
	Token       string `json:"token"`
	NamaSesi    string `json:"nama_sesi"`
	Lokasi      string `json:"lokasi"`
	Tanggal     string `json:"tanggal"`
	Gambar      string `json:"gambar"`
	Submit      int    `json:"submit"`
	TokenSubmit string `json:"token_submit"`
	StatusHasil int    `json:"status_hasil"`
	UrlResult   string `json:"url_result"`
	JsonUrl     string `json:"json_url"`
}

type QuizSesiMaster struct {
	IDSesiMaster   int32  `gorm:"column:id_sesi_master;primaryKey;autoIncrement:true" json:"id_sesi_master"`
	Kategori       string `gorm:"column:kategori" json:"kategori"`
	NamaSesiUjian  string `gorm:"column:nama_sesi_ujian" json:"nama_sesi_ujian"`
	Soal           string `gorm:"column:soal" json:"soal"`
	Mode           string `gorm:"column:mode" json:"mode"`
	Jawaban        int32  `gorm:"column:jawaban;comment:Jumlah Jawaban Ynag harus diisi / dikoreksi" json:"jawaban"` // Jumlah Jawaban Ynag harus diisi / dikoreksi
	PetunjukSesi   string `gorm:"column:petunjuk_sesi" json:"petunjuk_sesi"`
	MetodeSkoring  string `gorm:"column:metode_skoring" json:"metode_skoring"`
	UUID           string `gorm:"column:uuid" json:"uuid"`
	PanjangJawaban int32  `json:"panjang_jawaban"`
}

type QuizSesiDetil struct {
	IDQuizSesi   int32 `gorm:"column:id_quiz_sesi;primaryKey;autoIncrement:true" json:"id_quiz_sesi"`
	IDQuiz       int32 `gorm:"column:id_quiz" json:"id_quiz"`
	IDSesiMaster int32 `gorm:"column:id_sesi_master" json:"id_sesi_master"`
	Urutan       int32 `gorm:"column:urutan" json:"urutan"`
	Durasi       int32 `gorm:"column:durasi" json:"durasi"`
	KunciWaktu   int32 `gorm:"column:kunci_waktu" json:"kunci_waktu"`
}

type QuizSesiUser struct {
	IDQuizUser     int32     `json:"id_quiz_user"`
	IDUser         int32     `gorm:"column:id_user" json:"id_user"`
	IDQuiz         int32     `gorm:"column:id_quiz" json:"id_quiz"`
	StartAt        time.Time `gorm:"column:start_at" json:"start_at"`
	Submit         int32     `gorm:"column:submit" json:"submit"`
	SubmitAt       time.Time `gorm:"column:submit_at" json:"submit_at"`
	TokenSubmit    string    `gorm:"column:token_submit" json:"token_submit"`
	Jawaban        string    `gorm:"column:jawaban;comment:berisi JSON jawaban quiz" json:"jawaban"`                     // berisi JSON jawaban quiz
	StatusHasil    int32     `gorm:"column:status_hasil;comment:1=>Sudah Publish, 0=>Belum Publish" json:"status_hasil"` // 1=>Sudah Publish, 0=>Belum Publish
	SkoringAt      time.Time `gorm:"column:skoring_at" json:"skoring_at"`
	Skoring        int32     `gorm:"column:skoring;comment:1=>sudah skoring, 0=>Belum Skoring" json:"skoring"`                                 // 1=>sudah skoring, 0=>Belum Skoring
	JawabanSkoring string    `gorm:"column:jawaban_skoring;comment:Konversi Data Tabel quiz_user_sesi_jawaban ke JSON" json:"jawaban_skoring"` // Konversi Data Tabel quiz_user_sesi_jawaban ke JSON
	Saran          string    `gorm:"column:saran;comment:Ambil Dari Template Berdasarkan Nama Tabel Skoring" json:"saran"`                     // Ambil Dari Template Berdasarkan Nama Tabel Skoring
	UUID           string    `gorm:"column:uuid" json:"uuid"`
	NoSeri         string    `gorm:"column:no_seri" json:"no_seri"`
}

type QuizSesiInfo struct {
	Token          string   `json:"token"`
	Urutan         int      `json:"urutan"`
	NamaSesiUjian  string   `json:"nama_sesi_ujian"`
	Mode           string   `json:"mode"`
	Durasi         int      `json:"durasi"`
	KunciWaktu     int      `json:"kunci_waktu"`
	Kategori       string   `json:"kategori"`
	Play           int      `json:"play"`
	Finish         int      `json:"finish"`
	Soal           string   `json:"soal"`
	Jawaban        int      `json:"jawaban"`
	PanjangJawaban int      `json:"panjang_jawaban"`
	PetunjukSesi   string   `json:"petunjuk_sesi"`
	Sections       []string `json:"sections"`
}

type QuizFirebaseStorage struct {
	Session []*QuizSesiInfo `json:"session"`
	Soal    []*SoalSession  `json:"soal"`
}

type QuizSesiTemplate struct {
	IDQuizTemplate int32  `json:"id_quiz_template"`
	NamaSesi       string `json:"nama_sesi"`
	Gambar         string `json:"gambar"`
	SkoringTabel   string `json:"skoring_tabel"`
	UUID           string `json:"uuid"`
	Jenis          string `json:"jenis"` // demo atao quiz
}
