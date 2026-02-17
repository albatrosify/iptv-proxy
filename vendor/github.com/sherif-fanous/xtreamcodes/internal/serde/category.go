package serde

// Category is the internal representation used for marshaling/unmarshaling
type Category struct {
	CategoryID       IntegerAsInteger `json:"category_id"`
	CategoryName     string           `json:"category_name"`
	ParentCategoryID int              `json:"parent_id"`
}
