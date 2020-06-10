package domain

import "time"

type User struct {
	ID           int    `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"`
	SocialUserID string `gorm:"column:social_user_id" json:"social_user_id"`
	UserName     string `gorm:"column:user_name;not null" json:"user_name"`
	Auth         string `gorm:"not null" json:"auth"`
	UserImage    string `gorm:"column:user_image" json:"user_image"`

	CreatedAt time.Time `gorm:"column:user_created_at;not null" json:"user_created_at"`
	UpdatedAt time.Time `gorm:"column:user_updated_at;not null" json:"user_updated_at"`
}

func NewUser(socialID string, name string, auth string, image string) User {
	return User{
		SocialUserID: socialID,
		UserName:     name,
		Auth:         auth,
		UserImage:    image,
	}
}
