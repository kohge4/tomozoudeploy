package domain

// 対応したらこのタグを増やす
// 画像は Artist(Spotify の情報がメイン)のやつを使用する
type ArtistWebServiceTag struct {
	ID               int    `gorm:"column:id;not null;AUTO_INCREMENT" json:"id"`
	ArtistID         int    `gorm:"column:artist_id;not null" json:"artist_id"`
	WebServiceID     string `gorm:"column:webservice_id;not null" json:"webservice_id"`
	SocialArtistID   string `gorm:"column:social_artist_id" json:"social_artist_id"`
	SocialArtistURL  string `gorm:"column:social_artist_url" json:"social_artist_url"`
	SocialArtistName string `gorm:"column:social_artist_name" json:"social_artist_name"`

	Click int `gorm:"column:click" json:"click"`
}

type ArtistWithArtistWebServiceTag struct {
	Artist
	ArtistWebServiceTag
}

type ArtistWithArtistWebServiceTags struct {
	Artist
	WebServiceTags *[]ArtistWebServiceTag
}

/*
type ArtistWebServiceTag struct {
	ID            int    `gorm:"primary_key;column:id" json:"id"`
	TrackID       int    `gorm:"column:track_id;not null" json:"track_id"`
	WebServiceID  string `gorm:"column:webservice_id" json:"webservice_id"`
	SocialTrackID string `gorm:"column:sub_social_track_id" json:"sub_social_track_id"`
	SocialURL     string `gorm:"column:social_url" json:"social_url"`

	SearchBy string `gorm:"column:search_by" json:"search_by"`
	Click    int    `gorm:"column:click" json:"click"`
}
*/

/*
みたいな感じ
{1,1,"13dtgy2943uh","apple"}
*/
