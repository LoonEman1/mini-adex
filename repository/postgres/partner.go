package postgres

import (
	"context"
	"fmt"
	"mini-adex/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PartnerRepository struct {
	db *pgxpool.Pool
}

func NewPartnerRepository(db *pgxpool.Pool) *PartnerRepository {
	return &PartnerRepository{
		db: db,
	}
}

func (r *PartnerRepository) List(
	ctx context.Context,
) ([]domain.Partner, error) {
	const query = `
		SELECT 
			uid,
			name,
			endpoint,
			is_enabled,
			countries,
			device_types,
			min_bid_floor,
			blocked_categories
			FROM partners
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query partners: %w", err)
	}
	defer rows.Close()

	partners := make([]domain.Partner, 0)

	for rows.Next() {
		var partner domain.Partner

		if err := rows.Scan(
			&partner.UID,
			&partner.Name,
			&partner.Endpoint,
			&partner.IsEnabled,
			&partner.Countries,
			&partner.DeviceTypes,
			&partner.MinBidFloor,
			&partner.BlockedCategories,
		); err != nil {
			return nil, fmt.Errorf("scan partner: %w", err)
		}

		partners = append(partners, partner)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate partners: %w", err)
	}

	return partners, nil
}
