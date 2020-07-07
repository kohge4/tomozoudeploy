package itemrepoimpl

import "tomozou/domain"

func (repo *ItemRepositoryImpl) SaveTrack(track domain.Track) (int, error) {
	repo.DB.Create(&track)
	return track.ID, nil
}

func (repo *ItemRepositoryImpl) ReadTrackByTrackID(trackID int) (*domain.Track, error) {
	track := &domain.Track{}
	repo.DB.Where("id = ?", trackID).Find(&track)
	return track, nil
}

func (repo *ItemRepositoryImpl) ReadTrackBySocialTrackID(socialID string) (*domain.Track, error) {
	track := &domain.Track{}
	repo.DB.Where("social_track_id = ?", socialID).Find(track)
	return track, nil
}

func (repo *ItemRepositoryImpl) ReadTrackWithArtistListByTrackID(trackID int) (*domain.TrackWithArtistList, error) {
	track := &domain.Track{}
	repo.DB.Where("id = ?", trackID).Find(&track)

	aID := track.ArtistIDsList()
	artists := []domain.Artist{}
	for _, i := range aID {
		artist := domain.Artist{}
		repo.DB.Where("id = ?", i).Find(&artist)
		artists = append(artists, artist)
	}

	t := &domain.TrackWithArtistList{
		Track:   track,
		Artists: artists,
	}
	return t, nil
}
