// Definindo erros do dominio | Defining domain errors
// Antes de normalizar os dados do cliente, é preciso responder: | Before normalizing client data, we need to answer:
// Como o sistema comunica que algo deu errado? | How does the system communicate that something went wrong?

package client

import "errors"

// erros de domínio relacionados à normalização de dados do cliente | domain errors related to client data normalization.
// cada erro representa uma violação clara de regra de negócio | each error represents a clear business rule violation.
// isso ajuda a manter o código limpo e facilita o tratamento de erros | this helps keep the code clean and facilitates error handling.

var (
	ErrNameEmpty = errors.New("nome do cliente não pode ser vazio | client name cannot be empty")
	ErrCPFinvalid = errors.New("CPF inválido | invalid CPF")
	ErrPhoneInvalid = errors.New("telefone inválido | invalid phone number")
	ErrEmailInvalid = errors.New("email inválido | invalid email")
)

// esses erros podem ser usados em funções de validação | these errors can be used in validation functions.
// para sinalizar problemas específicos com os dados do cliente | to signal specific issues with client data.
// precisamos saber exatamente o que falhou, logs, metricas e alertas dependem disso | we need to know exactly what failed, logs, metrics, and alerts depend on it.