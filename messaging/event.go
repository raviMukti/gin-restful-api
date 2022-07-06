package messaging

type Event struct {
	Type string
}

type OrderEvent struct {
	Event
	Code string
}

type PurchaseEvent struct {
	Event
	Code string
}

func NewOrderEvent(orderCode string) OrderEvent {
	event := new(OrderEvent)
	event.Type = "OrderEvent"
	event.Code = orderCode
	return *event
}

func NewPurchaseEvent(purchaseCode string) PurchaseEvent {
	event := new(PurchaseEvent)
	event.Type = "PurchaseEvent"
	event.Code = purchaseCode
	return *event
}
