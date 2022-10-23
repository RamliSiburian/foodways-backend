package Repositories

import (
	"foodways/Models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user Models.User) (Models.User, error)
	CreateProfile(user Models.Profile) (Models.Profile, error)
	Login(email string) (Models.User, error)
}

func RepositoryAuth(db *gorm.DB) *users {
	return &users{db}
}
func (r *users) Register(user Models.User) (Models.User, error) {
	err := r.db.Preload("Profile").Create(&user).Error

	return user, err
}
func (r *users) CreateProfile(user Models.Profile) (Models.Profile, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *users) Login(email string) (Models.User, error) {
	var user Models.User
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}
