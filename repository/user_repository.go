package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

// インターフェースを作成
// golangではインターフェースはメソッドの一覧

type IUserRepository interface {
	// 第1引数にuserオブジェクトのポインタを受け取る
	// 返り値をerrorインターフェース型で定義
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
