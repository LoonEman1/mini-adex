package main

import (
	"context"
	"log"
	"mini-adex/internal/auction"
	"mini-adex/internal/dsp"
	httptransport "mini-adex/internal/transport/http"
	"mini-adex/repository/postgres"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	databaseURL := os.Getenv("DATABASE_URL")

	db, err := pgxpool.New(
		ctx,
		databaseURL,
	)

	if err != nil {
		log.Fatalf("create postgres pool: %v", err)
	}
	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		log.Fatalf("ping postgres: %v", err)
	}

	partnerRepo := postgres.NewPartnerRepository(db)

	dspClient := dsp.NewFakeClient()

	auctionTimeout := 200 * time.Millisecond

	if value := os.Getenv("AUCTION_TIMEOUT"); value != "" {
		auctionTimeout, err = time.ParseDuration(value)
		if err != nil {
			log.Fatalf("parse AUCTION_TIMEOUT: %v", err)
		}
	}

	service := auction.NewService(partnerRepo, dspClient, auctionTimeout)

	handler := httptransport.NewHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /auction", handler.Auction)
	mux.HandleFunc("GET /health", httptransport.Health)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("start HTTP server: %v", err)
	}

	log.Printf("server started on :8080")
}
