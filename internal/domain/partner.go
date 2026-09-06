package domain

type Partner struct {
	UID               string
	Name              string
	Endpoint          string
	IsEnabled         bool
	Countries         []string
	DeviceTypes       []string
	MinBidFloor       float64
	BlockedCategories []string
}
