package serde

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Float64AsFloat float64

func (f *Float64AsFloat) MarshalJSON() ([]byte, error) {
	return json.Marshal(float64(*f))
}

func (f *Float64AsFloat) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		var v string
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}

		if len(v) == 0 {
			v = "0"
		}

		v = strings.ReplaceAll(v, ",", ".")

		parsedValue, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}

		*f = Float64AsFloat(parsedValue)

		return nil
	}

	var v float64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*f = Float64AsFloat(v)

	return nil
}
