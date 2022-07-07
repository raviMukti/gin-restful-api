package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raviMukti/gin-restful-api/exception"
	"github.com/raviMukti/gin-restful-api/helper"
	"github.com/raviMukti/gin-restful-api/model/web"
	"github.com/raviMukti/gin-restful-api/model/web/dto"
	"github.com/raviMukti/gin-restful-api/service"
)

type EventProducerController interface {
	CreateEvent(ctx *gin.Context)
}

type EventProducerControllerImpl struct {
	EventProducerService service.EventProducerService
}

func NewEventProducerController(eventService service.EventProducerService) EventProducerController {
	return &EventProducerControllerImpl{
		EventProducerService: eventService,
	}
}

// CreateEvent implements EventProducerController
func (controller *EventProducerControllerImpl) CreateEvent(ctx *gin.Context) {
	var eventRequest *dto.EventRequest = &dto.EventRequest{}

	err := ctx.ShouldBindJSON(eventRequest)
	helper.PanicIfError(err)

	if eventRequest.Type == "Order" {
		controller.EventProducerService.OrderEventProducer(ctx, eventRequest)
	} else if eventRequest.Type == "Purchase" {
		controller.EventProducerService.PurchaseEventProducer(ctx, eventRequest)
	} else {
		panic(exception.NewNotFoundError("Event Type Not Found"))
	}

	webResponse := web.WebResponse{
		Error:       "",
		ErrorDetail: "",
		Code:        http.StatusOK,
		Message:     http.StatusText(http.StatusOK),
		Data:        "Event " + eventRequest.Type + " Created!!!",
	}

	ctx.JSON(http.StatusOK, webResponse)
}
