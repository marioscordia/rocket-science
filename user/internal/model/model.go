package model

import "time"

type User struct {
	ID                  string
	Username            string
	Email               string
	Password            string
	NotificationMethods []NotificationMethod
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type NotificationMethod struct {
	ProviderName string
	Target       string
}

type Session struct {
	ID        string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time
}
