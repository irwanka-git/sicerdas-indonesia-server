package controller

import (
	"encoding/json"
	"fmt"
	"irwanka/sicerdas/entity"
	"irwanka/sicerdas/helper"
	"irwanka/sicerdas/repository"
	"irwanka/sicerdas/service"
	"net/http"
	"os"

	"github.com/go-chi/jwtauth"
)

var (
	userRepository repository.UserRepository = repository.NewUserRepository()
	userService    service.UserService       = service.NewUserService(userRepository)
)

type UserController interface {
	SubmitLogin(w http.ResponseWriter, r *http.Request)
	GetMe(w http.ResponseWriter, r *http.Request)
	GetListInfoCerdas(w http.ResponseWriter, r *http.Request)
}

func NewUserController() UserController {
	return &controller{}
}

func (*controller) GetListInfoCerdas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	list, err := userService.GetTopTenInfoCerdas()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error(), Status: false})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.ResponseData{Status: true, Message: "", Data: list})
}

func (*controller) SubmitLogin(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	var credentials entity.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "Username dan Password Wajib Diisi", Status: false})
		return
	}

	userLogin, errLogin := userService.AuthLogin(&credentials)
	if errLogin != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errLogin.Error(), Status: false})
		return
	}
	var expire = true
	if credentials.Ref == os.Getenv("REF_MOBILE_LOGIN") {
		expire = false
	}

	access_token, errGenToken := helper.CreateJWTTokenLogin(userLogin, expire)
	if errGenToken != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errGenToken.Error()})
		return
	}
	if errGenToken != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errGenToken.Error()})
		return
	}

	json.NewEncoder(w).Encode(helper.ResponseData{
		Status:  true,
		Message: "Login Berhasil",
		Data: map[string]interface{}{
			"access_token":    access_token,
			"username":        userLogin.Username,
			"unit_organisasi": userLogin.UnitOrganisasi,
			"organisasi":      userLogin.Organisasi,
			"uuid":            userLogin.UUID,
			"nama_pengguna":   userLogin.NamaPengguna,
		},
	})
}

func (*controller) GetMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	_, claims, _ := jwtauth.FromContext(r.Context())

	sub := fmt.Sprintf("%v", claims["sub"])
	user, err := userService.GetlUserByUuid(sub)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.ResponseData{Status: true, Message: "", Data: user})
}
