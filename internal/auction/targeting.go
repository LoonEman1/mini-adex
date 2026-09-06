package auction

import (
	"log"
	"mini-adex/internal/domain"
)

type MatchResult struct {
	Matched bool
	Reason  string
}

func match(
	req domain.AuctionRequest,
	partner domain.Partner,
) MatchResult {
	if !partner.IsEnabled {
		return MatchResult{
			Matched: false,
			Reason:  "disabled",
		}
	}

	if !containsOrEmpty(partner.Countries, req.Country) {
		return MatchResult{
			Matched: false,
			Reason:  "country_mismatch",
		}
	}

	if !containsOrEmpty(partner.DeviceTypes, req.DeviceType) {
		return MatchResult{
			Matched: false,
			Reason:  "device_type_mismatch",
		}
	}

	if req.BidFloor < partner.MinBidFloor {
		return MatchResult{
			Matched: false,
			Reason:  "bid_floor_too_low",
		}
	}

	if containsBlockedCategory(req.Categories, partner.BlockedCategories) {
		return MatchResult{
			Matched: false,
			Reason:  "blocked_category",
		}
	}

	return MatchResult{
		Matched: true,
	}
}

func Match(
	req domain.AuctionRequest,
	partner domain.Partner,
) bool {
	return match(req, partner).Matched
}

func containsOrEmpty(items []string, targetValue string) bool {
	if len(items) == 0 {
		return true
	}

	for _, item := range items {
		if item == targetValue {
			return true
		}
	}

	return false
}

func containsBlockedCategory(reqCategories []string, blockedCategories []string) bool {

	for _, reqItem := range reqCategories {
		for _, partnerItem := range blockedCategories {
			if reqItem == partnerItem {
				return true
			}
		}
	}

	return false
}

func Filter(
	req domain.AuctionRequest,
	partners []domain.Partner,
) []domain.Partner {
	matchedPartners := make([]domain.Partner, 0, len(partners))

	for _, partner := range partners {

		result := match(req, partner)

		if result.Matched {
			matchedPartners = append(matchedPartners, partner)
			continue
		}

		log.Printf(
			"auction request_id=%s partner=%s filtered reason=%s",
			req.RequestID,
			partner.UID,
			result.Reason,
		)
	}

	return matchedPartners
}
