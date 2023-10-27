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
	tokenAuth      *jwtauth.JWTAuth
	userController controller.UserController = controller.NewUserController()
	quizController controller.QuizController = controller.NewQuizController()
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
		r.Get("/get-list-session-quiz/{token}", quizController.GetListQuizSessionInfo)
		r.Get("/get-list-quiz-user", quizController.GetListQuizSessionInfo)
		r.Post("/upload-quiz-to-firebase/{token}", quizController.UploadQuizJsonToFirebase)
	})

	http.ListenAndServe(port, r)
}
