package models

import "time"

type Cart struct {
	UserID     uint
	BookID     uint
	DateLoan   time.Time
	DateDue    time.Time
	DateReturn time.Time
	User       User
}
