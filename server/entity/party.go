package entity

import (
	"time"
)

type Party struct {
	CreatedAt  time.Time
	FinishedAt time.Time
	IsFinished bool
}
