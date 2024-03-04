package ports

import "time"

type UserRepo interface {
	FindUser(email string) (*UserEntity, error)
	FindUserDetail(userId string) (*UserEntity, error)
}

type UserEntity struct {
	UserId      string    `gorm:"column:userId;primaryKey"`
	Email       string    `gorm:"column:email"`
	Password    string    `gorm:"column:password"`
	Role        string    `gorm:"column:role"`
	AccountName string    `gorm:"column:accountName"`
	CreatedAt   time.Time `gorm:"column:createdAt"`
	UpdatedAt   time.Time `gorm:"column:updatedAt"`
}
