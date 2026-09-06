package auction

import (
	"context"
	"mini-adex/internal/domain"
	"testing"
	"time"
)

type fakePartnerRepository struct {
	partners []domain.Partner
	err      error
}

func (r *fakePartnerRepository) List(
	ctx context.Context,
) ([]domain.Partner, error) {
	return r.partners, r.err
}

type fakeDSPClient struct {
	err error
}

func (c *fakeDSPClient) Bid(
	ctx context.Context,
	partner domain.Partner,
	req domain.AuctionRequest,
) error {
	return c.err
}

type serviceFixture struct {
	request domain.AuctionRequest
	repo    *fakePartnerRepository
	dsp     *fakeDSPClient
}

func newServiceFixture() serviceFixture {
	return serviceFixture{
		request: domain.AuctionRequest{
			RequestID:  "request-1",
			Country:    "RU",
			DeviceType: "mobile",
			BidFloor:   1.5,
			Categories: []string{"news"},
		},
		repo: &fakePartnerRepository{
			partners: []domain.Partner{
				{
					UID:         "dsp-alpha",
					IsEnabled:   true,
					Countries:   []string{"RU"},
					DeviceTypes: []string{"mobile"},
				},
				{
					UID:         "dsp-beta",
					IsEnabled:   true,
					Countries:   []string{"RU"},
					DeviceTypes: []string{"mobile"},
				},
				{
					UID:         "dsp-us",
					IsEnabled:   true,
					Countries:   []string{"US"},
					DeviceTypes: []string{"mobile"},
				},
			},
		},
		dsp: &fakeDSPClient{},
	}
}

func (f serviceFixture) service() *Service {
	return NewService(
		f.repo,
		f.dsp,
		200*time.Millisecond,
	)
}

func TestService_Process(t *testing.T) {
	f := newServiceFixture()

	result, err := f.service().Process(
		context.Background(),
		f.request,
	)

	if err != nil {
		t.Fatalf("Process() returned error: %v", err)
	}

	if result.RequestID != f.request.RequestID {
		t.Errorf(
			"Process() RequestID = %q, want %q",
			result.RequestID,
			f.request.RequestID,
		)
	}

	if result.Sent != 2 {
		t.Errorf(
			"Process() Sent = %d, want 2",
			result.Sent,
		)
	}

	if result.Succeeded != 2 {
		t.Errorf(
			"Process() Succeeded = %d, want 2",
			result.Succeeded,
		)
	}

	if len(result.MatchedDSPs) != 2 {
		t.Errorf(
			"Process() matched %d DSPs, want 2",
			len(result.MatchedDSPs),
		)
	}
}
