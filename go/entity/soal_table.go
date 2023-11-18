package entity

type SoalKognitif struct {
	IDSoal           int32  `gorm:"column:id_soal;primaryKey;autoIncrement:true" json:"id_soal"`
	Bidang           string `gorm:"column:bidang" json:"bidang"`
	Urutan           int32  `gorm:"column:urutan" json:"urutan"`
	IDPetunjuk       int32  `gorm:"column:id_petunjuk" json:"id_petunjuk"`
	Pertanyaan       string `gorm:"column:pertanyaan" json:"pertanyaan"`
	PertanyaanGambar string `gorm:"column:pertanyaan_gambar" json:"pertanyaan_gambar"`
	PilihanA         string `gorm:"column:pilihan_a" json:"pilihan_a"`
	PilihanAGambar   string `gorm:"column:pilihan_a_gambar" json:"pilihan_a_gambar"`
	PilihanB         string `gorm:"column:pilihan_b" json:"pilihan_b"`
	PilihanBGambar   string `gorm:"column:pilihan_b_gambar" json:"pilihan_b_gambar"`
	PilihanC         string `gorm:"column:pilihan_c" json:"pilihan_c"`
	PilihanCGambar   string `gorm:"column:pilihan_c_gambar" json:"pilihan_c_gambar"`
	PilihanD         string `gorm:"column:pilihan_d" json:"pilihan_d"`
	PilihanDGambar   string `gorm:"column:pilihan_d_gambar" json:"pilihan_d_gambar"`
	PilihanE         string `gorm:"column:pilihan_e" json:"pilihan_e"`
	PilihanEGambar   string `gorm:"column:pilihan_e_gambar" json:"pilihan_e_gambar"`
	PilihanJawaban   string `gorm:"column:pilihan_jawaban" json:"pilihan_jawaban"`
	UUID             string `gorm:"column:uuid" json:"uuid"`
	Paket            string `gorm:"column:paket" json:"paket"`
}

type SoalMinatKuliahEksakta struct {
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

type SoalMinatKuliahSosial struct {
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

type SoalWlb struct {
	IDSoal   int16  `gorm:"column:id_soal;primaryKey" json:"id_soal"`
	IDModel  int16  `gorm:"column:id_model" json:"id_model"`
	Unsur    string `gorm:"column:unsur" json:"unsur"`
	Urutan   int16  `gorm:"column:urutan" json:"urutan"`
	Kategori string `gorm:"column:kategori;comment:P=> Mendukung, N=> Tidak Mendukung" json:"kategori"` // P=> Mendukung, N=> Tidak Mendukung
}
