package domain

type ItemChildRepository interface {
	SaveTrackComment(*TrackComment) error
	ReadTrackComment()
	ReadTrackCommentWithUserByTrackID(trackID int) ([]TrackCommentWithUser, error)

	SaveTrackWebServiceTag(tag *TrackWebServiceTag) error
	ReadTrackWithTrackWebServiceTagByTrackID(trackID int) (*[]TrackWithTrackWebServiceTag, error)
}
