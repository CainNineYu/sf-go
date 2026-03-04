package consts

import "time"

// 记录次数 前缀规范
const (
	CONUNTS_LIMIT = "CONUNTS_LIMIT_"
	LoginPrefix   = "Bearer "
)

const VALIDATE_KEY = "VALIDATE_CODE_"

// 前缀功能类型
const (
	SIGNUP          = "signup"
	UPDATE_PASSWORD = "update_password"
	TRADE_LINK      = "trade_link"
)

const (
	TokenTimeOut  = time.Hour * 24
	UploadTimeOut = time.Hour * 4
	UploadKey     = "Upload-"
	PingLimit     = time.Second * 1
)
