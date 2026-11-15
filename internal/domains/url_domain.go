package domain

import "time"

type Url struct {
	UrlId     string     `json:"url_id,omitempty"`
	Code      string     `json:"code,omitempty"`
	Url       string     `json:"url,omitempty"`
	ShortUrl  string     `json:"short_url,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
