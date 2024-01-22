package proxy

import (
	"auth-server/model"
)

type UserProxy struct{}

func (UserProxy) CreateUser(user *model.User) (err error) {
	db := storageEngine.GetStorageDB()
	err = db.Create(&user).Error
	return err
}

func (UserProxy) UpdateUser(userId int64, updateUser model.User) (err error) {
	db := storageEngine.GetStorageDB()
	var user model.User
	user.ID = userId
	if err = db.Model(&user).Updates(updateUser).Error; err != nil {
		return
	}
	return
}

func (UserProxy) GetUser(userId int64) (user *model.User, err error) {
	db := storageEngine.GetStorageDB()
	user.ID = userId
	err = db.First(&user).Error
	return
}

func (UserProxy) DeleteUser(userId int64) (err error) {
	db := storageEngine.GetStorageDB()
	err = db.Delete(&model.User{}, userId).Error
	return
}
