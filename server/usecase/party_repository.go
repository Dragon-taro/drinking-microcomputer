package usecase

import (
	"github.com/Dragon-taro/drinking-microcomputer/server/entity"
)

type PartyRepository interface {
	Store(entity.Party) error
	// FindAll() ([]entity.Party, error)
	FindLatest() (entity.Party, error)
}
