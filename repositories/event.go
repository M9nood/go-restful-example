package repositories

import (
	"fmt"
	eventModel "m9hub/models/event"

	"m9hub/database"

	"gorm.io/gorm"
)

type EventReposity struct {
	Db *gorm.DB
}

func NewEventReposity(Db *gorm.DB) *EventReposity {
	var db *gorm.DB
	if Db == nil {
		db = database.GetDB()
	} else {
		db = Db
	}
	return &EventReposity{
		Db: db,
	}
}

func (repo *EventReposity) GetAllEvents() (eventModel.EventListResponse, error) {
	var events = make([]eventModel.EventModel, 0)
	result := repo.Db.Table("events").Select(`
			id,
			event_name,
			description, 
			DATE_FORMAT(start_time,'%Y-%m-%d %H:%i') as start_time ,
			DATE_FORMAT(end_time,'%Y-%m-%d %H:%i') as end_time,
			total_seat`).Where("delete_flag = 0").Scan(&events)
	fmt.Println("result", result)
	if result.Error != nil {
		return eventModel.EventListResponse{
			Events:    events,
			TotalRows: 0,
		}, result.Error
	}

	return eventModel.EventListResponse{
		Events:    events,
		TotalRows: 2,
	}, nil
}
