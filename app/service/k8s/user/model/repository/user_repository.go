package repository

import (
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-pass/app/service/k8s/user/model"

	"gorm.io/gorm"
)

// 创建需要实现的接口
type IUserRepository interface {
	//初始化表
	InitTable() error
	//根据ID查处找数据
	FindUserByID(int64) (*model.User, error)
	//创建一条 svc 数据
	CreateUser(*model.User) (int64, error)
	//根据ID删除一条 svc 数据
	DeleteUserByID(int64) error
	//修改更新数据
	UpdateUser(*model.User) error
	//查找svc所有数据
	FindAll() ([]model.User, error)
}

// 创建svcRepository
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

// 初始化表
func (u *UserRepository) InitTable() error {
	return u.mysqlDb.AutoMigrate(&model.User{})
}

// 根据ID查找Svc信息
func (u *UserRepository) FindUserByID(svcID int64) (svc *model.User, err error) {
	svc = &model.User{}
	return svc, u.mysqlDb.Preload("SvcPort").First(svc, svcID).Error
}

// 创建Svc信息
func (u *UserRepository) CreateUser(User *model.User) (int64, error) {
	return User.ID, u.mysqlDb.Create(User).Error
}

// 根据ID删除Svc信息
func (u *UserRepository) DeleteUserByID(svcID int64) error {
	tx := u.mysqlDb.Begin()
	//遇到问题回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		logx.Error(tx.Error)
		return tx.Error
	}
	//删除svc
	if err := u.mysqlDb.Where("id = ?", svcID).Delete(&model.User{}).Error; err != nil {
		tx.Rollback()
		logx.Error(err)
		return err
	}
	//删除相关的port
	//if err := u.mysqlDb.Where("svc_id = ?", svcID).Delete(&model.SvcPort{}).Error; err != nil {
	//	tx.Rollback()
	//	logx.Error(err)
	//	return err
	//}
	return tx.Commit().Error
}

// 更新Svc信息
func (u *UserRepository) UpdateUser(svc *model.User) error {
	return u.mysqlDb.Model(svc).Updates(svc).Error
}

// 获取结果集
func (u *UserRepository) FindAll() (svcAll []model.User, err error) {
	return svcAll, u.mysqlDb.Preload("SvcPort").Find(&svcAll).Error
}
