package models

// URLRequest representa a estrutura de dados esperados na requisição
type URLRequest struct {
	URL string `json:"url" binding:"required"`
}
