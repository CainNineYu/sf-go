package models

import "time"

type Users struct {
	Id            uint64    `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
	Uuid          string    `gorm:"column:uuid;size:64;unique"`
	ParentId      string    `gorm:"column:parent_id;size:64"`
	User          string    `gorm:"column:user;size:32"`
	Name          string    `gorm:"column:name;size:32"`
	Email         string    `gorm:"column:email;size:32;unique;not null"`
	Password      string    `gorm:"column:password;size:128;not null"`
	Desc          string    `gorm:"column:desc;size:500" json:"desc"` // 个人描述
	Role          string    `gorm:"column:role;type:enum('admin','partner','leader','member','guest');not null" json:"role"`
	Level         int       `gorm:"column:level;default:0"`               // 等级
	LastLoginTime time.Time `gorm:"column:last_login_time;type:datetime"` // 最后登录时间
	Status        int       `gorm:"column:status;default:1"`              // 0.注销  1.正常  2.禁用
	RegisterTime  time.Time `gorm:"column:register_time;type:datetime"`   // 注册时间
}

// TableName 显式指定表名为 users
func (Users) TableName() string {
	return "users"
}
