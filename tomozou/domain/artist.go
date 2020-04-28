package domain

import "time"

type Artist struct {
	ID int `gorm:"column:id;not null;AUTO_INCREMENT" json:"id"`

	Name      string    `gorm:"column:name;not null" json:"name"`
	SocialID  string    `gorm:"column:social_id;not null" json:"social_id"`
	Image     string    `gorm:"column:image;not null" json:"image"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`

	WebService string `gorm:"column:webservice" json:"webservice"`
	// SocialService string : spotify, apple　とか書く done したかも書く
	// ArtistOption struct{} を作成して それに apple とかについて書いていく
}

func NewArtist(name string, socialID string, image string) Artist {
	return Artist{
		Name:     name,
		SocialID: socialID,
		Image:    image,
	}
}

type ArtistTrackTag struct {
	ID       int    `gorm:"column:id;not null;AUTO_INCREMENT" json:"id"`
	ArtistID int    `gorm:"column:artist_id;not null" json:"artist_id"`
	TrackID  int    `gorm:"column:track_id;not null" json:"track_id"`
	SocailID string `gorm:"column:social_id;not null" json:"social_id"`
	TagName  string `gorm:"column:tag_name;not null" json:"tag_name"`
}

// 対応したらこのタグを増やす
type ArtistWebServiceTag struct {
	ID         int    `gorm:"column:id;not null;AUTO_INCREMENT" json:"id"`
	ArtistID   int    `gorm:"column:artist_id;not null" json:"artist_id"`
	SocailID   string `gorm:"column:social_id;not null" json:"social_id"`
	TagName    string `gorm:"column:tag_name" json:"tag_name"`
	WebService string `gorm:"column:webservice;not null" json:"webservice"`
}
