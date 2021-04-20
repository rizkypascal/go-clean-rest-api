package repositories

import (
	"fmt"
	"reflect"

	"github.com/rizkypascal/go-clean-rest-api/interfaces"
	"github.com/rizkypascal/go-clean-rest-api/models"
)

type UserRepository struct {
	interfaces.IDbHandler
}

func (repository *UserRepository) Create(attributes map[string]interface{}) error {
	query := "INSERT INTO users (email, address, password) VALUES(?, ?, ?)"
	_, err := repository.PrepareAndExec(query, attributes["email"], attributes["address"], attributes["password"])
	if err != nil {
		return err
	}
	return nil
}

func (repository *UserRepository) Update(attributes map[string]interface{}) error {
	query := "UPDATE users SET"

	idx := 1
	userId := fmt.Sprintf("%v", attributes["id"])
	delete(attributes, "id")

	if len(attributes) > 0 {
		for key, element := range attributes {
			elementStr := fmt.Sprintf("%v", element)
			if reflect.ValueOf(element).Kind() == reflect.String {
				elementStr = fmt.Sprintf("'%v'", elementStr)
			}
			if idx == len(attributes) {
				query += " " + key + "=" + elementStr
			} else {
				query += " " + key + "=" + elementStr + ", "
			}
			idx++
		}
	}

	query += " WHERE id=" + userId

	_, err := repository.Execute(query)

	if err != nil {
		return err
	}
	return nil
}

func (repository *UserRepository) FetchUsers() ([]*models.User, error) {
	query := "SELECT id, email, address FROM users"

	rows, err := repository.Query(query)

	if err != nil {
		return []*models.User{}, err
	}

	list := make([]*models.User, 0)

	for rows.Next() {
		data := new(models.User)

		err := rows.Scan(&data.ID, &data.Email, &data.Address)

		if err != nil {
			return nil, err
		}
		list = append(list, data)
	}

	return list, nil
}

func (repository *UserRepository) FetchUser(attributes map[string]interface{}) (models.User, error) {
	query := "SELECT * FROM users"

	idx := 1
	if len(attributes) > 0 {
		query += " WHERE"
		for key, element := range attributes {
			elementStr := fmt.Sprintf("%v", element)
			if reflect.ValueOf(element).Kind() == reflect.String {
				elementStr = fmt.Sprintf("'%v'", elementStr)
			}
			if idx == len(attributes) {
				query += " " + key + "=" + elementStr
			} else {
				query += " " + key + "=" + elementStr + " AND"
			}
			idx++
		}
	}

	row, err := repository.Query(query)

	if err != nil {
		return models.User{}, err
	}

	var user models.User

	row.Next()
	err = row.Scan(&user.ID, &user.Email, &user.Address, &user.Password)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (repository *UserRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := repository.PrepareAndExec(query, id)
	if err != nil {
		return err
	}
	return nil
}
