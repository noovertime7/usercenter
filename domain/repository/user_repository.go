package repository

import (
	"gitee.com/noovertime/usercenter/domain/model"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	InitTable() error
	FindUserByName(string) (*model.User, error)
	FindUserByID(int64) (*model.User, error)
	CreateUser(*model.User) (int64, error)
	DeleteUserByID(int64) error
	UpdateUser(*model.User) error
	FindAll() ([]model.User, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqldb: db}
}

type UserRepository struct {
	mysqldb *gorm.DB
}

// 初始化表

func (u *UserRepository) InitTable() error {
	return u.mysqldb.CreateTable(&model.User{}).Error
}

//根据名称查找用户
func (u *UserRepository) FindUserByID(id int64) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqldb.First(user, id).Error
}

func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqldb.Where("user_name = ?", name).Find(user).Error
}
func (u *UserRepository) CreateUser(user *model.User) (userID int64, err error) {
	return user.ID, u.mysqldb.Create(user).Error
}

func (u *UserRepository) DeleteUserByID(userId int64) error {
	return u.mysqldb.Where("id = ", userId).Delete(&model.User{}).Error
}

func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqldb.Model(user).Update(&user).Error
}

func (u *UserRepository) FindAll() (userall []model.User, err error) {
	return userall, u.mysqldb.Find(&userall).Error
}
