package main

import (
	"fmt"
	"irwanka/sicerdas/controller"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/joho/godotenv"
)

var (
	tokenAuth        *jwtauth.JWTAuth
	userController   controller.UserController   = controller.NewUserController()
	quizController   controller.QuizController   = controller.NewQuizController()
	uploadController controller.UploadController = controller.NewUploadController()
	reportController controller.ReportController = controller.NewReportController()
	dummyController  controller.DummyController  = controller.NewDummyController()
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	signKey := os.Getenv("JWT_SIGN_KEY")
	tokenAuth = jwtauth.New("HS512", []byte(signKey), nil)

}

func main() {
	fmt.Println("Golang is Here!!")
	os.Setenv("TZ", "Asia/Jakarta")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(120 * time.Second))
	r.Post("/login", userController.SubmitLogin)
	//protected route
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/me", userController.GetMe)
		r.Post("/mobile-change-password", userController.ChangePasswordMobile)
		r.Get("/get-list-quiz-user", quizController.GetListQuizUser)
		r.Get("/get-detil-quiz-user/{token}", quizController.GetDetilQuizUser)
		r.Get("/get-salam-pembuka/{token}", quizController.GetSalamPembuka)
		r.Get("/get-list-info-cerdas", userController.GetListInfoCerdas)
		r.Get("/get-list-session-quiz/{token}", quizController.GetListQuizSessionInfo)
		r.Get("/get-detail-quiz/{token}", quizController.GetQuizDetil)
		r.Post("/upload-quiz-to-firebase/{token}", quizController.UploadQuizJsonToFirebase)
		r.Post("/upload-gambar-to-firebase/{filename}", uploadController.UploadGambarToFirebase)
		//sinkrob gambar
		r.Post("/sinkron-gambar-quiz-firebase", uploadController.SinkronGambarQuizTemplateToFirebase)
		r.Post("/sinkron-gambar-quiz-with-template", quizController.SinkronGambarQuizWithTemplate)
		r.Post("/sinkron-gambar-info-cerdas", uploadController.SinkronGambarInfoCerdasToFirebase)
		r.Post("/submit-jawaban-quiz/{token}", quizController.SubmitJawabanQuiz)

		//dummy generate
		r.Post("/cek-dummy-template/{id}", dummyController.CekExistDummyForTemplate)
		r.Post("/create-dummy-template/{id}", dummyController.CreateDummyForTemplate)
		r.Post("/export-report-peserta/{id_quiz}/{id_user}/{id_model}", reportController.ExportReportPDFToFirebase)
	})

	//render
	//report
	fileAsset := http.FileServer(http.Dir("./templates/assets/"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileAsset))

	fileKop := http.FileServer(http.Dir("/var/www/html/public/kop/"))
	r.Handle("/kop/*", http.StripPrefix("/kop/", fileKop))
	fileCover := http.FileServer(http.Dir("/var/www/html/public/cover/"))
	r.Handle("/cover/*", http.StripPrefix("/cover/", fileCover))
	fileGambar := http.FileServer(http.Dir("/var/www/html/public/gambar/"))
	r.Handle("/gambar/*", http.StripPrefix("/gambar/", fileGambar))
	fileImages := http.FileServer(http.Dir("/var/www/html/public/images/"))
	r.Handle("/images/*", http.StripPrefix("/images/", fileImages))

	fileIcon := http.FileServer(http.Dir("./templates/assets/icon"))
	r.Handle("/icon/*", http.StripPrefix("/icon/", fileIcon))

	r.Get("/preview-report-dummy/{uuid}", reportController.PreviewKomponenReportDummy)
	r.Get("/preview-lampiran-dummy/{uuid}", reportController.PreviewLampiranReportDummy)
	r.Get("/render-report-utama/{id_quiz}/{id_user}/{id_model}/{nomor_seri}", reportController.RenderReportUtama)
	r.Get("/render-report-lampiran/{id_quiz}/{id_user}/{id_report}/{nomor_seri}", reportController.RenderReportLampiran)
	r.Get("/render-cover/{id_quiz}/{id_user}", reportController.RenderCoverReport)
	http.ListenAndServe(port, r)
}
