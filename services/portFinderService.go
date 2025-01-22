package services

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ScanOpenPorts executa uma varredura de portas TCP em um domínio fornecido
// utilizando um limitador de goroutines para maior estabilidade.
func ScanOpenPorts(domain string, ports []int) []int {
	// Canal deve estar com capacidade para o mesmo tamanho do slice de portas.
	openPorts := make(chan int, len(ports))
	var wg sync.WaitGroup

	// Define o limite de goroutines rodando simultaneamente.
	const maxGoroutines = 5000
	guard := make(chan struct{}, maxGoroutines)

	for _, port := range ports {
		wg.Add(1)
		guard <- struct{}{} // Aumenta uma goroutine em execução.

		go func(port int) {
			defer wg.Done()
			// Libera o slot no final da execução desta goroutine.
			defer func() { <-guard }()

			address := fmt.Sprintf("%s:%d", domain, port)

			// Tenta conectar à porta usando TCP com um timeout.
			conn, err := net.DialTimeout("tcp", address, 3*time.Second)
			if err == nil {
				openPorts <- port // Se a porta está aberta, a envia para o canal.
				conn.Close()
			}
		}(port)
	}

	// Outra goroutine para fechar o canal após a execução de todas as goroutines principais.
	go func() {
		wg.Wait()
		close(openPorts)
	}()

	// Coleta os resultados do canal em um slice de inteiros.
	var result []int
	for port := range openPorts {
		result = append(result, port)
	}
	return result
}

// ParsePorts converte uma string como "80,443,22,10050-10080" em uma slice de inteiros.
func ParsePorts(portStr string) ([]int, error) {
	var ports []int
	segments := strings.Split(portStr, ",")

	for _, seg := range segments {
		seg = strings.TrimSpace(seg)

		if strings.Contains(seg, "-") { // Trata como intervalo.
			rangeBounds := strings.Split(seg, "-")
			if len(rangeBounds) != 2 {
				return nil, errors.New("invalid port range specification")
			}

			start, err1 := strconv.Atoi(rangeBounds[0])
			end, err2 := strconv.Atoi(rangeBounds[1])
			if err1 != nil || err2 != nil {
				return nil, errors.New("invalid port number")
			}

			if start > end {
				return nil, errors.New("range start cannot be greater than end")
			}

			for i := start; i <= end; i++ {
				ports = append(ports, i)
			}
		} else { // Porta única.
			port, err := strconv.Atoi(seg)
			if err != nil {
				return nil, errors.New("invalid port number")
			}
			ports = append(ports, port)
		}
	}

	return ports, nil
}
