package service

import "tomozou/domain"

type TrackService struct {
	Track      domain.Artist
	WebService domain.WebServiceAccount
}

func (s *TrackService) FetchWebService() {
	// TrackSubstitute 的なのを作る

	// UseCase との違いがわからなくなった
}
