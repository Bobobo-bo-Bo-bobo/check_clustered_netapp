package shared

import (
	"fmt"
)

// Version - package version
const Version = "1.0.0-20200515"
const name = "check_clustered_netapp"
const url = "https://git.ypbind.de/cgit/check_clustered_netapp"

// UserAgent - user agent for HTTP connection
var UserAgent = fmt.Sprintf("%s/%s (%s)", name, Version, url)

const (
	// Ok - Nagios state is OK
	Ok int = iota
	// Warning - Nagios state is warning
	Warning
	// Critical - Nagios state is critical
	Critical
	// Unknown - Nagios state is unknown
	Unknown
)

// RestVolumeListPath - URL path to volumes
const RestVolumeListPath = "/api/storage/volumes"
