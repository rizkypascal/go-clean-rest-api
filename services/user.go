package services

import (
	"github.com/rizkypascal/go-clean-rest-api/interfaces"
	"github.com/rizkypascal/go-clean-rest-api/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	interfaces.IUserRepository
}

func (service *UserService) CreateUser(req models.UserCreateReq) error {
	if err := req.Validate(); err != nil {
		return err
	}

	var user models.User
	user.Email = req.Email
	user.Address = req.Address
	user.Password = getHash([]byte(req.Password))

	if _, err := service.Create(user); err != nil {
		return err
	}

	return nil
}

func (service *UserService) AuthenticateUser(req models.UserLoginReq) (string, error) {
	if err := req.Validate(); err != nil {
		return "", err
	}

	userDB, err := service.GetByEmail(req.Email)

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
