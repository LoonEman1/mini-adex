package main

import (
	"context"
	"log"
	"mini-adex/internal/auction"
	"mini-adex/repository/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	db, err := pgxpool.New(
		ctx,
		"postgres://mini_adex:mini_adex@127.0.0.1:5433/mini_adex?sslmode=disable",
	)
	if err != nil {
		log.Fatalf("create postgres pool: %v", err)
	}
	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		log.Fatalf("ping postgres: %v", err)
	}

	partnerRepo := postgres.NewPartnerRepository(db)

	service := auction.NewService(partnerRepo)

	partners, err := partnerRepo.List(ctx)
}
