package subscriptions

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/objque/gohan/internal/api/errors"
	"github.com/objque/gohan/internal/api/httputils"
	repo "github.com/objque/gohan/internal/repositories/subscriptions"
)

type contextKey string

//nolint:gochecknoglobals
var ctxKey contextKey = "subscription"

func getSubscriptionFromContext(ctx context.Context) *repo.Subscription {
	subscription, _ := ctx.Value(ctxKey).(*repo.Subscription)
	return subscription //nolint:nlreturn
}

func decodeBodyMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet || r.Method == http.MethodHead || r.Method == http.MethodOptions {
			next.ServeHTTP(w, r)
			return //nolint:nlreturn
		}

		ctx := r.Context()
		subscription := repo.Subscription{}
		err := json.NewDecoder(r.Body).Decode(&subscription)
		if err != nil {
			httputils.WriteClientError(w, errors.NewIncorrectBodyError("subscription"))
			return //nolint:nlreturn
		}

		ctx = context.WithValue(ctx, ctxKey, &subscription)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func validateBodyMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			next.ServeHTTP(w, r)
			return //nolint:nlreturn
		}

		if err := validateCreate(getSubscriptionFromContext(r.Context())); err != nil {
			httputils.WriteClientError(w, err)
			return //nolint:nlreturn
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
