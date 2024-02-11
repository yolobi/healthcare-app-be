package drugsres

type DrugMaster struct {
	ID             uint64  `json:"id"`
	Name           string  `json:"name"`
	GenericName    string  `json:"generic_name"`
	Content        string  `json:"content"`
	Description    string  `json:"description,omitempty"`
	ManufactureID  uint64  `json:"manufacture_id"`
	Manufacture    string  `json:"manufacture"`
	CategoryID     uint64  `json:"category_id"`
	Category       string  `json:"category"`
	FormID         uint64  `json:"drug_form_id"`
	Form           string  `json:"drug_form"`
	UnitInPack     string  `json:"unit_in_pack"`
	Weight         uint64  `json:"weight"`
	Height         uint64  `json:"height"`
	Length         uint64  `json:"length"`
	Width          uint64  `json:"width"`
	Image          string  `json:"image"`
	MinSellingUnit float64 `json:"min_selling_unit"`
	MaxSellingUnit float64 `json:"max_selling_unit"`
	TotalStock     uint64  `json:"total_stock"`
}
