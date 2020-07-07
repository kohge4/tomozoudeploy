package usecase

/*

type ConnectorApplication struct {
	Handler        WebServiceConnector
	ItemRepository domain.ItemRepository
	//BaseWebService domain.WebService
}

type WebServiceConnector struct {
	AppleConnector   domain.WebServiceConnector
	SpotifyConnector domain.WebServiceConnector
}

func NewConnectorApplication() {}

func (a *ConnectorApplication) CreateTrackWebserviceTagByWebService(serviceName string, track *domain.Track) {
	searchObj := domain.SearchObj{
		SearchArtistName: track.ArtistName,
		SearchTrackName:  track.TrackName,
	}
	if serviceName == "applemusic" {
		// trackの名前 とかでいい感じに検索 => socialID を TrackWebServiceTag に追加
		//a.Handler.AppleConnector.SearchTrackAndSaveTrackInfo(&searchObj)

	} else if serviceName == "spotify" {
		//a.Handler.SpotifyConnector.SearchTrackAndSaveTrackInfo(&searchObj)
	}
}

func (a *ConnectorApplication) CreateArtistWebserviceTagByWebService(serviceName string, artist *domain.Artist) {
	// serviceにおけるTrack のURL (SocialID)を取得して保存する
	// a.ItemRepository.UpdateArtistByNewWebService()
	if serviceName == "applemusic" {
		// trackの名前 とかでいい感じに検索 => socialID を TrackWebServiceTag に追加

	} else if serviceName == "spotify" {
	}
}

*/
