package server

import "github.com/go-chi/chi/v5"

type containerResource struct {
	Router *chi.Mux
}

func NewContainerResource() Resource {
	c := containerResource{
		Router: chi.NewRouter(),
	}
	return c
}

func (cr containerResource) Routes() chi.Router	{
	return cr.Router
}
