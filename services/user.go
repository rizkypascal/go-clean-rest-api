package services

import (
	"encoding/json"

	"github.com/rizkypascal/go-clean-rest-api/interfaces"
	"github.com/rizkypascal/go-clean-rest-api/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	interfaces.IUserRepository
}

func (service *UserService) CreateUser(req models.UserCreateReq) error {
	var err error
	if err = req.Validate(); err != nil {
		return err
	}

	attributes := make(map[string]interface{})
	attributes["email"] = req.Email
	attributes["address"] = req.Address
	attributes["password"], err = getHash([]byte(req.Password))

	if err != nil {
		return err
	}

	if err = service.Create(attributes); err != nil {
		return err
	}

	return nil
}

func (service *UserService) UpdateUser(req models.UserUpdateReq) error {
	var err error
	if err = req.Validate(); err != nil {
		return err
	}

	attributes := make(map[string]interface{})

	jsonObj, _ := json.Marshal(req)
	json.Unmarshal(jsonObj, &attributes)

	attributes["id"] = req.ID

	if attributes["password"] != nil && attributes["password"] != "" {
		attributes["password"], err = getHash([]byte(req.Password))
	}

	if err != nil {
		return err
	}

	if err := service.Update(attributes); err != nil {
		return err
	}

	return nil
}

func (service *UserService) AuthenticateUser(req models.UserLoginReq) (string, error) {
	if err := req.Validate(); err != nil {
		return "", err
	}

	attributes := make(map[string]interface{})
	attributes["email"] = req.Email

	userDB, err := service.FetchUser(attributes)

	if err != nil {
		return "", err
	}

	reqPass := []byte(req.Password)
	dbPass := []byte(userDB.Password)

	err = bcrypt.CompareHashAndPassword(dbPass, reqPass)

	if err != nil {
		return "", err
	}

	token, err := GenerateJWT()

	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *UserService) FetchOneUser(id int) (models.UserProfileResp, error) {
	attributes := make(map[string]interface{})
	attributes["id"] = id

	user, err := service.FetchUser(attributes)

	var userProfile models.UserProfileResp

	if err != nil {
		return userProfile, err
	}

	userProfile.ID = user.ID
	userProfile.Email = user.Email
	userProfile.Address = user.Address

	return userProfile, nil
}

func (service *UserService) FetchAllUsers() ([]models.UserProfileResp, error) {
	users, err := service.FetchUsers()

	if err != nil {
		return nil, err
	}

	var userProfile []models.UserProfileResp
	var data models.UserProfileResp

	for _, element := range users {
		data.ID = element.ID
		data.Address = element.Address
		data.Email = element.Email

		userProfile = append(userProfile, data)
	}

	return userProfile, nil
}

func (service *UserService) DeleteUser(id int) error {
	err := service.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
