package itemrepoimpl

import (
	"tomozou/domain"

	"github.com/jinzhu/gorm"
)

// SpotifyHanlder が 構造体 に もつ リポジトリ
type ItemRepositoryImpl struct {
	DB *gorm.DB
}

func NewItemRepositoryImpl(db *gorm.DB) domain.ItemRepository {
	return &ItemRepositoryImpl{
		DB: db,
	}
}
