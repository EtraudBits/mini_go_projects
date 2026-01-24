package transaction

// transação representa uma transação financeira | Transaction represents a financial transaction
type Transaction struct {
	ID string 
	Customer string
	Amount float64
}