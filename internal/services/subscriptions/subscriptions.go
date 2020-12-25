package subscriptions

import (
	"errors"

	"github.com/objque/gohan/internal/guard"
	repo "github.com/objque/gohan/internal/repositories/subscriptions"
)

type Service struct{}

func New() *Service {
	return new(Service)
}

func (s *Service) GetSubscriptions(opts *repo.GetSubscriptionsOpts) ([]*repo.Subscription, error) {
	return nil, guard.NewInternalError(errors.New("not implemented"))
}

func (s *Service) CreateSubscription(subscription *repo.Subscription) error {
	return guard.NewInternalError(errors.New("not implemented"))
}

func (s *Service) DeleteSubscription(subscription *repo.Subscription) error {
	return guard.NewInternalError(errors.New("not implemented"))
}
