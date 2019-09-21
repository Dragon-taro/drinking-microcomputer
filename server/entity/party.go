package entity

import (
	"time"
)

type Party struct {
	CreatedAt  time.Time `json:"createdAt`
	FinishedAt time.Time `json:"finishedAt"`
	IsFinished bool `json:"isFinished"`
}
