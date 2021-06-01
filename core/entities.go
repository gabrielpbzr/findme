package core

import (
	"time"

	"github.com/google/uuid"
)

type Position struct {
	Longitude float64   `json:"longitude"`
	Latitude  float64   `json:"latitude"`
	Timestamp time.Time `json:"timestamp"`
	Id        uuid.UUID `json:"id"`
}

func CreatePosition(longitude float64, latitude float64) *Position {
	position := &Position{Longitude: longitude, Latitude: latitude, Timestamp: time.Now().UTC()}
	position.Id, _ = uuid.NewUUID()
	return position
}
