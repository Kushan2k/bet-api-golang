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

	bet_service:=services.NewBetService(bets, balances, &mu)

	http.HandleFunc("/place_bet", bet_service.PlaceBetHandler)
	http.HandleFunc("/settle_bet", bet_service.SettleBetHandler)
	http.HandleFunc("/balance", bet_service.BalanceHandler)

	log.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
