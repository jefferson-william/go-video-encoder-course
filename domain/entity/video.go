package entity

import "time"

type Video struct {
	ID string
	ResourceId string
	FilePath string
	CreatedAt time.Time
}