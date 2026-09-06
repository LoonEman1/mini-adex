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
	"os/signal"
	"syscall"
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

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		log.Printf("server started on :8080")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("start HTTP server: %v", err)
		}
	}()

	stopSignal := make(chan os.Signal, 1)

	signal.Notify(
		stopSignal,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-stopSignal

	log.Printf("shutting down server")

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("server shutdown: %v", err)

	}

	log.Printf("server stopped")
}
