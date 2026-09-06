package auction

import (
	"context"
	"fmt"
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

func (s *Service) Process(
	ctx context.Context,
	req domain.AuctionRequest,
) (domain.AuctionResult, error) {
	partners, err := s.partners.List(ctx)

	if err != nil {
		return domain.AuctionResult{}, fmt.Errorf("list partners: %w", err)
	}

	matched := Filter(req, partners)

	return domain.AuctionResult{}, nil
}
