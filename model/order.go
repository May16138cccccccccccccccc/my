package model

import (
	"go.uber.org/zap"
	"time"
)

// 订单实体
type Order struct {
	Id          int64
	UserId      int64
	OrderNo     int64
	ProductName string
	Price       int
	Quantity    int
	Amount      int
	Status      int8
	OrderType   int8
	ClientType  int8
	Ip          string
	CreatedAt   *time.Time
	ClosedAt    *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}

func (Order) TableName() string {
	return tablePrefix + "order"
}

// CreateOrder 创建order
func (m *DBModel) CreateOrder(order *Order) (err error) {
	err = m.db.Create(order).Error
	if err != nil {
		m.logger.Error("CreateOrder", zap.Error(err))
		return
	}
	return
}
