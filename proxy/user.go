package proxy

import (
	"auth-server/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserProxy struct{}

func (UserProxy) CreateUser(user *model.User) (err model.ErrorCode) {
	db := storageEngine.GetStorageDB()
	dbErr := db.Create(&user).Error
	if dbErr != nil {
		err = model.ErrDb.AddErr(dbErr)
	}
	return
}

func (UserProxy) UpdateUser(userId int64, updateUser model.User) (err model.ErrorCode) {
	db := storageEngine.GetStorageDB()
	var user model.User
	user.ID = userId
	if dbErr := db.Model(&user).Updates(updateUser).Error; dbErr != nil {
		err = model.ErrDb.AddErr(dbErr)
	}
	return
}

func (UserProxy) GetUserById(userId int64) (user model.User, found bool, err model.ErrorCode) {
	db := storageEngine.GetStorageDB()
	dbErr := db.First(&user, userId).Error
	if errors.Is(dbErr, gorm.ErrRecordNotFound) {
	} else if dbErr != nil {
		err = model.ErrDb.AddErr(dbErr)
	} else {
		found = true
	}
	return
}

func (UserProxy) GetUserByUsername(userName string) (user model.User, found bool, err model.ErrorCode) {
	db := storageEngine.GetStorageDB()
	user.Username = userName
	dbErr := db.First(&user).Error
	if errors.Is(dbErr, gorm.ErrRecordNotFound) {
	} else if dbErr != nil {
		err = model.ErrDb.AddErr(dbErr)
	} else {
		found = true
	}

	return
}

func (UserProxy) DeleteUser(userId int64) (err model.ErrorCode) {
	db := storageEngine.GetStorageDB()
	dbErr := db.Delete(&model.User{}, userId).Error
	if dbErr != nil {
		err = model.ErrDb.AddErr(dbErr)
	}
	return
}
