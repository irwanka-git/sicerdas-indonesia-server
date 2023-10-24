package repository

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

const (
	TIMESTAMP = "20060102150405"
)

type UploadFirebaseRepository interface {
	UploadFileToFirebase(src string, bucket string) (string, error)
}

func NewUploadFirebaseRepository() UploadFirebaseRepository {
	return &repo{}
}

const (
	YYYMMDD = "2006-01-02"
	HHII    = "15.04"
)

func (*repo) UploadFileToFirebase(src string, directory string) (string, error) {

	ctx := context.Background()

	now := time.Now().UTC().Local()
	filename := now.Format("200601021504") + "-" + uuid.NewString() + "-" + filepath.Base(src)

	bucket := os.Getenv("FIREBASE_BUCKET")
	object := directory + "/" + filename

	wc := storageClient.Bucket(bucket).Object(object).NewWriter(ctx)
	wc.Metadata = map[string]string{
		"firebaseStorageDownloadTokens": uuid.NewString(),
	}
	wc.PredefinedACL = "publicRead"

	file, errOpen := os.Open(src)
	if errOpen != nil {
		log.Panicln("File not found ")
	}
	_, errUpload := io.Copy(wc, file)
	if errUpload != nil {
		return "", errors.New("Gagal upload ke firebase")
	}
	// fmt.Println("Berhasil Upload File ke Firebase")
	result_url := fmt.Sprintf(os.Getenv("URL_FIREBASE_PREFIX"), bucket, directory, filename)
	// fmt.Println(result_url)
	errRemove := os.Remove(src)
	if errRemove != nil {
		return "", errors.New("Gagal hapus file asli")
	}
	// fmt.Println("Berhasil Hapus File Backup (Original)")
	defer wc.Close()
	return result_url, nil
}