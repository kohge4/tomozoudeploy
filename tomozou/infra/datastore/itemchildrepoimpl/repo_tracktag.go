package itemchildrepoimpl

import "tomozou/domain"

func (repo *ItemChildRepositoryImpl) SaveTrackWebServiceTag(tag *domain.TrackWebServiceTag) error {
	repo.DB.Create(tag)
	//repo.DB.Last(&tag)
	return nil
}

func (repo *ItemChildRepositoryImpl) ReadTrackWithTrackWebServiceTagByTrackID(tarckID int) (*[]domain.TrackWithTrackWebServiceTag, error) {
	return nil, nil
}
