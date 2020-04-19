package userpost

type UserPost struct {
	UserID        int
	ReceiveUserID []int
	ArtistID      []int
	TrackID       []int
	PlaylistID    []int
	Comment       string
	Content       []interface{}
}

// これを全部ドメインに書いたら解決する　=> ドメインに書くことが適切なのか考える必要 => アプリケーションが表示したり保存する形式だからドメインでいい
func (p *UserPost) UserChat() *UserChat {
	return &UserChat{}
}

func (p *UserPost) ArtistComment() *ArtistComment {
	return &ArtistComment{}
}

func (p *UserPost) TrackComment() *TrackComment {
	return &TrackComment{}
}

func (p *UserPost) PlaylistComment() *PlaylistComment {
	return &PlaylistComment{}
}
