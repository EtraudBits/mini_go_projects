package transaction

import "errors"

// Erro de transação já processada | Error for already processed transaction
var ErrTransactionAlreadyProcessed = errors.New(
	"transaction already processed",
)