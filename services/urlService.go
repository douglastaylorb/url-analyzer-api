package services

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
	"golang.org/x/net/html"
)

func AnalyzeURL(inputUrl string) (map[string]interface{}, error) {
	parsedUrl, err := url.Parse(inputUrl)
	if err != nil {
		return nil, fmt.Errorf("erro ao analisar a URL: %v", err)
	}

	host := parsedUrl.Host
	ipAddresses, err := net.LookupIP(host)
	if err != nil {
		return nil, fmt.Errorf("erro ao resolver o IP: %v", err)
	}

	ipAddress := "IP não encontrado"
	if len(ipAddresses) > 0 {
		ipAddress = ipAddresses[0].String()
	}

	client := &http.Client{}

	// Executar requisição GET para coletar informações principais
	req, err := http.NewRequest("GET", inputUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição GET: %v", err)
	}

	startTime := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar requisição GET: %v", err)
	}
	defer resp.Body.Close()
	responseTime := time.Since(startTime).Milliseconds()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o corpo HTML: %v", err)
	}

	hrefs, staticResources, err := extractInfo(bodyBytes, inputUrl)
	if err != nil {
		return nil, fmt.Errorf("erro ao analisar o conteúdo HTML: %v", err)
	}

	totalContentSize := len(bodyBytes)
	contentTypeSize := map[string]int{resp.Header.Get("Content-Type"): len(bodyBytes)}

	for _, resourceURL := range staticResources {
		resourceSize, resourceType := fetchResource(resourceURL)
		if resourceSize > 0 {
			totalContentSize += resourceSize
			contentTypeSize[resourceType] += resourceSize
		}
	}

	// Realizar requisição OPTIONS para capturar métodos permitidos
	reqOptions, err := http.NewRequest("OPTIONS", inputUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição OPTIONS: %v", err)
	}

	respOptions, err := client.Do(reqOptions)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar requisição OPTIONS: %v", err)
	}
	defer respOptions.Body.Close()

	allowedMethods := respOptions.Header.Get("Allow")

	// Cálculo das porcentagens de tamanho do conteúdo
	contentTypePercentage := calculatePercentages(contentTypeSize, totalContentSize)

	serverInfo := resp.Header.Get("Server")
	performanceMetrics := fmt.Sprintf("%d ms", responseTime)

	results := map[string]interface{}{
		"ip":                 ipAddress,
		"serverInfo":         serverInfo,
		"performanceMetrics": performanceMetrics,
		"allowedMethods":     allowedMethods,
		"hrefs":              hrefs,
		"contentType":        contentTypePercentage,
	}
	return results, nil
}

// Função para criar o PDF a partir dos resultados
func GeneratePDF(results map[string]interface{}) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Adicionando a fonte Roboto
	fontDir := "./assets"
	pdf.AddUTF8Font("Roboto", "", fontDir+"/Roboto-Regular.ttf")
	pdf.AddUTF8Font("Roboto", "B", fontDir+"/Roboto-Bold.ttf")

	pdf.SetFont("Roboto", "", 16)
	pdf.AddPage()
	pdf.Cell(0, 10, "Análise de URL")
	pdf.Ln(12)

	// Seção Informações Gerais
	pdf.SetFont("Roboto", "", 12)
	pdf.Cell(0, 10, "Informações Gerais")
	pdf.Ln(10)
	info := []struct {
		Label string
		Value interface{}
	}{
		{"IP", results["ip"]},
		{"Servidor", results["serverInfo"]},
		{"Tempo de Resposta", results["performanceMetrics"]},
		{"Métodos Permitidos", results["allowedMethods"]},
	}

	for _, item := range info {
		pdf.CellFormat(0, 10, fmt.Sprintf("%s: %v", item.Label, item.Value), "", 1, "", false, 0, "")
	}
	pdf.Ln(5)

	// Seção Percentual de Conteúdo
	pdf.SetFont("Roboto", "B", 12)
	pdf.Cell(0, 10, "Percentual de Tamanho de Conteúdo por Tipo")
	pdf.Ln(10)
	pdf.SetFont("Roboto", "", 11)
	contentData := results["contentType"].(map[string]string)
	for contentType, percentage := range contentData {
		pdf.CellFormat(0, 8, fmt.Sprintf("%s: %s", contentType, percentage), "", 1, "", false, 0, "")
	}
	pdf.Ln(5)

	// Seção Hrefs
	pdf.SetFont("Roboto", "B", 12)
	pdf.Cell(0, 10, "Hrefs")
	pdf.Ln(10)
	pdf.SetFont("Roboto", "", 11)
	hrefs := results["hrefs"].([]string)
	for _, href := range hrefs {
		pdf.MultiCell(0, 8, href, "", "", false)
		pdf.Ln(2)
	}

	// Buffer para gerar o PDF final
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar PDF: %v", err)
	}
	return buf.Bytes(), nil
}

func extractInfo(content []byte, baseUrl string) ([]string, []string, error) {
	var hrefs []string
	var staticResources []string
	tokens := html.NewTokenizer(bytes.NewReader(content))

	for {
		tt := tokens.Next()

		if tt == html.ErrorToken {
			return hrefs, staticResources, nil
		}

		token := tokens.Token()
		switch token.Data {
		case "a":
			for _, attr := range token.Attr {
				if attr.Key == "href" && isValidHref(attr.Val) {
					hrefs = append(hrefs, attr.Val)
				}
			}
		case "script", "link":
			for _, attr := range token.Attr {
				switch attr.Key {
				case "src":
					staticResources = append(staticResources, resolveURL(attr.Val, baseUrl))
				case "href":
					if token.Data == "link" {
						staticResources = append(staticResources, resolveURL(attr.Val, baseUrl))
					}
				}
			}
		}
	}
}

func isValidHref(href string) bool {
	return strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://")
}

func resolveURL(relativeURL, baseURL string) string {
	u, err := url.Parse(relativeURL)
	if err != nil || u.IsAbs() {
		return relativeURL
	}
	base, err := url.Parse(baseURL)
	if err != nil {
		return relativeURL
	}
	return base.ResolveReference(u).String()
}

func fetchResource(resourceURL string) (int, string) {
	resp, err := http.Get(resourceURL)
	if err != nil {
		return 0, ""
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, ""
	}

	contentType := resp.Header.Get("Content-Type")
	contentTypeSplit := strings.Split(contentType, ";")
	return len(data), contentTypeSplit[0]
}

func calculatePercentages(contentTypeSize map[string]int, totalSize int) map[string]string {
	percentages := make(map[string]string)
	for contentType, size := range contentTypeSize {
		percent := (float64(size) / float64(totalSize)) * 100
		percentages[contentType] = fmt.Sprintf("%.2f%%", percent)
	}
	return percentages
}
