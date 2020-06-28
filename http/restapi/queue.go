package restapi

import (
	"github.com/appointment/http/models"
	"github.com/appointment/http/restapi/operations/queue"
	rec "github.com/appointment/resources"
	middleware "github.com/go-openapi/runtime/middleware"
)

type queueGetter interface {
	GetQueue() (rec.Queue, error)
}

func getQueue(stg queueGetter) queue.GetQueueHandlerFunc {
	return func(params queue.GetQueueParams) middleware.Responder {
		result, err := stg.GetQueue()
		if err != nil {
			return queue.NewGetQueueInternalServerError().WithPayload(newRestApiError(err))
		}

		response := &models.Queue{}
		for _, item := range result.Items {
			response.Items = append(response.Items, &models.QueueItem{
				ID:          item.ID,
				IDPatient:   &item.IDPatient,
				IDSpecialty: &item.IDSpecialty,
				Status:      item.Status,
			})

		}

		return queue.NewGetQueueOK().WithPayload(response)
	}
}

type queueAdderItem interface {
	AddItem(rec.QueueItem) error
}

func addItem(stg queueAdderItem) queue.AddPatientToQueueHandlerFunc {
	return func(params queue.AddPatientToQueueParams) middleware.Responder {
		queueItem := rec.QueueItem{
			IDPatient:   *params.Item.IDPatient,
			IDSpecialty: *params.Item.IDSpecialty,
		}
		err := stg.AddItem(queueItem)
		if err != nil {
			return queue.NewAddPatientToQueueInternalServerError().WithPayload(newRestApiError(rec.PatientIsAlreadyWaiting))
		}

		return queue.NewAddPatientToQueueOK()
	}
}
