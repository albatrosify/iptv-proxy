package xtreamcodes

import (
	"encoding/json"

	"github.com/sherif-fanous/xtreamcodes/internal/serde"
)

// Category is the public representation with clean and simple types
type Category struct {
	CategoryID       int
	CategoryName     string
	ParentCategoryID int
}

type (
	LiveCategory   = Category
	SeriesCategory = Category
	VODCategory    = Category
)

// fromInternal populates a Category from an internal models.Category
func (c *Category) fromInternal(internal *serde.Category) {
	c.CategoryID = int(internal.CategoryID)
	c.CategoryName = internal.CategoryName
	c.ParentCategoryID = internal.ParentCategoryID
}

// toInternal converts a Category to the internal models.Category representation
func (c *Category) toInternal() *serde.Category {
	return &serde.Category{
		CategoryID:       serde.IntegerAsInteger(c.CategoryID),
		CategoryName:     c.CategoryName,
		ParentCategoryID: c.ParentCategoryID,
	}
}

// MarshalJSON implements the json.Marshaler interface
func (c *Category) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.toInternal())
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (c *Category) UnmarshalJSON(data []byte) error {
	var internal serde.Category
	if err := json.Unmarshal(data, &internal); err != nil {
		return err
	}

	c.fromInternal(&internal)

	return nil
}
