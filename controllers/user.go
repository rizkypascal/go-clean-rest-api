package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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

func (controller *UserController) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	result, err := controller.FetchOneUser(id)

	if err != nil {
		respondWithError(w, http.StatusNotFound, "User not found")
	} else {
		respondwithJSON(w, http.StatusAccepted, result)
	}
}

func (controller *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := controller.FetchAllUsers()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	} else {
		respondwithJSON(w, http.StatusAccepted, result)
	}
}

func (controller *UserController) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	userUpdateReq := models.UserUpdateReq{}
	userUpdateReq.ID = id
	json.NewDecoder(r.Body).Decode(&userUpdateReq)

	err := controller.UpdateUser(userUpdateReq)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	} else {
		respondwithJSON(w, http.StatusCreated, map[string]string{"message": "User Updated Successfully"})
	}
}

func (controller *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := controller.DeleteUser(id)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	} else {
		respondwithJSON(w, http.StatusAccepted, map[string]string{"message": "User Deleted Successfully"})
	}
}

func (controller *UserController) SignIn(w http.ResponseWriter, r *http.Request) {
	userLoginReq := models.UserLoginReq{}

	json.NewDecoder(r.Body).Decode(&userLoginReq)

	token, err := controller.AuthenticateUser(userLoginReq)

	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusUnauthorized, "Unauthorized")
	} else {
		respondwithJSON(w, http.StatusAccepted, map[string]string{"token": token})
	}

}
