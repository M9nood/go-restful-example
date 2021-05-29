package event

type EventListResponse struct {
	Events    []EventModel `gorm:"column:events" json:"events"`
	TotalRows int          `json:"totalRows"`
}
