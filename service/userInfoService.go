package service

import (
	"GinHello/entity"
	"GinHello/mapper"
)

func InsertUser(userInfo *entity.User_info) (err error) {
	if err = mapper.SqlSession.Create(userInfo).Error; err != nil {
		return err
	}
	return
}

func FindAllUser() (userList []*entity.User_info, err error) {
	if err := mapper.SqlSession.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

func DeleteUserById(id string) (err error) {
	err = mapper.SqlSession.Where("id=?", id).Delete(&entity.User_info{}).Error
	return
}

func GetUserById(id string) (user *entity.User_info, err error) {
	if err = mapper.SqlSession.Where("id=?", id).First(user).Error; err != nil {
		return nil, err
	}
	return
}
func UpdateUser(user *entity.User_info) (err error) {
	err = mapper.SqlSession.Save(user).Error
	return
}
