package auction

import "mini-adex/internal/domain"

func Match(req domain.AuctionRequest, partner domain.Partner) bool {
	if !partner.IsEnabled {
		return false
	}

	if !containsOrEmpty(partner.Countries, req.Country) {
		return false
	}

	if !containsOrEmpty(partner.DeviceTypes, req.DeviceType) {
		return false
	}

	if req.BidFloor < partner.MinBidFloor {
		return false
	}

	if containsBlockedCategory(req.Categories, partner.BlockedCategories) {
		return false
	}

	return true
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
		if Match(req, partner) {
			matchedPartners = append(matchedPartners, partner)
		}
	}

	return matchedPartners
}
