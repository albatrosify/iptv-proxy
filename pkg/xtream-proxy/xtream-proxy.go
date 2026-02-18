/*
 * Iptv-Proxy is a project to proxyfie an m3u file and to proxyfie an Xtream iptv service (client API).
 * Copyright (C) 2020  Pierre-Emmanuel Jacquier
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package xtreamproxy

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/pierre-emmanuelJ/iptv-proxy/pkg/config"
	"github.com/pierre-emmanuelJ/iptv-proxy/pkg/utils"
	"github.com/sherif-fanous/xmltv"
	xtream "github.com/sherif-fanous/xtreamcodes"
)

const (
	getLiveCategories   = "get_live_categories"
	getLiveStreams      = "get_live_streams"
	getVodCategories    = "get_vod_categories"
	getVodStreams       = "get_vod_streams"
	getVodInfo          = "get_vod_info"
	getSeriesCategories = "get_series_categories"
	getSeries           = "get_series"
	getSerieInfo        = "get_series_info"
	getShortEPG         = "get_short_epg"
	getSimpleDataTable  = "get_simple_data_table"
)

// Client represent an xtream client
type Client struct {
	*xtream.Client
	baseURL   string
	username  string
	password  string
	userAgent string
}

// New new xtream client
func New(user, password, baseURL, userAgent string) (*Client, error) {
	cli := xtream.NewClient(baseURL, user, password, xtream.WithUserAgent(userAgent))
	return &Client{
		Client:    cli,
		baseURL:   baseURL,
		username:  user,
		password:  password,
		userAgent: userAgent,
	}, nil
}

type login struct {
	UserInfo   xtream.UserInfo   `json:"user_info"`
	ServerInfo xtream.ServerInfo `json:"server_info"`
}

// Login xtream login
func (c *Client) login(proxyUser, proxyPassword, proxyURL string, proxyPort int, protocol string) (login, error) {
	// Note: The new library returns specific types. We need to map them back to what the proxy expects
	// or update the proxy to use the new types.
	// This function seems to construct a response object to return to the client.
	// We might need to manually construct the struct since we don't have the "internal" one.

	// For now, let's assume we can construct the clean structs directly.
	// However, we need to populate them.
	// Since we are mocking the response for the proxy, we can just set values.

	// But wait, where does c.UserInfo come from?
	// The original code accessed c.UserInfo. It seems the old client struct exposed the UserInfo directly?
	// The new Client struct doesn't expose UserInfo/ServerInfo directly. We have to fetch it.

	ctx := context.Background()
	authInfo, err := c.GetAuthInfo(ctx)
	if err != nil {
		return login{}, err
	}

	req := login{
		UserInfo: xtream.UserInfo{
			Username:             proxyUser,
			Password:             proxyPassword,
			Message:              authInfo.UserInfo.Message,
			IsAuthorized:         authInfo.UserInfo.IsAuthorized, // Mapped from Auth
			Status:               authInfo.UserInfo.Status,
			ExpiresAt:            authInfo.UserInfo.ExpiresAt, // Mapped from ExpDate
			IsTrial:              authInfo.UserInfo.IsTrial,
			ActiveConnections:    authInfo.UserInfo.ActiveConnections,
			CreatedAt:            authInfo.UserInfo.CreatedAt,
			MaxConnections:       authInfo.UserInfo.MaxConnections,
			AllowedOutputFormats: authInfo.UserInfo.AllowedOutputFormats,
		},
		ServerInfo: xtream.ServerInfo{
			URL:            proxyURL,
			HTTPPort:       proxyPort, // Mapped from Port
			HTTPSPort:      proxyPort,
			ServerProtocol: protocol, // Mapped from Protocol
			RTMPPort:       proxyPort,
			Timezone:       authInfo.ServerInfo.Timezone,
			TimestampNow:   authInfo.ServerInfo.TimestampNow,
			TimeNow:        authInfo.ServerInfo.TimeNow,
		},
	}

	return req, nil
}

// Action execute an xtream action.
func (c *Client) Action(config *config.ProxyConfig, action string, q url.Values) (respBody interface{}, httpcode int, contentType string, err error) {
	protocol := "http"
	if config.HTTPS {
		protocol = "https"
	}

	// Default content type for most responses
	contentType = "application/json"
	ctx := context.Background()

	switch action {
	case getLiveCategories:
		respBody, err = c.ListLiveCategories(ctx)
		if err != nil {
			err = utils.PrintErrorAndReturn(err)
		}
	case getLiveStreams:
		categoryID := ""
		if len(q["category_id"]) > 0 {
			categoryID = q["category_id"][0]
		}
		if categoryID != "" {
			respBody, err = c.ListLiveStreamsInCategory(ctx, categoryID)
		} else {
			respBody, err = c.ListLiveStreams(ctx)
		}

		if err != nil {
			err = utils.PrintErrorAndReturn(err)
		}
	case getVodCategories:
		respBody, err = c.ListVODCategories(ctx)
		if err != nil {
			err = utils.PrintErrorAndReturn(err)
		}
	case getVodStreams:
		categoryID := ""
		if len(q["category_id"]) > 0 {
			categoryID = q["category_id"][0]
		}
		if categoryID != "" {
			respBody, err = c.ListVODStreamsInCategory(ctx, categoryID)
		} else {
			respBody, err = c.ListVODStreams(ctx)
		}

		if err != nil {
			err = utils.PrintErrorAndReturn(err)
		}
	case getVodInfo:
		httpcode, err = validateParams(q, "vod_id")
		if err != nil {
			err = utils.PrintErrorAndReturn(err)
			return
		}
		respBody, err = c.GetVODInfo(ctx, q["vod_id"][0])
		if err != nil {
			err = utils.PrintErrorAndReturn(err)
		}
	case getSeriesCategories:
		respBody, err = c.ListSeriesCategories(ctx)
		if err != nil {
			err = utils.PrintErrorAndReturn(err)
		}
	case getSeries:
		categoryID := ""
		if len(q["category_id"]) > 0 {
			categoryID = q["category_id"][0]
		}
		if categoryID != "" {
			respBody, err = c.ListSeriesStreamsInCategory(ctx, categoryID)
		} else {
			respBody, err = c.ListSeriesStreams(ctx)
		}

		if err != nil {
			err = utils.PrintErrorAndReturn(err)
		}
	case getSerieInfo:
		httpcode, err = validateParams(q, "series_id")
		if err != nil {
			err = utils.PrintErrorAndReturn(err)
			return
		}
		respBody, err = c.GetSeriesInfo(ctx, q["series_id"][0])
		if err != nil {
			err = utils.PrintErrorAndReturn(err)
		}
	case getShortEPG:
		limit := 0

		httpcode, err = validateParams(q, "stream_id")
		if err != nil {
			err = utils.PrintErrorAndReturn(err)
			return
		}
		if len(q["limit"]) > 0 {
			limit, err = strconv.Atoi(q["limit"][0])
			if err != nil {
				httpcode = http.StatusInternalServerError
				err = utils.PrintErrorAndReturn(err)
				return
			}
		}
		if limit > 0 {
			respBody, err = c.GetShortEPGWithLimits(ctx, q["stream_id"][0], limit)
		} else {
			respBody, err = c.GetShortEPG(ctx, q["stream_id"][0])
		}

		if err != nil {
			err = utils.PrintErrorAndReturn(err)
		}
	case getSimpleDataTable:
		httpcode, err = validateParams(q, "stream_id")
		if err != nil {
			err = utils.PrintErrorAndReturn(err)
			return
		}
		respBody, err = c.GetAllEPG(ctx, q["stream_id"][0])
		if err != nil {
			err = utils.PrintErrorAndReturn(err)
		}
	default:
		respBody, err = c.login(config.User.String(), config.Password.String(), protocol+"://"+config.HostConfig.Hostname, config.AdvertisedPort, protocol)
		if err != nil {
			err = utils.PrintErrorAndReturn(err)
		}
	}

	return
}

func validateParams(u url.Values, params ...string) (int, error) {
	for _, p := range params {
		if len(u[p]) < 1 {
			return http.StatusBadRequest, fmt.Errorf("missing %q", p)
		}

	}

	return 0, nil
}

// GetXMLTV retrieves the XMLTV EPG data.
// Overrides the default xtream.Client.GetXMLTV method to handle date parsing issues.
func (c *Client) GetXMLTV(ctx context.Context) (*xmltv.EPG, error) {
	// 1. Construct the URL manually as we need to customize the request or handle raw response
	// The original client uses internal buildURL logic which we can replicate simply.
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, err
	}
	u.Path = "/xmltv.php" //Standard path for XMLTV in Xtream Codes

	q := u.Query()
	q.Set("username", c.username)
	q.Set("password", c.password)
	u.RawQuery = q.Encode()

	// 2. Create Request
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.userAgent)

	// 3. Execute Request
	// Note: We use default client here, ideally we should reuse the one from xtream.Client but it is unexported.
	// We can create a new one.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("XMLTV request failed with status: %d", resp.StatusCode)
	}

	// 4. Read Body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	bodyStr := string(bodyBytes)

	// 5. Sanitize Body
	// Replace <date>YYYY</date> with <date>YYYY0101</date>
	// Case insensitive just in case
	re := regexp.MustCompile(`(?i)<date>\s*(\d{4})\s*</date>`)
	fixedBodyStr := re.ReplaceAllString(bodyStr, "<date>${1}0101</date>")

	// 6. Decode XML
	var epg xmltv.EPG
	if err := xml.Unmarshal([]byte(fixedBodyStr), &epg); err != nil {
		return nil, fmt.Errorf("failed to decode XMLTV response: %w", err)
	}

	return &epg, nil
}
