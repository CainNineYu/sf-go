package api

import (
	"sf-go/docs"
	"sf-go/internal/common"
	"sf-go/internal/config"
	"sf-go/internal/dao/db"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const SwaggerHost = "http://api.sohofreelancer.com"

func Router(
	dbInstance *db.DB,
	baseCfg *config.ApiSrvCfg,
) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	api := r.Group("/api")

	index := api.Group("/")
	{
		index.POST("/login", func(ctx *gin.Context) {
			handle.Login(ctx,
				dbInstance,
				redis,
			)
		})
	}
	r.Use(common.CrossDomainMiddleware())
	// Swagger 文档路由
	registerSwagger(r)

	return r
}

func registerSwagger(r gin.IRouter) {
	// API文档访问地址: http://host/swagger/index.html
	// 注解定义可参考 https://github.com/swaggo/swag#declarative-comments-format
	// 样例 https://github.com/swaggo/swag/blob/master/example/basic/api/api.go
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Title = "量化后台接口" // 必填
	docs.SwaggerInfo.Description = "量化后台内部服务接口"
	docs.SwaggerInfo.Version = "1.0" // 必填
	docs.SwaggerInfo.Host = SwaggerHost
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
