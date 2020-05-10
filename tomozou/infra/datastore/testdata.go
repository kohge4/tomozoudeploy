package datastore

import "tomozou/domain"

var TestUser = domain.User{
	ID:       0,
	SocialID: "test",
	Name:     "taro",
	Auth:     "spotify",
	Image:    "",
}

var TestArtist = domain.Artist{
	ID:         0,
	Name:       "testMans",
	SocialID:   "socialartist",
	Image:      "https://i.scdn.co/image/ab6775700000ee85b2bd4f64bd8c250aedd13123",
	WebService: "spotofy",
}

var TestTrack = domain.Track{
	ID:         1,
	SocialID:   "socialtrack",
	Name:       "testtrack",
	ArtistName: "testMans",
	ArtistID:   0,
}

var TestUserArtistTag = domain.UserArtistTag{
	ID:         0,
	UserID:     0,
	ArtistID:   0,
	TagName:    "recently_favorite",
	ArtistName: "testMans",
	URL:        "",
	Image:      "",
}

var TestUserTrackTag = domain.UserTrackTag{
	ID:            0,
	UserID:        0,
	TrackID:       0,
	ArtistID:      0,
	TagName:       "nowplaying",
	ArtistName:    "testMans",
	TrackName:     "testtrack",
	TrackSocialID: "socialtrack",
}

var TestTrackComment = domain.TrackComment{
	UserID:  1,
	TrackID: 1,
	Comment: "pop",
}

var TestUserChat = domain.UserChat{}

var TestUserToken = domain.UserToken{}