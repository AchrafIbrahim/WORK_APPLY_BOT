package models

import "time"

type Company struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Website   *string   `json:"website,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
