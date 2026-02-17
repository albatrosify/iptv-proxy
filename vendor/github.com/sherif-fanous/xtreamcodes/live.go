package xtreamcodes

import (
	"encoding/json"
	"time"

	"github.com/sherif-fanous/xtreamcodes/internal/serde"
)

// LiveStream is the public representation with clean and simple types
type LiveStream struct {
	AddedOn             time.Time
	CatchupDurationDays int
	CategoryID          *int
	CategoryIDs         []int
	CustomSID           *string
	DirectSource        string
	EPGChannelID        *string
	HasCatchup          bool
	IsAdult             bool
	Name                string
	Number              int
	StreamIcon          string
	StreamID            int
	StreamType          string
}

// fromInternal populates a LiveStream from an internal models.LiveStream
func (l *LiveStream) fromInternal(internal *serde.LiveStream) {
	l.AddedOn = time.Time(internal.AddedOn.Time)
	l.CatchupDurationDays = int(internal.CatchupDurationDays)
	l.CategoryID = (*int)(internal.CategoryID)
	l.CategoryIDs = internal.CategoryIDs
	l.CustomSID = internal.CustomSID
	l.DirectSource = internal.DirectSource
	l.EPGChannelID = internal.EPGChannelID
	l.HasCatchup = bool(internal.HasCatchup)
	l.IsAdult = bool(internal.IsAdult)
	l.Name = internal.Name
	l.Number = internal.Number
	l.StreamIcon = internal.StreamIcon
	l.StreamID = internal.StreamID
	l.StreamType = internal.StreamType
}

// toInternal converts a LiveStream to the internal models.LiveStream representation
func (l *LiveStream) toInternal() *serde.LiveStream {
	return &serde.LiveStream{
		AddedOn:             serde.UnixTimeAsInteger{Time: l.AddedOn},
		CatchupDurationDays: serde.IntegerAsInteger(l.CatchupDurationDays),
		CategoryID:          (*serde.IntegerAsInteger)(l.CategoryID),
		CategoryIDs:         l.CategoryIDs,
		CustomSID:           l.CustomSID,
		DirectSource:        l.DirectSource,
		EPGChannelID:        l.EPGChannelID,
		HasCatchup:          serde.BooleanAsInteger(l.HasCatchup),
		IsAdult:             serde.BooleanAsInteger(l.IsAdult),
		Name:                l.Name,
		Number:              l.Number,
		StreamIcon:          l.StreamIcon,
		StreamID:            l.StreamID,
		StreamType:          l.StreamType,
	}
}

// MarshalJSON implements the json.Marshaler interface
func (l *LiveStream) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.toInternal())
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (l *LiveStream) UnmarshalJSON(data []byte) error {
	var internal serde.LiveStream
	if err := json.Unmarshal(data, &internal); err != nil {
		return err
	}

	l.fromInternal(&internal)

	return nil
}
