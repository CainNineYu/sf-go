package models

type Users struct {
	Id       int64  `gorm:"column:id;type:varchar(64);primaryKey;autoIncrement;not null;comment:主键" json:"id"`
	Uuid     string `gorm:"column:uuid;type:varchar(64);not null;unique;comment:用户唯一标识" json:"uuid"`
	ParentId string `gorm:"column:parent_id;type:varchar(64);not null;default:'000000';comment:策略id" json:"parentId"`

	User      string `gorm:"column:user;type:varchar(32)" json:"user"`
	Name      string `gorm:"column:name;type:varchar(32)" json:"name"`
	Email     string `gorm:"column:email;type:varchar(32);unique;not null" json:"email"`
	Password  string `gorm:"column:password;type:varchar(128);not null" json:"password"`
	Desc      string `gorm:"column:desc;type:varchar(500)" json:"desc"` // 个人描述
	Role      string `gorm:"column:role;type:enum('admin','partner','leader','member','guest');not null" json:"role"`
	Level     int    `gorm:"column:level;type:int;default:0" json:"level"`   // 等级
	Status    int    `gorm:"column:status;type:int;default:1" json:"status"` // 0.注销  1.正常  2.禁用
	LastAt    int64  `gorm:"column:last_at;type:bigint unsigned;not null;default:0;comment:创建时间" json:"lastAt"`
	CreatedAt int64  `gorm:"column:created_at;type:bigint unsigned;not null;default:0;comment:创建时间" json:"createdAt"`
	UpdatedAt int64  `gorm:"column:updated_at;type:bigint unsigned;not null;default:0;comment:更新时间" json:"updatedAt"`
}

// TableName 显式指定表名为 users
func (Users) TableName() string {
	return "users"
}
