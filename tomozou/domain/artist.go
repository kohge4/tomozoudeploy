package domain

import "time"

type Artist struct {
	ID int `gorm:"column:id;not null;AUTO_INCREMENT" json:"id"`

	Name      string    `gorm:"column:name;not null" json:"name"`
	SocialID  string    `gorm:"column:social_id;not null" json:"social_id"`
	Image     string    `gorm:"column:image;not null" json:"image"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`

	Webservice    string `gorm:"column:webservice" json:"webservice"`
	NameOpt       string `gorm:"column:name_opt" json:"name_opt"`
	ArtistNameOpt string `gorm:"column:artist_name_opt" json:"artist_name_opt"`
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
	ID             int    `gorm:"column:id;not null;AUTO_INCREMENT" json:"id"`
	ArtistID       int    `gorm:"column:artist_id;not null" json:"artist_id"`
	SocailArtistID string `gorm:"column:social_artist_id;not null" json:"social_artist_id"`
	WebServiceID   string `gorm:"column:webservice_id;not null" json:"webservice_id"`
	//TagName        string `gorm:"column:tag_name" json:"tag_name"`
}

/*
みたいな感じ
{1,1,"13dtgy2943uh","apple"}
*/
