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
	//SearchTrackAndSaveTrackInfo(searchObj *SearchObj) error
}

type SearchObj struct {
	SearchKey        string
	SearchArtistName string
	SearchTrackName  string
	ArtistNameOption string
	TrackNameOption  string
	ItemType         string
	ItemID           int
	Status           int
	Results          []SearchResult
}

type SearchResult struct {
	ResultKey   string
	ArtistName  string
	TrackName   string
	SocialID    string
	URL         string
	OtherResult string
	Accuracy    int
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

func (s *SearchObj) GetAccuracy(result *SearchResult) int {
	respArtist := result.ArtistName
	respTrack := result.TrackName
	if s.SearchArtistName == respArtist || s.ArtistNameOption == respArtist {
		if s.SearchTrackName == respTrack || s.TrackNameOption == respTrack {
			// 部分一致率みたいなのがあれば求めたい
			// 最初から二文字目まで一致してればOK？？
			return 100
		}
		// 変な文字を削除してどうなるかを判別したい(部分一致がどの程度か)
		// 最初から二文字目まで一致してればOK？？
		return 90
	}
	return 50
}
