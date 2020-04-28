package itemrepoimpl

import "tomozou/domain"

func (repo *ItemRepositoryImpl) ReadUserIDByArtistID(artistID int) ([]int, error) {
	var users []int
	flag := 0

	userArtistTags := []domain.UserArtistTag{}
	repo.DB.Where("artist_id = ?", artistID).Find(&userArtistTags)

	// 並び替えた後
	for _, tag := range userArtistTags {
		if flag != tag.UserID {
			users = append(users, tag.UserID)
			flag = tag.UserID
		}
	}
	return users, nil
}

func (repo *ItemRepositoryImpl) ReadUserIDByArtistName(artistName string) ([]int, error) {
	var users []int
	flag := 0

	userArtistTags := []domain.UserArtistTag{}
	repo.DB.Where("artist_name = ?", artistName).Find(&userArtistTags)

	// 並び替えた後
	for _, tag := range userArtistTags {
		if flag != tag.UserID {
			users = append(users, tag.UserID)
			flag = tag.UserID
		}
	}
	return users, nil
}
