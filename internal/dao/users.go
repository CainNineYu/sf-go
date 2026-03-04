package dao

import (
	"errors"
	"sf-go/internal/dao/db"
	"sf-go/internal/dao/models"
	"time"

	"github.com/jinzhu/gorm"
)

type Users struct {
	Id            int64     `gorm:"AUTO_INCREMENT;PRIMARY_KEY;COMMENT:'主键ID'"`
	Uuid          string    `gorm:"varchar(64);NOT NULL;COMMENT:'用户帐号'"`
	ParentId      string    `gorm:"varchar(64);COMMENT:'上级'"`
	Name          string    `gorm:"varchar(32);COMMENT:'姓名'"`
	Phone         string    `gorm:"varchar(15);COMMENT:'手机号'"`
	Password      string    `gorm:"varchar(128);COMMENT:'密码'"`
	PubKey        string    `gorm:"varchar(64);COMMENT:'pubKey'"`
	PriKey        string    `gorm:"varchar(64);COMMENT:'priKey'"`
	APIPassword   string    `gorm:"varchar(128);COMMENT:'priKey'"`
	DeviceId      string    `gorm:"varchar(64);COMMENT:'priKey'"`
	Status        int       `gorm:"varchar(64);COMMENT:'状态'"`
	LastloginTime time.Time `gorm:";COMMENT:'最后登录'"`
	UpdateTime    time.Time `gorm:";COMMENT:'更新时间'"`
	RegisterTime  time.Time `gorm:";COMMENT:'注册时间'"`
	db            *gorm.DB
}

type UsersDAO struct {
	db *db.DB
}

type UserBasic struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Desc  string `json:"desc"`
	Role  string `json:"role"`
}

func NewUsersDAO(db *db.DB) *UsersDAO {
	return &UsersDAO{db: db}
}

// 根据 user 字段查找用户
func (u *UsersDAO) UserByUser(user string) (models.Users, error) {
	var users models.Users
	if err := u.db.ReadDB.Where("user = ? and status = 1", user).Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

// 根据 user 和 password 查找用户
func (u *UsersDAO) UserByPwd(email string, password string) (models.Users, error) {
	var users models.Users
	if err := u.db.ReadDB.Where("email = ? AND password = ?", email, password).Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

// 更新最后登录时间
func (u *UsersDAO) UpUserTime(user string) error {
	res := u.db.WriteDB.Model(&models.Users{}).Where("user = ?", user).Update("last_login_time", time.Now())
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("update users error")
	}
	return nil
}

// 更新密码
func (u *UsersDAO) UpPassword(user string, password string) error {
	res := u.db.WriteDB.Model(&models.Users{}).Where("user = ?", user).Update("password", password)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("update users error")
	}
	return nil
}

func (u *UsersDAO) ListUserBasicsAll() ([]UserBasic, error) {
	var list []UserBasic
	q := u.db.ReadDB.Model(&models.Users{}).
		Select("id, name, email, `desc` as `desc`, role").
		Where("status = 1")
	if err := q.Order("id desc").Scan(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (u *UsersDAO) ListUserBasicsExcludeAdmin() ([]UserBasic, error) {
	var list []UserBasic
	q := u.db.ReadDB.Model(&models.Users{}).
		Select("id, name, email, `desc` as `desc`, role").
		Where("status = 1").
		Where("role <> ?", string(RoleAdmin))
	if err := q.Order("id desc").Scan(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (u *UsersDAO) ListUserBasicsByParentUUID(parentUUID string) ([]UserBasic, error) {
	var list []UserBasic
	q := u.db.ReadDB.Model(&models.Users{}).
		Select("id, name, email, `desc` as `desc`, role").
		Where("status = 1").
		Where("parent_id = ?", parentUUID)
	if err := q.Order("id desc").Scan(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

//func UserByUser(user string) ([]interface{}, error) {
//	sql := "select * from users where user = ?"
//	rows, err := db.ReadDB.Query(sql, user)
//	if err != nil {
//		return nil, err
//	}
//	res, err := db.SqlRowsToJSON(rows)
//	if err != nil {
//		return nil, err
//	}
//	return res, nil
//}
//
//func UserByPwd(email string, password string) ([]interface{}, error) {
//	sql := "select * from users where user = ? and password =?"
//	rows, err := db.ReadDB.Query(sql, email, password)
//	if err != nil {
//		return nil, err
//	}
//	res, err := db.SqlRowsToJSON(rows)
//	if err != nil {
//		return nil, err
//	}
//	return res, nil
//}
//
//func UpUserTime(id string) error {
//	sql := "update users set last_time = ? where user = ?"
//	res, err := db.WriterDb.Exec(sql, time.Now().Unix(), id)
//	if err != nil {
//		return err
//	}
//	rowAccount, err := res.RowsAffected()
//	if err != nil || rowAccount == 0 {
//		return errors.New("update users error")
//	}
//	return nil
//}
//
//func UpPassword(user string, password string) error {
//	sql := "update users set password = ? where user = ?"
//	res, err := db.WriterDb.Exec(sql, password, user)
//	if err != nil {
//		return err
//	}
//	rowAccount, err := res.RowsAffected()
//	if err != nil || rowAccount == 0 {
//		return errors.New("update users error")
//	}
//	return nil
//}
