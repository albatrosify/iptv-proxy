package xtreamcodes

import (
	"encoding/json"
	"time"

	"github.com/sherif-fanous/xtreamcodes/internal/serde"
)

// AuthInfo is the public representation with clean and simple types
type AuthInfo struct {
	UserInfo   UserInfo
	ServerInfo ServerInfo
}

// MarshalJSON implements the json.Marshaler interface
func (a *AuthInfo) MarshalJSON() ([]byte, error) {
	internal := &serde.AuthInfo{
		UserInfo:   *a.UserInfo.toInternal(),
		ServerInfo: *a.ServerInfo.toInternal(),
	}

	return json.Marshal(internal)
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (a *AuthInfo) UnmarshalJSON(data []byte) error {
	var internal serde.AuthInfo
	if err := json.Unmarshal(data, &internal); err != nil {
		return err
	}

	a.UserInfo.fromInternal(&internal.UserInfo)
	a.ServerInfo.fromInternal(&internal.ServerInfo)

	return nil
}

// UserInfo is the public representation with clean and simple types
type UserInfo struct {
	ActiveConnections    int
	AllowedOutputFormats []string
	CreatedAt            time.Time
	ExpiresAt            time.Time
	IsAuthorized         bool
	IsTrial              bool
	MaxConnections       int
	Message              string
	Password             string
	Status               string
	Username             string
}

// fromInternal populates a UserInfo from an internal models.UserInfo
func (u *UserInfo) fromInternal(internal *serde.UserInfo) {
	u.ActiveConnections = int(internal.ActiveConnections)
	u.AllowedOutputFormats = internal.AllowedOutputFormats
	u.CreatedAt = time.Time(internal.CreatedAt.Time)
	u.ExpiresAt = time.Time(internal.ExpiresAt.Time)
	u.IsAuthorized = bool(internal.IsAuthorized)
	u.IsTrial = bool(internal.IsTrial)
	u.MaxConnections = int(internal.MaxConnections)
	u.Message = internal.Message
	u.Password = internal.Password
	u.Status = internal.Status
	u.Username = internal.Username
}

// toInternal converts a UserInfo to the internal models.UserInfo representation
func (u *UserInfo) toInternal() *serde.UserInfo {
	return &serde.UserInfo{
		ActiveConnections:    serde.IntegerAsInteger(u.ActiveConnections),
		AllowedOutputFormats: u.AllowedOutputFormats,
		CreatedAt:            serde.UnixTimeAsInteger{Time: u.CreatedAt},
		ExpiresAt:            serde.UnixTimeAsInteger{Time: u.ExpiresAt},
		IsAuthorized:         serde.BooleanAsInteger(u.IsAuthorized),
		IsTrial:              serde.BooleanAsInteger(u.IsTrial),
		MaxConnections:       serde.IntegerAsInteger(u.MaxConnections),
		Message:              u.Message,
		Password:             u.Password,
		Status:               u.Status,
		Username:             u.Username,
	}
}

// ServerInfo is the public representation with clean and simple types
type ServerInfo struct {
	HTTPPort       int
	HTTPSPort      int
	Process        bool
	RTMPPort       int
	ServerProtocol string
	TimeNow        time.Time
	TimestampNow   time.Time
	Timezone       string
	URL            string
}

// fromInternal populates a ServerInfo from an internal models.ServerInfo
func (s *ServerInfo) fromInternal(internal *serde.ServerInfo) {
	s.HTTPPort = int(internal.HTTPPort)
	s.HTTPSPort = int(internal.HTTPSPort)
	s.Process = internal.Process
	s.RTMPPort = int(internal.RTMPPort)
	s.ServerProtocol = internal.ServerProtocol
	s.TimeNow = time.Time(internal.TimeNow.Time)
	if location, err := time.LoadLocation(internal.Timezone); err == nil {
		s.TimeNow = time.Date(internal.TimeNow.Year(), internal.TimeNow.Month(),
			internal.TimeNow.Day(), internal.TimeNow.Hour(),
			internal.TimeNow.Minute(), internal.TimeNow.Second(),
			internal.TimeNow.Nanosecond(), location)
	}
	s.TimestampNow = time.Time(internal.TimestampNow.Time)
	s.Timezone = internal.Timezone
	s.URL = internal.URL
}

// toInternal converts a ServerInfo to the internal models.ServerInfo representation
func (s *ServerInfo) toInternal() *serde.ServerInfo {
	return &serde.ServerInfo{
		HTTPPort:       serde.IntegerAsInteger(s.HTTPPort),
		HTTPSPort:      serde.IntegerAsInteger(s.HTTPSPort),
		Process:        s.Process,
		RTMPPort:       serde.IntegerAsInteger(s.RTMPPort),
		ServerProtocol: s.ServerProtocol,
		TimeNow:        serde.DateTimeAsString{Time: s.TimeNow},
		TimestampNow:   serde.UnixTimeAsInteger{Time: s.TimestampNow},
		Timezone:       s.Timezone,
		URL:            s.URL,
	}
}
