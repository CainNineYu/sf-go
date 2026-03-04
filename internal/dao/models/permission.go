package models

type Permission struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
	Code     string `gorm:"column:code;size:64;not null" json:"code"`
	Path     string `gorm:"column:path;type:enum('admin','partner','leader','member','guest');not null" json:"path"`
	PermType string `gorm:"column:perm_type;type:enum('page','button');not null" json:"permType"`
}

func (Permission) TableName() string {
	return "permissions"
}
