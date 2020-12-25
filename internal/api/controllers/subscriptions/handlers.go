package subscriptions

import (
	"net/http"

	"github.com/objque/gohan/internal/api/errors"
	"github.com/objque/gohan/internal/api/httputils"
	repo "github.com/objque/gohan/internal/repositories/subscriptions"
)

const defaultSubscriptionsLimit = 100

type Controller struct {
	repository repo.Repository
}

func New(repository repo.Repository) *Controller {
	return &Controller{repository: repository}
}

func (c *Controller) List(w http.ResponseWriter, r *http.Request) {
	opts := repo.GetSubscriptionsOpts{
		UserName: httputils.GetUserName(r),
		Limit: defaultSubscriptionsLimit,
	}

	if values, exists := r.URL.Query()["offset"]; exists {
		rawOffset := values[0]
		offset, err := httputils.GetUint16FromQuery(rawOffset)
		if err != nil {
			httputils.WriteClientError(w, errors.NewWrongQueryValueError("uint16", "offset", rawOffset))
			return
		}

		opts.Offset = offset
	}

	subscriptions, err := c.repository.GetSubscriptions(&opts)
	if err != nil {
		httputils.WriteInternalError(w, err)
		return
	}

	_ = httputils.WriteJSON(w, http.StatusOK, subscriptions)
}
