package calculator

import (
	"testing" // importa o pacote de testes do Go
)

// TestCalculate testa a função Calculate usando table-driven tests
func TestCalculate(t *testing.T) {
	// define uma slice de casos de teste (table-driven tests)
	tests := []struct {
		name     string  // nome descritivo do teste
		a        float64 // primeiro número
		b        float64 // segundo número
		op       string  // operação matemática
		expected float64 // resultado esperado
		wantErr  error   // erro esperado (nil se não houver erro)
	}{
		// casos de teste para soma
		{
			name:     "soma de números positivos",
			a:        5,
			b:        3,
			op:       "+",
			expected: 8,
			wantErr:  nil,
		},
		{
			name:     "soma com número negativo",
			a:        10,
			b:        -5,
			op:       "+",
			expected: 5,
			wantErr:  nil,
		},
		{
			name:     "soma de decimais",
			a:        2.5,
			b:        3.7,
			op:       "+",
			expected: 6.2,
			wantErr:  nil,
		},
		// casos de teste para subtração
		{
			name:     "subtração de números positivos",
			a:        10,
			b:        4,
			op:       "-",
			expected: 6,
			wantErr:  nil,
		},
		{
			name:     "subtração com resultado negativo",
			a:        5,
			b:        10,
			op:       "-",
			expected: -5,
			wantErr:  nil,
		},
		{
			name:     "subtração de decimais",
			a:        7.5,
			b:        2.3,
			op:       "-",
			expected: 5.2,
			wantErr:  nil,
		},
		// casos de teste para multiplicação
		{
			name:     "multiplicação de números positivos",
			a:        4,
			b:        5,
			op:       "*",
			expected: 20,
			wantErr:  nil,
		},
		{
			name:     "multiplicação com número negativo",
			a:        6,
			b:        -3,
			op:       "*",
			expected: -18,
			wantErr:  nil,
		},
		{
			name:     "multiplicação por zero",
			a:        100,
			b:        0,
			op:       "*",
			expected: 0,
			wantErr:  nil,
		},
		{
			name:     "multiplicação de decimais",
			a:        2.5,
			b:        4.0,
			op:       "*",
			expected: 10.0,
			wantErr:  nil,
		},
		// casos de teste para divisão
		{
			name:     "divisão de números positivos",
			a:        10,
			b:        2,
			op:       "/",
			expected: 5,
			wantErr:  nil,
		},
		{
			name:     "divisão com resultado decimal",
			a:        7,
			b:        2,
			op:       "/",
			expected: 3.5,
			wantErr:  nil,
		},
		{
			name:     "divisão com número negativo",
			a:        15,
			b:        -3,
			op:       "/",
			expected: -5,
			wantErr:  nil,
		},
		{
			name:     "divisão por zero",
			a:        10,
			b:        0,
			op:       "/",
			expected: 0,
			wantErr:  ErrDivisionByZero, // espera o erro de divisão por zero
		},
		// casos de teste para operações inválidas
		{
			name:     "operação inválida - módulo",
			a:        10,
			b:        3,
			op:       "%",
			expected: 0,
			wantErr:  ErrInvalidOperation, // espera o erro de operação inválida
		},
		{
			name:     "operação inválida - string vazia",
			a:        5,
			b:        3,
			op:       "",
			expected: 0,
			wantErr:  ErrInvalidOperation, // espera o erro de operação inválida
		},
		{
			name:     "operação inválida - caractere desconhecido",
			a:        8,
			b:        4,
			op:       "^",
			expected: 0,
			wantErr:  ErrInvalidOperation, // espera o erro de operação inválida
		},
	}

	// loop através de todos os casos de teste
	for _, tt := range tests {
		// executa cada caso de teste como um sub-teste
		t.Run(tt.name, func(t *testing.T) {
			// chama a função Calculate com os valores do teste
			got, err := Calculate(tt.a, tt.b, tt.op)

			// verifica se o erro retornado é o esperado
			if err != tt.wantErr {
				// se o erro não for o esperado, marca o teste como falho
				t.Errorf("Calculate(%v, %v, %q) erro = %v, esperado %v",
					tt.a, tt.b, tt.op, err, tt.wantErr)
				return // retorna sem verificar o resultado
			}

			// se não esperamos erro, verifica o resultado
			if tt.wantErr == nil {
				// compara o resultado obtido com o esperado
				// usa uma pequena margem de erro para comparar floats
				if !floatEquals(got, tt.expected) {
					// se o resultado não for o esperado, marca o teste como falho
					t.Errorf("Calculate(%v, %v, %q) = %v, esperado %v",
						tt.a, tt.b, tt.op, got, tt.expected)
				}
			}
		})
	}
}

// floatEquals compara dois números float64 com uma margem de erro
// isso é necessário porque comparações diretas de floats podem falhar devido a imprecisão
func floatEquals(a, b float64) bool {
	const epsilon = 1e-9 // margem de erro muito pequena
	// calcula a diferença absoluta entre os números
	diff := a - b
	// se a diferença for negativa, torna positiva
	if diff < 0 {
		diff = -diff
	}
	// retorna true se a diferença for menor que a margem de erro
	return diff < epsilon
}

// TestCalculateEdgeCases testa casos extremos adicionais
func TestCalculateEdgeCases(t *testing.T) {
	// define casos extremos para testar
	tests := []struct {
		name     string
		a        float64
		b        float64
		op       string
		expected float64
		wantErr  error
	}{
		{
			name:     "zero dividido por número",
			a:        0,
			b:        5,
			op:       "/",
			expected: 0,
			wantErr:  nil,
		},
		{
			name:     "números muito grandes",
			a:        1e10,
			b:        1e10,
			op:       "+",
			expected: 2e10,
			wantErr:  nil,
		},
		{
			name:     "números muito pequenos",
			a:        0.0001,
			b:        0.0002,
			op:       "+",
			expected: 0.0003,
			wantErr:  nil,
		},
		{
			name:     "subtração resultando em zero",
			a:        5.5,
			b:        5.5,
			op:       "-",
			expected: 0,
			wantErr:  nil,
		},
	}

	// loop através dos casos extremos
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// chama a função Calculate
			got, err := Calculate(tt.a, tt.b, tt.op)

			// verifica o erro
			if err != tt.wantErr {
				t.Errorf("Calculate(%v, %v, %q) erro = %v, esperado %v",
					tt.a, tt.b, tt.op, err, tt.wantErr)
				return
			}

			// verifica o resultado se não há erro esperado
			if tt.wantErr == nil && !floatEquals(got, tt.expected) {
				t.Errorf("Calculate(%v, %v, %q) = %v, esperado %v",
					tt.a, tt.b, tt.op, got, tt.expected)
			}
		})
	}
}
