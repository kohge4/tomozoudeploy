package service

import "tomozou/domain"

type ArtistService struct {
	Artist     domain.Artist
	WebService domain.WebServiceAccount
}

func (s *ArtistService) FetchWebService() {

}
