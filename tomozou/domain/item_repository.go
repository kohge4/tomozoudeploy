package domain

// webservoceaccountImpl が 構造体依存する
type ItemRepository interface {
	ReadArtistBySocialID(socialID string) (*Artist, error)

	SaveArtist(Artist) (int, error)
	SaveTrack(Track) (int, error)
	SaveUserArtistTag(UserArtistTag) error
	SaveUserTrackTag(UserTrackTag) error

	ReadTagByUser(userID int) (interface{}, error)
	ReadUserArtistTagByUserID(userID int) (interface{}, error)
	ReadUserArtistTagByTagID(tagID int) (interface{}, error)

	ReadUserTrackTagByUserID(userID int) ([]UserTrackTagFull, error)
	ReadUserTrackTagByTagName(tagName string) ([]UserTrackTagFull, error)
	ReadUserTrackTagByUserIDANDTagName(userID int, tagName string) ([]UserTrackTagFull, error)

	ReadUserIDByArtistID(artistID int) ([]int, error)
	ReadUserIDByArtistName(artistName string) ([]int, error)

	//DeleteAllArtistByUserID(userID int) error
	DeleteAllUserArtistTagsByUserID(userID int) error
}
