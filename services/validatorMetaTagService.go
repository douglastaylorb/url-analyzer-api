package services

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/douglastaylorb/url-analyzer-api/models"
)

func ValidateMetaTags(url string) (*models.MetaTags, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.New("Erro ao acessar a URL")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Erro ao acessar a URL: Não retornou 200")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, errors.New("Erro ao extrair meta tags")
	}

	metaTags := &models.MetaTags{
		Thumbnail:   doc.Find("meta[property='og:image']").AttrOr("content", ""),
		Title:       doc.Find("meta[property='og:title']").AttrOr("content", "Título não encontrado."),
		Description: doc.Find("meta[property='og:description']").AttrOr("content", "Descrição não encontrada."),
	}

	return metaTags, nil
}
