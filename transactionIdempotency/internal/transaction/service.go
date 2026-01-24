package transaction

// Service fornece operações para processar transações com idempotência | Service provides operations to process transactions with idempotency
type Service struct {
	repo *Repository
}

// NewService cria um novo serviço de transações | NewService creates a new transaction service
func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo, // 1º repo = campo da struct Service; 2º repo = parâmetro da função NewService 
	}
}

// Process processa uma transação garantindo idempotência | Process processes a transaction ensuring idempotency
func (s *Service) Process(t Transaction) error {
	_, exists := s.repo.processed[t.ID]

	if exists {
		return ErrTransactionAlreadyProcessed
	}

	s.repo.processed[t.ID] = t
	return nil
}