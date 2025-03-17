package serve

import (
	"moredoc/biz"
	"moredoc/middleware/auth"
	"moredoc/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RegisterGinRouter 注册gin路由
func RegisterGinRouter(app *gin.Engine, dbModel *model.DBModel, logger *zap.Logger, auth *auth.Auth) (err error) {
	// 初始化业务服务（比如文件上传的API）
	attachmentAPIService := biz.NewAttachmentAPIService(dbModel, logger)

	// 初始化订单支付服务（支付宝沙箱测试）
	appID := "9021000145623685"
	privateKey := "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCuV/a/FvjDSCghNE42hrff8fUMvC3wHAVwBzagednh9UreEmisWE0ZIlO66cyzMfuEgLW19N72AVuYacuLDdhXVJ14/dLAKA0mMzvxvS8+G7ykrmxeHZhnw28e2py7AxMneN5hrBA6OyhUbsSf8QhpEktaW/1Fequ6JmG+eh6N8oZpBEJefcL8ffChAhacWrb3vdNHH6Mck9yO+VRSJ/xlAShJIeJxP9Jxj74cMRDcrFo0Xt/4aKD6tEO1YxyiWuV2aACO0AwzopZ5cNTwHR9cAO+HyCGyXmUGnQUT/44peJoaOKcMcurIymrki2eL1mTjb5XXlBtZ7PgxYMIjZzjpAgMBAAECggEAfPLMA226QXKgWBO9jjSE6XnmDwd6wN/EQeLZkq9hqSB8VKXK4OGz97RzX70aGL9UrET2df6WIKcedyAzYWg9yXD6HvGCrnbF3b/QUVMt3YxRaZEcV2NMi+kz5V/1/c2ZV6u3bTa8to8ZO1Hbl2lOhsc2u/67iKT/GD1TxgiNh1q6Be7f3H0LIO37txvkFfAeBQYMS1KIHeLHyY0NnJzUGRGUbZ37e/dELdQPH2XAUFgd/hscYxXPrIbcOhEBkPpmMg1VebkDYVGTIwhAbDaeI3IN8hLHwxYetbMD5wtZWLMKqhEWDhdzmN5lCjw7GIcuO6SMBRrMfr6srfQHT9XJZQKBgQD/nfFYisBgC5iaKzW6c2GEUsXUGawTW2vT47DmYOG77/watAZR4voPv7rvOvmT8X8oga3I9YCk4vj85EzzQpQGwj1lfdx0HLRtyoDaAkXkxra64Aq5DdgECslIkV8ZZk3qxHqRQFS7mZtGbKzH++iZUcNezEcPXhQnwPou6VXmawKBgQCumtgEd3nTS7ceyKWE9Q6mngkTnVJngwyySKijm+WbV9hxuiKuri+b+twlBr5TuryITyUgVJnotY+5Sr5mtorARCnHo6mvoi1uH6ZPYAVhhT+PFQnO1qXAokG0tHgnWDYbYhVmYtMCM0ybmuzWDRy8QR3whJAB19MEo71YFdJq+wKBgQD474grAnV3E9P50Dry6Yr+ot2mQZoi/9Vy9TtCIPe1X3HVT50YAkQlMkOK15RXT8jWWfQ00WIHMkPphSyrextNubyxKnGXYp0UjzINHkmTEzLBq47FyYWl6hs8YNaMleUrzUzQSCY8zMs7OnrKB7nuFoJ9v9M0VBqs3HLJXNQbIQKBgDZMp1M0SKmiSU+EfJ0NjMHaMd5stXzFb4RePjurNsuJlSFBDpoeR2YNrLrDDxuA3jLR0izswnRxSkIIcZAFwr6qNAgG6LVbDR/LLfBPXz0FsE9x7nvpmuB+VUq6OQebUjBP33HY1+A+TduyYWbr6vzMj2AGCbbZBKHTvgWPu8FTAoGBAML5xHUUGZV+qt2Xb8YA4h61X3sX1zVimHwMGn5gn7uH4ceAGqZWjuvNO9xw/pJvGVSm+HqD0RAwWTc8Nk9rWa9npbE2n8LgmYAZqWBUdIeTQkl4CXJLNAAnRpPMYho8a/UQUYi5kvqkagNmfUIlc63LsnvOADcbNs1y+1KBrKr8"
	alipayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAj0PuMZcyV7+Cjl/Z5EcRSj3WPJP+S3uRLFwcT0rLJBTrJpkcB3c7gd/QSMApgOMHpOwi7QPnSv1ADS4KDHAUGnNR+63Fhdpm4tvIwSIctFkbX2jPT7g/F1e+ofJVa2NZ/rvR44zy59PyI5Lp9/sVAETZBFvCzFaysoATrxVpVlkS1MtGfVEiSXzHBZraqd5P66Gj0laqwKFrxAFf7ESn8qhK+aXaWIcfH/v5Me5BrdfGBVCtujxyGKbGwCBZfc2tdUrbH7Zv6f/HIh/uVnKuYTMmaUOhcmgCz7IKcm11eOkOkO8NWU2p5V+yIgQYCG7p77jDe6VAYupsydCN8Ac99wIDAQAB"
	paymentService := biz.NewPaymentService(appID, privateKey, alipayPublicKey)

	app.POST("/order_notify", paymentService.HandleAlipayNotify)

	// 注册不需要权限的GET路由（比如图标、静态资源）
	app.GET("/favicon.ico", attachmentAPIService.Favicon)
	app.GET("/static/images/logo.png", attachmentAPIService.Logo)
	app.GET("/sitemap.xml", func(ctx *gin.Context) {
		ctx.File("./sitemap/sitemap.xml")
	})

	// 其他GET路由（比如查看文档、下载文件）
	app.GET("/view/page/:hash/:page", attachmentAPIService.ViewDocumentPages)
	app.GET("/view/cover/:hash", attachmentAPIService.ViewDocumentCover)
	app.GET("/download/:jwt", attachmentAPIService.DownloadDocument)

	// 创建需要权限校验的路由组（比如上传相关API）
	checkPermissionGroup := app.Group("/api/v1/upload") // 路由组前缀
	checkPermissionGroup.Use(auth.AuthGin())            // 应用权限中间件
	{
		// 组内所有路由都需要权限
		checkPermissionGroup.POST("avatar", attachmentAPIService.UploadAvatar)
		checkPermissionGroup.POST("config", attachmentAPIService.UploadConfig)
		checkPermissionGroup.POST("banner", attachmentAPIService.UploadBanner)
		checkPermissionGroup.POST("document", attachmentAPIService.UploadDocument)
		checkPermissionGroup.POST("category", attachmentAPIService.UploadCategory)
		checkPermissionGroup.POST("article", attachmentAPIService.UploadArticle)
	}

	return
}
