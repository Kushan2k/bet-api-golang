package main

import (
	"github/bet-api/models"
	"github/bet-api/services"
	"log"
	"net/http"
	"sync"
)
var (
		bets     = make(map[string][]*models.Bet) // key: eventID
		balances = make(map[string]float64)      // key: userID
		mu       sync.RWMutex
	)

func main() {

	server:=http.NewServeMux()
	bet_service:=services.NewBetService(bets, balances, &mu)

	server.HandleFunc("/place_bet", bet_service.PlaceBetHandler)
	server.HandleFunc("/settle_bet", bet_service.SettleBetHandler)
	server.HandleFunc("/balance", bet_service.BalanceHandler)

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Println("Server stopped.")
}
