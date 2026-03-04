package models

type Vips struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
	User       string `gorm:"column:user;size:32;unique;not null"`
	Email      string `gorm:"column:email;size:32;unique;not null"`
	Level      int    `gorm:"column:level;not null"`
	ExpireTime int64  `gorm:"column:expire_time;not null"`
	CreateTime int64  `gorm:"column:create_time;not null"`
}

// TableName 显式指定表名为 vips（如果你的表名是 vips）
func (Vips) TableName() string {
	return "vips"
}
