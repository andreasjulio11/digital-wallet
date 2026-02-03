package dto

type TransactionRequest struct {
	Amount          float64 `json:"amount"`
	TypeTransaction string  `json:"type"`
}
