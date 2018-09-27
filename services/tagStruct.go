package services

import "time"

type TagShow struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
