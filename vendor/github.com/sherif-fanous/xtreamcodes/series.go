package xtreamcodes

import (
	"encoding/json"
	"time"

	"github.com/sherif-fanous/xtreamcodes/internal/serde"
)

// SeriesStream is the public representation with clean and simple types
type SeriesStream struct {
	BackdropPath   []string
	Cast           string
	CategoryID     *int
	CategoryIDs    []int
	Cover          string
	Director       string
	EpisodeRunTime int
	Genre          string
	LastModifiedOn time.Time
	Name           string
	Number         int
	Plot           string
	Rating         float64
	Rating5Based   float64
	ReleaseDate    time.Time
	SeriesID       int
	TMDBID         *int
	YoutubeTrailer string
}

// fromInternal populates a SeriesStream from an internal models.SeriesStream
func (s *SeriesStream) fromInternal(internal *serde.SeriesStream) {
	s.BackdropPath = internal.BackdropPath
	s.Cast = internal.Cast
	s.CategoryID = (*int)(internal.CategoryID)
	s.CategoryIDs = internal.CategoryIDs
	s.Cover = internal.Cover
	s.Director = internal.Director
	s.EpisodeRunTime = int(internal.EpisodeRunTime)
	s.Genre = internal.Genre
	s.LastModifiedOn = time.Time(internal.LastModifiedOn.Time)
	s.Name = internal.Name
	s.Number = internal.Number
	s.Plot = internal.Plot
	s.Rating = float64(internal.Rating)
	s.Rating5Based = float64(internal.Rating5Based)
	s.ReleaseDate = time.Time(internal.ReleaseDate.Time)
	s.SeriesID = internal.SeriesID
	s.TMDBID = (*int)(internal.TMDBID)
	s.YoutubeTrailer = internal.YoutubeTrailer
}

// toInternal converts a SeriesStream to the internal models.SeriesStream representation
func (s *SeriesStream) toInternal() *serde.SeriesStream {
	return &serde.SeriesStream{
		BackdropPath:   s.BackdropPath,
		Cast:           s.Cast,
		CategoryID:     (*serde.IntegerAsInteger)(s.CategoryID),
		CategoryIDs:    s.CategoryIDs,
		Cover:          s.Cover,
		Director:       s.Director,
		EpisodeRunTime: serde.IntegerAsInteger(s.EpisodeRunTime),
		Genre:          s.Genre,
		LastModifiedOn: serde.UnixTimeAsInteger{Time: s.LastModifiedOn},
		Name:           s.Name,
		Number:         s.Number,
		Plot:           s.Plot,
		Rating:         serde.Float64AsFloat(s.Rating),
		Rating5Based:   serde.Float64AsFloat(s.Rating5Based),
		ReleaseDate:    serde.DateAsString{Time: s.ReleaseDate},
		SeriesID:       s.SeriesID,
		TMDBID:         (*serde.IntegerAsInteger)(s.TMDBID),
		YoutubeTrailer: s.YoutubeTrailer,
	}
}

// MarshalJSON implements the json.Marshaler interface
func (s *SeriesStream) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.toInternal())
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (s *SeriesStream) UnmarshalJSON(data []byte) error {
	var internal serde.SeriesStream
	if err := json.Unmarshal(data, &internal); err != nil {
		return err
	}

	s.fromInternal(&internal)

	return nil
}

// Series is the public representation with clean and simple types
type Series struct {
	Episodes map[string][]Episode
	Info     SeriesInfo
	Seasons  []Season
}

// fromInternal populates Series from its internal representation
func (s *Series) fromInternal(internal *serde.Series) {
	s.Episodes = make(map[string][]Episode)
	for seasonKey, episodes := range internal.Episodes {
		publicEpisodes := make([]Episode, len(episodes))
		for i, episode := range episodes {
			publicEpisodes[i].fromInternal(&episode)
		}
		s.Episodes[seasonKey] = publicEpisodes
	}

	s.Info.fromInternal(&internal.Info)

	s.Seasons = make([]Season, len(internal.Seasons))
	for i, season := range internal.Seasons {
		s.Seasons[i].fromInternal(&season)
	}
}

// toInternal converts Series to its internal representation
func (s *Series) toInternal() *serde.Series {
	internal := &serde.Series{
		Episodes: make(map[string][]serde.Episode),
		Info:     *s.Info.toInternal(),
		Seasons:  make([]serde.Season, len(s.Seasons)),
	}

	for seasonKey, episodes := range s.Episodes {
		internalEpisodes := make([]serde.Episode, len(episodes))
		for i, episode := range episodes {
			internalEpisodes[i] = *episode.toInternal()
		}
		internal.Episodes[seasonKey] = internalEpisodes
	}

	for i, season := range s.Seasons {
		internal.Seasons[i] = *season.toInternal()
	}

	return internal
}

// MarshalJSON implements the json.Marshaler interface
func (s *Series) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.toInternal())
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (s *Series) UnmarshalJSON(data []byte) error {
	var internal serde.Series
	if err := json.Unmarshal(data, &internal); err != nil {
		return err
	}

	s.fromInternal(&internal)

	return nil
}

// Episode is the public representation with clean and simple types
type Episode struct {
	AddedOn            time.Time
	ContainerExtension string
	CustomSID          *string
	DirectSource       string
	EpisodeInfo        map[string]any
	EpisodeNumber      int
	ID                 int
	Season             int
	Title              string
}

// fromInternal populates Episode from its internal representation
func (e *Episode) fromInternal(internal *serde.Episode) {
	e.AddedOn = time.Time(internal.AddedOn.Time)
	e.ContainerExtension = internal.ContainerExtension
	e.CustomSID = internal.CustomSID
	e.DirectSource = internal.DirectSource
	e.EpisodeInfo = internal.EpisodeInfo
	e.EpisodeNumber = internal.EpisodeNumber
	e.ID = int(internal.ID)
	e.Season = internal.Season
	e.Title = internal.Title
}

// toInternal converts Episode to its internal representation
func (e *Episode) toInternal() *serde.Episode {
	return &serde.Episode{
		AddedOn:            serde.UnixTimeAsInteger{Time: e.AddedOn},
		ContainerExtension: e.ContainerExtension,
		CustomSID:          e.CustomSID,
		DirectSource:       e.DirectSource,
		EpisodeInfo:        e.EpisodeInfo,
		EpisodeNumber:      e.EpisodeNumber,
		ID:                 serde.IntegerAsInteger(e.ID),
		Season:             e.Season,
		Title:              e.Title,
	}
}

// SeriesInfo is the public representation with clean and simple types
type SeriesInfo struct {
	BackdropPath   []string
	Cast           string
	CategoryID     *int
	CategoryIDs    []int
	Cover          string
	Director       string
	EpisodeRunTime int
	Genre          string
	LastModifiedOn time.Time
	Name           string
	Plot           string
	Rating         float64
	Rating5Based   float64
	ReleaseDate    time.Time
	TMDBID         *int
	YoutubeTrailer string
}

// fromInternal populates SeriesInfo from its internal representation
func (s *SeriesInfo) fromInternal(internal *serde.SeriesInfo) {
	s.BackdropPath = ([]string)(internal.BackdropPath)
	s.Cast = internal.Cast
	s.CategoryID = (*int)(internal.CategoryID)
	s.CategoryIDs = internal.CategoryIDs
	s.Cover = internal.Cover
	s.Director = internal.Director
	s.EpisodeRunTime = int(internal.EpisodeRunTime)
	s.Genre = internal.Genre
	s.LastModifiedOn = time.Time(internal.LastModifiedOn.Time)
	s.Name = internal.Name
	s.Plot = internal.Plot
	s.Rating = float64(internal.Rating)
	s.Rating5Based = float64(internal.Rating5Based)
	s.ReleaseDate = time.Time(internal.ReleaseDate.Time)
	s.TMDBID = (*int)(internal.TMDBID)
	s.YoutubeTrailer = internal.YoutubeTrailer
}

// toInternal converts SeriesInfo to its internal representation
func (s *SeriesInfo) toInternal() *serde.SeriesInfo {
	return &serde.SeriesInfo{
		BackdropPath:   (serde.SliceStringSliceString)(s.BackdropPath),
		Cast:           s.Cast,
		CategoryID:     (*serde.IntegerAsInteger)(s.CategoryID),
		CategoryIDs:    s.CategoryIDs,
		Cover:          s.Cover,
		Director:       s.Director,
		EpisodeRunTime: serde.IntegerAsInteger(s.EpisodeRunTime),
		Genre:          s.Genre,
		LastModifiedOn: serde.UnixTimeAsInteger{Time: s.LastModifiedOn},
		Name:           s.Name,
		Plot:           s.Plot,
		Rating:         serde.Float64AsFloat(s.Rating),
		Rating5Based:   serde.Float64AsFloat(s.Rating5Based),
		ReleaseDate:    serde.DateAsString{Time: s.ReleaseDate},
		TMDBID:         (*serde.IntegerAsInteger)(s.TMDBID),
		YoutubeTrailer: s.YoutubeTrailer,
	}
}

// Season is the public representation with clean and simple types
type Season struct {
	AirDate      time.Time
	Cover        string
	CoverBig     string
	CoverTMDB    *string
	Duration     *int
	EpisodeCount int
	ID           *int
	Name         string
	Overview     string
	ReleaseDate  *time.Time
	SeasonNumber int
}

// fromInternal populates Season from its internal representation
func (s *Season) fromInternal(internal *serde.Season) {
	s.AirDate = time.Time(internal.AirDate.Time)
	s.Cover = internal.Cover
	s.CoverBig = internal.CoverBig
	s.CoverTMDB = internal.CoverTMDB
	s.Duration = (*int)(internal.Duration)
	s.EpisodeCount = int(internal.EpisodeCount)
	s.ID = internal.ID
	s.Name = internal.Name
	s.Overview = internal.Overview
	if internal.ReleaseDate != nil {
		s.ReleaseDate = (*time.Time)(&internal.ReleaseDate.Time)
	}
	s.SeasonNumber = internal.SeasonNumber
}

// toInternal converts Season to its internal representation
func (s *Season) toInternal() *serde.Season {
	var releaseDate *serde.DateAsString
	if s.ReleaseDate != nil {
		releaseDate = &serde.DateAsString{Time: *s.ReleaseDate}
	}

	return &serde.Season{
		AirDate:      serde.DateAsString{Time: s.AirDate},
		Cover:        s.Cover,
		CoverBig:     s.CoverBig,
		CoverTMDB:    s.CoverTMDB,
		Duration:     (*serde.IntegerAsInteger)(s.Duration),
		EpisodeCount: serde.IntegerAsInteger(s.EpisodeCount),
		ID:           s.ID,
		Name:         s.Name,
		Overview:     s.Overview,
		ReleaseDate:  releaseDate,
		SeasonNumber: s.SeasonNumber,
	}
}
