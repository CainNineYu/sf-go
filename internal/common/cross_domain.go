package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sf-go/internal/common/consts"
	"sf-go/internal/dao/db"
	"sf-go/internal/dao/dto"
	"sf-go/pkg/common"

	//"sf-go/pkg/common"
	"strings"
)

func CrossDomainMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		language := c.GetHeader("Accept-Language")
		if language == "CN" {
			c.Set("language", "CN")
		} else if language == "TC" {
			c.Set("language", "TC")
		} else {
			c.Set("language", "EN")
		}

		// 处理请求
		c.Next()
	}
}

func AuthMiddleware(rdb *db.RDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//if c.Request.URL.Path == "/trade/create-order" {
		//	c.Next()
		//	return
		//}
		app := dto.Gin{C: c}
		authorizationStr := c.GetHeader("Authorization")
		if authorizationStr == "" {
			c.Abort()
			app.Response(http.StatusUnauthorized, dto.UNAUTHORIZED_ERROR, nil)
			return
		}
		authorization := strings.Split(authorizationStr, consts.LoginPrefix)[1]
		claims, err := common.ParseToken(authorization)
		if err != nil {
			c.Abort()
			app.Response(http.StatusUnauthorized, dto.UNAUTHORIZED_ERROR, nil)
			return
		}
		user, ok := claims["user"]
		if !ok {
			c.Abort()
			app.Response(http.StatusUnauthorized, dto.UNAUTHORIZED_ERROR, nil)
			return
		}
		priceCmd := rdb.Rdb.Get(common.GetMD5Encode(consts.LoginPrefix + user.(string))).Val()
		if priceCmd != authorization {
			c.Abort()
			app.Response(http.StatusUnauthorized, dto.UNAUTHORIZED_ERROR, nil)
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
