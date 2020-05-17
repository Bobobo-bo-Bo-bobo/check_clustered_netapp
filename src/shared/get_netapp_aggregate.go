package shared

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GetNetAppAggregateList - Get list of aggregates from NetApp fileserver in clustered mode
func GetNetAppAggregateList(host string, fields string, user string, password string, caFile string, insecure bool, timeout time.Duration) (AggregateList, error) {
	var aggrs AggregateList
	var url string

	if fields != "" {
		url = "https://" + host + RestAggregateListPath + "?fields=" + fields
	} else {
		url = "https://" + host + RestAggregateListPath
	}

	reply, err := HTTPRequest(url, "GET", nil, nil, user, password, caFile, insecure, timeout)
	if err != nil {
		return aggrs, err
	}

	if reply.StatusCode != http.StatusOK {
		return aggrs, fmt.Errorf("Unexpected HTTP status %s", reply.Status)
	}
	err = json.Unmarshal(reply.Content, &aggrs)
	if err != nil {
		return aggrs, err
	}
	return aggrs, nil
}

// GetNetAppAggregateRecord - Get details of volume from NetApp fileserver in clustered mode
func GetNetAppAggregateRecord(host string, aggr string, fields string, user string, password string, caFile string, insecure bool, timeout time.Duration) (AggregateRecord, error) {
	var aggrr AggregateRecord
	var url string

	if fields != "" {
		url = "https://" + host + RestAggregateListPath + "/" + aggr + "?fields=" + fields
	} else {
		url = "https://" + host + RestAggregateListPath + "/" + aggr
	}

	reply, err := HTTPRequest(url, "GET", nil, nil, user, password, caFile, insecure, timeout)
	if err != nil {
		return aggrr, err
	}

	if reply.StatusCode != http.StatusOK {
		return aggrr, fmt.Errorf("Unexpected HTTP status %s", reply.Status)
	}
	err = json.Unmarshal(reply.Content, &aggrr)
	if err != nil {
		return aggrr, err
	}
	return aggrr, nil
}
