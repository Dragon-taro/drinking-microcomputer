package usecase

import (
	"github.com/Dragon-taro/drinking-microcomputer/server/entity"
)

type DataRepository interface {
	Store(entity.Data) error
	FindAll() (entity.Data, error)
	FindLatest() (entity.Data, error)
}
