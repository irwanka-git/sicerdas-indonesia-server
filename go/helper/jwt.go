package helper

import (
	"irwanka/sicerdas/entity"
	"log"
	"os"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func CreateJWTTokenLogin(user *entity.User, expire bool) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	signKey := os.Getenv("JWT_SIGN_KEY")
	iss := os.Getenv("URL_ISS")

	tokenAuth := jwtauth.New("HS512", []byte(signKey), nil)

	sub := user.UUID
	jti := uuid.NewString()
	claims := map[string]interface{}{
		"sub":      sub,
		"jti":      jti,
		"iss":      iss,
		"name":     user.NamaPengguna,
		"username": user.Username,
	}
	if expire {
		jwtauth.SetExpiryIn(claims, time.Hour*12)
	}
	jwtauth.SetIssuedAt(claims, time.Now())
	_, tokenString, _ := tokenAuth.Encode(claims)

	return tokenString, nil
}
