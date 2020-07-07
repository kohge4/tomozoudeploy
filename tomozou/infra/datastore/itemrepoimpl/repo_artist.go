package itemrepoimpl

import (
	"fmt"
	"tomozou/domain"
)

// 以下は SpotifyHandler から 保存するときガンガン使用する
func (repo *ItemRepositoryImpl) SaveArtist(artist domain.Artist) (int, error) {
	repo.DB.Create(&artist)
	return artist.ID, nil
}

func (repo *ItemRepositoryImpl) ReadArtistBySocialID(socialID string) (*domain.Artist, error) {
	artists := []*domain.Artist{}
	repo.DB.Where("social_id = ?", socialID).Find(&artists)
	if len(artists) == 0 {
		return nil, nil
	}
	if len(artists) == 1 {
		return artists[0], nil
	}
	fmt.Println("DUPLICATED ARTIST")
	return artists[0], nil
}

func (repo *ItemRepositoryImpl) ReadArtistByArtistID(artistID int) (*domain.Artist, error) {
	artists := []*domain.Artist{}
	repo.DB.Where("id = ?", artistID).Find(&artists)
	if len(artists) == 0 {
		return nil, nil
	}
	if len(artists) == 1 {
		return artists[0], nil
	}
	fmt.Println("DUPLICATED ARTIST")
	return artists[0], nil
}
