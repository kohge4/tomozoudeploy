package chatrepoimpl

import (
	"tomozou/domain"

	"github.com/jinzhu/gorm"
)

type ChatDBRepository struct {
	DB *gorm.DB
}

func NewChatDBRepository(db *gorm.DB) domain.ChatRepository {
	return &ChatDBRepository{
		DB: db,
	}
}
