package event

type EventListResponse struct {
	Events    []EventModel `gorm:"column:events" json:"events"`
	TotalRows int64        `json:"totalRows"`
}
