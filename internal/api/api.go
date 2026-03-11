package api

import (
	"sf-go/docs"
	"sf-go/internal/api/handle"
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
	redis *db.RDB,
) *gin.Engine {
	r := gin.Default()

	r.Use(common.CrossDomainMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	api := r.Group("/api")

	index := api.Group("/")
	{
		//index.GET("/:email/:type", handlers.SendEmail)
		index.POST("/:email/:type", func(ctx *gin.Context) {
			handle.SendEmail(ctx,
				dbInstance,
				redis,
			)
		})
		index.POST("/register/email", func(ctx *gin.Context) {
			handle.EmailRegister(ctx,
				dbInstance,
				redis,
			)
		})
		index.POST("/login", func(ctx *gin.Context) {
			handle.Login(ctx,
				dbInstance,
				redis,
			)
		})
	}
	api.Use(common.AuthMiddleware(redis))
	user := api.Group("/user")
	{
		user.POST("/logout", func(ctx *gin.Context) {
			handle.Logout(ctx, dbInstance, redis)
		})
	}
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
