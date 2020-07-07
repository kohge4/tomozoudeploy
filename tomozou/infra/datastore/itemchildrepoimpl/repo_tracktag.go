package itemchildrepoimpl

import (
	"fmt"
	"tomozou/domain"
)

func (repo *ItemChildRepositoryImpl) SaveTrackWebServiceTag(tag *domain.TrackWebServiceTag) error {
	repo.DB.Create(tag)
	//repo.DB.Last(&tag)
	return nil
}

func (repo *ItemChildRepositoryImpl) ReadTrackWebServiceTagByTrackID(trackID int) (*domain.TrackWebServiceTag, error) {
	tag := &domain.TrackWebServiceTag{}
	repo.DB.Where("track_id = ?", trackID).First(&tag)
	if tag.ID == 0 {
		return nil, nil
	}
	return tag, nil
}

func (repo *ItemChildRepositoryImpl) ReadTrackWebServiceTagAllByTrackID(trackID int) (*[]domain.TrackWebServiceTag, error) {
	tags := []domain.TrackWebServiceTag{}
	sql := "SELECT * FROM track_web_service_tags WHERE track_web_service_tags.track_id = ?"
	repo.DB.Raw(sql, trackID).Scan(&tags)
	if len(tags) == 0 {
		return nil, fmt.Errorf("nil error")
	}
	return &tags, nil
}

func (repo *ItemChildRepositoryImpl) ReadTrackWithTrackWebServiceTagByTrackID(trackID int) (*[]domain.TrackWithTrackWebServiceTag, error) {
	tags := []domain.TrackWithTrackWebServiceTag{}
	sql := "SELECT * FROM track_web_service_tags JOIN tracks ON track_web_service_tags.track_id = tracks.id WHERE track_web_service_tags.track_id = ?"
	repo.DB.Raw(sql, trackID).Scan(&tags)
	if len(tags) == 0 {
		return nil, fmt.Errorf("nil error")
	}
	return &tags, nil
}
