package repositories

import (
	"context"
	"fmt"
	"github.com/AchrafIbrahim/WORK_APPLY_BOT/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ApplicationRepository struct {
	db *pgxpool.Pool
}

func NewApplicationRepository(db *pgxpool.Pool) *ApplicationRepository {
	return &ApplicationRepository{
		db: db,
	}
}

func (r *ApplicationRepository) CreateApplication(
	ctx context.Context,
	application *models.Application,
) error {
	query := `
		INSERT INTO applications (
			user_id,
			company_id,
			position,
			source,
			job_url,
			location,
			employment_type,
			salary_min,
			salary_max,
			applied_at,
			status,
			notes
		)
		VALUES (
			$1, $2, $3, $4, $5, $6,
			$7, $8, $9, $10, $11, $12
		)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRow(
		ctx,
		query,
		application.UserID,
		application.CompanyID,
		application.Position,
		application.Source,
		application.JobURL,
		application.Location,
		application.EmploymentType,
		application.SalaryMin,
		application.SalaryMax,
		application.AppliedAt,
		application.Status,
		application.Notes,
	).Scan(
		&application.ID,
		&application.CreatedAt,
		&application.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf(
			"failed to create application: %w",
			err,
		)
	}

	return nil
}
