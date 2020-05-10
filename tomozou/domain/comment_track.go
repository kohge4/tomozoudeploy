package domain

import (
	"time"
)

type TrackComment struct {
	ID        int       `gorm:"primary_key;column:id" json:"id"`
	UserID    int       `gorm:"column:user_id;not null" json:"user_id"`
	TrackID   int       `gorm:"column:track_id;not null" json:"track_id"`
	Comment   string    `gorm:"column:comment;not null" json:"comment"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

type TrackCommentFull struct {
	ID        int       `gorm:"primary_key;column:id" json:"id"`
	UserID    int       `gorm:"column:user_id;not null" json:"user_id"`
	TrackID   int       `gorm:"column:track_id;not null" json:"track_id"`
	Comment   string    `gorm:"column:comment;not null" json:"comment"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`

	SocialID   string
	Name       string
	ArtistName string
	ArtistID   int
}
