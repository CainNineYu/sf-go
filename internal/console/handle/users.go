package handle

import (
	"fmt"
	"net/http"
	"sf-go/internal/common/consts"
	consts2 "sf-go/internal/common/consts"

	"sf-go/internal/dao"
	"sf-go/internal/dao/db"
	"sf-go/internal/dto"
	"sf-go/pkg/common"
	"time"

	"github.com/gin-gonic/gin"
)

//func EmailRegister(c *gin.Context, db *db.DB, rdb *db.RDB) {
//	app := dto.Gin{C: c}
//	req := dto.NewEmailRegisterReq()
//	if err := req.Bind(&app); err != nil {
//		return
//	}
//
//	pwd := common.GetMD5Encode(req.Password)
//	//TODO: 取消注释
//	//redisKey := consts.VALIDATE_KEY + consts.SIGNUP + "_" + req.Email
//	//captchaBool := redis.RedisClient.Rdb.Get(redisKey)
//	//if captchaBool.Val() != req.Captcha {
//	//	app.Response(http.StatusInternalServerError, resUtils.CAPTCHA_ERROR, nil)
//	//	return
//	//}
//	usersDAO := dao.NewUsersDAO(db)
//	userInfo, err := usersDAO.UserByEmail(req.Email)
//	if err != nil {
//		app.Response(http.StatusInternalServerError, resUtils.NETWORK_ERROR, nil)
//		return
//	}
//	if len(userInfo) != 0 {
//		app.Response(http.StatusInternalServerError, resUtils.HAVE_REGISTERED, nil)
//		return
//	}
//
//	err = db.AddUsers(req.Email, req.Email, "", pwd)
//	if err != nil {
//		app.Response(http.StatusInternalServerError, resUtils.NETWORK_ERROR, nil)
//		return
//	}
//	app.Response(http.StatusOK, resUtils.SUCCESS, nil)
//}

func Login(c *gin.Context, db *db.DB, rdb *db.RDB) {
	app := dto.Gin{C: c}
	req := dto.NewLoginReq()
	if err := req.Bind(&app); err != nil {
		return
	}

	pwd := common.GetMD5Encode(req.Password)

	usersDAO := dao.NewUsersDAO(db)
	userInfo, err := usersDAO.UserByPwd(req.Email, pwd)
	if err != nil {
		app.Response(http.StatusInternalServerError, dto.NETWORK_ERROR, nil)
		return
	}
	if userInfo.Id == 0 {
		app.Response(http.StatusInternalServerError, dto.ACCOUNT_PASSWORD_ERROR, nil)
		return
	}
	lastLoginTime := time.Now().Unix()
	err = usersDAO.UpUserTime(userInfo.User)
	if err != nil {
		app.Response(http.StatusInternalServerError, dto.NETWORK_ERROR, nil)
		return
	}

	accessToken, err := common.CreateToken(userInfo.User, lastLoginTime)
	if err != nil {
		app.Response(http.StatusInternalServerError, dto.NETWORK_ERROR, nil)
		return
	}
	err = rdb.Rdb.Set(common.GetMD5Encode(consts2.LoginPrefix+userInfo.User), accessToken, consts.TokenTimeOut).Err()
	if err != nil {
		app.Response(http.StatusInternalServerError, dto.NETWORK_ERROR, nil)
		return
	}

	results := make(map[string]interface{}, 10)
	results["id"] = userInfo.Id
	results["uuid"] = userInfo.Uuid
	results["user"] = userInfo.User
	results["level"] = userInfo.Level
	results["email"] = userInfo.Email
	results["image"] = fmt.Sprint("/images/avatars/1.png")

	results["invite"] = ""

	//用户可以访问的路径和权限
	permissions := make([]string, 0)
	permDAO := dao.NewPermissionDAO(db)
	paths, err := permDAO.ListPermissionPathsByRole(userInfo.Role)
	if err != nil {
		app.Response(http.StatusInternalServerError, dto.NETWORK_ERROR, nil)
		return
	}
	permissions = paths

	//VIP
	vip, err := dao.NewVipsDAO(db).VIPByUser(userInfo.User)
	if err != nil {
		app.Response(http.StatusInternalServerError, dto.NETWORK_ERROR, nil)
		return
	}
	if vip.ID != 0 {
		expireTime := vip.ExpireTime
		if vip.ID != 0 && expireTime > time.Now().Unix() {
			results["isVip"] = 1
			results["expireTime"] = vip.ExpireTime
		} else {
			results["isVip"] = 0
			results["expireTime"] = 0
		}
	} else {
		results["isVip"] = 0
		results["expireTime"] = 0
	}
	res := map[string]interface{}{
		"userInfo":    results,
		"token":       accessToken,
		"permissions": permissions,
	}
	app.Response(http.StatusOK, dto.SUCCESS, res)

}

func Logout(c *gin.Context, db *db.DB, rdb *db.RDB) {
	app := dto.Gin{C: c}
	user, _ := c.Get("user")
	if user == nil {
		app.Response(http.StatusInternalServerError, dto.NETWORK_ERROR, nil)
		return
	}
	lastLoginTime := time.Now().Unix()
	err := dao.NewUsersDAO(db).UpUserTime(user.(string))
	if err != nil {
		app.Response(http.StatusInternalServerError, dto.NETWORK_ERROR, nil)
		return
	}
	_, err = common.CreateToken(user.(string), lastLoginTime)
	if err != nil {
		app.Response(http.StatusInternalServerError, dto.NETWORK_ERROR, nil)
		return
	}
	res := rdb.Rdb.Del(common.GetMD5Encode(consts2.LoginPrefix + user.(string))).Val()
	if res != 1 {
		app.Response(http.StatusInternalServerError, dto.NETWORK_ERROR, nil)
		return
	}
	app.Response(http.StatusOK, dto.SUCCESS, nil)
}

func UpPassword(c *gin.Context, db *db.DB) {
	app := dto.Gin{C: c}
	user, _ := c.Get("user")

	req := dto.NewUpdatePasswordReq()
	if err := req.Bind(&app); err != nil {
		return
	}

	pwd := common.GetMD5Encode(req.OldPassword)
	newPwd := common.GetMD5Encode(req.NewPassword)

	usersDAO := dao.NewUsersDAO(db)
	userInfo, err := usersDAO.UserByPwd(user.(string), pwd)
	if err != nil {
		app.Response(http.StatusInternalServerError, dto.NETWORK_ERROR, nil)
		return
	}
	if userInfo.Id == 0 {
		app.Response(http.StatusInternalServerError, dto.ACCOUNT_PASSWORD_ERROR, nil)
		return
	}
	err = usersDAO.UpPassword(user.(string), newPwd)
	if err != nil {
		app.Response(http.StatusInternalServerError, dto.NETWORK_ERROR, nil)
		return
	}
	app.Response(http.StatusOK, dto.SUCCESS, nil)
	return

}
