package domain

type Track struct {
	ID            int    `gorm:"primary_key;column:id" json:"id"`
	SocialTrackID string `gorm:"column:social_track_id;not null" json:"social_track_id"`
	TrackName     string `gorm:"column:track_name;not null" json:"track_name"`
	AlbumName     string `gorm:"column:album_name" json:"album_name"`

	ArtistName string `gorm:"column:arttist_name;not null" json:"artist_name"`
	ArtistID   int    `gorm:"column:arttist_id;not null" json:"artist_id"`
	Webservice string `gorm:"column:webservice;not null" json:"webservice"`

	NameOpt       string `gorm:"column:name_opt" json:"name_opt"`
	TrackNameOpt  string `gorm:"column:track_name_opt" json:"track_name_opt"`
	ArtistNameOpt string `gorm:"column:artist_name_opt" json:"artist_name_opt"`
	// "spap", "sp", "ap"  とかで対応数を 長さで判断したい
}

func (t *Track) UserTrackTag(userID int, tagName string) *UserTrackTag {
	return NewUserTrackTag(t, userID, tagName)
}

/*
みたいな感じ
{1,1,"13dtgy2943uh","apple"}
*/
/*
多分track 側で何に対応してるかわかったほうがいい
==>　属性値の追加
*/
