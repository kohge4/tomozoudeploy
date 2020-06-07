package domain

type WebServiceAccount interface {
	// ログイン内容に基づいて Social アカウントの情報を domain.User 形式に 出力
	User() (*User, error)
	//Link(User) error

	// SpotifyHandler.SaveUserItem で 必要な情報を 全部保存
	SaveUserItem(userID int) error

	// 再連携時の Item 周りの処理
	UpdateUserItem(userID int) error
	UpdateUserItemOpt(userID int, opt string) error
}

type WebServiceConnector interface {
	SearchWebServiceItem(searchObj *SearchObj) error
	SearchWebServiceItemAndCreateItemTag(searchObj *SearchObj) error
	SearchTrackAndSaveTrackInfo(searchObj *SearchObj) error
}

type SearchObj struct {
	SearchKey        string
	SearchArtistName string
	SearchTrackName  string
	ItemType         string
	ItemID           int
	Results          []SearchResult
}

type SearchResult struct {
	ResultKey   string
	ArtistName  string
	TrackName   string
	SocialID    string
	URL         string
	OtherResult string
	Accuracy    float64
	Options     string
}

type WebService struct {
	ServiceName       string
	WebServiceAccount WebServiceAccount
}

func NewWebService(name string, wSA WebServiceAccount) *WebService {
	return &WebService{
		ServiceName:       name,
		WebServiceAccount: wSA,
	}
}

func (s *SearchObj) GetAccuracy(result *SearchResult) float64 {
	originArtist := s.SearchArtistName
	respArtist := result.ArtistName

	originTrack := s.SearchTrackName
	respTrack := result.TrackName
	if originArtist == respArtist {
		if originTrack == respTrack {
			return 1.0
		}
		return 0.9
	}
	return 0.5
}
