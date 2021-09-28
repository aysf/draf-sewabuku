package models

import "time"

type Transfers struct {
	ID            uint
	ToAccountId   uint
	FromAccountId uint
	Amount        uint
	CreatedAt     time.Time
}
