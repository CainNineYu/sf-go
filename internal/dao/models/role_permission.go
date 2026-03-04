package models

type RolePermission struct {
	ID   uint64 `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
	Role string `gorm:"column:role;type:enum('admin','partner','leader','member','guest');not null" json:"role"`
	Code string `gorm:"column:code;size:64;not null" json:"code"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
