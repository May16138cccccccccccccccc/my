package biz

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	pb "moredoc/api/v1"
	"moredoc/middleware/auth"
	"moredoc/model"
	"moredoc/util"
	"net/http"
	"time"
)

type PaymentService struct {
	client *alipay.Client

	pb.UnimplementedPaymentServiceServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

// 新建支付宝客户端并初始化参数
func NewPaymentService(appID, privateKey, alipayPublicKey string) *PaymentService {
	client, _ := alipay.New(appID, privateKey, false)

	client.LoadAliPayPublicKey(alipayPublicKey)
	return &PaymentService{client: client}
}

func (s *PaymentService) checkPermission(ctx context.Context) (userClaims *auth.UserClaims, err error) {
	return checkGRPCPermission(s.dbModel, ctx)
}

// 创建订单
func (s *PaymentService) CreateOrder(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	// 鉴权
	userCliams, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	// 初始化订单实体
	order := &model.Order{}
	// 前端传递参数复制给订单实体字段
	if err = util.CopyStruct(req, order); err != nil {
		s.logger.Error("CreateOrder", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "创建订单失败:"+err.Error())
	}

	// 字段判空

	// 调用model层方法，写入数据到数据库
	order.UserId = userCliams.UserId
	if err = s.dbModel.CreateOrder(order); err != nil {
		s.logger.Error("CreateOrder", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "创建广告失败:"+err.Error())
	}

	// 返回响应结果 需要更改结构体定义
	res := &pb.CreatePaymentResponse{
		PaymentUrl: "创建订单成功",
	}

	return res, nil
}

func (s *PaymentService) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	// 创建订单（数据库）
	// 创建支付请求
	p := alipay.TradePagePay{
		Trade: alipay.Trade{
			// 订单号
			OutTradeNo: generateOrderNo(),
			// 订单金额
			TotalAmount: formatAmount(req.Amount),
			// 订单标题
			Subject: req.Subject,
			// 支付产品模式代码（代表网页支付）
			ProductCode: "FAST_INSTANT_TRADE_PAY",
			// 支付宝异步回调地址（返回订单状态到服务器）
			NotifyURL: "https://902c-171-212-106-136.ngrok-free.app/order_notify",
			// 支付宝同步回调地址（支付后页面跳转地址）
			ReturnURL: "http://192.168.199.1:5555/order",
		},
	}

	url, err := s.client.TradePagePay(p)
	if err != nil {
		return nil, status.Error(codes.Internal, "支付创建失败")
	}

	return &pb.CreatePaymentResponse{
		PaymentUrl: url.String(),
	}, nil
}

// 接收支付宝异步回调请求（请求地址必须为公网地址，可用内网穿透工具）
func (s *PaymentService) HandleAlipayNotify(ctx *gin.Context) {
	// 解析请求参数
	ctx.Request.ParseForm()

	// 验证签名
	// DecodeNotification 内部已调用 VerifySign 方法验证签名
	noti, err := s.client.DecodeNotification(ctx.Request.Form)
	if err != nil {
		fmt.Println("验签失败！")
		ctx.String(http.StatusBadRequest, "failure")
		return
	}

	// 业务处理
	if noti.TradeStatus == alipay.TradeStatusSuccess {
		// 更新订单状态（需幂等处理）
		// err = UpdateOrderStatus(noti.OutTradeNo, "PAID")
		fmt.Println("正在更新订单...")
		if err != nil {
			ctx.String(http.StatusOK, "success") // 即使失败也要返回success
			return
		}
	}

	// 必须返回success告知支付宝已处理
	ctx.String(200, "success")
}

// 生成唯一订单号（示例实现）
func generateOrderNo() string {
	// 时间戳（纳秒）+ 随机数
	return fmt.Sprintf("%d%04d",
		time.Now().UnixNano(),
		rand.Intn(10000))
}

// 金额格式化（保留两位小数）
func formatAmount(amount float64) string {
	return fmt.Sprintf("%.2f", amount)
}
