package requests

type GlobalQuery struct {
	Page       int
	PerPage    int
	OrderedBy  string
	Conditions []*Condition
	With       []string
	Search     string
	Locked     bool
}

func NewQuery(req *GlobalFilter) *GlobalQuery {
	if req.Sort != "desc" && req.Sort != "asc" {
		req.Sort = "desc"
	}
	query := &GlobalQuery{
		Page:      req.Page,
		PerPage:   req.Limit,
		OrderedBy: req.OrderBy + " " + req.Sort,
		Search:    "'%" + req.Search + "%'",
	}
	query.Conditions = make([]*Condition, 0)
	return query
}

func (q *GlobalQuery) AddCondition(field string, operator Operator, value any) *GlobalQuery {
	condition := NewCondition(field, operator, value)
	q.Conditions = append(q.Conditions, condition)
	return q
}

func (q *GlobalQuery) GetPagination() (int, int) {
	limit := q.PerPage
	if limit <= 0 {
		limit = 10
	}
	page := q.Page
	if page < 0 {
		page = 1
	}
	offset := (page - 1) * limit
	return limit, offset
}
