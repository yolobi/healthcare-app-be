package responses

import "github.com/gin-gonic/gin"

type Pagination struct {
	Items       any   `json:"items"`
	TotalItems  int64 `json:"total_items"`
	TotalPages  int64 `json:"total_pages"`
	CurrentPage int64 `json:"current_page"`
}

func NewPagination(p *Pagination, key string) map[string]interface{} {
	return gin.H{
		key:             p.Items,
		"total_items":   p.TotalItems,
		"total_pages":   p.TotalPages,
		"current_pages": p.CurrentPage,
	}
}
