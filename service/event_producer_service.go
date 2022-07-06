package service

import (
	"github.com/gin-gonic/gin"
	"github.com/raviMukti/gin-restful-api/helper"
	"github.com/raviMukti/gin-restful-api/messaging"
	"github.com/raviMukti/gin-restful-api/model/web/dto"
)

type EventProducerService interface {
	OrderEventProducer(ctx *gin.Context, order *dto.EventRequest)
	PurchaseEventProducer(ctx *gin.Context, purchase *dto.EventRequest)
}

type EventProducerServiceImpl struct {
}

func NewEventServiceProducer() EventProducerService {
	return &EventProducerServiceImpl{}
}

// OrderEventProducer implements EventServiceProducer
func (service *EventProducerServiceImpl) OrderEventProducer(ctx *gin.Context, order *dto.EventRequest) {
	kafka := messaging.NewKafkaSyncProducer()
	event := messaging.NewOrderEvent(order.Code)
	err := messaging.SendMsg(kafka, event)
	helper.PanicIfError(err)
}

// PurchaseEventProducer implements EventServiceProducer
func (service *EventProducerServiceImpl) PurchaseEventProducer(ctx *gin.Context, purchase *dto.EventRequest) {
	kafka := messaging.NewKafkaSyncProducer()
	event := messaging.NewPurchaseEvent(purchase.Code)
	err := messaging.SendMsg(kafka, event)
	helper.PanicIfError(err)
}
