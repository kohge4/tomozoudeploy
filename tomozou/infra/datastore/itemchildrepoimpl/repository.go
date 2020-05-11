package itemchildrepoimpl

import (
	"tomozou/domain"

	"github.com/jinzhu/gorm"
)

// SpotifyHanlder が 構造体 に もつ リポジトリ
type ItemChildRepositoryImpl struct {
	DB *gorm.DB
}

func NewItemChildRepositoryImpl(db *gorm.DB) domain.ItemChildRepository {
	return &ItemChildRepositoryImpl{
		DB: db,
	}
}
