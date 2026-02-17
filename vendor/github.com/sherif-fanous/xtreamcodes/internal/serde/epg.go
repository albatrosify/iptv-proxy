package serde

// EPG is the internal representation used for marshaling/unmarshaling
type EPG struct {
	EPGListings []EPGListing `json:"epg_listings"`
}

// EPGListing is the internal representation used for marshaling/unmarshaling
type EPGListing struct {
	ChannelID      string                      `json:"channel_id"`
	Description    StringAsBase64EncodedString `json:"description"`
	EndDateTime    DateTimeAsString            `json:"end"`
	EPGID          IntegerAsInteger            `json:"epg_id"`
	HasArchive     *BooleanAsInteger           `json:"has_archive,omitempty"`
	ID             IntegerAsInteger            `json:"id"`
	Language       string                      `json:"lang"`
	NowPlaying     *BooleanAsInteger           `json:"now_playing,omitempty"`
	StartDateTime  DateTimeAsString            `json:"start"`
	StartTimestamp UnixTimeAsInteger           `json:"start_timestamp"`
	StopTimestamp  UnixTimeAsInteger           `json:"stop_timestamp"`
	StreamID       *IntegerAsInteger           `json:"stream_id,omitempty"`
	Title          StringAsBase64EncodedString `json:"title"`
}
