package models

import "time"

type Transfers struct {
	ID            uint
	ToAccountId   string
	FromAccountId string
	Amount        int
	CreatedAt     time.Time
}
