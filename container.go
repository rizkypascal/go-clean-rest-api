package main

import (
	"sync"

	"github.com/rizkypascal/go-clean-rest-api/controllers"
	"github.com/rizkypascal/go-clean-rest-api/infrastructures"
	"github.com/rizkypascal/go-clean-rest-api/repositories"
	"github.com/rizkypascal/go-clean-rest-api/services"

	"database/sql"
)

type IContainer interface {
	InjectUserController() controllers.UserController
}

type kernel struct{}

func (k *kernel) InjectUserController() controllers.UserController {

	sqlConn, _ := sql.Open("sqlite3", "/var/tmp/user.db")
	sqliteHandler := &infrastructures.SQLiteHandler{}
	sqliteHandler.Conn = sqlConn

	userRepository := &repositories.UserRepository{IDbHandler: sqliteHandler}
	userService := &services.UserService{IUserRepository: userRepository}
	userController := controllers.UserController{IUserService: userService}

	return userController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func Container() IContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
