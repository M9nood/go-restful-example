package event

type EventModel struct {
	Id          int    `gorm:"column:id" json:"id"`
	EventName   string `gorm:"column:event_name" json:"event_name"`
	Description string `gorm:"column:description" json:"description"`
	StartTime   string `gorm:"column:start_time" json:"start_time"`
	EndTime     string `gorm:"column:end_time" json:"end_time"`
	TotalSeat   string `gorm:"column:total_seat" json:"total_seat"`
}
