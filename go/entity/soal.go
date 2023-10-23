package entity

type SoalSession struct {
	Nomor           string   `json:"nomor"`
	Uuid            string   `json:"uuid"`
	Token           string   `json:"token"`
	Pernyataan      string   `json:"pernyataan"`
	PernyataanMulti []string `json:"pernyataan_multi"`
	Section         string   `json:"section"`

	Pilihan  []*PilihanJawaban `json:"pilihan"`
	Kategori string            `json:"kategori"`
	Gambar   string            `json:"gambar"`

	Sn1 string `json:"sn1"`
	Sn2 string `json:"sn2"`
	Sn3 string `json:"sn3"`
	Sp1 string `json:"sp1"`
	Sp2 string `json:"sp2"`
	Sp3 string `json:"sp3"`

	MaxSikap int `json:"max_sikap"`
	MinSikap int `json:"min_sikap"`

	Mode     string `json:"mode"`
	Petunjuk string `json:"petunjuk"`
}

type PilihanJawaban struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}
