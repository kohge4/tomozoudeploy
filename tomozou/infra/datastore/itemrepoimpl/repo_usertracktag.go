package itemrepoimpl

import (
	"fmt"
	"tomozou/domain"
)

func (repo *ItemRepositoryImpl) SaveUserTrackTag(tag domain.UserTrackTag) error {
	repo.DB.Create(&tag)
	return nil
}

func (repo *ItemRepositoryImpl) ReadUserTrackTagByUserID(userID int) ([]domain.UserTrackTagFull, error) {
	// nowplaying の 表示用
	userTrackTags := []domain.UserTrackTagFull{}
	//repo.DB.Where("user_id = ?", userID).Find(&userTrackTags)
	sql := "SELECT * FROM user_track_tags JOIN tracks ON user_track_tags.track_id = tracks.id JOIN users ON user_track_tags.user_id = users.id WHERE user_track_tags.user_id = ?"
	repo.DB.Raw(sql, userID).Scan(&userTrackTags)
	if len(userTrackTags) == 0 {
		return userTrackTags, fmt.Errorf("nil error")
	}
	return userTrackTags, nil
}

func (repo *ItemRepositoryImpl) ReadUserTrackTagByTagName(tagName string) ([]domain.UserTrackTagFull, error) {
	userTrackTags := []domain.UserTrackTagFull{}
	//repo.DB.Where("tag_name = ?", tagName).Find(&userTrackTags)
	sql := "SELECT * FROM user_track_tags JOIN tracks ON user_track_tags.track_id = tracks.id JOIN users ON user_track_tags.user_id = users.id WHERE tag_name = ?"
	repo.DB.Raw(sql, tagName).Scan(&userTrackTags)
	if len(userTrackTags) == 0 {
		return userTrackTags, fmt.Errorf("nil error")
	}
	return userTrackTags, nil
}

func (repo *ItemRepositoryImpl) ReadUserTrackTagByUserIDANDTagName(userID int, tagName string) ([]domain.UserTrackTagFull, error) {
	// nowplaying の 表示用
	userTrackTags := []domain.UserTrackTagFull{}
	//repo.DB.Where("user_id = ? AND tag_name = ?", userID, tagName).Find(&userTrackTags)
	sql := "SELECT * FROM user_track_tags JOIN tracks ON user_track_tags.track_id = tracks.id JOIN users ON user_track_tags.user_id = users.id WHERE user_track_tags.user_id = ? AND user_track_tags.tag_name = ?"
	repo.DB.Raw(sql, userID, tagName).Scan(&userTrackTags)
	if len(userTrackTags) == 0 {
		return userTrackTags, fmt.Errorf("nil error")
	}
	return userTrackTags, nil
}
