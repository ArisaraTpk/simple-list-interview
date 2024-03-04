package repository

import (
	"gorm.io/gorm"
	"simple-list-interview/internal/core/ports"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) ports.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r userRepo) FindUser(email string) (*ports.UserEntity, error) {
	var result ports.UserEntity
	res := r.db.Where("email = ?", email).First(&result)
	return &result, res.Error
}

func (r userRepo) FindUserDetail(userId string) (*ports.UserEntity, error) {
	var result ports.UserEntity
	res := r.db.Where("userId = ?", userId).First(&result)
	return &result, res.Error
}
