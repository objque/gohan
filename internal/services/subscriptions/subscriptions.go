package subscriptions

import (
	"errors"

	"github.com/objque/gohan/internal/guard"
	"github.com/objque/gohan/internal/log"
	repo "github.com/objque/gohan/internal/repositories/subscriptions"
)

type Service struct{}

func New() *Service {
	return new(Service)
}

func (s *Service) GetSubscriptions(opts *repo.GetSubscriptionsOpts) ([]*repo.Subscription, error) {
	// TODO (m.kalinin): replace mock with database call
	subs := []*repo.Subscription{
		{
			ID: 100, //nolint:gomnd
			Artist: repo.Artist{
				ID:     123456,                       //nolint:gomnd
				Name:   "Skrillex",                   //nolint:gomnd
				Poster: "http://posters.io/skrillex", //nolint:gomnd
			},
		},
	}

	return subs, nil
}

func (s *Service) CreateSubscription(subscription *repo.Subscription) error {
	subscription.ID = 100
	subscription.Artist.Name = "Skrillex"
	subscription.Artist.Poster = "http://posters.io/skrillex"
	log.Debug("subscription created")

	return nil
}

func (s *Service) DeleteSubscription(subscription *repo.Subscription) error {
	return guard.NewInternalError(errors.New("not implemented")) //nolint:err113
}
