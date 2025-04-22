package models

type Bet struct {
	UserID  string  `json:"user_id"`
	EventID string  `json:"event_id"`
	Odds    float64 `json:"odds"`
	Amount  float64 `json:"amount"`
	Won     *bool   `json:"won,omitempty"`
}
