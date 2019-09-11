package entity

import (
	"time"
)

type Data struct {
	CreatedAt   time.Time
	Amount      int
	TotalAmount int
}
