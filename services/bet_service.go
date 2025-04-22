package services

import (
	"encoding/json"
	"github/bet-api/models"
	"log"
	"net/http"
	"sync"
)

type BetService struct {
	bets     map[string][]*models.Bet // key: eventID
	balances map[string]float64       // key: userID
	mu       *sync.RWMutex
}

func NewBetService(bets map[string][]*models.Bet,balances map[string]float64,mu *sync.RWMutex) *BetService {
	return &BetService{
		bets:     bets,
		balances: balances,
		mu:       mu,
	}
}


func (s *BetService) PlaceBetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var bet models.Bet
	if err := json.NewDecoder(r.Body).Decode(&bet); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.bets[bet.EventID] = append(s.bets[bet.EventID], &bet)
	s.balances[bet.UserID] -= bet.Amount

	log.Printf("Placed bet: %+v\n", bet)
	w.WriteHeader(http.StatusCreated)
}

func (s *BetService) SettleBetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	type SettleRequest struct {
		EventID string `json:"event_id"`
		Result  string `json:"result"` // "win" or "lose"
	}

	var req SettleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	betsForEvent, exists := s.bets[req.EventID]
	if !exists {
		http.Error(w, "No bets found for this event", http.StatusNotFound)
		return
	}

	for _, bet := range betsForEvent {
		if req.Result == "win" {
			win := true
			bet.Won = &win
			winnings := bet.Amount * bet.Odds
			s.balances[bet.UserID] += winnings
		} else {
			lose := false
			bet.Won = &lose
		}
	}

	log.Printf("Settled bets for event %s with result: %s\n", req.EventID, req.Result)
	w.WriteHeader(http.StatusOK)
}

func (s *BetService) BalanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id", http.StatusBadRequest)
		return
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	balance := s.balances[userID]
	json.NewEncoder(w).Encode(models.UserBalance{Balance: balance})
}
