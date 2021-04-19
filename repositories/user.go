package repositories

import (
	"github.com/rizkypascal/go-clean-rest-api/interfaces"
	"github.com/rizkypascal/go-clean-rest-api/models"
)

type UserRepository struct {
	interfaces.IDbHandler
}

func (repository *UserRepository) Create(u models.User) (int64, error) {
	query := "INSERT INTO users (email, address, password) VALUES(?, ?, ?)"
	result, err := repository.PrepareAndExec(query,
		u.Email, u.Address, u.Password)

	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func (repository *UserRepository) GetByEmail(email string) (models.User, error) {
	query := "SELECT * FROM users WHERE email = ?"
	row, err := repository.Query(query, email)

	if err != nil {
		return models.User{}, err
	}

	var user models.User

	row.Next()
	row.Scan(&user.ID, &user.Email, &user.Address, &user.Password)

	return user, nil
}
