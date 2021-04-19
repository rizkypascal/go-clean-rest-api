package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rizkypascal/go-clean-rest-api/interfaces"
	"github.com/rizkypascal/go-clean-rest-api/models"
)

type UserController struct {
	interfaces.IUserService
}

func (controller *UserController) SignUp(w http.ResponseWriter, r *http.Request) {
	userCreateReq := models.UserCreateReq{}

	json.NewDecoder(r.Body).Decode(&userCreateReq)
	err := controller.CreateUser(userCreateReq)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	} else {
		respondwithJSON(w, http.StatusCreated, map[string]string{"message": "User Created Successfully"})
	}
}

func (controller *UserController) SignIn(w http.ResponseWriter, r *http.Request) {
	userLoginReq := models.UserLoginReq{}

	json.NewDecoder(r.Body).Decode(&userLoginReq)

	token, err := controller.AuthenticateUser(userLoginReq)

	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unaothorized")
	} else {
		respondwithJSON(w, http.StatusAccepted, map[string]string{"token": token})
	}

}
