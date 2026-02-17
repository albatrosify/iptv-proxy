package serde

import (
	"encoding/json"
	"strconv"
)

type IntegerAsInteger int

func (i *IntegerAsInteger) MarshalJSON() ([]byte, error) {
	if i == nil {
		return json.Marshal(nil)
	}

	return json.Marshal(int(*i))
}

func (i *IntegerAsInteger) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		var v string
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}

		parsedValue, err := strconv.Atoi(v)
		if err != nil {
			parsedValue = -1
		}

		*i = IntegerAsInteger(parsedValue)

		return nil
	}

	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = IntegerAsInteger(v)

	return nil
}
