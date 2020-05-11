package domain

type ItemChildRepository interface {
	SaveTrackComment(*TrackComment) error
	ReadTrackComment()
	ReadTrackCommentWithUserByTrackID(trackID int) ([]TrackCommentWithUser, error)
}
