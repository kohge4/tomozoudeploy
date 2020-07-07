package domain

// webservoceaccountImpl が 構造体依存する
type ItemRepository interface {
	ReadArtistBySocialID(socialID string) (*Artist, error)
	ReadArtistByArtistID(artistID int) (*Artist, error)

	SaveArtist(Artist) (int, error)
	SaveTrack(Track) (int, error)
	SaveUserArtistTag(UserArtistTag) error
	SaveUserTrackTag(UserTrackTag) error

	ReadTagByUser(userID int) (interface{}, error)
	ReadUserArtistTagByUserID(userID int) (interface{}, error)
	ReadUserArtistTagByTagID(tagID int) (interface{}, error)

	ReadTrackByTrackID(trackID int) (*Track, error)
	ReadTrackBySocialTrackID(socialID string) (*Track, error)

	ReadTrackWithArtistListByTrackID(trackID int) (*TrackWithArtistList, error)

	ReadUserTrackTagByUserID(userID int) ([]UserTrackTagFull, error)
	ReadUserTrackTagByTagName(tagName string) ([]UserTrackTagFull, error)
	ReadUserTrackTagByUserIDANDTagName(userID int, tagName string) ([]UserTrackTagFull, error)
	ReadUserTrackTagByUserIDANDTagNameANDTrackID(userID int, tagName string, trackID int) ([]UserTrackTagFull, error)
	UpdateUserTrackTagByUserIDANDTagNameANDTrackID(userID int, tagName string, trackID int) ([]UserTrackTagFull, error)

	ReadUserIDByArtistID(artistID int) ([]int, error)
	ReadUserIDByArtistName(artistName string) ([]int, error)

	//DeleteAllArtistByUserID(userID int) error
	DeleteAllUserArtistTagsByUserID(userID int) error
	DeleteUserTrackTag(UserTrackTag) error
}
