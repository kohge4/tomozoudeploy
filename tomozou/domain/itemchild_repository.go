package domain

type ItemChildRepository interface {
	SaveTrackComment(*TrackComment) error
	ReadTrackComment()
}
