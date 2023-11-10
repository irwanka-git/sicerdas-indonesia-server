package entity

type QuizSesiReport struct {
	IDReport       int32  `gorm:"column:id_report;not null" json:"id_report"`
	NamaReport     string `gorm:"column:nama_report" json:"nama_report"`
	TabelReferensi string `gorm:"column:tabel_referensi" json:"tabel_referensi"`
	Blade          string `gorm:"column:blade" json:"blade"`
	TabelTerkait   string `gorm:"column:tabel_terkait" json:"tabel_terkait"`
	Jenis          int32  `gorm:"column:jenis;comment:1=> utama, 2=> lampiran" json:"jenis"` // 1=> utama, 2=> lampiran
	Orientasi      string `json:"orientasi"`
	Urutan         int    `json:"urutan"`
}

type ModelReport struct {
	ID        string `gorm:"column:id;primaryKey" json:"id"`
	Nama      string `gorm:"column:nama" json:"nama"`
	Direktori string `gorm:"column:direktori" json:"direktori"`
}

type QuizReportKomponenUtama struct {
	TabelReferensi string `json:"tabel_referensi"`
	Blade          string `json:"blade"`
	Urutan         int    `json:"urutan"`
}

type QuizReportRender struct {
	TabelReferensi string
	Blade          string
	Urutan         int
	Skoring        any
}

type SaranReport struct {
	JudulSaran string `json:"judul_saran"`
	IsiSaran   string `json:"isi_saran"`
}

type QuizSesiReportAndTemplate struct {
	IDReport       int32  `json:"id_report"`
	NamaReport     string `json:"nama_report"`
	TabelReferensi string `json:"tabel_referensi"`
	Blade          string `json:"blade"`
	TabelTerkait   string `json:"tabel_terkait"`
	Jenis          int32  `json:"jenis"`
	Model          string `json:"model"`
	IDQuizTemplate int32  `json:"id_quiz_template"`
	Urutan         int    `json:"urutan"`
}

type RefKuliahAlam struct {
	IDSoal           int32  `gorm:"column:id_soal;primaryKey;autoIncrement:true" json:"id_soal"`
	Urutan           int32  `gorm:"column:urutan" json:"urutan"`
	Indikator        string `gorm:"column:indikator" json:"indikator"`
	Minat            string `gorm:"column:minat" json:"minat"`
	DeskripsiMinat   string `gorm:"column:deskripsi_minat" json:"deskripsi_minat"`
	Jurusan          string `gorm:"column:jurusan" json:"jurusan"`
	DeskripsiJurusan string `gorm:"column:deskripsi_jurusan" json:"deskripsi_jurusan"`
	Matakuliah       string `gorm:"column:matakuliah" json:"matakuliah"`
	PeluangKarier    string `gorm:"column:peluang_karier" json:"peluang_karier"`
	TersediaDi       string `gorm:"column:tersedia_di" json:"tersedia_di"`
	UUID             string `gorm:"column:uuid" json:"uuid"`
	IDKelompok       int16  `gorm:"column:id_kelompok;comment:SKORING PSIKOTES LENGKAP" json:"id_kelompok"` // SKORING PSIKOTES LENGKAP
	IDKelas          int16  `gorm:"column:id_kelas;comment:SKORING MINAT SMA V2" json:"id_kelas"`           // SKORING MINAT SMA V2
	Gambar           string `gorm:"column:gambar;default:'default.png'" json:"gambar"`
}

type RefKuliahSosial struct {
	IDSoal           int32  `gorm:"column:id_soal;primaryKey;autoIncrement:true" json:"id_soal"`
	Urutan           int32  `gorm:"column:urutan" json:"urutan"`
	Indikator        string `gorm:"column:indikator" json:"indikator"`
	Minat            string `gorm:"column:minat" json:"minat"`
	DeskripsiMinat   string `gorm:"column:deskripsi_minat" json:"deskripsi_minat"`
	Jurusan          string `gorm:"column:jurusan" json:"jurusan"`
	DeskripsiJurusan string `gorm:"column:deskripsi_jurusan" json:"deskripsi_jurusan"`
	Matakuliah       string `gorm:"column:matakuliah" json:"matakuliah"`
	PeluangKarier    string `gorm:"column:peluang_karier" json:"peluang_karier"`
	TersediaDi       string `gorm:"column:tersedia_di" json:"tersedia_di"`
	UUID             string `gorm:"column:uuid" json:"uuid"`
	IDKelompok       int16  `gorm:"column:id_kelompok;comment:MINAT PSIKOTES LENGKAP" json:"id_kelompok"` // MINAT PSIKOTES LENGKAP
	IDKelas          int16  `gorm:"column:id_kelas;comment:MINAT SMA V2" json:"id_kelas"`                 // MINAT SMA V2
	Gambar           string `gorm:"column:gambar;default:'default.png'" json:"gambar"`
}

type RefKuliahAgama struct {
	IDSoal    int32  `gorm:"column:id_soal;primaryKey;autoIncrement:true" json:"id_soal"`
	Urutan    int32  `gorm:"column:urutan" json:"urutan"`
	Indikator string `gorm:"column:indikator" json:"indikator"`
	Jurusan   string `gorm:"column:jurusan" json:"jurusan"`
	UUID      string `gorm:"column:uuid" json:"uuid"`
	Gambar    string `gorm:"column:gambar;default:'default.png'" json:"gambar"`
}

type RefSuasanaKerja struct {
	IDKegiatan int32  `gorm:"column:id_kegiatan;primaryKey;autoIncrement:true" json:"id_kegiatan"`
	Nomor      string `gorm:"column:nomor" json:"nomor"`
	Kegiatan   string `gorm:"column:kegiatan" json:"kegiatan"`
	Gambar     string `gorm:"column:gambar" json:"gambar"`
	Keterangan string `gorm:"column:keterangan" json:"keterangan"`
	Deskripsi  string `gorm:"column:deskripsi" json:"deskripsi"`
	UUID       string `gorm:"column:uuid" json:"uuid"`
}

type RefSikapPelajaran struct {
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

type ResultSikapPelajaran struct {
	Urutan      int32  `gorm:"column:urutan" json:"urutan"`
	Pelajaran   string `gorm:"column:pelajaran" json:"pelajaran"`
	Klasifikasi string `json:"klasifikasi"`
	Kelompok    string `json:"kelompok"`
}

type RefSikapPelajaranMK struct {
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

type ResultSikapPelajaranMK struct {
	Urutan      int32  `gorm:"column:urutan" json:"urutan"`
	Pelajaran   string `gorm:"column:pelajaran" json:"pelajaran"`
	Klasifikasi string `json:"klasifikasi"`
	Kelompok    string `json:"kelompok"`
}

type RefPilihanMinatSma struct {
	ID           int32  `json:"id"`
	KdPilihan    string `json:"kd_pilihan"`
	NamaPilihan  string `json:"nama_pilihan"`
	FieldSkoring string `json:"field_skoring"`
	Keterangan   string `json:"keterangan"`
	Gambar       string `json:"gambar"`
}

type RefPilihanMinatMan struct {
	ID           int32  `json:"id"`
	KdPilihan    string `json:"kd_pilihan"`
	NamaPilihan  string `json:"nama_pilihan"`
	FieldSkoring string `json:"field_skoring"`
	Keterangan   string `json:"keterangan"`
	Gambar       string `json:"gambar"`
}

type RefPilihanMinatTMI struct {
	IDSoal     int32  `gorm:"column:id_soal;primaryKey;autoIncrement:true" json:"id_soal"`
	Urutan     int32  `gorm:"column:urutan" json:"urutan"`
	Pernyataan string `gorm:"column:pernyataan" json:"pernyataan"`
	Kelompok   string `gorm:"column:kelompok" json:"kelompok"`
	Minat      string `gorm:"column:minat" json:"minat"`
	Keterangan string `gorm:"column:keterangan" json:"keterangan"`
	UUID       string `gorm:"column:uuid" json:"uuid"`
}

type ResultMinatTMI struct {
	Minat      string `gorm:"column:minat" json:"minat"`
	Keterangan string `gorm:"column:keterangan" json:"keterangan"`
}

type InterprestasiTipologiJung struct {
	IDInterprestasi int32  `gorm:"column:id_interprestasi;primaryKey;autoIncrement:true" json:"id_interprestasi"`
	Urutan          int32  `gorm:"column:urutan" json:"urutan"`
	Kode            string `gorm:"column:kode" json:"kode"`
	Nama            string `gorm:"column:nama" json:"nama"`
	Keterangan      string `gorm:"column:keterangan" json:"keterangan"`
	Deskripsi       string `gorm:"column:deskripsi" json:"deskripsi"`
	UUID            string `gorm:"column:uuid" json:"uuid"`
}

type RefKomponenKarakteristikPribadi struct {
	IDKomponen   int32  `gorm:"column:id_komponen;primaryKey;autoIncrement:true" json:"id_komponen"`
	NamaKomponen string `gorm:"column:nama_komponen" json:"nama_komponen"`
	Keterangan   string `gorm:"column:keterangan" json:"keterangan"`
	FieldSkoring string `gorm:"column:field_skoring" json:"field_skoring"`
	UUID         string `gorm:"column:uuid" json:"uuid"`
	Icon         string `gorm:"column:icon" json:"icon"`
}

type ResultKarakteristikPribadi struct {
	Urutan       int32  `json:"urutan"`
	NamaKomponen string `json:"nama_komponen"`
	Keterangan   string `json:"keterangan"`
	Klasifikasi  string `json:"klasifikasi"`
	Icon         string `gorm:"column:icon" json:"icon"`
}

type ResultGayaPekerjaan struct {
	ID            int32  `json:"id"`
	Rangking      int32  `json:"rangking"`
	Kode          string `json:"kode"`
	NamaKomponen  string `json:"nama_komponen"`
	CetakKomponen string `json:"cetak_komponen"`
	Skor          int    `json:"skor"`
	Klasifikasi   string `json:"klasifikasi"`
	Deskripsi     string `json:"deskripsi"`
	Pekerjaan     string `json:"pekerjaan"`
}

type ResultGayaBelajar struct {
	Kode            string `json:"kode"`
	Nama            string `json:"nama"`
	Skor            int    `json:"skor"`
	Klasifikasi     string `json:"klasifikasi"`
	Deskripsi       string `json:"deskripsi"`
	Gambar          string `json:"gambar"`
	FieldName       string `json:"field_name"`
	KlasifikasiName string `json:"klasifikasi_name"`
}

type ResultPeminatanSMK struct {
	Urutan     int    `json:"urutan"`
	Nomor      string `json:"nomor"`
	Keterangan string `json:"keterangan"`
	Deskripsi  string `json:"deskripsi"`
	Gambar     string `json:"gambar"`
}

type QuizSesiMappingSmk struct {
	IDQuiz     int32  `gorm:"column:id_quiz" json:"id_quiz"`
	IDKegiatan int32  `gorm:"column:id_kegiatan;" json:"id_kegiatan"`
	UUID       string `gorm:"column:uuid" json:"uuid"`
}

type ResultModeBelajar struct {
	Urutan    int    `json:"urutan"`
	Deskripsi string `json:"deskripsi"`
	Suasana   string `json:"suasana"`

	PilihanA string `json:"pilihan_a"`
	PilihanB string `json:"pilihan_b"`
	PilihanC string `json:"pilihan_c"`
	PilihanD string `json:"pilihan_d"`
	PilihanE string `json:"pilihan_e"`

	P1 string `json:"p1"`
	P2 string `json:"p2"`
	P3 string `json:"p3"`
	P4 string `json:"p4"`
	P5 string `json:"p5"`
}

type RefModelKesehatanMental struct {
	ID           int32  `gorm:"column:id;primaryKey" json:"id"`
	Nama         string `gorm:"column:nama" json:"nama"`
	FieldSkoring string `gorm:"column:field_skoring" json:"field_skoring"`
}

type ResultSkorKesehatanMental struct {
	Id          int32  `json:"id"`
	Nama        string `json:"nama"`
	Klasifikasi int    `json:"klasifikasi"`
	Skor        int    `json:"skor"`
}
