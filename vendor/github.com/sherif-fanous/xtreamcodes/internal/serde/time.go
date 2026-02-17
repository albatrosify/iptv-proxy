package serde

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type DateAsString struct {
	time.Time
}

func (d *DateAsString) MarshalJSON() ([]byte, error) {
	if d == nil || d.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(d.Format(time.DateOnly))
}

func (d *DateAsString) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t, err := time.Parse(time.DateOnly, v)
	if err != nil {
		d.Time = time.Time{}

		return nil
	}

	d.Time = t

	return nil
}

type DateTimeAsString struct {
	time.Time
}

func (d *DateTimeAsString) MarshalJSON() ([]byte, error) {
	if d == nil || d.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(d.Format(time.DateTime))
}

func (d *DateTimeAsString) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t, err := time.Parse(time.DateTime, v)
	if err != nil {
		return err
	}

	d.Time = t

	return nil
}

type DurationAsString struct {
	time.Duration
}

func (d *DurationAsString) MarshalJSON() ([]byte, error) {
	if d == nil {
		return json.Marshal(nil)
	}

	return json.Marshal(
		fmt.Sprintf("%02d:%02d:%02d", int(d.Hours()), int(d.Minutes())%60, int(d.Seconds())%60),
	)
}

func (d *DurationAsString) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	durationParts := strings.Split(v, ":")

	parsedDuration, err := time.ParseDuration(
		fmt.Sprintf("%sh%sm%ss", durationParts[0], durationParts[1], durationParts[2]),
	)
	if err != nil {
		parsedDuration = time.Duration(0)
	}

	d.Duration = parsedDuration

	return nil
}

type UnixTimeAsInteger struct {
	time.Time
}

func (u *UnixTimeAsInteger) MarshalJSON() ([]byte, error) {
	if u == nil {
		return json.Marshal(nil)
	}

	return json.Marshal(u.Unix())
}

func (u *UnixTimeAsInteger) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		var v string
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}

		n, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}

		u.Time = time.Unix(n, 0)

		return nil
	}

	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	u.Time = time.Unix(v, 0)

	return nil
}
