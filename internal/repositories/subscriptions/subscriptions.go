package subscriptions

type Artist struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Poster string `json:"poster"`
}

type Subscription struct {
	ID       uint64 `json:"id"`
	Artist   Artist `json:"artist"`
	UserName string `json:"-"`
}

type GetSubscriptionsOpts struct {
	UserName string
	Limit    uint16
	Offset   uint16
}

type Repository interface {
	GetSubscriptions(opts *GetSubscriptionsOpts) ([]*Subscription, error)
	CreateSubscription(subscription *Subscription) error
	DeleteSubscription(subscription *Subscription) error
}
