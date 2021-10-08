package server

import "github.com/go-chi/chi/v5"

type Resource interface {
	Routes() chi.Router
}
