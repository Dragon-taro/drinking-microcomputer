package entity

import (
	"time"
)

type Data struct {
	CreatedAt   time.Time `json:"createdAt"`
	Amount      int `json:"amount"`
	TotalAmount int `json:"totalAmount"`
	Diff        int `json:"diff"`
}

type DataReqest struct {
	Amount int `validate:"required"`
}
