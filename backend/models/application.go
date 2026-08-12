package models

import "time"

type Application struct {
	ID             int64     `json:"id"`
	UserID         string    `json:"user_id"`
	CompanyID      int64     `json:"company_id"`
	Position       string    `json:"position"`
	Source         string    `json:"source"`
	JobURL         *string   `json:"job_url,omitempty"`
	Location       *string   `json:"location,omitempty"`
	EmploymentType *string   `json:"employment_type,omitempty"`
	SalaryMin      *float64  `json:"salary_min,omitempty"`
	SalaryMax      *float64  `json:"salary_max,omitempty"`
	AppliedAt      time.Time `json:"applied_at"`
	Status         string    `json:"status"`
	Notes          *string   `json:"notes,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
