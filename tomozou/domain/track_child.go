package domain

import "time"

type TrackComment struct {
	//  column: track_comment_id とかにしたらうまく動くのでは
	ID        int       `gorm:"primary_key;column:id" json:"id"`
	UserID    int       `gorm:"column:user_id;not null" json:"user_id"`
	TrackID   int       `gorm:"column:track_id;not null" json:"track_id"`
	Comment   string    `gorm:"column:comment;not null" json:"comment"`
	CreatedAt time.Time `gorm:"column:track_comment_created_at" json:"track_comment_created_at"`
}

type TrackCommentFull struct {
	TrackComment
	Track
	User
}

type TrackCommentWithUser struct {
	TrackComment
	User
}

// 対応したらこのタグを増やす
// SearchBy は TrackName + ArtistName で 変になった時に TrackName だけで検索する場合を考える
type TrackWebServiceTag struct {
	ID             int    `gorm:"primary_key;column:id" json:"id"`
	TrackID        int    `gorm:"column:track_id;not null" json:"track_id"`
	WebServiceID   string `gorm:"column:webservice_id" json:"webservice_id"`
	SocialTrackID  string `gorm:"column:sub_social_track_id" json:"sub_social_track_id"`
	SocialTrackURL string `gorm:"column:social_track_url" json:"social_track_url"`
	SocialArtistID string `gorm:"column:sub_social_artist_id" json:"sub_social_artist_id"`

	SearchBy string `gorm:"column:search_by" json:"search_by"`
	Click    int    `gorm:"column:click" json:"click"`
}

type TrackWithTrackWebServiceTag struct {
	Track
	TrackWebServiceTag
}

type TrackWithTrackWebServiceTags struct {
	Track
	WebServiceTags *[]TrackWebServiceTag
}

func (t *TrackWebServiceTag) EmbedURL() string {
	if len(t.SocialTrackURL) < 9 {
		return t.SocialTrackURL
	}
	u := "https://embed." + t.SocialTrackURL[8:]
	t.SocialTrackURL = u
	return t.SocialTrackURL
}
