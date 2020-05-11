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
