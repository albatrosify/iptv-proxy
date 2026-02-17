package xtreamcodes

import (
	"encoding/json"
	"time"

	"github.com/sherif-fanous/xtreamcodes/internal/serde"
)

// VODStream is the public representation with clean and simple types
type VODStream struct {
	AddedOn            time.Time
	CategoryID         *int
	CategoryIDs        []int
	ContainerExtension string
	CustomSID          *string
	DirectSource       string
	IsAdult            bool
	Name               string
	Number             int
	Rating             float64
	Rating5Based       float64
	StreamIcon         string
	StreamID           int
	StreamType         string
	TMDBID             *int
	Trailer            *string
}

// fromInternal populates a VODStream from an internal models.VODStream
func (v *VODStream) fromInternal(internal *serde.VODStream) {
	v.AddedOn = time.Time(internal.AddedOn.Time)
	v.CategoryID = (*int)(internal.CategoryID)
	v.CategoryIDs = internal.CategoryIDs
	v.ContainerExtension = internal.ContainerExtension
	v.CustomSID = internal.CustomSID
	v.DirectSource = internal.DirectSource
	v.IsAdult = bool(internal.IsAdult)
	v.Name = internal.Name
	v.Number = internal.Number
	v.Rating = float64(internal.Rating)
	v.Rating5Based = float64(internal.Rating5Based)
	v.StreamIcon = internal.StreamIcon
	v.StreamID = internal.StreamID
	v.StreamType = internal.StreamType
	v.TMDBID = (*int)(internal.TMDBID)
	v.Trailer = internal.Trailer
}

// toInternal converts a VODStream to the internal models.VODStream representation
func (v *VODStream) toInternal() *serde.VODStream {
	return &serde.VODStream{
		AddedOn:            serde.UnixTimeAsInteger{Time: v.AddedOn},
		CategoryID:         (*serde.IntegerAsInteger)(v.CategoryID),
		CategoryIDs:        v.CategoryIDs,
		ContainerExtension: v.ContainerExtension,
		CustomSID:          v.CustomSID,
		DirectSource:       v.DirectSource,
		IsAdult:            serde.BooleanAsInteger(v.IsAdult),
		Name:               v.Name,
		Number:             v.Number,
		Rating:             serde.Float64AsFloat(v.Rating),
		Rating5Based:       serde.Float64AsFloat(v.Rating5Based),
		StreamIcon:         v.StreamIcon,
		StreamID:           v.StreamID,
		StreamType:         v.StreamType,
		TMDBID:             (*serde.IntegerAsInteger)(v.TMDBID),
		Trailer:            v.Trailer,
	}
}

// MarshalJSON implements the json.Marshaler interface
func (v *VODStream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.toInternal())
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (v *VODStream) UnmarshalJSON(data []byte) error {
	var internal serde.VODStream
	if err := json.Unmarshal(data, &internal); err != nil {
		return err
	}

	v.fromInternal(&internal)

	return nil
}

// VOD is the public representation of a VOD with clean and simple types
type VOD struct {
	Info      VODInfo
	MovieData MovieData
}

// fromInternal populates a VOD from an internal serde.VOD
func (v *VOD) fromInternal(internal *serde.VOD) {
	v.Info.fromInternal(&internal.Info)
	v.MovieData.fromInternal(&internal.MovieData)
}

// toInternal converts a VOD to the internal serde.VOD representation
func (v *VOD) toInternal() *serde.VOD {
	return &serde.VOD{
		Info:      *v.Info.toInternal(),
		MovieData: *v.MovieData.toInternal(),
	}
}

// MarshalJSON implements the json.Marshaler interface
func (v *VOD) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.toInternal())
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (v *VOD) UnmarshalJSON(data []byte) error {
	var internal serde.VOD
	if err := json.Unmarshal(data, &internal); err != nil {
		return err
	}

	v.fromInternal(&internal)

	return nil
}

// VODInfo is the public representation of VOD info with clean and simple types
type VODInfo struct {
	Actors          *string
	Age             *string
	Audio           map[string]any
	Backdrop        *string
	BackdropPath    []string
	Bitrate         int
	Cast            string
	Country         *string
	CoverBig        *string
	Description     *string
	Director        string
	Duration        time.Duration
	DurationSeconds int
	Genre           string
	MovieImage      string
	Name            *string
	OriginalName    *string
	Plot            string
	Rating          float64
	ReleaseDate     time.Time
	Runtime         *int
	Status          *string
	TMDBID          int
	Video           map[string]any
	YoutubeTrailer  string
}

// fromInternal populates a VODInfo from an internal serde.VODInfo
func (v *VODInfo) fromInternal(internal *serde.VODInfo) {
	v.Actors = internal.Actors
	v.Age = internal.Age
	v.Audio = internal.Audio
	v.Backdrop = internal.Backdrop
	v.BackdropPath = []string(internal.BackdropPath)
	v.Bitrate = internal.Bitrate
	v.Cast = internal.Cast
	v.Country = internal.Country
	v.CoverBig = internal.CoverBig
	v.Description = internal.Description
	v.Director = internal.Director
	v.Duration = (time.Duration)(internal.Duration.Duration)
	v.DurationSeconds = int(internal.DurationSeconds)
	v.Genre = internal.Genre
	v.MovieImage = internal.MovieImage
	v.Name = internal.Name
	v.OriginalName = internal.OriginalName
	v.Plot = internal.Plot
	v.Rating = float64(internal.Rating)
	v.ReleaseDate = time.Time(internal.ReleaseDate.Time)
	v.Runtime = (*int)(internal.Runtime)
	v.Status = internal.Status
	v.TMDBID = int(internal.TMDBID)
	v.Video = internal.Video
	v.YoutubeTrailer = internal.YoutubeTrailer
}

// toInternal converts a VODInfo to the internal serde.VODInfo representation
func (v *VODInfo) toInternal() *serde.VODInfo {
	return &serde.VODInfo{
		Actors:          v.Actors,
		Age:             v.Age,
		Audio:           v.Audio,
		Backdrop:        v.Backdrop,
		BackdropPath:    serde.SliceStringSliceString(v.BackdropPath),
		Bitrate:         v.Bitrate,
		Cast:            v.Cast,
		Country:         v.Country,
		CoverBig:        v.CoverBig,
		Description:     v.Description,
		Director:        v.Director,
		Duration:        serde.DurationAsString{Duration: v.Duration},
		DurationSeconds: serde.IntegerAsInteger(v.DurationSeconds),
		Genre:           v.Genre,
		MovieImage:      v.MovieImage,
		Name:            v.Name,
		OriginalName:    v.OriginalName,
		Plot:            v.Plot,
		Rating:          serde.Float64AsFloat(v.Rating),
		ReleaseDate:     serde.DateAsString{Time: v.ReleaseDate},
		Runtime:         (*serde.IntegerAsInteger)(v.Runtime),
		Status:          v.Status,
		TMDBID:          serde.IntegerAsInteger(v.TMDBID),
		Video:           v.Video,
		YoutubeTrailer:  v.YoutubeTrailer,
	}
}

// MovieData is the public representation of movie data with clean and simple types
type MovieData struct {
	AddedOn            time.Time
	CategoryID         *int
	CategoryIDs        []int
	ContainerExtension string
	CustomSID          *string
	DirectSource       string
	Name               string
	StreamID           int
}

// fromInternal populates a MovieData from an internal serde.MovieData
func (m *MovieData) fromInternal(internal *serde.MovieData) {
	m.AddedOn = time.Time(internal.AddedOn.Time)
	m.CategoryID = (*int)(internal.CategoryID)
	m.CategoryIDs = internal.CategoryIDs
	m.ContainerExtension = internal.ContainerExtension
	m.CustomSID = internal.CustomSID
	m.DirectSource = internal.DirectSource
	m.Name = internal.Name
	m.StreamID = internal.StreamID
}

// toInternal converts a MovieData to the internal serde.MovieData representation
func (m *MovieData) toInternal() *serde.MovieData {
	return &serde.MovieData{
		AddedOn:            serde.UnixTimeAsInteger{Time: m.AddedOn},
		CategoryID:         (*serde.IntegerAsInteger)(m.CategoryID),
		CategoryIDs:        m.CategoryIDs,
		ContainerExtension: m.ContainerExtension,
		CustomSID:          m.CustomSID,
		DirectSource:       m.DirectSource,
		Name:               m.Name,
		StreamID:           m.StreamID,
	}
}
