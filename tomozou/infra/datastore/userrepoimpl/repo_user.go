package userrepoimpl

import "tomozou/domain"

func (repo *UserRepositoryImpl) Save(user domain.User) (int, error) {
	repo.DB.Create(&user)
	return user.ID, nil
}

func (repo *UserRepositoryImpl) Update(id int) (domain.User, error) {
	user := domain.User{}
	repo.DB.Where("ID = ?", id).Find(&user)
	return user, nil
}

func (repo *UserRepositoryImpl) ReadAll() []domain.User {
	users := []domain.User{}
	repo.DB.Find(&users)
	return users
}

func (repo *UserRepositoryImpl) ReadBySocialID(socialID string) ([]domain.User, error) {
	users := []domain.User{}
	repo.DB.Where("social_user_id = ?", socialID).Find(&users)
	return users, nil
}

func (repo *UserRepositoryImpl) ReadByID(id int) (domain.User, error) {
	user := domain.User{}
	repo.DB.Where("id = ?", id).Find(&user)
	return user, nil
}

// 以下開発時の デバッグ用

func (repo *UserRepositoryImpl) CheckUser() (interface{}, error) {
	user := []domain.User{}
	repo.DB.Find(&user)
	return user, nil
}
