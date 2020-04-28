package itemrepoimpl

import "tomozou/domain"

func (repo *ItemRepositoryImpl) SaveUserArtistTag(tag domain.UserArtistTag) error {
	repo.DB.Create(&tag)
	return nil
}

func (repo *ItemRepositoryImpl) ReadTagByUser(userID int) (interface{}, error) {
	var tag []domain.UserArtistTag
	repo.DB.Find(&tag)
	return tag, nil
}

func (repo *ItemRepositoryImpl) ReadUserArtistTagByUserID(userID int) (interface{}, error) {
	userArtistTags := []domain.UserArtistTag{}
	repo.DB.Where("user_id = ?", userID).Find(&userArtistTags)
	return userArtistTags, nil
}

func (repo *ItemRepositoryImpl) ReadUserArtistTagByTagID(tagID int) (interface{}, error) {
	userArtistTags := []domain.UserArtistTag{}
	repo.DB.Where("tag_id = ?", tagID).Find(&userArtistTags)
	return userArtistTags, nil
}

func (repo *ItemRepositoryImpl) DeleteAllUserArtistTagsByUserID(userID int) error {
	tag := domain.UserArtistTag{}
	repo.DB.Where("user_id LIKE ?", userID).Delete(&tag)
	return nil
}
