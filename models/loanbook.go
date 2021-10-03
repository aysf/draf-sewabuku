package models

import "time"

type LoanBook struct {
	ID         uint `gorm:"primaryKey"`
	BookDataID uint
	BookData   BookData
	CreatedAt  time.Time
	MaxReturn  time.Time
}
