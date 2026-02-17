# Xtream Codes API Client

A Go library for interacting with the Xtream Codes API. This client provides a clean, type-safe interface to the Xtream Codes API, handling all the complexity of serialization and deserialization between Go's strongly-typed system and the sometimes inconsistent API responses.

## Installation

```bash
go get github.com/sherif-fanous/xtreamcodes
```

## Usage

### Initializing the Client

```go
import (
    "github.com/sherif-fanous/xtreamcodes"
)

// Basic client initialization
client := xtreamcodes.NewClient("https://your-iptv-server.com", "username", "password")

// With a custom HTTP client
httpClient := &http.Client{
    Timeout: 30 * time.Second,
}
client := xtreamcodes.NewClient(
    "https://your-iptv-server.com",
    "username",
    "password",
    xtreamcodes.WithHTTPClient(httpClient),
)
```

### Authentication Information

```go
authInfo, err := client.GetAuthInfo(context.Background())
if err != nil {
    // Handle error
}
fmt.Printf("User info: %+v\n", authInfo.UserInfo)
fmt.Printf("Server info: %+v\n", authInfo.ServerInfo)
```

### Live Streams

```go
// List all live categories
liveCategories, err := client.ListLiveCategories(context.Background())
if err != nil {
    // Handle error
}

// List all live streams
liveStreams, err := client.ListLiveStreams(context.Background())
if err != nil {
    // Handle error
}

// List live streams in a specific category
categoryID := "1"  // Category ID as a string
liveStreamsInCategory, err := client.ListLiveStreamsInCategory(context.Background(), categoryID)
if err != nil {
    // Handle error
}
```

### Series

```go
// List all series categories
seriesCategories, err := client.ListSeriesCategories(context.Background())
if err != nil {
    // Handle error
}

// List all series
seriesStreams, err := client.ListSeriesStreams(context.Background())
if err != nil {
    // Handle error
}

// Get detailed information about a specific series
seriesID := "456"  // Series ID as a string
seriesInfo, err := client.GetSeriesInfo(context.Background(), seriesID)
if err != nil {
    // Handle error
}
```

### Video on Demand (VOD)

```go
// List all VOD categories
vodCategories, err := client.ListVODCategories(context.Background())
if err != nil {
    // Handle error
}

// List all VOD streams
vodStreams, err := client.ListVODStreams(context.Background())
if err != nil {
    // Handle error
}

// Get detailed information about a specific VOD
vodID := "123"  // VOD ID as a string
vodInfo, err := client.GetVODInfo(context.Background(), vodID)
if err != nil {
    // Handle error
}
```

### EPG (Electronic Program Guide)

```go
// Get EPG data for a specific stream
streamID := "789"  // Stream ID as a string

// Get short EPG data
shortEPG, err := client.GetShortEPG(context.Background(), streamID)
if err != nil {
    // Handle error
}

// Get short EPG data with a limit
limit := 5  // Number of EPG entries to retrieve
shortEPGWithLimit, err := client.GetShortEPGWithLimits(context.Background(), streamID, limit)
if err != nil {
    // Handle error
}

// Get all EPG data for a stream
allEPG, err := client.GetAllEPG(context.Background(), streamID)
if err != nil {
    // Handle error
}

// Get XMLTV data
xmltv, err := client.GetXMLTV(context.Background())
if err != nil {
    // Handle error
}
```

### M3U Playlist

```go
// Get M3U playlist
playlist, err := client.GetPlaylist(context.Background(), "m3u_plus", "ts")
if err != nil {
    // Handle error
}
```
