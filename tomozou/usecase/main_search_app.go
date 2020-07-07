package usecase

import (
	"tomozou/domain"
)

func (u *UserProfileApplication) NewSearchObjByTrack(t *domain.TrackWithTrackWebServiceTags) *domain.SearchObj {
	a, _ := u.ItemRepository.ReadArtistByArtistID(t.Track.ArtistID)
	if a.NameOption == "" {
		a.NameOption = t.Track.ArtistName
	}
	if t.TrackNameOption == "" {
		t.TrackNameOption = t.Track.TrackName
	}
	searchObj := &domain.SearchObj{
		SearchKey:        "track",
		SearchArtistName: t.Track.ArtistName,
		SearchTrackName:  t.Track.TrackName,
		ItemID:           t.Track.ID,
		ArtistNameOption: a.NameOption,
		TrackNameOption:  t.TrackNameOption,
	}
	return searchObj
}

func (u *UserProfileApplication) NewSearchObjByArtist(a *domain.ArtistWithArtistWebServiceTags) *domain.SearchObj {
	searchObj := &domain.SearchObj{
		SearchKey:        "artist",
		SearchArtistName: a.Name,
		ItemID:           a.ID,
		ArtistNameOption: a.NameOption,
	}
	return searchObj
}
