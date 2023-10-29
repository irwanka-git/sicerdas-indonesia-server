package main

import (
	"fmt"
	"irwanka/sicerdas/helper"
	"irwanka/sicerdas/repository"
	"irwanka/sicerdas/service"
)

func main() {
	sinkronRepo := repository.NewSinkronRepository()
	uploadRepo := repository.NewUploadFirebaseRepository()
	uploadService := service.NewUploadFirebaseRepository(uploadRepo)

	list, _ := sinkronRepo.GetAllSoalMinatKuliahEksakta()
	for _, soal := range list {
		if soal.Gambar != "" {
			filename := soal.Gambar
			path := fmt.Sprintf("/var/www/html/public/gambar/%v", filename)
			directory := "gambar"
			if helper.FileExists((path)) {
				url, errUplaod := uploadService.UploadGambarToFirebase(path, directory)
				if errUplaod == nil {
					fmt.Println(url)
					sinkronRepo.UpdateGambarSoalMinatKuliahEksakta(soal.IDSoal, url)
				}
			} else {
				fmt.Println("not found : " + path)
			}
		}
	}
}
