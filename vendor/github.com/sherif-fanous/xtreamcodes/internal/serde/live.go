package serde

// LiveStream is the internal representation used for marshaling/unmarshaling
type LiveStream struct {
	AddedOn             UnixTimeAsInteger `json:"added"`
	CatchupDurationDays IntegerAsInteger  `json:"tv_archive_duration"`
	CategoryID          *IntegerAsInteger `json:"category_id"`
	CategoryIDs         []int             `json:"category_ids,omitempty"`
	CustomSID           *string           `json:"custom_sid"`
	DirectSource        string            `json:"direct_source"`
	EPGChannelID        *string           `json:"epg_channel_id"`
	HasCatchup          BooleanAsInteger  `json:"tv_archive"`
	IsAdult             BooleanAsInteger  `json:"is_adult"`
	Name                string            `json:"name"`
	Number              int               `json:"num"`
	StreamIcon          string            `json:"stream_icon"`
	StreamID            int               `json:"stream_id"`
	StreamType          string            `json:"stream_type"`
}
