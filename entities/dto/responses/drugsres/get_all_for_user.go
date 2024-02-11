package drugsres

type DrugUserRes struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	UnitInPack  string  `json:"unit_in_pack"`
	MaxPrice    float64 `json:"max_selling_unit"`
	MinPrice    float64 `json:"min_selling_unit"`
	Image       string  `json:"image"`
	Category    string  `json:"category"`
	Manufacture string  `json:"manufacture"`
	Stock       int64   `json:"stock"`
}
