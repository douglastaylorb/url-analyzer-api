package services

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func AnalyzeURL(inputUrl string) (map[string]interface{}, error) {
	// Analisar a URL para extrair o host
	parsedUrl, err := url.Parse(inputUrl)
	if err != nil {
		return nil, fmt.Errorf("erro ao analisar a URL: %v", err)
	}

	// Obter o host a partir da URL
	host := parsedUrl.Host

	// Realizar a consulta DNS para obter endereço IP
	ipAddresses, err := net.LookupIP(host)
	if err != nil {
		return nil, fmt.Errorf("erro ao resolver o IP: %v", err)
	}

	var ipAddress string
	if len(ipAddresses) > 0 {
		ipAddress = ipAddresses[0].String()
	} else {
		ipAddress = "IP não encontrado"
	}

	// Realizar a requisição OPTIONS para encontrar métodos suportados
	client := &http.Client{}
	req, err := http.NewRequest("OPTIONS", inputUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição OPTIONS: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar requisição OPTIONS: %v", err)
	}
	defer resp.Body.Close()

	// Obter métodos suportados a partir de "Allow" header
	allowedMethods := resp.Header.Get("Allow")

	// Download da página HTML
	htmlResp, err := http.Get(inputUrl)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter HTML da URL: %v", err)
	}
	defer htmlResp.Body.Close()

	// Parse do HTML para extrair hrefs
	hrefs := extractHrefs(htmlResp)

	serverInfo := "MockServer/1.0"                // Isso ainda é fictício
	performanceMetrics := "Very fast"             // Fictício
	missingHeaders := []string{"X-Frame-Options"} // Fictício

	return map[string]interface{}{
		"ip":                 ipAddress,
		"serverInfo":         serverInfo,
		"performanceMetrics": performanceMetrics,
		"missingHeaders":     missingHeaders,
		"allowedMethods":     allowedMethods,
		"hrefs":              hrefs,
	}, nil
}

func extractHrefs(resp *http.Response) []string {
	var hrefs []string
	tokens := html.NewTokenizer(resp.Body)

	for {
		tt := tokens.Next()

		switch {
		case tt == html.ErrorToken:
			// Fim do documento
			return hrefs
		case tt == html.StartTagToken:
			t := tokens.Token()

			// Verifica se é uma tag <a>
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}

			// Extrai o atributo href
			for _, a := range t.Attr {
				if a.Key == "href" {
					href := a.Val
					if isValidHref(href) {
						hrefs = append(hrefs, href)
					}
				}
			}
		}
	}
}

func isValidHref(href string) bool {
	// Simplisticamente, certificar-se de que os hrefs não são referências locais
	// Pode-se melhorar para validar caso sejam URLs válidas
	return strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://")
}
