package utils

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// HeaderValues represents a slice of header values
type HeaderValues []string

// Contains checks if a value exists in the header values
func (vs HeaderValues) Contains(s string) bool {
	for _, v := range vs {
		if v == s {
			return true
		}
	}
	return false
}

// HeaderFilter is a function that determines whether a header should be included
type HeaderFilter func(key string) bool

// DefaultHeaderFilter excludes common problematic headers
func DefaultHeaderFilter(key string) bool {
	return key != "Host" && key != "Connection" && key != "User-Agent" && key != "Accept" && key != "Accept-Encoding" && key != "Accept-Language" && key != "Cache-Control" && key != "Pragma" && key != "If-Modified-Since" && key != "If-None-Match"
}

// ExcludeHeaders creates a HeaderFilter that excludes the specified header names
func ExcludeHeaders(headers ...string) HeaderFilter {
	excludeMap := make(map[string]bool)
	for _, h := range headers {
		excludeMap[h] = true
	}
	return func(key string) bool {
		return !excludeMap[key]
	}
}

// MergeHttpHeader merges HTTP headers from source to destination
// If a header value already exists in the destination, it will not be duplicated
// Multiple filter functions can be provided - a header must pass all filters to be included
// Example usage:
//
//	MergeHttpHeader(dst, src, DefaultHeaderFilter)
//	MergeHttpHeader(dst, src, ExcludeHeaders("Host", "Connection"))
func MergeHttpHeader(dst, src http.Header, filters ...HeaderFilter) {
	if os.Getenv("MERGE_HEADERS") != "true" {
		return
	}

	for k, vv := range src {
		// Check all filters - header must pass all of them
		shouldInclude := true
		for _, filter := range filters {
			if filter != nil && !filter(k) {
				shouldInclude = false
				break
			}
		}
		if !shouldInclude {
			continue
		}

		for _, v := range vv {
			if HeaderValues(dst.Values(k)).Contains(v) {
				continue
			}
			dst.Add(k, v)
		}
	}
}

// LogHeaders formats HTTP request/response headers if DEBUG_HEADERS environment variable is set to true
func LogHeaders(prefix string, headers http.Header) string {
	if os.Getenv("DEBUG_HEADERS") != "true" {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\n=== %s Headers ===\n", prefix))
	for key, values := range headers {
		for _, value := range values {
			sb.WriteString(fmt.Sprintf("%s: %s\n", key, value))
		}
	}
	sb.WriteString("==================\n\n")
	return sb.String()
}
