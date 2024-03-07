package model

import "time"

type Holiday struct {
	Date time.Time `json:"date"`
	Name string    `json:"name"`
}
