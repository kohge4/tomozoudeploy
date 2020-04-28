package itemrepoimpl

import (
	"fmt"
	"tomozou/domain"
)

func (repo *ItemRepositoryImpl) SaveUserTrackTag(tag domain.UserTrackTag) error {
	repo.DB.Create(&tag)
	return nil
}

func (repo *ItemRepositoryImpl) ReadUserTrackTagByUserID(userID int) ([]domain.UserTrackTag, error) {
	// nowplaying の 表示用
	userTrackTags := []domain.UserTrackTag{}
	repo.DB.Where("user_id = ?", userID).Find(&userTrackTags)
	if len(userTrackTags) == 0 {
		return userTrackTags, fmt.Errorf("nil error")
	}
	return userTrackTags, nil
}

func (repo *ItemRepositoryImpl) ReadUserTrackTagByTagName(tagName string) ([]domain.UserTrackTag, error) {
	userTrackTags := []domain.UserTrackTag{}
	repo.DB.Where("tag_name = ?", tagName).Find(&userTrackTags)
	if len(userTrackTags) == 0 {
		return userTrackTags, fmt.Errorf("nil error")
	}
	return userTrackTags, nil
}

func (repo *ItemRepositoryImpl) ReadUserTrackTagByUserIDANDTagName(userID int, tagName string) ([]domain.UserTrackTag, error) {
	// nowplaying の 表示用
	userTrackTags := []domain.UserTrackTag{}
	repo.DB.Where("user_id = ? AND tag_name = ?", userID, tagName).Find(&userTrackTags)
	if len(userTrackTags) == 0 {
		return userTrackTags, fmt.Errorf("nil error")
	}
	return userTrackTags, nil
}
