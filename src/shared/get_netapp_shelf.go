package shared

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GetNetAppShelfList - Get list of shelves from NetApp fileserver in clustered mode
func GetNetAppShelfList(host string, fields string, user string, password string, caFile string, insecure bool, timeout time.Duration) (ShelfList, error) {
	var shlvs ShelfList
	var url string

	if fields != "" {
		url = "https://" + host + RestShelfListPath + "?fields=" + fields
	} else {
		url = "https://" + host + RestShelfListPath
	}

	reply, err := HTTPRequest(url, "GET", nil, nil, user, password, caFile, insecure, timeout)
	if err != nil {
		return shlvs, err
	}

	if reply.StatusCode != http.StatusOK {
		return shlvs, fmt.Errorf("Unexpected HTTP status %s", reply.Status)
	}
	err = json.Unmarshal(reply.Content, &shlvs)
	if err != nil {
		return shlvs, err
	}
	return shlvs, nil
}

// GetNetAppShelfRecord - Get details of volume from NetApp fileserver in clustered mode
func GetNetAppShelfRecord(host string, shl string, fields string, user string, password string, caFile string, insecure bool, timeout time.Duration) (ShelfRecord, error) {
	var shlr ShelfRecord
	var url string

	if fields != "" {
		url = "https://" + host + RestShelfListPath + "/" + shl + "?fields=" + fields
	} else {
		url = "https://" + host + RestShelfListPath + "/" + shl
	}

	reply, err := HTTPRequest(url, "GET", nil, nil, user, password, caFile, insecure, timeout)
	if err != nil {
		return shlr, err
	}

	if reply.StatusCode != http.StatusOK {
		return shlr, fmt.Errorf("Unexpected HTTP status %s", reply.Status)
	}
	err = json.Unmarshal(reply.Content, &shlr)
	if err != nil {
		return shlr, err
	}
	return shlr, nil
}
