package main

import (
	"fmt"
	"irwanka/sicerdas/helper"
	"irwanka/sicerdas/repository"
	"irwanka/sicerdas/service"
	"os"
	"os/exec"
)

func main() {
	reportRepository := repository.NewReportRepository()
	userRepository := repository.NewUserRepository()
	uploadRepository := repository.NewUploadFirebaseRepository()
	uploadService := service.NewUploadFirebaseRepository(uploadRepository)

	id_quiz := reportRepository.GetStatusForRunPublish()
	fmt.Println(id_quiz)
	if id_quiz > 0 {

		url_docker := os.Getenv("URL_HOST_DOCKER")
		quiz, _ := reportRepository.GetDetilQuizCetak(id_quiz)
		model, _ := reportRepository.GetModelReport(quiz.ModelReport)
		fmt.Println(model)
		nomor_seri := reportRepository.GenerateQRCodeNomorSeriCetak()

		users := reportRepository.GetIDUserForPublish(id_quiz)
		quizTemplate, _ := reportRepository.GetDetilQuizTemplate(int(quiz.IDQuizTemplate))
		//set running job
		reportRepository.SetRunPublishIDQuiz(id_quiz)

		for i := 0; i < len(users); i++ {
			id_user := users[i]
			fmt.Printf("ID USER: %v", id_user)
			fmt.Println()

			user, errUser := userRepository.GetDataUserById(id_user)
			if errUser == nil {
				//render cover
				path_export_cover := fmt.Sprintf("/app/export/pdf/%v.%v.cover.pdf", user.ID, quiz.Token)
				url_render_cover := fmt.Sprintf("%v/render-cover/%v/%v", url_docker, id_quiz, id_user)
				exCover := exec.Command("/app/export.sh", url_render_cover, path_export_cover, "Portrait", quizTemplate.Kertas, nomor_seri, "cover")
				_, err2 := exCover.Output()
				// fmt.Println(string(out))
				if err2 != nil {
					fmt.Println(err2.Error())
					return
				}
				//render repot utama
				url_render_utama := fmt.Sprintf("%v/render-report-utama/%v/%v/%v/%v", url_docker, id_quiz, id_user, model.ID, nomor_seri)
				path_export_utama := fmt.Sprintf("/app/export/pdf/%v.%v.%v.pdf", user.ID, quiz.Token, model.Direktori)
				ex := exec.Command("/app/export.sh", url_render_utama, path_export_utama, "Portrait", quizTemplate.Kertas, nomor_seri, "utama")
				// fmt.Println(ex.Args)
				_, err3 := ex.Output()
				// fmt.Println(string(out))
				if err3 != nil {
					fmt.Println(err3.Error())
					return
				}

				//render lampiran
				listLampiran := reportRepository.GetListLampiranReport(int(quizTemplate.IDQuizTemplate))
				var listExportLampiran = []string{}
				for i := 0; i < len(listLampiran); i++ {
					url_render_lampiran := fmt.Sprintf("%v/render-report-lampiran/%v/%v/%v/%v", url_docker, id_quiz, id_user, listLampiran[i].IDReport, nomor_seri)
					path_export_lampiran := fmt.Sprintf("/app/export/pdf/%v.%v.lampiran.%v.pdf", user.ID, quiz.Token, listLampiran[i].IDReport)
					// fmt.Println(url_render_lampiran)
					listExportLampiran = append(listExportLampiran, path_export_lampiran)
					exLampiran := exec.Command("/app/export.sh", url_render_lampiran, path_export_lampiran, listLampiran[i].Orientasi, quizTemplate.Kertas, nomor_seri, "lampiran")
					// fmt.Println(ex.Args)
					_, err4 := exLampiran.Output()
					if err4 != nil {
						fmt.Println(err4.Error())
						return
					}
				}

				//buat list pdf untuk pasing ke pdfunite
				var listPDF = []string{}
				listPDF = append(listPDF, path_export_cover)
				listPDF = append(listPDF, path_export_utama)
				for i := 0; i < len(listExportLampiran); i++ {
					listPDF = append(listPDF, listExportLampiran[i])
				}

				//merge cover + utama + lampiran => final;
				path_export_report_final := fmt.Sprintf("/app/export/pdf/%v-%v.pdf", helper.CleanNamaFileOnly(user.NamaPengguna), nomor_seri)
				listPDF = append(listPDF, path_export_report_final)
				exMergePDF := exec.Command("pdfunite", listPDF...)
				_, err5 := exMergePDF.Output()
				if err5 != nil {
					fmt.Println(err5.Error())
					return
				}

				//upload firebase
				directory := fmt.Sprintf("report-individu/%v/%v", quiz.Token, model.Direktori)
				url_firebase_result, err := uploadService.UploadReportPDFToFirebase(path_export_report_final, directory)
				if err != nil {
					fmt.Println(err.Error())
					return
				}

				for i := 0; i < len(listPDF); i++ {
					os.Remove(listPDF[i])
				}
				os.Remove(fmt.Sprintf("templates/assets/qrcode/%v.png", nomor_seri))
				//update no_seri
				reportRepository.UpdateNomorSeriCetak(id_quiz, id_user, nomor_seri, url_firebase_result)
				fmt.Printf("Upload: %v", url_firebase_result)
				fmt.Println()
			}
		}
		//set finihsh running job clear
		reportRepository.ClearPublishCronjobIDQUiz(id_quiz)
	}
}
