package services

import (
	"m9hub/database"

	eventModel "m9hub/models/event"
	"m9hub/repositories"

	"gorm.io/gorm"
)

type EventService struct {
	Db *gorm.DB
}

var eventServiceInstance *EventService

func NewEventService() *EventService {
	if eventServiceInstance == nil {
		db := database.GetDB()
		eventServiceInstance = &EventService{
			Db: db,
		}
	}
	return eventServiceInstance
}

func (ev EventService) GetAllEventsService() (eventModel.EventListResponse, error) {
	eventServiceInstance.Db = nil
	eventRepo := repositories.NewEventReposity(eventServiceInstance.Db)
	result, err := eventRepo.GetAllEvents()
	if err != nil {
		return eventModel.EventListResponse{}, err
	}
	return result, nil
}
