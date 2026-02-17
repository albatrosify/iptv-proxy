package serde

import "encoding/json"

// VODStream is the internal representation used for marshaling/unmarshaling
type VODStream struct {
	AddedOn            UnixTimeAsInteger `json:"added"`
	CategoryID         *IntegerAsInteger `json:"category_id"`
	CategoryIDs        []int             `json:"category_ids,omitempty"`
	ContainerExtension string            `json:"container_extension"`
	CustomSID          *string           `json:"custom_sid"`
	DirectSource       string            `json:"direct_source"`
	IsAdult            BooleanAsInteger  `json:"is_adult"`
	Name               string            `json:"name"`
	Number             int               `json:"num"`
	Rating             Float64AsFloat    `json:"rating"`
	Rating5Based       Float64AsFloat    `json:"rating_5based"`
	StreamIcon         string            `json:"stream_icon"`
	StreamID           int               `json:"stream_id"`
	StreamType         string            `json:"stream_type"`
	TMDBID             *IntegerAsInteger `json:"tmdb,omitempty"`
	Trailer            *string           `json:"trailer,omitempty"`
}

// VOD is the internal representation used for marshaling/unmarshaling
type VOD struct {
	Info      VODInfo   `json:"info"`
	MovieData MovieData `json:"movie_data"`
}

// VODInfo is the internal representation used for marshaling/unmarshaling
type VODInfo struct {
	Actors          *string                `json:"actors,omitempty"`
	Age             *string                `json:"age,omitempty"`
	Audio           map[string]any         `json:"audio"`
	Backdrop        *string                `json:"backdrop,omitempty"`
	BackdropPath    SliceStringSliceString `json:"backdrop_path,omitempty"`
	Bitrate         int                    `json:"bitrate"`
	Cast            string                 `json:"cast"`
	Country         *string                `json:"country,omitempty"`
	CoverBig        *string                `json:"cover_big,omitempty"`
	Description     *string                `json:"description,omitempty"`
	Director        string                 `json:"director"`
	Duration        DurationAsString       `json:"duration"`
	DurationSeconds IntegerAsInteger       `json:"duration_secs"`
	Genre           string                 `json:"genre"`
	MovieImage      string                 `json:"movie_image"`
	Name            *string                `json:"name,omitempty"`
	OriginalName    *string                `json:"o_name,omitempty"`
	Plot            string                 `json:"plot"`
	Rating          Float64AsFloat         `json:"rating"`
	ReleaseDate     DateAsString           `json:"release_date"`
	Runtime         *IntegerAsInteger      `json:"runtime,omitempty"`
	Status          *string                `json:"status,omitempty"`
	TMDBID          IntegerAsInteger       `json:"tmdb_id"`
	Video           map[string]any         `json:"video"`
	YoutubeTrailer  string                 `json:"youtube_trailer"`
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (s *VODInfo) UnmarshalJSON(data []byte) error {
	type vodInfoWithAltFields struct {
		Actors               *string                `json:"actors,omitempty"`
		Age                  *string                `json:"age,omitempty"`
		Audio                map[string]any         `json:"audio"`
		Backdrop             *string                `json:"backdrop,omitempty"`
		BackdropPath         SliceStringSliceString `json:"backdrop_path,omitempty"`
		Bitrate              int                    `json:"bitrate"`
		Cast                 string                 `json:"cast"`
		Country              *string                `json:"country,omitempty"`
		CoverBig             *string                `json:"cover_big,omitempty"`
		Description          *string                `json:"description,omitempty"`
		Director             string                 `json:"director"`
		Duration             DurationAsString       `json:"duration"`
		DurationSeconds      IntegerAsInteger       `json:"duration_secs"`
		Genre                string                 `json:"genre"`
		MovieImage           string                 `json:"movie_image"`
		Name                 *string                `json:"name,omitempty"`
		OriginalName         *string                `json:"o_name,omitempty"`
		Plot                 string                 `json:"plot"`
		Rating               Float64AsFloat         `json:"rating"`
		ReleaseDateCamelCase DateAsString           `json:"releasedate"`
		ReleaseDateSnakeCase DateAsString           `json:"release_date"`
		Runtime              *IntegerAsInteger      `json:"runtime,omitempty"`
		Status               *string                `json:"status,omitempty"`
		TMDBID               IntegerAsInteger       `json:"tmdb_id"`
		Video                map[string]any         `json:"video"`
		YoutubeTrailer       string                 `json:"youtube_trailer"`
	}

	var v vodInfoWithAltFields
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	s.Actors = v.Actors
	s.Age = v.Age
	s.Audio = v.Audio
	s.Backdrop = v.Backdrop
	s.BackdropPath = v.BackdropPath
	s.Bitrate = v.Bitrate
	s.Cast = v.Cast
	s.Country = v.Country
	s.CoverBig = v.CoverBig
	s.Description = v.Description
	s.Director = v.Director
	s.Duration = v.Duration
	s.DurationSeconds = v.DurationSeconds
	s.Genre = v.Genre
	s.MovieImage = v.MovieImage
	s.Name = v.Name
	s.OriginalName = v.OriginalName
	s.Plot = v.Plot
	s.Rating = v.Rating
	s.Runtime = v.Runtime
	s.Status = v.Status
	s.TMDBID = v.TMDBID
	s.Video = v.Video
	s.YoutubeTrailer = v.YoutubeTrailer

	if !v.ReleaseDateSnakeCase.IsZero() {
		s.ReleaseDate = v.ReleaseDateSnakeCase
	} else if !v.ReleaseDateCamelCase.IsZero() {
		s.ReleaseDate = v.ReleaseDateCamelCase
	}

	return nil
}

// MovieData is the internal representation used for marshaling/unmarshaling
type MovieData struct {
	AddedOn            UnixTimeAsInteger `json:"added"`
	CategoryID         *IntegerAsInteger `json:"category_id"`
	CategoryIDs        []int             `json:"category_ids,omitempty"`
	ContainerExtension string            `json:"container_extension"`
	CustomSID          *string           `json:"custom_sid"`
	DirectSource       string            `json:"direct_source"`
	Name               string            `json:"name"`
	StreamID           int               `json:"stream_id"`
}
