package service

import (
	"errors"
	"gitee.com/noovertime/usercenter/domain/model"
	"gitee.com/noovertime/usercenter/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserDataService interface {
	AddUser(user *model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(user *model.User, isChangePwd bool) (err error)
	FindUserByName(string) (*model.User, error)
	CheckPwd(username string, pwd string) (isOK bool, err error)
}

func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository: userRepository}
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

//用户名密码加密
func GeneratePassword(userPasswd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPasswd), bcrypt.DefaultCost)
}

//验证用户密码
func ValidatePassword(userpassword string, hashed string) (isOK bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userpassword)); err != nil {
		return false, errors.New("密码比对错误")
	}
	return true, nil
}

func (u *UserDataService) AddUser(user *model.User) (userid int64, err error) {
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return userid, err
	}
	user.HashPassword = string(pwdByte)
	return u.UserRepository.CreateUser(user)
}

//删除用户
func (u *UserDataService) DeleteUser(userid int64) error {
	return u.UserRepository.DeleteUserByID(userid)
}

func (u *UserDataService) UpdateUser(user *model.User, isChangePwd bool) (err error) {
	if isChangePwd {
		pwdByte, err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdByte)
	}
	return u.UserRepository.UpdateUser(user)
}

func (u *UserDataService) FindUserByName(username string) (*model.User, error) {
	return u.UserRepository.FindUserByName(username)
}

func (u *UserDataService) CheckPwd(username string, pwd string) (isOK bool, err error) {
	user, err := u.UserRepository.FindUserByName(username)
	if err != nil {
		return false, err
	}
	return ValidatePassword(pwd, user.HashPassword)
}
