package main // declara que este arquivo faz parte do pacote main (executável)

import (
	"bufio"   // importa o pacote para leitura bufferizada (ler linha por linha)
	"fmt"     // importa o pacote de formatação para entrada/saída
	"os"      // importa o pacote do sistema operacional
	"strconv" // importa o pacote para converter strings em números
	"strings" // importa o pacote para manipulação de strings

	"calculadoraBasica/internal/calculator" // importa o pacote calculator do projeto
)

// Run inicia a aplicação da calculadora
func Run() error {
	// cria um scanner para ler do terminal (stdin) linha por linha
	scanner := bufio.NewScanner(os.Stdin)
	
	// exibe o cabeçalho inicial da aplicação
	printHeader()
	
	// loop infinito que processa cálculos até o usuário sair
	for {
		// processa um cálculo completo (lê números, operação e calcula)
		// se retornar false, o usuário quer sair ou houve erro fatal
		if !processCalculation(scanner) {
			break // sai do loop infinito
		}
	}

	// retorna qualquer erro que possa ter ocorrido no scanner
	return scanner.Err()
}

// printHeader exibe o cabeçalho da aplicação
func printHeader() {
	// imprime o título da calculadora
	fmt.Println("=== Calculadora Básica ===")
	// imprime a instrução de como sair
	fmt.Println("Digite 'sair' para encerrar")
	// imprime uma linha em branco para melhor visual
	fmt.Println()
}

// readInput lê e valida uma entrada do usuário
// retorna a string lida e um bool indicando se deve continuar (true) ou sair (false)
func readInput(scanner *bufio.Scanner, prompt string) (string, bool) {
	// exibe a mensagem solicitando entrada (ex: "Digite o número: ")
	fmt.Print(prompt)
	// tenta ler uma linha do terminal
	if !scanner.Scan() {
		// se falhar (EOF ou erro), retorna vazio e false para encerrar
		return "", false
	}
	
	// pega o texto lido e remove espaços em branco do início e fim
	input := strings.TrimSpace(scanner.Text())
	// verifica se o usuário digitou "sair" (em qualquer capitalização)
	if strings.ToLower(input) == "sair" {
		// exibe mensagem de encerramento
		fmt.Println("Encerrando...")
		// retorna vazio e false para indicar que deve sair
		return "", false
	}
	
	// retorna a entrada lida e true para indicar que deve continuar
	return input, true
}

// readNumber lê e converte um número
// retorna o número convertido e um bool indicando se deve continuar
func readNumber(scanner *bufio.Scanner, prompt string) (float64, bool) {
	// lê a entrada do usuário usando readInput
	input, ok := readInput(scanner, prompt)
	// se ok for false, o usuário quer sair ou houve erro
	if !ok {
		// retorna 0 e false para indicar que deve encerrar
		return 0, false
	}
	
	// tenta converter a string para float64 (número decimal)
	num, err := strconv.ParseFloat(input, 64)
	// se houve erro na conversão (entrada inválida)
	if err != nil {
		// exibe mensagem de erro
		fmt.Println("Erro: número inválido. Tente novamente.")
		// imprime linha em branco para melhor leitura
		fmt.Println()
		// retorna 0 e true para continuar o loop (tentar novamente)
		return 0, true
	}
	
	// retorna o número convertido e true para continuar
	return num, true
}

// processCalculation processa uma operação completa
// retorna true se deve continuar o loop, false se deve encerrar
func processCalculation(scanner *bufio.Scanner) bool {
	// Lê o primeiro número do usuário
	a, ok := readNumber(scanner, "Digite o primeiro número: ")
	// se ok for false, o usuário quer sair
	if !ok {
		// retorna false para encerrar a aplicação
		return false
	}
	// se a for 0 e ok for false, houve erro de parsing
	if a == 0 && !ok {
		// retorna true para continuar o loop (tentar novamente)
		return true // erro de parsing, continua o loop
	}

	// Lê a operação matemática (+, -, *, /)
	op, ok := readInput(scanner, "Digite a operação (+, -, *, /): ")
	// se ok for false, o usuário quer sair
	if !ok {
		// retorna false para encerrar a aplicação
		return false
	}

	// Lê o segundo número do usuário
	b, ok := readNumber(scanner, "Digite o segundo número: ")
	// se ok for false, o usuário quer sair
	if !ok {
		// retorna false para encerrar a aplicação
		return false
	}
	// se b for 0 e ok for false, houve erro de parsing
	if b == 0 && !ok {
		// retorna true para continuar o loop (tentar novamente)
		return true // erro de parsing, continua o loop
	}

	// Chama a função que calcula e exibe o resultado
	displayResult(a, b, op)
	// imprime uma linha em branco para separar os cálculos
	fmt.Println()
	
	// retorna true para continuar processando mais cálculos
	return true
}

// displayResult calcula e exibe o resultado da operação
func displayResult(a, b float64, op string) {
	// chama a função Calculate do pacote calculator para fazer o cálculo
	resultado, err := calculator.Calculate(a, b, op)
	// se houve erro (ex: divisão por zero ou operação inválida)
	if err != nil {
		// exibe a mensagem de erro
		fmt.Printf("Erro: %v\n", err)
	} else { // se não houve erro
		// exibe o resultado formatado (ex: "2.00 + 3.00 = 5.00")
		fmt.Printf("Resultado: %.2f %s %.2f = %.2f\n", a, op, b, resultado)
	}
}
