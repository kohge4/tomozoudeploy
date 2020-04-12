package datastore

import (
	"tomozou/domain"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GormConn(driver string, dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	if !db.HasTable(&domain.User{}) {
		db.CreateTable(&domain.User{})
	}
	if !db.HasTable(&domain.Artist{}) {
		db.CreateTable(&domain.Artist{})
	}
	if !db.HasTable(&domain.UserArtistTag{}) {
		db.CreateTable(&domain.UserArtistTag{})
	}
	if !db.HasTable(&domain.Track{}) {
		db.CreateTable(&domain.Track{})
	}

	if !db.HasTable(&domain.UserTrackTag{}) {
		db.CreateTable(&domain.UserTrackTag{})
	}
	if !db.HasTable(&domain.UserChat{}) {
		db.CreateTable(&domain.UserChat{})
	}
	if !db.HasTable(&domain.UserToken{}) {
		db.CreateTable(&domain.UserToken{})
	}
	return db, nil
}
