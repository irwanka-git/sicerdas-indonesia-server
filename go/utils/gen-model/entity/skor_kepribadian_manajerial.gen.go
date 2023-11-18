// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameSkorKepribadianManajerial = "skor_kepribadian_manajerial"

// SkorKepribadianManajerial mapped from table <skor_kepribadian_manajerial>
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

// TableName SkorKepribadianManajerial's table name
func (*SkorKepribadianManajerial) TableName() string {
	return TableNameSkorKepribadianManajerial
}
