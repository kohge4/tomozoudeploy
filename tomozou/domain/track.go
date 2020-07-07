package domain

import (
	"strconv"
	"strings"
)

type Track struct {
	ID            int    `gorm:"column:id" json:"id"`
	SocialTrackID string `gorm:"column:social_track_id;not null" json:"social_track_id"`
	TrackName     string `gorm:"column:track_name;not null" json:"track_name"`
	AlbumName     string `gorm:"column:album_name" json:"album_name"`

	ArtistName string `gorm:"column:arttist_name;not null" json:"artist_name"`
	ArtistID   int    `gorm:"column:arttist_id;not null" json:"artist_id"`
	Webservice string `gorm:"column:webservice;not null" json:"webservice"`
	ArtistIDs  string `gorm:"column:arttist_ids" json:"artist_ids"`

	TrackNameOption string `gorm:"column:track_name_option" json:"track_name_option"`
	// "spap", "sp", "ap"  とかで対応数を 長さで判断したい
}

func (t *Track) UserTrackTag(userID int, tagName string, count int) *UserTrackTag {
	return NewUserTrackTag(t, userID, tagName, count)
}

func (t *Track) ArtistIDsList() []int {
	idStringList := strings.Split(t.ArtistIDs, ",")
	idList := make([]int, 0)
	for _, s := range idStringList {
		i, _ := strconv.Atoi(s)
		idList = append(idList, i)
	}
	return idList
}

// TrackWithArtistList : TrackIDs から track のlist を入手
type TrackWithArtistList struct {
	*Track
	Artists []Artist `json:"artists"`
}

/*
みたいな感じ
{1,1,"13dtgy2943uh","apple"}
*/
/*
多分track 側で何に対応してるかわかったほうがいい
==>　属性値の追加
*/
