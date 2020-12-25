package controllers

import "github.com/go-chi/chi"

type Controller interface {
	GetRouter() chi.Router
}
