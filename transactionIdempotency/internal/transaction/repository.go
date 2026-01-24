package transaction

// Repository armazena transações processadas para garantir idempotência | Repository stores processed transactions to ensure idempotency
type Repository struct {
	processed map[string]Transaction
}

// NewRepository cria um novo repositório de transações | NewRepository creates a new transaction repository
func NewRepository() *Repository {
	return &Repository{
		processed: make(map[string]Transaction),
	}
}