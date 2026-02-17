package xtreamcodes

import (
	"encoding/json"
	"time"

	"github.com/sherif-fanous/xtreamcodes/internal/serde"
)

// EPG is the public representation with clean and simple types
type EPG struct {
	EPGListings []EPGListing
}

// MarshalJSON implements the json.Marshaler interface
func (e *EPG) MarshalJSON() ([]byte, error) {
	internal := &serde.EPG{
		EPGListings: make([]serde.EPGListing, len(e.EPGListings)),
	}

	// Convert each EPGListing to its internal representation
	for i := range e.EPGListings {
		internal.EPGListings[i] = *e.EPGListings[i].toInternal()
	}

	return json.Marshal(internal)
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (e *EPG) UnmarshalJSON(data []byte) error {
	var internal serde.EPG
	if err := json.Unmarshal(data, &internal); err != nil {
		return err
	}

	e.EPGListings = make([]EPGListing, len(internal.EPGListings))

	for i, listing := range internal.EPGListings {
		e.EPGListings[i].fromInternal(&listing)
	}

	return nil
}

// EPGListing is the public representation with clean and simple types
type EPGListing struct {
	ChannelID      string
	Description    string
	EndDateTime    time.Time
	EPGID          int
	HasArchive     *bool
	ID             int
	Language       string
	NowPlaying     *bool
	StartDateTime  time.Time
	StartTimestamp time.Time
	StopTimestamp  time.Time
	StreamID       *int
	Title          string
}

// fromInternal populates an EPGListing from an internal models.EPGListing
func (e *EPGListing) fromInternal(internal *serde.EPGListing) {
	e.ChannelID = internal.ChannelID
	e.Description = string(internal.Description)
	e.EndDateTime = time.Time(internal.EndDateTime.Time)
	e.EPGID = int(internal.EPGID)
	e.HasArchive = (*bool)(internal.HasArchive)
	e.ID = int(internal.ID)
	e.Language = internal.Language
	e.NowPlaying = (*bool)(internal.NowPlaying)
	e.StartDateTime = time.Time(internal.StartDateTime.Time)
	e.StartTimestamp = time.Time(internal.StartTimestamp.Time)
	e.StopTimestamp = time.Time(internal.StopTimestamp.Time)
	e.StreamID = (*int)(internal.StreamID)
	e.Title = string(internal.Title)
}

// toInternal converts an EPGListing to the internal models.EPGListing representation
func (e *EPGListing) toInternal() *serde.EPGListing {
	return &serde.EPGListing{
		ChannelID:      e.ChannelID,
		Description:    serde.StringAsBase64EncodedString(e.Description),
		EndDateTime:    serde.DateTimeAsString{Time: e.EndDateTime},
		EPGID:          serde.IntegerAsInteger(e.EPGID),
		HasArchive:     (*serde.BooleanAsInteger)(e.HasArchive),
		ID:             serde.IntegerAsInteger(e.ID),
		Language:       e.Language,
		NowPlaying:     (*serde.BooleanAsInteger)(e.NowPlaying),
		StartDateTime:  serde.DateTimeAsString{Time: e.StartDateTime},
		StartTimestamp: serde.UnixTimeAsInteger{Time: e.StartTimestamp},
		StopTimestamp:  serde.UnixTimeAsInteger{Time: e.StopTimestamp},
		StreamID:       (*serde.IntegerAsInteger)(e.StreamID),
		Title:          serde.StringAsBase64EncodedString(e.Title),
	}
}
