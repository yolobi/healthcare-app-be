package requests

import "time"

type GlobalFilter struct {
	Search      string    `form:"search"`
	PharmacyId  int       `form:"pharmacy_id"`
	Pharmacy    string    `form:"pharmacy"`
	Limit       int       `form:"limit,default=10"`
	Page        int       `form:"page,default=1"`
	OrderBy     string    `form:"order_by,default=updated_at"`
	Sort        string    `form:"sort,default=desc"`
	StartDate   time.Time `form:"start_date" binding:"ltefield=EndDate" time_format:"2006-01-02"`
	EndDate     time.Time `form:"end_date" binding:"lt" time_format:"2006-01-02"`
	OrderStatus string    `form:"order_status,default="`
}
