package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/AchrafIbrahim/WORK_APPLY_BOT/models"
	"github.com/AchrafIbrahim/WORK_APPLY_BOT/repositories"
)

type ApplicationService struct {
	repository *repositories.ApplicationRepository
}

func NewApplicationService(
	repository *repositories.ApplicationRepository,
) *ApplicationService {
	return &ApplicationService{
		repository: repository,
	}
}

func (s *ApplicationService) CreateApplication(
	ctx context.Context,
	application *models.Application,
) error {
	if application.UserID == "" {
		return fmt.Errorf("user_id is required")
	}

	if application.CompanyID <= 0 {
		return fmt.Errorf("company_id is required")
	}

	if strings.TrimSpace(application.Position) == "" {
		return fmt.Errorf("position is required")
	}

	if strings.TrimSpace(application.Source) == "" {
		return fmt.Errorf("source is required")
	}

	if application.AppliedAt.IsZero() {
		application.AppliedAt = time.Now()
	}

	if strings.TrimSpace(application.Status) == "" {
		application.Status = "Applied"
	}

	return s.repository.CreateApplication(ctx, application)
}
