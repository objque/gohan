package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/objque/gohan/internal/api/controllers/subscriptions"
	repo "github.com/objque/gohan/internal/repositories/subscriptions"
)

type RouterOpts struct {
	SubscriptionsRepository repo.Repository
}

func GetRouter(opts RouterOpts) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		// use logger inside /v1 sub-route only
		// to avoid logging of healthz/metrics requests
		r.Use(middleware.Logger)

		r.Mount("/subscriptions", subscriptions.New(opts.SubscriptionsRepository).GetRouter())
	})

	return r
}
