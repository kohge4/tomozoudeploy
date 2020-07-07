package itemchildrepoimpl

import (
	"fmt"
	"tomozou/domain"
)

func (repo *ItemChildRepositoryImpl) SaveArtistWebServiceTag(tag *domain.ArtistWebServiceTag) error {
	repo.DB.Create(tag)
	//repo.DB.Last(&tag)
	return nil
}

func (repo *ItemChildRepositoryImpl) ReadArtistWebServiceTagByArtistID(artistID int) (*domain.ArtistWebServiceTag, error) {
	tag := &domain.ArtistWebServiceTag{}
	repo.DB.Where("artist_id = ?", artistID).First(&tag)
	if tag.ID == 0 {
		return nil, nil
	}
	return tag, nil
}

func (repo *ItemChildRepositoryImpl) ReadArtistWebServiceTagAllByArtistID(artistID int) (*[]domain.ArtistWebServiceTag, error) {
	tags := []domain.ArtistWebServiceTag{}
	sql := "SELECT * FROM artist_web_service_tags WHERE artist_web_service_tags.artist_id = ?"
	repo.DB.Raw(sql, artistID).Scan(&tags)
	if len(tags) == 0 {
		return nil, fmt.Errorf("nil error")
	}
	return &tags, nil
}

func (repo *ItemChildRepositoryImpl) ReadArtistWithArtistWebServiceTagByArtistID(artistID int) (*[]domain.ArtistWithArtistWebServiceTag, error) {
	tags := []domain.ArtistWithArtistWebServiceTag{}
	sql := "SELECT * FROM artist_web_service_tags JOIN artists ON artist_web_service_tags.artist_id = artists.id WHERE artist_web_service_tags.artist_id = ?"
	repo.DB.Raw(sql, artistID).Scan(&tags)
	if len(tags) == 0 {
		return nil, fmt.Errorf("nil error")
	}
	return &tags, nil
}
