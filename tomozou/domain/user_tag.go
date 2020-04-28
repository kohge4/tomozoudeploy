package domain

import "time"

type UserArtistTag struct {
	ID int `gorm:"column:id;not null;AUTO_INCREMENT" json:"id"`

	UserID    int       `gorm:"column:user_id;not null" json:"user_id"`
	ArtistID  int       `gorm:"column:artist_id;not null" json:"artist_id"`
	TagName   string    `gorm:"column:tag_name;not null" json:"tag_name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	//ArtistComment string    `gorm:"column:comment" json:"comment"`

	ArtistName string `gorm:"column:artist_name" json:"artist_name"`
	URL        string `gorm:"column:url" json:"url"`
	Image      string `gorm:"column:image" json:"image"`
}

type UserArtistTagIn struct {
	ID int `gorm:"column:id;not null;AUTO_INCREMENT" json:"id"`

	UserID    int       `gorm:"column:user_id;not null" json:"user_id"`
	ArtistID  int       `gorm:"column:artist_id;not null" json:"artist_id"`
	TagName   string    `gorm:"column:tag_name;not null" json:"tag_name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func NewUserArtistTag(userID int, artistID int, tagName string) UserArtistTag {
	return UserArtistTag{
		UserID:   userID,
		ArtistID: artistID,
		TagName:  tagName,
	}
}

type UserTrackTag struct {
	ID int `gorm:"column:id;not null;AUTO_INCREMENT" json:"id"`

	UserID    int       `gorm:"column:user_id;not null" json:"user_id"`
	TrackID   int       `gorm:"column:track_id;not null" json:"track_id"`
	ArtistID  int       `gorm:"column:artist_id;not null" json:"artist_id"`
	TagName   string    `gorm:"column:tag_name;not null" json:"tag_name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	//TrackComment string    `gorm:"column:track_comment" json:"track_comment"`

	ArtistName    string `gorm:"column:artist_name" json:"artist_name"`
	TrackName     string `gorm:"column:track_name" json:"track_name"`
	TrackSocialID string `gorm:"column:track_social_id" json:"track_social_id"`
}

type UserTrackTagIn struct {
	ID int `gorm:"column:id;not null;AUTO_INCREMENT" json:"id"`

	UserID    int       `gorm:"column:user_id;not null" json:"user_id"`
	ArtistID  int       `gorm:"column:artist_id;not null" json:"artist_id"`
	TagName   string    `gorm:"column:tag_name;not null" json:"tag_name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func NewUserTrackTag(track *Track, userID int, tagName string) *UserTrackTag {
	return &UserTrackTag{
		UserID:        userID,
		TrackID:       track.ID,
		ArtistID:      track.ArtistID,
		TagName:       tagName,
		ArtistName:    track.ArtistName,
		TrackName:     track.Name,
		TrackSocialID: track.SocialID,
	}
}
