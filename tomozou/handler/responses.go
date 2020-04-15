package handler

import (
	"time"
	"tomozou/domain"
	"tomozou/settings"
)

type Response struct {
	Status int    `json:"status"`
	URL    string `json:"url"`
}

type MyProfileResponse struct {
	Me      interface{} `json:"me"`
	Artists interface{} `json:"artists"`
	Tracks  interface{} `json:"tracks"`
}

type MyTrackResponse struct {
	UserID       int
	TrackID      int
	TrackURL     string
	TrackName    string
	TrackComment string
	ArtistID     int
	Artistname   string
}

func NewMyTrackResponse(track *domain.UserTrackTag) *MyTrackResponse {
	return &MyTrackResponse{
		TrackURL:     newTrackURL(track.TrackSocialID),
		TrackName:    track.TrackName,
		TrackComment: track.TrackComment,
	}
}

type TrackTimeLineResponse struct {
	Items []TrackResponse `json:"items"`
	//Offset int
	//Limit  int
	//LastID int
	// 本来は 順番に表示させるやつやりたい
	Length int `json:"length"`
	//Filter string
}

type TrackResponse struct {
	TrackID      int       `json:"track_id"`
	TrackURL     string    `json:"track_url"`
	SpotifyID    string    `json:"spotify_id"`
	AppleID      string    `json:"apple_id"`
	UserID       int       `json:"user_id"`
	UserName     string    `json:"user_name"`
	UserImageURL string    `json:"user_image_url"`
	CreatedAt    time.Time `json:"created_at"`
}

func NewTrackTimeLineResponse(u *UserProfileApplicationImpl, trackTags []domain.UserTrackTag) *TrackTimeLineResponse {
	var items []TrackResponse
	for _, tag := range trackTags {
		items = append(items, NewTrackResponse(u, tag))
	}
	return &TrackTimeLineResponse{
		Items:  items,
		Length: len(trackTags),
	}
}

func NewTrackResponse(u *UserProfileApplicationImpl, trackTag domain.UserTrackTag) TrackResponse {
	user, _ := u.UseCase.UserRepository.ReadByID(trackTag.UserID)
	return TrackResponse{
		TrackID:      trackTag.TrackID,
		TrackURL:     newTrackURL(trackTag.TrackSocialID),
		SpotifyID:    trackTag.TrackSocialID,
		AppleID:      "",
		UserID:       trackTag.UserID,
		UserName:     user.Name,
		UserImageURL: user.Image,
		CreatedAt:    trackTag.CreatedAt,
	}
}

func newTrackURL(socialID string) string {
	url := settings.SpotifyTrackURL + socialID
	return url
}

type ChatResponse struct {
	ID        int       `json:"id"`
	Comment   string    `json:"comment"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	UserID    int       `json:"user_id"`
	ArtistID  int       `json:"artist_id"`
	CreatedAt time.Time `json:"created_at"`
}

type MyChatListResponse struct {
	Artists     interface{} `json:"artists"`
	ArtistsInfo interface{} `json:"artists_info"`
}
