package auction

import (
	"mini-adex/internal/domain"
	"testing"
)

type targetingFixture struct {
	request domain.AuctionRequest
	partner domain.Partner
}

func newTargetingFixture() targetingFixture {
	return targetingFixture{
		request: domain.AuctionRequest{
			RequestID:  "request-1",
			Country:    "RU",
			DeviceType: "mobile",
			BidFloor:   1.5,
			Categories: []string{"news", "sport"},
		},
		partner: domain.Partner{
			UID:               "dsp-alpha",
			IsEnabled:         true,
			Countries:         []string{"RU", "KZ"},
			DeviceTypes:       []string{"mobile", "desktop"},
			MinBidFloor:       0.5,
			BlockedCategories: []string{"gambling"},
		},
	}
}

func TestMatch_DisabledPartner(t *testing.T) {
	f := newTargetingFixture()

	f.partner.IsEnabled = false

	isMatched := Match(f.request, f.partner)

	if isMatched {
		t.Errorf(
			"Match() = true for partner %q with IsEnabled=%v, want false",
			f.partner.UID,
			f.partner.IsEnabled,
		)
	}
}

func TestMatch_CountryMismatch(t *testing.T) {
	f := newTargetingFixture()

	f.request.Country = "US"

	isMatched := Match(f.request, f.partner)

	if isMatched {
		t.Errorf(
			"Match() = true for request country %q, partner countries %v, want false",
			f.request.Country,
			f.partner.Countries,
		)
	}
}

func TestMatch_DevicesMismatch(t *testing.T) {
	f := newTargetingFixture()

	f.request.DeviceType = "tv"

	isMatched := Match(f.request, f.partner)

	if isMatched {
		t.Errorf(
			"Match() = true for device type %q, partner supports %v, want false",
			f.request.DeviceType,
			f.partner.DeviceTypes,
		)
	}
}

func TestMatch_BidFloorBelowMinimum(t *testing.T) {
	f := newTargetingFixture()

	f.request.BidFloor = 0.4

	isMatched := Match(f.request, f.partner)

	if isMatched {
		t.Errorf(
			"Match() = true when BidFloor %f is below partner minimum %f, want false",
			f.request.BidFloor,
			f.partner.MinBidFloor,
		)
	}
}

func TestMatch_BlockedCategories(t *testing.T) {
	f := newTargetingFixture()

	blockedCategory := "gambling"

	f.request.Categories = append(f.request.Categories, blockedCategory)

	isMatched := Match(f.request, f.partner)

	if isMatched {
		t.Errorf(
			"Match() = true for blocked category %q, want false",
			blockedCategory,
		)
	}
}

func TestMatch_EmptyCountriesMatchesAnyCountry(t *testing.T) {
	f := newTargetingFixture()

	f.partner.Countries = nil
	f.request.Country = "US"

	if !Match(f.request, f.partner) {
		t.Error("Match() = false for empty partner countries, want true")
	}
}

func TestMatch_EmptyDeviceTypesMatchesAnyDevice(t *testing.T) {
	f := newTargetingFixture()

	f.partner.DeviceTypes = nil
	f.request.DeviceType = "tv"

	if !Match(f.request, f.partner) {
		t.Error("Match() = false for empty partner device types, want true")
	}
}

func TestMatch_BidFloorEqualsMinimum(t *testing.T) {
	f := newTargetingFixture()

	f.request.BidFloor = f.partner.MinBidFloor

	if !Match(f.request, f.partner) {
		t.Errorf(
			"Match() = false when BidFloor %f equals partner minimum %f, want true",
			f.request.BidFloor,
			f.partner.MinBidFloor,
		)
	}
}

func TestMatch_EmptyRequestCategories(t *testing.T) {
	f := newTargetingFixture()

	f.request.Categories = nil

	if !Match(f.request, f.partner) {
		t.Error("Match() = false for empty request categories, want true")
	}
}

func TestMatch_EmptyBlockedCategories(t *testing.T) {
	f := newTargetingFixture()

	f.partner.BlockedCategories = nil

	if !Match(f.request, f.partner) {
		t.Error("Match() = false for partner with empty blocked categories, want true")
	}
}

func TestMatch_ValidPartner(t *testing.T) {
	f := newTargetingFixture()

	isMatched := Match(f.request, f.partner)

	if !isMatched {
		t.Errorf("Match() = false for valid partner, want true")
	}
}
