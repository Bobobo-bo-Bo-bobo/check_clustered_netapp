package shared

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GetNetAppDiskList - Get list of disks from NetApp fileserver in clustered mode
func GetNetAppDiskList(host string, fields string, user string, password string, caFile string, insecure bool, timeout time.Duration) (DiskList, error) {
	var disks DiskList
	var url string

	if fields != "" {
		url = "https://" + host + RestDiskListPath + "?fields=" + fields
	} else {
		url = "https://" + host + RestDiskListPath
	}

	reply, err := HTTPRequest(url, "GET", nil, nil, user, password, caFile, insecure, timeout)
	if err != nil {
		return disks, err
	}

	if reply.StatusCode != http.StatusOK {
		return disks, fmt.Errorf("Unexpected HTTP status %s", reply.Status)
	}
	err = json.Unmarshal(reply.Content, &disks)
	if err != nil {
		return disks, err
	}
	return disks, nil
}

// GetNetAppDiskRecord - Get details of volume from NetApp fileserver in clustered mode
func GetNetAppDiskRecord(host string, dsk string, fields string, user string, password string, caFile string, insecure bool, timeout time.Duration) (DiskRecord, error) {
	var dskr DiskRecord
	var url string

	if fields != "" {
		url = "https://" + host + RestDiskListPath + "/" + dsk + "?fields=" + fields
	} else {
		url = "https://" + host + RestDiskListPath + "/" + dsk
	}

	reply, err := HTTPRequest(url, "GET", nil, nil, user, password, caFile, insecure, timeout)
	if err != nil {
		return dskr, err
	}

	if reply.StatusCode != http.StatusOK {
		return dskr, fmt.Errorf("Unexpected HTTP status %s", reply.Status)
	}
	err = json.Unmarshal(reply.Content, &dskr)
	if err != nil {
		return dskr, err
	}
	return dskr, nil
}
