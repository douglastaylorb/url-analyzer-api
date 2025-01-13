package models

type URLRequest struct {
	URL string `json:"url" binding:"required"`
}
