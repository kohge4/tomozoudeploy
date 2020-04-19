package userpost

import "time"

// これ全部ドメインに書こう
// UserPostModel と UserPostRepository に変更する(chatIn とかを)
type UserChat struct {
	ID            int       `gorm:"not null;AUTO_INCREMENT" json:"id"`
	Comment       string    `gorm:"column:comment;not null" json:"comment"`
	UserID        int       `gorm:"column:user_id;not null" json:"user_id"`
	ReceiveUserID int       `gorm:"column:receive_user_id;not null" json:"receive_user_id"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`

	// グループチャットの時
	UserChatID int `gorm:"AUTO_INCREMENT" json:"id"`
}

type ArtistComment struct {
	ID        int       `gorm:"not null;AUTO_INCREMENT" json:"id"`
	Comment   string    `gorm:"column:comment;not null" json:"comment"`
	UserID    int       `gorm:"column:user_id;not null" json:"user_id"`
	ArtistID  int       `gorm:"column:artist_id;not null" json:"artist_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

type TrackComment struct {
	ID        int       `gorm:"not null;AUTO_INCREMENT" json:"id"`
	Comment   string    `gorm:"column:comment;not null" json:"comment"`
	UserID    int       `gorm:"column:user_id;not null" json:"user_id"`
	TrackID   int       `gorm:"column:track_id;not null" json:"track_id"`
	ArtistID  int       `gorm:"column:artist_id" json:"artist_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

type PlaylistComment struct {
	ID         int       `gorm:"not null;AUTO_INCREMENT" json:"id"`
	Comment    string    `gorm:"column:comment;not null" json:"comment"`
	UserID     int       `gorm:"column:user_id;not null" json:"user_id"`
	PlaylistID int       `gorm:"column:playlist_id;not null" json:"playlist_id"`
	ArtistID   int       `gorm:"column:artist_id" json:"artist_id"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
}
