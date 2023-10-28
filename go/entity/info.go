package entity

import "time"

type InfoCerdas struct {
	IDInfo    int32     `json:"id_info"`
	Judul     string    `json:"judul"`
	Isi       string    `json:"isi"`
	Gambar    string    `json:"gambar"`
	CreatedAt time.Time `json:"created_at"`
	UUID      string    `json:"uuid"`
	URL       string    `json:"url"`
}
