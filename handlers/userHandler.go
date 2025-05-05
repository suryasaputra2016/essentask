package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/suryasaputra2016/essentask/model"
	"github.com/suryasaputra2016/essentask/repo"
	"github.com/suryasaputra2016/essentask/utils"
)

type UserHandler struct {
	ur *repo.UserRepo
}

func NewUserHandler(ur *repo.UserRepo) *UserHandler {
	return &UserHandler{ur: ur}
}

func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user model.UserRegisterRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("register user: malformed json")
		http.Error(w, "malformed json", http.StatusBadRequest)
		return
	}

	err = utils.CheckEmailFormat(user.Email)
	if err == nil {
		log.Println("register user: email not well formatted")
		http.Error(w, "email is not well formatted", http.StatusBadRequest)
		return
	}

	_, err = uh.ur.GetByEmail(user.Email)
	if err == nil {
		log.Println("register user: email used")
		http.Error(w, "email is already in used", http.StatusBadRequest)
		return
	}

	if user.Password == "" {
		log.Println("register user: password empty")
		http.Error(w, "password is empty", http.StatusBadRequest)
		return
	}

	passwordHash, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Printf("register user: %v\n", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	newUser := model.User{
		Name:         user.Email,
		Email:        user.Email,
		PasswordHash: passwordHash,
	}

	err = uh.ur.Create(&newUser)
	if err != nil {
		log.Printf("register user: %v\n", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	response := model.UserRegisterResponse{
		Message: "user successfully created",
		ID:      newUser.ID,
		Name:    newUser.Name,
		Email:   newUser.Email,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("register user: %v\n", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
