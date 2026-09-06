package domain

type AuctionRequest struct {
	RequestID string
	Country string
	DeviceType string
	BidFloor float64
	Categories []string
}

type AuctionResult struct {
	RequestID string
	MatchedDSPs []string
	Sent int
	Succeeded int
	DurationMS int64 
}
