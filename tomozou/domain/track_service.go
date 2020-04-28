package domain

type TrackService struct {
	//Track      Artist
	WebService WebServiceAccount
}

func (s *TrackService) FetchWebServiceID(track *Track) {
	// TrackSubstitute 的なのを作る

	// UseCase との違いがわからなくなった
}
