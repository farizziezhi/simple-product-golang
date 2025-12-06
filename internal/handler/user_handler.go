package handler

import (
	"net/http"

	"github.com/farizziezhi/layered/internal/domain"
	"github.com/farizziezhi/layered/internal/service"
	"github.com/farizziezhi/layered/pkg/helper"
)

type UserHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type UserHandlerImpl struct {
	Service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return &UserHandlerImpl{
		Service: service,
	}
}

func (handler *UserHandlerImpl) Register(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	helper.ParseBody(r, &user)

	createdUser, err := handler.Service.Register(r.Context(), user)
	if err != nil {
		helper.SendJSON(w, http.StatusInternalServerError, domain.Response{
			Message: "Gagal registrasi user",
			Data:    nil,
		})
		return
	}

	helper.SendJSON(w, http.StatusCreated, domain.Response{
		Message: "User berhasil diregistrasi",
		Data:    createdUser,
	})
}

func (handler *UserHandlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	helper.ParseBody(r, &user)

	token, err := handler.Service.Login(r.Context(), user)
	if err != nil {
		helper.SendJSON(w, http.StatusUnauthorized, domain.Response{
			Message: "Username atau password salah",
			Data:    nil,
		})
		return
	}

	helper.SendJSON(w, http.StatusOK, domain.Response{
		Message: "Login berhasil",
		Data:    map[string]string{"token": token},
	})
}
