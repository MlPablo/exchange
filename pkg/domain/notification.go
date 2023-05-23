package domain

import "context"

type Notification struct {
	UserGroup string `json:"user_group"`
}

func DefaultNotification() *Notification {
	return &Notification{}
}

type INotificationService interface {
	Notify(ctx context.Context, n *Notification) error
}
