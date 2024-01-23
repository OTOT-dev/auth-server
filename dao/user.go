package dao

import (
	"auth-server/model"
)

type UserDao struct{}

func (UserDao) CreateUser(user *model.UserProps) (err error) {
	db := storageEngine.GetStorageDB()
	err = db.Create(&user).Error
	return err
}

func (UserDao) UpdateUser(userId int64, updateUser model.UserProps) (err error) {
	db := storageEngine.GetStorageDB()
	var user model.UserProps
	user.ID = userId
	if err = db.Model(&user).Updates(updateUser).Error; err != nil {
		return
	}
	return
}

func (UserDao) GetUser(userId int64) (user *model.UserProps, err error) {
	db := storageEngine.GetStorageDB()
	err = db.First(&user, userId).Error
	return
}

func (UserDao) DeleteUser(userId int64) (err error) {
	db := storageEngine.GetStorageDB()
	err = db.Delete(&model.UserProps{}, userId).Error
	return
}
