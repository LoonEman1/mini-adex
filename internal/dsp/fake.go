package dsp

import (
	"context"
	"errors"
	"math/rand/v2"
	"mini-adex/internal/domain"
	"time"
)

type FakeClient struct{}

func NewFakeClient() *FakeClient {
	return &FakeClient{}
}

func (c *FakeClient) Bid(
	ctx context.Context,
	partner domain.Partner,
	req domain.AuctionRequest,
) error {
	delay := time.Duration(rand.IntN(400)) * time.Millisecond

	timer := time.NewTimer(delay)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()

	case <-timer.C:
		if rand.IntN(100) < 20 {
			return errors.New("DSP request failed")
		}

		return nil
	}
}
