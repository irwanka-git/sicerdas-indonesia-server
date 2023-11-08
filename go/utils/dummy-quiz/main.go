package main

import (
	"encoding/json"
	"fmt"
	"irwanka/sicerdas/entity"
	"irwanka/sicerdas/helper"
	"irwanka/sicerdas/repository"
	"irwanka/sicerdas/service"
	"os"
	"time"
)

func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	dummyRepo := repository.NewDummyQuizUserRepository()
	skoringRepo := repository.NewSkoringRepository()
	skoringService := service.NewSkoringService(skoringRepo)

	quiz, err := dummyRepo.CekDummyQuizUser(44)
	if err != nil {
		fmt.Println("Belum ada sesi dummy untuk template ini..")
		fmt.Println("Buat Sesi Dummy terlebih dahulu untuk jenis tes ini")
		return
	}
	id_quiz := quiz.IDQuiz
	//fmt.Println(id_quiz)
	randomJawaban, errJawaban := dummyRepo.GenerateJawaban(int(id_quiz))
	if errJawaban != nil {
		fmt.Println("Error gagal buat jawaban sesi dummy untuk template ini..")
		return
	}
	var jawabanString = []entity.JawabanString{}
	for i := 0; i < len(randomJawaban); i++ {
		//fmt.Println(randomJawaban[i].Jawaban)
		//fmt.Println(randomJawaban[i].Kategori)
		tmp := entity.JawabanString{}
		tmp.Kategori = randomJawaban[i].Kategori
		jws, _ := json.Marshal(randomJawaban[i].Jawaban)
		tmp.Jawaban = string(jws)
		jawabanString = append(jawabanString, tmp)
	}
	j, _ := json.Marshal(jawabanString)
	jsonString := string(j)
	// log.Println(string(jsonString))

	dummyRepo.SubmitJawabanUser(int(id_quiz), jsonString)

	kategori_tabel, _ := skoringRepo.GetTabelSkoring(id_quiz)
	skoringRepo.ClearTabelTemporaryJawabanUser(id_quiz, 100)
	skoringRepo.GenerateTabelTemporaryJawabanUser(id_quiz, 100)
	skoringService.SkoringAllKategori(kategori_tabel, id_quiz, 100)
	skoringRepo.ClearTabelTemporaryJawabanUser(id_quiz, 100)
	now := helper.StringTimeYMDHIS(time.Now())
	skoringRepo.StartRunningSkoring()
	skoringRepo.FinishSkoring(id_quiz, 100, now)
	skoringRepo.StopRunningSkoring()
}
