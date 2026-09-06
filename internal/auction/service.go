package auction

import (
	"context"
	"fmt"
	"mini-adex/internal/domain"
	"sync"
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
	) error
}

type Service struct {
	partners PartnerRepository
	dsp      DSPClient
	timeout  time.Duration
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
	started := time.Now()

	partners, err := s.partners.List(ctx)

	if err != nil {
		return domain.AuctionResult{}, fmt.Errorf("list partners: %w", err)
	}

	matched := Filter(req, partners)
	successful := s.sendToPartners(ctx, req, partners)

	return domain.AuctionResult{
		RequestID:   req.RequestID,
		MatchedDSPs: partnerUIDs(matched),
		Sent:        len(matched),
		Succeeded:   successful,
		DurationMS:  int64(time.Since(started).Milliseconds()),
	}, nil
}

func (s *Service) sendToPartners(
	ctx context.Context,
	req domain.AuctionRequest,
	partners []domain.Partner,
) int {
	auctionCtx, cancel := context.WithTimeout(ctx, s.timeout)

	defer cancel()

	var wg sync.WaitGroup

	resultsChar := make(chan error, len(partners))

	for _, partner := range partners {
		wg.Add(1)

		go func() {
			defer wg.Done()

			resultsChar <- s.dsp.Bid(auctionCtx, partner, req)

		}()
	}

	wg.Wait()
	close(resultsChar)

	successful := 0

	for err := range resultsChar {
		if err == nil {
			successful++
		}
	}

	return successful
}

func partnerUIDs(partners []domain.Partner) []string {
	uids := make([]string, 0, len(partners))

	for _, partner := range partners {
		uids = append(uids, partner.UID)
	}

	return uids
}
