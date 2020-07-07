package domain

type ItemChildRepository interface {
	SaveTrackComment(*TrackComment) error
	ReadTrackComment()
	ReadTrackCommentWithUserByTrackID(trackID int) ([]TrackCommentWithUser, error)

	SaveTrackWebServiceTag(tag *TrackWebServiceTag) error
	ReadTrackWebServiceTagByTrackID(trackID int) (*TrackWebServiceTag, error)
	ReadTrackWebServiceTagAllByTrackID(trackID int) (*[]TrackWebServiceTag, error)
	ReadTrackWithTrackWebServiceTagByTrackID(trackID int) (*[]TrackWithTrackWebServiceTag, error)

	SaveArtistWebServiceTag(tag *ArtistWebServiceTag) error
	ReadArtistWebServiceTagByArtistID(artistID int) (*ArtistWebServiceTag, error)
	ReadArtistWebServiceTagAllByArtistID(artistID int) (*[]ArtistWebServiceTag, error)
	ReadArtistWithArtistWebServiceTagByArtistID(trackID int) (*[]ArtistWithArtistWebServiceTag, error)
}
