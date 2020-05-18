package domain

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
