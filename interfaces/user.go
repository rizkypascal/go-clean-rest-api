package interfaces

import "github.com/rizkypascal/go-clean-rest-api/models"

type IUserRepository interface {
	Create(attributes map[string]interface{}) error
	Update(attributes map[string]interface{}) error
	Delete(id int) error
	FetchUser(attributes map[string]interface{}) (models.User, error)
	FetchUsers() ([]*models.User, error)
}

type IUserService interface {
	CreateUser(userCreateReq models.UserCreateReq) error
	UpdateUser(userUpdateReq models.UserUpdateReq) error
	DeleteUser(id int) error
	AuthenticateUser(userLoginReq models.UserLoginReq) (string, error)
	FetchOneUser(id int) (models.UserProfileResp, error)
	FetchAllUsers() ([]models.UserProfileResp, error)
}
