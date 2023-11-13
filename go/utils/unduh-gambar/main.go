package main

import (
	"fmt"
	"io"
	"irwanka/sicerdas/repository"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	fileName    string
	fullURLFile string
)

func main() {
	err := os.RemoveAll("templates/assets/icon")
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	os.Mkdir("templates/assets/icon", os.ModePerm)

	repo := repository.NewReportRepository()

	//gambar sains
	gambarSains, _ := repo.GetReferensiKuliahIlmuAlam()
	for i := 0; i < len(gambarSains); i++ {
		fullURLFile = gambarSains[i].Gambar
		// Build fileName from fullPath
		fileURL, err := url.Parse(fullURLFile)
		if err != nil {
			log.Fatal(err)
		}
		path := fileURL.Path
		segments := strings.Split(path, "/")
		fileName = segments[len(segments)-1]

		// Create blank file
		var path_save = fmt.Sprintf("templates/assets/icon/%v", fileName)
		file, err := os.Create(path_save)
		if err != nil {
			log.Fatal(err)
		}
		client := http.Client{
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				r.URL.Opaque = r.URL.Path
				return nil
			},
		}
		// Put content on file
		resp, err := client.Get(fullURLFile)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		size, _ := io.Copy(file, resp.Body)
		defer file.Close()
		fmt.Printf("Downloaded a file %s with size %d", fileName, size)
		fmt.Println()
	}

	gambarSosial, _ := repo.GetReferensiKuliahIlmuSosial()
	for i := 0; i < len(gambarSosial); i++ {
		fullURLFile = gambarSosial[i].Gambar
		// Build fileName from fullPath
		fileURL, err := url.Parse(fullURLFile)
		if err != nil {
			log.Fatal(err)
		}
		path := fileURL.Path
		segments := strings.Split(path, "/")
		fileName = segments[len(segments)-1]

		// Create blank file
		var path_save = fmt.Sprintf("templates/assets/icon/%v", fileName)
		file, err := os.Create(path_save)
		if err != nil {
			log.Fatal(err)
		}
		client := http.Client{
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				r.URL.Opaque = r.URL.Path
				return nil
			},
		}
		// Put content on file
		resp, err := client.Get(fullURLFile)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		size, _ := io.Copy(file, resp.Body)
		defer file.Close()
		fmt.Printf("Downloaded a file %s with size %d", fileName, size)
		fmt.Println()
	}

}
