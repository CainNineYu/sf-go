package dao

import (
	"sf-go/internal/dao/db"
	"sf-go/internal/dao/models"
)

// func VIPByUser(user string) ([]interface{}, error) {
// 	sql := "select * from vips where user= ? "
// 	rows, err := db.ReaderDb.Query(sql, user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res, err := db.SqlRowsToJSON(rows)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

type VipsDAO struct {
	db *db.DB
}

func NewVipsDAO(db *db.DB) *VipsDAO {
	return &VipsDAO{db: db}
}

func (v *VipsDAO) VIPByUser(user string) (models.Vips, error) {
	var vips models.Vips
	dbConn := v.db.ReadDB // 或 db.ReaderDb，确保是 *gorm.DB
	if err := dbConn.Where("user = ?", user).Find(&vips).Error; err != nil {
		return vips, err
	}
	return vips, nil
}
