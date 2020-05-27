package shared

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GetNetAppSnapmirrorRelationshipList - Get list of disks from NetApp fileserver in clustered mode
func GetNetAppSnapmirrorRelationshipList(host string, fields string, user string, password string, caFile string, insecure bool, timeout time.Duration) (SnapmirrorRelationshipList, error) {
	var snaps SnapmirrorRelationshipList
	var url string

	if fields != "" {
		url = "https://" + host + RestSnapmirrorRelationshipListPath + "?fields=" + fields
	} else {
		url = "https://" + host + RestSnapmirrorRelationshipListPath
	}

	reply, err := HTTPRequest(url, "GET", nil, nil, user, password, caFile, insecure, timeout)
	if err != nil {
		return snaps, err
	}

	if reply.StatusCode != http.StatusOK {
		return snaps, fmt.Errorf("Unexpected HTTP status %s", reply.Status)
	}
	err = json.Unmarshal(reply.Content, &snaps)
	if err != nil {
		return snaps, err
	}
	return snaps, nil
}

// GetNetAppSnapmirrorRelationshipRecord - Get details of volume from NetApp fileserver in clustered mode
func GetNetAppSnapmirrorRelationshipRecord(host string, snp string, fields string, user string, password string, caFile string, insecure bool, timeout time.Duration) (SnapmirrorRelationshipRecord, error) {
	var snap SnapmirrorRelationshipRecord
	var url string

	if fields != "" {
		url = "https://" + host + RestSnapmirrorRelationshipListPath + "/" + snp + "?fields=" + fields
	} else {
		url = "https://" + host + RestSnapmirrorRelationshipListPath + "/" + snp
	}

	reply, err := HTTPRequest(url, "GET", nil, nil, user, password, caFile, insecure, timeout)
	if err != nil {
		return snap, err
	}

	if reply.StatusCode != http.StatusOK {
		return snap, fmt.Errorf("Unexpected HTTP status %s", reply.Status)
	}
	err = json.Unmarshal(reply.Content, &snap)
	if err != nil {
		return snap, err
	}
	return snap, nil
}
