package shared

import (
	"net/http"
)

// HTTPResult - result of the http_request calls
type HTTPResult struct {
	URL        string
	StatusCode int
	Status     string
	Header     http.Header
	Content    []byte
}

// NagiosState - Nagios states
type NagiosState struct {
	OK       []string
	Warning  []string
	Critical []string
	Unknown  []string
	PerfData []string
}
