package model

/**
 * 封装操作用户表的方法
 */

import (
	"fmt"

	"go-api-example/pkg/auth"
	"go-api-example/pkg/constvar"

	validator "gopkg.in/go-playground/validator.v9"
)

// 用户model
type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (c *UserModel) TableName() string {
	return "tb_users"
}

// 封装插入一个用户
func (u *UserModel) Create() error {
	return DB.LocalDB.Create(&u).Error
}

// 根据id删除用户
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id
	return DB.LocalDB.Delete(&user).Error
}

// 更新用户
func (u *UserModel) Update() error {
	return DB.LocalDB.Save(u).Error
}

// 获取用户
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.LocalDB.Where("username = ?", username).First(&u)
	return u, d.Error
}

// 获取用户列表
func ListUser(username string, offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count uint64

	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := DB.LocalDB.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.LocalDB.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

// 校验密码
func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// 密码加密
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// 校验字段
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
