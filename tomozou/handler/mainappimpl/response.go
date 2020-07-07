package mainappimpl

import (
	"time"
	"tomozou/domain"
	"tomozou/settings"

	"github.com/rs/zerolog/log"
)

type Response struct {
	Status int    `json:"status"`
	URL    string `json:"url"`
}

type MyProfileResponse struct {
	Me         interface{} `json:"me"`
	Artists    interface{} `json:"artists"`
	Tracks     interface{} `json:"tracks"`
	Nowplaying interface{} `json:"nowplaying`
}

type MyTrackResponse struct {
	UserID       int    `json:"user_id"`
	TrackID      int    `json:"track_id"`
	TrackURL     string `json:"track_url`
	TrackName    string `json:"track_name"`
	TrackComment string `json:"track_commemnt"`
	ArtistID     int    `json:"artist_id"`
	Artistname   string `json:"artost_name"`
}

func NewMyTrackResponse(track *domain.UserTrackTag) *MyTrackResponse {
	return &MyTrackResponse{
		//TrackURL:  newTrackURL(track.TrackSocialID),
		//TrackName: track.TrackName,
		//TrackComment: track.TrackComment,
	}
}

type TrackTimeLineResponse struct {
	Items  []TrackResponse `json:"items"`
	Offset int             `json:"offset"`
	Length int             `json:"length"`
	//Limit  int
	//LastID int
	// 本来は 順番に表示させるやつやりたい
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

func NewTrackTimeLineResponse(u *UserProfileApplicationImpl, trackTags []domain.UserTrackTagFull) *TrackTimeLineResponse {
	var items []TrackResponse
	for _, tag := range trackTags {
		items = append(items, NewTrackResponse(u, &tag))
	}
	// offset と length はなんかいい方法ありそう(option でまとめて引数とか)
	return &TrackTimeLineResponse{
		Items:  items,
		Length: len(trackTags),
		Offset: 0,
	}
}

type ConnectedTrackResponse struct {
	ArtistName string                       `json:"artist_name"`
	Artists    []domain.ArtistWebServiceTag `json:"artists"`
	TrackName  string                       `json:"track_name"`
	Tracks     []domain.TrackWebServiceTag  `json:"tracks"`
	Image      string                       `json:"image"`
}

func NewConnectedTracksResponse(tags *[]domain.TrackWebServiceTag) *ConnectedTrackResponse {
	resp := &ConnectedTrackResponse{}
	for i := 0; i < len(*tags); i++ {
		log.Info().Interface("CONNECTED_RESPONSE_WEBSERVICETAG_LIST", (*tags)[i]).Msg("mainappimpl/ShowAppleMusic ")
		if (*tags)[i].SearchBy == "artist" {
			//resp.Artists = append(resp.Artists, (*tags)[i])
		} else if (*tags)[i].SearchBy == "track" {
			resp.Tracks = append(resp.Tracks, (*tags)[i])
		}
	}
	return resp
}

// Artisr の方だけ 画像と名前を保時できるようにしたい
func NewConnectedTracksAndArtistResponse(tags *[]domain.TrackWebServiceTag, a *domain.ArtistWithArtistWebServiceTags) *ConnectedTrackResponse {
	resp := &ConnectedTrackResponse{}
	resp.ArtistName = a.Name
	resp.Image = a.Image

	for i := 0; i < len(*tags); i++ {
		log.Info().Interface("CONNECTED_RESPONSE_WEBSERVICETAG_LIST", (*tags)[i]).Msg("mainappimpl/ShowAppleMusic ")
		if (*tags)[i].SearchBy == "track" {
			(*tags)[i].EmbedURL()
			resp.Tracks = append(resp.Tracks, (*tags)[i])
		}
	}
	for i := 0; i < len(*(a.WebServiceTags)); i++ {
		resp.Artists = append(resp.Artists, (*a.WebServiceTags)[i])
	}
	return resp
}

func NewTrackResponse(u *UserProfileApplicationImpl, trackTag *domain.UserTrackTagFull) TrackResponse {
	if trackTag == nil {
		return TrackResponse{}
	}
	user, _ := u.UseCase.UserRepository.ReadByID(trackTag.UserID)
	return TrackResponse{
		TrackID:      trackTag.TrackID,
		TrackURL:     newTrackURL(trackTag.Track.SocialTrackID),
		SpotifyID:    trackTag.Track.SocialTrackID,
		AppleID:      "",
		UserID:       trackTag.UserID,
		UserName:     user.UserName,
		UserImageURL: user.UserImage,
		CreatedAt:    trackTag.UserTrackTag.CreatedAt,
	}
}

func newTrackURL(socialID string) string {
	url := settings.SpotifyTrackURL + socialID
	return url
}
