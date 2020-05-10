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
	if !db.HasTable(&domain.TrackComment{}) {
		db.CreateTable(&domain.TrackComment{})
		//db.Model(&domain.TrackComment{}).Related(&domain.Track{})
	}

	Constructor(db)
	return db, nil
}

func Constructor(db *gorm.DB) error {
	Users := []domain.User{}
	Artists := []domain.Artist{}
	UserArtistTags := []domain.UserArtistTag{}
	Tracks := []domain.Track{}
	UserTrackTags := []domain.UserTrackTag{}
	UserChats := []domain.UserChat{}
	UserTokens := []domain.UserToken{}
	TrackComments := []domain.TrackComment{}

	//db.Find(&[]domain.User{})
	db.Find(&Users)
	if len(Users) == 0 {
		db.Create(&TestUser)
	}
	db.Find(&Artists)
	if len(Artists) == 0 {
		db.Create(&TestArtist)
	}
	db.Find(&UserArtistTags)
	if len(UserArtistTags) == 0 {
		db.Create(&TestUserArtistTag)
	}
	db.Find(&Tracks)
	if len(Tracks) == 0 {
		db.Create(&TestTrack)
	}
	db.Find(&UserTrackTags)
	if len(UserTrackTags) == 0 {
		db.Create(&TestUserTrackTag)
	}
	db.Find(&UserChats)
	if len(UserChats) == 0 {
		db.Create(&TestUserChat)
	}
	db.Find(&UserTokens)
	if len(UserTokens) == 0 {
		db.Create(&TestUserToken)
	}
	db.Find(&TrackComments)
	if len(TrackComments) == 0 {
		//db.Model(&TestTrackComment).Related(&TestTrack)
		//trackIn := domain.Track{}
		//db.First(&trackIn)
		//TestTrackComment.Track = trackIn
		db.Create(&TestTrackComment)
	}
	return nil
}
