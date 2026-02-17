package serde

import (
	"encoding/json"
)

type BooleanAsInteger bool

func (b *BooleanAsInteger) MarshalJSON() ([]byte, error) {
	if b == nil {
		return json.Marshal(nil)
	}

	if *b {
		return json.Marshal(1)
	}

	return json.Marshal(0)
}

func (b *BooleanAsInteger) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		var v string
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}

		*b = false

		if v == "1" {
			*b = true
		}

		return nil
	}

	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*b = false

	if v == 1 {
		*b = true
	}

	return nil
}
