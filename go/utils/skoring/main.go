package main

import (
	"fmt"
	"irwanka/sicerdas/helper"
	"irwanka/sicerdas/repository"
	"irwanka/sicerdas/service"
	"os"
	"time"
)

func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	skoringRepo := repository.NewSkoringRepository()
	skoringService := service.NewSkoringService(skoringRepo)
	cek, _ := skoringRepo.GetStatusRunningSkoring()
	if cek.Status == 1 {
		fmt.Println("Skoring sedang berjalan..")
		return
	}
	mulai := helper.StringTimeYMDHIS(time.Now())
	skoringRepo.StartRunningSkoring(mulai)
	list_belum_skoring, _ := skoringRepo.GetUserSesiBelumSkoring()
	for _, user := range list_belum_skoring {

		kategori_tabel, _ := skoringRepo.GetTabelSkoring(user.IDQuiz)
		skoringRepo.ClearTabelTemporaryJawabanUser(user.IDQuiz, user.IDUser)
		skoringRepo.GenerateTabelTemporaryJawabanUser(user.IDQuiz, user.IDUser)
		skoringService.SkoringAllKategori(kategori_tabel, user.IDQuiz, user.IDUser)
		skoringRepo.ClearTabelTemporaryJawabanUser(user.IDQuiz, user.IDUser)
		now := helper.StringTimeYMDHIS(time.Now())
		skoringRepo.FinishSkoring(user.IDQuiz, user.IDUser, now)
		text := fmt.Sprintf("[%v] \tID USER : %v   ID QUIZ : %v  Sesi: %v (%v)  Kota: %v", now, user.IDUser, user.IDQuiz, user.NamaSesi, user.IDQuizTemplate, user.Kota)
		fmt.Println(text)
	}
	selesai := helper.StringTimeYMDHIS(time.Now())
	skoringRepo.StopRunningSkoring(selesai, len(list_belum_skoring))
}
