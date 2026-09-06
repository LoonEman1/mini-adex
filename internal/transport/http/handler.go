package httptransport

import (
	"context"
	"encoding/json"
	"log"
	"mini-adex/internal/domain"
	"net/http"
)

type AuctionService interface {
	Process(
		ctx context.Context,
		req domain.AuctionRequest,
	) (domain.AuctionResult, error)
}

type Handler struct {
	service AuctionService
}

func NewHandler(service AuctionService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Auction(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req auctionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	if err := req.validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.service.Process(r.Context(), req.toDomain())

	if err != nil {
		log.Printf("process auction: %v", err)

		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(newAuctionResponse(result)); err != nil {
		log.Printf("encode auction response: %v", err)
	}
}
