package interfaces

import "github.com/rizkypascal/go-clean-rest-api/models"

type IUserRepository interface {
	Create(user models.User) (int64, error)
	GetByEmail(email string) (models.User, error)
}

type IUserService interface {
	CreateUser(userCreateReq models.UserCreateReq) error
	AuthenticateUser(userLoginReq models.UserLoginReq) (string, error)
}
