package subscriptions

import (
	"errors"

	"github.com/objque/gohan/internal/repositories/subscriptions"
)

var errEmptyArtistID = errors.New("artist.id not provided")

func validateCreate(subscription *subscriptions.Subscription) error {
	if subscription.Artist.ID == 0 {
		return errEmptyArtistID
	}

	return nil
}
