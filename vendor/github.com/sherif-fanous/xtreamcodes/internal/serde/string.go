package serde

import (
	"encoding/base64"
	"encoding/json"
)

type SliceStringSliceString []string

func (s *SliceStringSliceString) MarshalJSON() ([]byte, error) {
	if s == nil {
		return json.Marshal(nil)
	}

	return json.Marshal(*s)
}

func (s *SliceStringSliceString) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		var v string
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}

		if v == "" {
			*s = []string{}

			return nil
		}

		*s = []string{v}

		return nil
	}

	var v []string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = v

	return nil
}

type StringAsBase64EncodedString string

func (s *StringAsBase64EncodedString) MarshalJSON() ([]byte, error) {
	if s == nil {
		return json.Marshal(nil)
	}

	return json.Marshal(base64.StdEncoding.EncodeToString([]byte(*s)))
}

func (s *StringAsBase64EncodedString) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	decoded, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return err
	}

	*s = StringAsBase64EncodedString(decoded)

	return nil
}
