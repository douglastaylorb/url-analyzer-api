package models

type MetaTags struct {
	Thumbnail   string `json:"thumbnail"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type MetaTagRequest struct {
	URL string `json:"url" binding:"required"`
}
