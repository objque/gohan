package subscriptions

import (
	"errors"

	repo "github.com/objque/gohan/internal/repositories/subscriptions"
)

type Service struct{}

func New() *Service {
	return new(Service)
}

func (s *Service) GetSubscriptions(opts *repo.GetSubscriptionsOpts) ([]*repo.Subscription, error) {
	return nil, errors.New("not implemented")
}

func (s *Service) CreateSubscription(subscription *repo.Subscription) error {
	return errors.New("not implemented")
}

func (s *Service) DeleteSubscription(subscription *repo.Subscription) error {
	return errors.New("not implemented")
}
