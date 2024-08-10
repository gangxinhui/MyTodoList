package dao

import (
	"MyTodoList/model"
	"context"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

// FindUserByUserName 根据用户名找到用户

func (dao *UserDao) FindUserByUserName(userName string) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("user_name = ?", userName).First(&user).Error
	return

}
func (dao *UserDao) CreateUser(user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).Create(user).Error
	return

}
func (dao *UserDao) FindUserByUserId(id uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	return

}
func (s *TaskDao) FindTaskByIdAndUserID(uId, id uint) (r *model.Task, err error) {
	err = s.Model(&model.Task{}).Where("id=? AND uid = ?", id, uId).First(&r).Error
	return
}
