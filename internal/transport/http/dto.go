package httptransport

import (
	"errors"
	"mini-adex/internal/domain"
)

type auctionRequest struct {
	RequestID  string   `json:"request_id"`
	Country    string   `json:"country"`
	DeviceType string   `json:"device_type"`
	BidFloor   float64  `json:"bid_floor"`
	Categories []string `json:"categories"`
}

type auctionResponse struct {
	RequestID   string   `json:"request_id"`
	MatchedDSPs []string `json:"matched_dsps"`
	Sent        int      `json:"sent"`
	Succeeded   int      `json:"succeeded"`
	DurationMS  int64    `json:"duration_ms"`
}

func (r auctionRequest) validate() error {
	if r.RequestID == "" {
		return errors.New("request_id is required")
	}

	if len(r.Country) != 2 {
		return errors.New("country must be a two-letter code")
	}

	if !isValidDeviceType(r.DeviceType) {
		return errors.New("device_type must be mobile, desktop or tv")
	}

	if r.BidFloor < 0 {
		return errors.New("bid_floor must be greater than or equal to zero")
	}

	return nil
}

func isValidDeviceType(deviceType string) bool {
	return deviceType == "mobile" || deviceType == "desktop" || deviceType == "tv"
}

func (r auctionRequest) toDomain() domain.AuctionRequest {
	return domain.AuctionRequest{
		RequestID:  r.RequestID,
		Country:    r.Country,
		DeviceType: r.DeviceType,
		BidFloor:   r.BidFloor,
		Categories: r.Categories,
	}
}

func newAuctionResponse(result domain.AuctionResult) auctionResponse {
	return auctionResponse{
		RequestID:   result.RequestID,
		MatchedDSPs: result.MatchedDSPs,
		Sent:        result.Sent,
		Succeeded:   result.Succeeded,
		DurationMS:  result.DurationMS,
	}
}
