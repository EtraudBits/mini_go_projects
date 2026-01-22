package client

// Tipos nomeados para documentação automática e type safety | Named types for automatic documentation and type safety
// Use os factories para criar instâncias válidas | Use factories to create valid instances
// Métodos de apresentação estão em methods.go | Presentation methods are in methods.go

// CPF representa um CPF válido e normalizado (apenas dígitos) | CPF represents a valid and normalized CPF (digits only)
// Garante que apenas CPFs válidos existam no sistema | Ensures only valid CPFs exist in the system
// Use NewCPF() no factoryCPF.go para criar um CPF válido | Use NewCPF() in factoryCPF.go to create a valid CPF
type CPF string

// Name representa um nome válido e normalizado | Name represents a valid and normalized name
// Sempre com capitalização correta (primeira letra maiúscula) | Always with correct capitalization (first letter uppercase)
// Use NewName() no factoryName.go para criar um nome válido | Use NewName() in factoryName.go to create a valid name
type Name string

// Email representa um email válido | Email represents a valid email
// Use NewEmail() no factoryEmail.go para criar um email válido | Use NewEmail() in factoryEmail.go to create a valid email
type Email string

// Phone representa um telefone válido e normalizado | Phone represents a valid and normalized phone
// Use NewPhone() no factoryPhone.go para criar um telefone válido | Use NewPhone() in factoryPhone.go to create a valid phone
type Phone string

