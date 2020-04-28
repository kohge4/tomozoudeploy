package itemrepoimpl

import "tomozou/domain"

func (repo *ItemRepositoryImpl) SaveTrack(track domain.Track) (int, error) {
	repo.DB.Create(&track)
	return track.ID, nil
}
