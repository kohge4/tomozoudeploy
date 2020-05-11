package itemchildrepoimpl

import "tomozou/domain"

func (repo *ItemChildRepositoryImpl) SaveTrackComment(trackComment *domain.TrackComment) error {
	repo.DB.Create(trackComment)
	// いらないかも
	repo.DB.Last(&trackComment)
	return nil
}

func (repo *ItemChildRepositoryImpl) ReadTrackComment() {
	return
}

func (repo *ItemChildRepositoryImpl) ReadTrackCommentWithUserByTrackID(trackID int) ([]domain.TrackCommentWithUser, error) {
	trackCommentWithUser := []domain.TrackCommentWithUser{}
	sql := "SELECT * FROM track_comments JOIN users ON track_comments.user_id = users.id WHERE track_id = ?"
	repo.DB.Raw(sql, trackID).Scan(&trackCommentWithUser)
	return trackCommentWithUser, nil
}
