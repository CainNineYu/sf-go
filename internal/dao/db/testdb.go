package db

func NewMockDB() (*DB, error) {
	dbCfg, err := NewDB(DatabaseConfig{
		User:         "root",
		Database:     "test",
		Host:         "13.115.25.26",
		Port:         "3306",
		Password:     "root",
		MaxOpenConns: 50,
		MaxIdleConns: 10,
	}, DatabaseConfig{
		User:         "root",
		Database:     "test",
		Host:         "13.115.25.26",
		Port:         "3306",
		Password:     "root",
		MaxOpenConns: 50,
		MaxIdleConns: 10,
	})
	if err != nil {
		return nil, err
	}

	//// 自动迁移表结构（确保表存在且字段正确）
	//err = dbCfg.WriteDB.AutoMigrate(&models.ExternalOrderRecord{}, &models.ExternalOrderRecordPatch{})
	return dbCfg, nil
}
