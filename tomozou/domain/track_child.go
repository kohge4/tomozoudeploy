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
type TrackWebServiceTag struct {
	ID            int    `gorm:"primary_key;column:id" json:"id"`
	TrackID       int    `gorm:"column:track_id;not null" json:"track_id"`
	WebServiceID  string `gorm:"column:webservice_id" json:"webservice_id"`
	SocialTrackID string `gorm:"column:sub_social_track_id" json:"sub_social_track_id"`
	SocialURL     string `gorm:"column:social_url" json:"social_url"`
}

type TrackWithTrackWebServiceTag struct {
	Track
	TrackWebServiceTag
}
