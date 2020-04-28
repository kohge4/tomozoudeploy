package userrepoimpl

import (
	"tomozou/domain"

	"github.com/jinzhu/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) domain.UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func NewDevUserRepo(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		DB: db,
	}
}
