package shared

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GetNetAppVolumeList - Get list of volumes from NetApp fileserver in clustered mode
func GetNetAppVolumeList(host string, fields string, user string, password string, caFile string, insecure bool, timeout time.Duration) (VolumeList, error) {
	var vols VolumeList
	var url string

	if fields != "" {
		url = "https://" + host + RestVolumeListPath + "?fields=" + fields
	} else {
		url = "https://" + host + RestVolumeListPath
	}

	reply, err := HTTPRequest(url, "GET", nil, nil, user, password, caFile, insecure, timeout)
	if err != nil {
		return vols, err
	}

	if reply.StatusCode != http.StatusOK {
		return vols, fmt.Errorf("Unexpected HTTP status %s", reply.Status)
	}
	err = json.Unmarshal(reply.Content, &vols)
	if err != nil {
		return vols, err
	}
	return vols, nil
}
