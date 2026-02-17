package serde

import (
	"encoding/json"
)

// SeriesStream is the internal representation used for marshaling/unmarshaling
type SeriesStream struct {
	BackdropPath   SliceStringSliceString `json:"backdrop_path,omitempty"`
	Cast           string                 `json:"cast"`
	CategoryID     *IntegerAsInteger      `json:"category_id"`
	CategoryIDs    []int                  `json:"category_ids,omitempty"`
	Cover          string                 `json:"cover"`
	Director       string                 `json:"director"`
	EpisodeRunTime IntegerAsInteger       `json:"episode_run_time"`
	Genre          string                 `json:"genre"`
	LastModifiedOn UnixTimeAsInteger      `json:"last_modified"`
	Name           string                 `json:"name"`
	Number         int                    `json:"num"`
	Plot           string                 `json:"plot"`
	Rating         Float64AsFloat         `json:"rating"`
	Rating5Based   Float64AsFloat         `json:"rating_5based"`
	ReleaseDate    DateAsString           `json:"release_date"`
	SeriesID       int                    `json:"series_id"`
	TMDBID         *IntegerAsInteger      `json:"tmdb,omitempty"`
	YoutubeTrailer string                 `json:"youtube_trailer"`
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (s *SeriesStream) UnmarshalJSON(data []byte) error {
	type seriesStreamWithAltFields struct {
		BackdropPath         SliceStringSliceString `json:"backdrop_path"`
		Cast                 string                 `json:"cast"`
		CategoryID           *IntegerAsInteger      `json:"category_id"`
		CategoryIDs          []int                  `json:"category_ids"`
		Cover                string                 `json:"cover"`
		Director             string                 `json:"director"`
		EpisodeRunTime       IntegerAsInteger       `json:"episode_run_time"`
		Genre                string                 `json:"genre"`
		LastModifiedOn       UnixTimeAsInteger      `json:"last_modified"`
		Name                 string                 `json:"name"`
		Number               int                    `json:"num"`
		Plot                 string                 `json:"plot"`
		Rating               Float64AsFloat         `json:"rating"`
		Rating5Based         Float64AsFloat         `json:"rating_5based"`
		ReleaseDateCamelCase DateAsString           `json:"releaseDate"`
		ReleaseDateSnakeCase DateAsString           `json:"release_date"`
		SeriesID             int                    `json:"series_id"`
		TMDBID               *IntegerAsInteger      `json:"tmdb"`
		YoutubeTrailer       string                 `json:"youtube_trailer"`
	}

	var v seriesStreamWithAltFields
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	s.BackdropPath = v.BackdropPath
	s.Cast = v.Cast
	s.CategoryID = v.CategoryID
	s.CategoryIDs = v.CategoryIDs
	s.Cover = v.Cover
	s.Director = v.Director
	s.EpisodeRunTime = v.EpisodeRunTime
	s.Genre = v.Genre
	s.LastModifiedOn = v.LastModifiedOn
	s.Name = v.Name
	s.Number = v.Number
	s.Plot = v.Plot
	s.Rating = v.Rating
	s.Rating5Based = v.Rating5Based
	s.SeriesID = v.SeriesID
	s.TMDBID = v.TMDBID
	s.YoutubeTrailer = v.YoutubeTrailer

	if !v.ReleaseDateSnakeCase.IsZero() {
		s.ReleaseDate = v.ReleaseDateSnakeCase
	} else if !v.ReleaseDateCamelCase.IsZero() {
		s.ReleaseDate = v.ReleaseDateCamelCase
	}

	return nil
}

// Series is the internal representation used for marshaling/unmarshaling
type Series struct {
	Episodes map[string][]Episode `json:"episodes"`
	Info     SeriesInfo           `json:"info"`
	Seasons  []Season             `json:"seasons"`
}

// Episode is the internal representation used for marshaling/unmarshaling
type Episode struct {
	AddedOn            UnixTimeAsInteger `json:"added"`
	ContainerExtension string            `json:"container_extension"`
	CustomSID          *string           `json:"custom_sid"`
	DirectSource       string            `json:"direct_source"`
	EpisodeInfo        map[string]any    `json:"info"`
	EpisodeNumber      int               `json:"episode_num"`
	ID                 IntegerAsInteger  `json:"id"`
	Season             int               `json:"season"`
	Title              string            `json:"title"`
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (e *Episode) UnmarshalJSON(data []byte) error {
	type EpisodeAlias Episode

	v := struct {
		*EpisodeAlias
		Info json.RawMessage `json:"info"`
	}{
		EpisodeAlias: (*EpisodeAlias)(e),
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	if len(v.Info) == 0 || string(v.Info) == "[]" || string(v.Info) == "null" {
		e.EpisodeInfo = make(map[string]any, 0)

		return nil
	}

	return json.Unmarshal(v.Info, &e.EpisodeInfo)
}

// SeriesInfo is the internal representation used for marshaling/unmarshaling
type SeriesInfo struct {
	BackdropPath   SliceStringSliceString `json:"backdrop_path,omitempty"`
	Cast           string                 `json:"cast"`
	CategoryID     *IntegerAsInteger      `json:"category_id"`
	CategoryIDs    []int                  `json:"category_ids,omitempty"`
	Cover          string                 `json:"cover"`
	Director       string                 `json:"director"`
	EpisodeRunTime IntegerAsInteger       `json:"episode_run_time"`
	Genre          string                 `json:"genre"`
	LastModifiedOn UnixTimeAsInteger      `json:"last_modified"`
	Name           string                 `json:"name"`
	Plot           string                 `json:"plot"`
	Rating         Float64AsFloat         `json:"rating"`
	Rating5Based   Float64AsFloat         `json:"rating_5based"`
	ReleaseDate    DateAsString           `json:"release_date"`
	TMDBID         *IntegerAsInteger      `json:"tmdb,omitempty"`
	YoutubeTrailer string                 `json:"youtube_trailer"`
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (s *SeriesInfo) UnmarshalJSON(data []byte) error {
	type seriesInfoWithAltFields struct {
		BackdropPath         SliceStringSliceString `json:"backdrop_path,omitempty"`
		Cast                 string                 `json:"cast"`
		CategoryID           *IntegerAsInteger      `json:"category_id"`
		CategoryIDs          []int                  `json:"category_ids,omitempty"`
		Cover                string                 `json:"cover"`
		Director             string                 `json:"director"`
		EpisodeRunTime       IntegerAsInteger       `json:"episode_run_time"`
		Genre                string                 `json:"genre"`
		LastModifiedOn       UnixTimeAsInteger      `json:"last_modified"`
		Name                 string                 `json:"name"`
		Plot                 string                 `json:"plot"`
		Rating               Float64AsFloat         `json:"rating"`
		Rating5Based         Float64AsFloat         `json:"rating_5based"`
		ReleaseDateCamelCase DateAsString           `json:"releaseDate"`
		ReleaseDateSnakeCase DateAsString           `json:"release_date"`
		TMDBID               *IntegerAsInteger      `json:"tmdb,omitempty"`
		YoutubeTrailer       string                 `json:"youtube_trailer"`
	}

	var v seriesInfoWithAltFields
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	s.BackdropPath = v.BackdropPath
	s.Cast = v.Cast
	s.CategoryID = v.CategoryID
	s.CategoryIDs = v.CategoryIDs
	s.Cover = v.Cover
	s.Director = v.Director
	s.EpisodeRunTime = v.EpisodeRunTime
	s.Genre = v.Genre
	s.LastModifiedOn = v.LastModifiedOn
	s.Name = v.Name
	s.Plot = v.Plot
	s.Rating = v.Rating
	s.Rating5Based = v.Rating5Based
	s.TMDBID = v.TMDBID
	s.YoutubeTrailer = v.YoutubeTrailer

	if !v.ReleaseDateSnakeCase.IsZero() {
		s.ReleaseDate = v.ReleaseDateSnakeCase
	} else if !v.ReleaseDateCamelCase.IsZero() {
		s.ReleaseDate = v.ReleaseDateCamelCase
	}

	return nil
}

// Season is the internal representation used for marshaling/unmarshaling
type Season struct {
	AirDate      DateAsString      `json:"air_date"`
	Cover        string            `json:"cover"`
	CoverBig     string            `json:"cover_big"`
	CoverTMDB    *string           `json:"cover_tmdb,omitempty"`
	Duration     *IntegerAsInteger `json:"duration,omitempty"`
	EpisodeCount IntegerAsInteger  `json:"episode_count"`
	ID           *int              `json:"id,omitempty"`
	Name         string            `json:"name"`
	Overview     string            `json:"overview"`
	ReleaseDate  *DateAsString     `json:"release_date,omitempty"`
	SeasonNumber int               `json:"season_number"`
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (s *Season) UnmarshalJSON(data []byte) error {
	type seasonWithAltFields struct {
		AirDate              DateAsString      `json:"air_date"`
		Cover                string            `json:"cover"`
		CoverBig             string            `json:"cover_big"`
		CoverTMDB            *string           `json:"cover_tmdb,omitempty"`
		Duration             *IntegerAsInteger `json:"duration,omitempty"`
		EpisodeCount         IntegerAsInteger  `json:"episode_count"`
		ID                   *int              `json:"id,omitempty"`
		Name                 string            `json:"name"`
		Overview             string            `json:"overview"`
		ReleaseDateCamelCase *DateAsString     `json:"releaseDate"`
		ReleaseDateSnakeCase *DateAsString     `json:"release_date"`
		SeasonNumber         int               `json:"season_number"`
	}

	var v seasonWithAltFields
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	s.AirDate = v.AirDate
	s.Cover = v.Cover
	s.CoverBig = v.CoverBig
	s.CoverTMDB = v.CoverTMDB
	s.Duration = v.Duration
	s.EpisodeCount = v.EpisodeCount
	s.ID = v.ID
	s.Name = v.Name
	s.Overview = v.Overview
	s.SeasonNumber = v.SeasonNumber

	if v.ReleaseDateSnakeCase != nil {
		s.ReleaseDate = v.ReleaseDateSnakeCase
	} else if v.ReleaseDateCamelCase != nil {
		s.ReleaseDate = v.ReleaseDateCamelCase
	}

	return nil
}
