package usecase

import "tomozou/domain"

type ConnectorApplication struct {
	WebServices    []domain.WebService
	BaseWebService domain.WebService
	ItemRepository domain.ItemRepository
}

func NewConnectorApplication() {}

func (a *ConnectorApplication) SetBase(serviceName string) {
	a.BaseWebService = domain.WebService{ServiceName: serviceName}
}

func (a *ConnectorApplication) Track(serviceName string, track *domain.Track) {
	// serviceにおけるTrack のURL (SocialID)を取得して保存する
	// a.ItemRepository.UpdateTrackByNewWebService()
}

func (a *ConnectorApplication) Artist(serviceName string, artist *domain.Artist) {
	// serviceにおけるTrack のURL (SocialID)を取得して保存する
	// a.ItemRepository.UpdateArtistByNewWebService()
}
