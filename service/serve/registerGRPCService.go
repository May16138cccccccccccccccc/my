package serve

import (
	"context"
	v1 "moredoc/api/v1"
	"moredoc/biz"
	"moredoc/middleware/auth"
	"moredoc/model"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RegisterGRPCService 注册grpc服务
func RegisterGRPCService(dbModel *model.DBModel, logger *zap.Logger, endpoint string, authMiddleWare *auth.Auth, grpcServer *grpc.Server, gwmux *runtime.ServeMux, dialOpts ...grpc.DialOption) (err error) {
	// 用户API接口服务
	userAPIService := biz.NewUserAPIService(dbModel, logger, authMiddleWare)
	v1.RegisterUserAPIServer(grpcServer, userAPIService)
	err = v1.RegisterUserAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterUserAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 分组API接口服务
	groupAPIService := biz.NewGroupAPIService(dbModel, logger)
	v1.RegisterGroupAPIServer(grpcServer, groupAPIService)
	err = v1.RegisterGroupAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterGroupAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 友链API接口服务
	friendlinkAPIService := biz.NewFriendlinkAPIService(dbModel, logger)
	v1.RegisterFriendlinkAPIServer(grpcServer, friendlinkAPIService)
	err = v1.RegisterFriendlinkAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterFriendlinkAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 附件API接口服务
	attachmentAPIService := biz.NewAttachmentAPIService(dbModel, logger)
	v1.RegisterAttachmentAPIServer(grpcServer, attachmentAPIService)
	err = v1.RegisterAttachmentAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterAttachmentAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 轮播图API接口服务
	bannerAPIService := biz.NewBannerAPIService(dbModel, logger)
	v1.RegisterBannerAPIServer(grpcServer, bannerAPIService)
	err = v1.RegisterBannerAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterBannerAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 权限API接口服务
	permissionAPIService := biz.NewPermissionAPIService(dbModel, logger)
	v1.RegisterPermissionAPIServer(grpcServer, permissionAPIService)
	err = v1.RegisterPermissionAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterPermissionAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// Config API接口服务
	configAPIService := biz.NewConfigAPIService(dbModel, logger)
	v1.RegisterConfigAPIServer(grpcServer, configAPIService)
	err = v1.RegisterConfigAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterConfigAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 注册分类服务
	categoryAPIService := biz.NewCategoryAPIService(dbModel, logger)
	v1.RegisterCategoryAPIServer(grpcServer, categoryAPIService)
	err = v1.RegisterCategoryAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterCategoryAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 注册文档服务
	documentAPIService := biz.NewDocumentAPIService(dbModel, logger)
	v1.RegisterDocumentAPIServer(grpcServer, documentAPIService)
	err = v1.RegisterDocumentAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterDocumentAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 文档回收站服务
	v1.RegisterRecycleAPIServer(grpcServer, documentAPIService)
	err = v1.RegisterRecycleAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterRecycleAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 文章服务
	articleAPIService := biz.NewArticleAPIService(dbModel, logger)
	v1.RegisterArticleAPIServer(grpcServer, articleAPIService)
	err = v1.RegisterArticleAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterArticleAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 收藏服务
	favoriteAPIService := biz.NewFavoriteAPIService(dbModel, logger)
	v1.RegisterFavoriteAPIServer(grpcServer, favoriteAPIService)
	err = v1.RegisterFavoriteAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterFavoriteAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 评论服务
	commentAPIService := biz.NewCommentAPIService(dbModel, logger)
	v1.RegisterCommentAPIServer(grpcServer, commentAPIService)
	err = v1.RegisterCommentAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterCommentAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 举报服务
	reportAPIService := biz.NewReportAPIService(dbModel, logger)
	v1.RegisterReportAPIServer(grpcServer, reportAPIService)
	err = v1.RegisterReportAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterReportAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 导航服务
	navgationAPIService := biz.NewNavigationAPIService(dbModel, logger)
	v1.RegisterNavigationAPIServer(grpcServer, navgationAPIService)
	err = v1.RegisterNavigationAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterNavigationAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 惩罚服务
	punishmentAPIService := biz.NewPunishmentAPIService(dbModel, logger)
	v1.RegisterPunishmentAPIServer(grpcServer, punishmentAPIService)
	err = v1.RegisterPunishmentAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterPunishmentAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 广告服务
	advertisementAPIService := biz.NewAdvertisementAPIService(dbModel, logger)
	v1.RegisterAdvertisementAPIServer(grpcServer, advertisementAPIService)
	err = v1.RegisterAdvertisementAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterAdvertisementAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 搜索记录服务
	searchRecordAPIService := biz.NewSearchRecordAPIService(dbModel, logger)
	v1.RegisterSearchRecordAPIServer(grpcServer, searchRecordAPIService)
	err = v1.RegisterSearchRecordAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterSearchRecordAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 语言服务
	languageAPIService := biz.NewLanguageAPIService(dbModel, logger)
	v1.RegisterLanguageAPIServer(grpcServer, languageAPIService)
	err = v1.RegisterLanguageAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterLanguageAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 注册订单服务（支付宝测试版）
	// 初始化支付宝client
	appID := "9021000145623685"
	privateKey := "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCuV/a/FvjDSCghNE42hrff8fUMvC3wHAVwBzagednh9UreEmisWE0ZIlO66cyzMfuEgLW19N72AVuYacuLDdhXVJ14/dLAKA0mMzvxvS8+G7ykrmxeHZhnw28e2py7AxMneN5hrBA6OyhUbsSf8QhpEktaW/1Fequ6JmG+eh6N8oZpBEJefcL8ffChAhacWrb3vdNHH6Mck9yO+VRSJ/xlAShJIeJxP9Jxj74cMRDcrFo0Xt/4aKD6tEO1YxyiWuV2aACO0AwzopZ5cNTwHR9cAO+HyCGyXmUGnQUT/44peJoaOKcMcurIymrki2eL1mTjb5XXlBtZ7PgxYMIjZzjpAgMBAAECggEAfPLMA226QXKgWBO9jjSE6XnmDwd6wN/EQeLZkq9hqSB8VKXK4OGz97RzX70aGL9UrET2df6WIKcedyAzYWg9yXD6HvGCrnbF3b/QUVMt3YxRaZEcV2NMi+kz5V/1/c2ZV6u3bTa8to8ZO1Hbl2lOhsc2u/67iKT/GD1TxgiNh1q6Be7f3H0LIO37txvkFfAeBQYMS1KIHeLHyY0NnJzUGRGUbZ37e/dELdQPH2XAUFgd/hscYxXPrIbcOhEBkPpmMg1VebkDYVGTIwhAbDaeI3IN8hLHwxYetbMD5wtZWLMKqhEWDhdzmN5lCjw7GIcuO6SMBRrMfr6srfQHT9XJZQKBgQD/nfFYisBgC5iaKzW6c2GEUsXUGawTW2vT47DmYOG77/watAZR4voPv7rvOvmT8X8oga3I9YCk4vj85EzzQpQGwj1lfdx0HLRtyoDaAkXkxra64Aq5DdgECslIkV8ZZk3qxHqRQFS7mZtGbKzH++iZUcNezEcPXhQnwPou6VXmawKBgQCumtgEd3nTS7ceyKWE9Q6mngkTnVJngwyySKijm+WbV9hxuiKuri+b+twlBr5TuryITyUgVJnotY+5Sr5mtorARCnHo6mvoi1uH6ZPYAVhhT+PFQnO1qXAokG0tHgnWDYbYhVmYtMCM0ybmuzWDRy8QR3whJAB19MEo71YFdJq+wKBgQD474grAnV3E9P50Dry6Yr+ot2mQZoi/9Vy9TtCIPe1X3HVT50YAkQlMkOK15RXT8jWWfQ00WIHMkPphSyrextNubyxKnGXYp0UjzINHkmTEzLBq47FyYWl6hs8YNaMleUrzUzQSCY8zMs7OnrKB7nuFoJ9v9M0VBqs3HLJXNQbIQKBgDZMp1M0SKmiSU+EfJ0NjMHaMd5stXzFb4RePjurNsuJlSFBDpoeR2YNrLrDDxuA3jLR0izswnRxSkIIcZAFwr6qNAgG6LVbDR/LLfBPXz0FsE9x7nvpmuB+VUq6OQebUjBP33HY1+A+TduyYWbr6vzMj2AGCbbZBKHTvgWPu8FTAoGBAML5xHUUGZV+qt2Xb8YA4h61X3sX1zVimHwMGn5gn7uH4ceAGqZWjuvNO9xw/pJvGVSm+HqD0RAwWTc8Nk9rWa9npbE2n8LgmYAZqWBUdIeTQkl4CXJLNAAnRpPMYho8a/UQUYi5kvqkagNmfUIlc63LsnvOADcbNs1y+1KBrKr8"
	alipayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAj0PuMZcyV7+Cjl/Z5EcRSj3WPJP+S3uRLFwcT0rLJBTrJpkcB3c7gd/QSMApgOMHpOwi7QPnSv1ADS4KDHAUGnNR+63Fhdpm4tvIwSIctFkbX2jPT7g/F1e+ofJVa2NZ/rvR44zy59PyI5Lp9/sVAETZBFvCzFaysoATrxVpVlkS1MtGfVEiSXzHBZraqd5P66Gj0laqwKFrxAFf7ESn8qhK+aXaWIcfH/v5Me5BrdfGBVCtujxyGKbGwCBZfc2tdUrbH7Zv6f/HIh/uVnKuYTMmaUOhcmgCz7IKcm11eOkOkO8NWU2p5V+yIgQYCG7p77jDe6VAYupsydCN8Ac99wIDAQAB"
	paymentService := biz.NewPaymentService(appID, privateKey, alipayPublicKey)
	v1.RegisterPaymentServiceServer(grpcServer, paymentService)
	err = v1.RegisterPaymentServiceHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterPaymentServiceHandlerFromEndpoint", zap.Error(err))
		return
	}

	return
}
