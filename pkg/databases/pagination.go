package databases

type Pagination struct {
	Limit   string `form:"limit"`
	Page    string `form:"page"`
	OrderBy string `form:"order_by"`
	Sort    string `form:"sort"`
}
