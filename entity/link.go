package entity

type CreateLinkRequest struct {
	LongUrl  string `json:"long_url" binding:"required,url"`
	ShortUrl string `json:"short_url" binding:"required,url"`
}
