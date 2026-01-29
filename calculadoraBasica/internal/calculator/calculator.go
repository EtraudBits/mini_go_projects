package calculator

//calulate recebe dois números e uma operacao
// retorna o resultado ou erro

func Calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil // soma e retorna nil para erro
	case "-":
		return a - b, nil // subtração e retorna nil para erro
	case "*":
		return a * b, nil // multiplicação e retorna nil para erro
	case "/":
		if b == 0 {
			return 0, ErrDivisionByZero // verifica divisão por zero
		}
		return a / b, nil // divisão e retorna nil para erro
	default:
		return 0, ErrInvalidOperation // operação inválida
	}
}