package subscriptions

import "github.com/go-chi/chi"

func (c *Controller) GetRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(decodeBodyMiddleware)
	r.Use(validateBodyMiddleware)

	r.Get("/", c.List)
	r.Post("/", c.Create)

	return r
}
