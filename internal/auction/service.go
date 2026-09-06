package auction

import (
	"context"
	"fmt"
	"mini-adex/internal/domain"
	"time"
)

type PartnerRepository interface {
	List(ctx context.Context) ([]domain.Partner, error)
}

type DSPClient interface {
	Bid(
		ctx context.Context,
		partner domain.Partner,
		req domain.AuctionRequest,
	)
}

type Service struct {
	partners PartnerRepository
	dsp      DSPClient
	timeout time.Duration
}

func NewService(
	partners PartnerRepository,
	dsp DSPClient,
	timeout time.Duration,
	) *Service {
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
