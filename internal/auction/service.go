package auction

import (
	"context"
	"mini-adex/internal/domain"
)

type PartnerRepository interface {
	List(ctx context.Context) ([]domain.Partner, error)
}

type Service struct {
	partners PartnerRepository
}

func NewService(partners PartnerRepository) *Service {
	return &Service{
		partners: partners,
	}
}
