package services

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateValidCPF gera um CPF válido formatado como string
func GenerateValidCPF() string {
	// Inicializa a fonte de geração de números aleatórios com base no tempo
	rand.Seed(time.Now().UnixNano())

	// Cria um slice para armazenar os primeiros 9 dígitos do CPF
	baseCPF := make([]int, 9)
	for i := range baseCPF {
		// Preenche cada posição com um número aleatório entre 0 e 9
		baseCPF[i] = rand.Intn(10)
	}

	// Calcula o primeiro dígito verificador, usando os primeiros 9 dígitos
	dv1 := calculateCPFVerifierDigit(baseCPF)

	// Adiciona o primeiro dígito verificador ao slice baseCPF
	baseCPF = append(baseCPF, dv1)

	// Calcula o segundo dígito verificador, agora usando os 10 dígitos
	dv2 := calculateCPFVerifierDigit(baseCPF)

	// Adiciona o segundo dígito verificador ao slice baseCPF
	baseCPF = append(baseCPF, dv2)

	// Formata o CPF em uma string do jeito ###.###.###-## usando os valores no slice
	cpf := fmt.Sprintf(
		"%d%d%d.%d%d%d.%d%d%d-%d%d",
		baseCPF[0], baseCPF[1], baseCPF[2],
		baseCPF[3], baseCPF[4], baseCPF[5],
		baseCPF[6], baseCPF[7], baseCPF[8],
		baseCPF[9], baseCPF[10])

	return cpf // Retorna o CPF gerado e formatado
}

// calculateCPFVerifierDigit calcula um dígito verificador do CPF
func calculateCPFVerifierDigit(digits []int) int {
	weight := len(digits) + 1 // Define o peso inicial baseado no comprimento do slice
	sum := 0
	// Calcula a soma ponderada dos dígitos
	for i, digit := range digits {
		sum += digit * (weight - i) // Multiplica cada dígito pelo peso decrescente e soma
	}

	remainder := sum % 11 // Calcula o resto da divisão por 11
	if remainder < 2 {
		return 0 // Se o resto for menor que 2, o DV é 0
	}
	return 11 - remainder // Caso contrário, subtrai o resto de 11 para obter o DV
}
