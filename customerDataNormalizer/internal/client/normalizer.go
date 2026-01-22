package client

// Antes de escrever os campos, definimos o contrato | Before writing the fields, we define the contract
// Pergunta: O que entra e o que sai | Question: What goes in and what comes out
// Decisão-> entra: dados crus(string) e sai: client e lista de erros | Decision-> in: raw data(string) and out: client and list of errors

// Nunca retorna client parcielamente válido sem avisar | Never returns partially valid client without warning

import (
	"strings"
)

// função principal de normalização | main normalization function
// Normalize recebe dados crus de um cliente e retorna | Normalize receives raw client data and returns
// Um cliente com dados normalizados ou uma lista de erros | A client with normalized data or a list of errors
func Normalize(name, cpf, phone, email string) (Client, []error) { // name, cpf, phone, email são strings crus e retorna client (struct no model.go) normalizados e lista de erros | name, cpf, phone, email are raw strings and returns normalized client and list of errors
	
	var errs []error // lista de erros | list of errors

	// (&errs) errs é um ponteiro para a lista de erros | (&errs) errs is a pointer to the list of errors
	normalizedName := normalizeName(name, &errs) 
	normalizedCPF := normalizeCPF(cpf, &errs)
	normalizedPhone := normalizePhone(phone, &errs)
	normalizedEmail := normalizeEmail(email, &errs)

	if len(errs) > 0 { // se tamanho da lista de erros for maior que 0 | if size of error list is greater than 0
		return Client{}, errs // retorna client vazio e lista de erros | returns empty client and list of errors
	}

	// atribui valores aos campos da struct (NomeDoCampo: valor) | assigns values to struct fields (FieldName: value)
	client := Client { // cria client com dados normalizados | creates client with normalized data
		Name: normalizedName,
		CPF: normalizedCPF, 
		Phone: normalizedPhone,
		Email: normalizedEmail,
	}

	return client, nil // retorna client normalizado e nil (sem erros) | return normalized client and nil (no errors)
}

// funções auxiliares de normalização | auxiliary normalization functions
// 1 - NormalizeName
// 2 - NormalizeCPF
// 3 - NormalizePhone
// 4 - NormalizeEmail

// 1 - NormalizeName (normalização do nome) | 1 - NormalizeName (name normalization)
// recebe valor cru e ponteiro para lista de erros | receives raw value and pointer to error list
// Retorna o tipo Name para garantir type safety | Returns the Name type to ensure type safety
func normalizeName(value string, errs *[]error) Name { 
	name := strings.TrimSpace(value) // remove espaços em branco no inicio e no fim | removes leading and trailing whitespace

	if name == "" { // se name for vazio | if name is empty
		*errs = append(*errs, ErrNameEmpty) // Ponteiro para lista de erros recebe erro de nome vazio e retorna o erro da pasta errors.go | Pointer to error list receives empty name error and returns the error from errors.go folder
		return "" // retorna string vazia | return empty string
	}

	words := strings.Fields(strings.ToLower(name)) // variavel words (palavras) divide o nome em palavras e converte para minusculo  (Fields separa palavras) | variable words splits the name into words and converts to lowercase (Fields separates words)

	for i, w := range words { // parcorre cada palavra separadamente | iterates through each word separately
		words[i] = strings.Title(w) // converte a primeira letra de cada palavra para maiusculo (strings.Title esta depreciado em Go) | converts the first letter of each word to uppercase (strings.Title is deprecated in Go)
	}
	return Name(strings.Join(words, " ")) // junta as palavras com espaço entre elas e retorna o nome normalizado como tipo Name | joins the words with space between them and returns the normalized name as Name type
}

// 2 - NormalizeCPF (normalização e validação completa do CPF) | 2 - NormalizeCPF (complete CPF normalization and validation)
// recebe valor cru e ponteiro para lista de erros | receives raw value and pointer to error list
// Retorna o tipo CPF para garantir type safety | Returns the CPF type to ensure type safety
func normalizeCPF(value string, errs *[]error) CPF {
	cpf, err := NewCPF(value) // usa o factory que valida tudo: tamanho, dígitos repetidos e checksum | uses the factory that validates everything: length, repeated digits and checksum
	if err != nil {
		*errs = append(*errs, err) // adiciona erro à lista | adds error to the list
		return "" // retorna CPF vazio | returns empty CPF
	}
	return cpf // retorna o cpf validado e normalizado como tipo CPF | returns the validated and normalized cpf as CPF type
}

// 3 - NormalizePhone (normalização e validação do telefone) | 3 - NormalizePhone (phone normalization and validation)
// recebe valor cru e ponteiro para lista de erros | receives raw value and pointer to error list
// Retorna o tipo Phone para garantir type safety | Returns the Phone type to ensure type safety
func normalizePhone(value string, errs *[]error) Phone {
	phone, err := NewPhone(value) // usa o factory que valida tudo: tamanho, DDD, nono digito | uses the factory that validates everything: length, DDD, ninth digit
	if err != nil {
		*errs = append(*errs, err) // adiciona erro à lista | adds error to the list
		return "" // retorna Phone vazio | returns empty Phone
	}
	return phone // retorna o telefone validado e normalizado como tipo Phone | returns the validated and normalized phone as Phone type
}		

// 4 - NormalizeEmail (normalização e validação do email) | 4 - NormalizeEmail (email normalization and validation)
// recebe valor cru e ponteiro para lista de erros | receives raw value and pointer to error list
// Retorna o tipo Email para garantir type safety | Returns the Email type to ensure type safety
func normalizeEmail(value string, errs *[]error) Email {
	email, err := NewEmail(value) // usa o factory que valida tudo: formato, @, domínio | uses the factory that validates everything: format, @, domain
	if err != nil {
		*errs = append(*errs, err) // adiciona erro à lista | adds error to the list
		return "" // retorna Email vazio | returns empty Email
	}
	return email // retorna o email validado e normalizado como tipo Email | returns the validated and normalized email as Email type
}
