package main

import (
	"sync"

	"github.com/go-chi/chi"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {
	userController := Container().InjectUserController()

	r := chi.NewRouter()
	r.Post("/users", userController.SignUp)
	r.Post("/users/login", userController.SignIn)
	r.Get("/users/{id:[0-9]+}", userController.Get)
	r.Get("/users", userController.GetUsers)
	r.Patch("/users/{id:[0-9]+}", userController.Update)
	r.Delete("/users/{id:[0-9]+}", userController.Delete)

	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
