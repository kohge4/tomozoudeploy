package domain

import "time"

// UserArtistTagFull
// やらなくていい気もする
type UserArtistTag struct {
	ID int `gorm:"column:id;not null;AUTO_INCREMENT" json:"id"`

	UserID        int       `gorm:"column:user_id;not null" json:"user_id"`
	ArtistID      int       `gorm:"column:artist_id;not null" json:"artist_id"`
	TagName       string    `gorm:"column:tag_name;not null" json:"tag_name"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	ArtistComment string    `gorm:"column:comment" json:"comment"`

	ArtistName string `gorm:"column:artist_name" json:"artist_name"`
	URL        string `gorm:"column:url" json:"url"`
	Image      string `gorm:"column:image" json:"image"`
}

func NewUserArtistTag(userID int, artistID int, tagName string) UserArtistTag {
	return UserArtistTag{
		UserID:   userID,
		ArtistID: artistID,
		TagName:  tagName,
	}
}

type UserTrackTag struct {
	ID      int `gorm:"column:id;not null;AUTO_INCREMENT" json:"id"`
	UserID  int `gorm:"column:user_id;not null" json:"user_id"`
	TrackID int `gorm:"column:track_id;not null" json:"track_id"`
	//ArtistID  int       `gorm:"column:artist_id;not null" json:"artist_id"`
	TagName   string    `gorm:"column:tag_name;not null" json:"tag_name"`
	CreatedAt time.Time `gorm:"column:tag_created_at" json:"tag_created_at"`
	Count     int       `gorm:"column:count;not null" json:"count"`
}

func NewUserTrackTag(track *Track, userID int, tagName string, count int) *UserTrackTag {
	return &UserTrackTag{
		UserID:  userID,
		TrackID: track.ID,
		TagName: tagName,
		Count:   count,
	}
}

type UserTrackTagFull struct {
	UserTrackTag
	User
	Track
}
