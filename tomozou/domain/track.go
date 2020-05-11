package domain

type Track struct {
	ID            int    `gorm:"primary_key;column:id" json:"id"`
	SocialTrackID string `gorm:"column:social_id;not null" json:"social_track_id"`
	TrackName     string `gorm:"column:track_name;not null" json:"track_name"`

	ArtistName string `gorm:"column:arttist_name;not null" json:"artist_name"`
	ArtistID   int    `gorm:"column:arttist_id;not null" json:"artist_id"`
}

func (t *Track) UserTrackTag(userID int, tagName string) *UserTrackTag {
	return NewUserTrackTag(t, userID, tagName)
}

// 対応したらこのタグを増やす
type TrackWebServiceTag struct {
	ID           int
	TrackID      int
	WebServiceID string
}
