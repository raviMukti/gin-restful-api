package service

import (
	"context"

	"github.com/raviMukti/gin-restful-api/messaging"
	"github.com/raviMukti/gin-restful-api/model/domain"
	"github.com/raviMukti/gin-restful-api/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProcessEventService interface {
	ProcessOrderEvent(e messaging.OrderEvent) error
	ProcessPurchaseEvent(e messaging.PurchaseEvent) error
}

type ProcessEventServiceImpl struct {
	EventLogRepository repository.EventLogRepository
}

func NewProcessEvent(repository repository.EventLogRepository) ProcessEventService {
	return &ProcessEventServiceImpl{
		EventLogRepository: repository,
	}
}

// ProcessOrderEvent implements ProcessEvent
func (process *ProcessEventServiceImpl) ProcessOrderEvent(e messaging.OrderEvent) error {
	order := domain.EventLog{
		Id:   primitive.NewObjectID(),
		Type: e.Type,
		Code: e.Code,
	}

	err := process.EventLogRepository.Save(context.TODO(), order)
	return err
}

// ProcessPurchaseEvent implements ProcessEvent
func (process *ProcessEventServiceImpl) ProcessPurchaseEvent(e messaging.PurchaseEvent) error {
	purchase := domain.EventLog{
		Id:   primitive.NewObjectID(),
		Type: e.Type,
		Code: e.Code,
	}

	err := process.EventLogRepository.Save(context.TODO(), purchase)
	return err
}
