package service

import (
	"MyTodoList/dao"
	"MyTodoList/model"
	"MyTodoList/pkg/ctl"
	"MyTodoList/pkg/util"
	"MyTodoList/types"
	"context"
	"errors"
	"gorm.io/gorm"
	"sync"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
}

// UserSrvOnce: 用于确保单例实例只被初始化一次。sync.Once是Go标准库提供的一个类型，用于实现只执行一次的功能。
func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns

}
func (s *UserSrv) Register(ctx context.Context, req *types.UserServiceReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	//查找数据库用户名
	u, err := userDao.FindUserByUserName(req.UserName)
	switch err {
	//如果错误是gorm.ErrRecordNotFound，表示没有找到用户，继续进行用户注册流程。
	case gorm.ErrRecordNotFound:
		//创建一个新的User实例，设置用户名为请求中的用户名
		u = &model.User{
			UserName: req.UserName,
		}
		// 密码加密存储
		if err = u.SetPassword(req.Password); err != nil {
			//util.LogrusObj.Info(err)
			return
		}
		//在数据库中创建用户
		if err = userDao.CreateUser(u); err != nil {
			//util.LogrusObj.Info(err)
			return
		}
		//如果用户创建成功，返回成功响应ctl.RespSuccess()，错误为nil
		return ctl.RespSuccess(), nil
	case nil:
		err = errors.New("用户已存在")
		return
	default:
		return
	}
}
func (s *UserSrv) Login(ctx context.Context, req *types.UserServiceReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	//查找数据库用户名
	user, err := userDao.FindUserByUserName(req.UserName)
	if err == gorm.ErrRecordNotFound {
		err = errors.New("用户不存在")
	}
	//校验密码
	if !user.CheckPassword(req.Password) {
		err = errors.New("账号/密码错误")
		//util.LogrusObj.Info(err)
		return
	}
	// token 签发
	token, err := util.GenerateToken(user.ID, req.UserName, 0)
	if err != nil {
		//util.LogrusObj.Info(err)
		//return
	}
	u := &types.UserResp{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
	uResp := &types.TokenData{
		User:  u,
		Token: token,
	}
	return ctl.RespSuccessWithData(uResp), nil

}
