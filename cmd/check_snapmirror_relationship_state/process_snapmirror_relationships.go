package main

import (
	"fmt"
	"math"
	"shared"
	"strings"
)

func processSnapmirrorRelationships(sl shared.SnapmirrorRelationshipList) shared.NagiosState {
	var snpCount uint64
	var result shared.NagiosState
	var msg string
	var ok uint64
	var nok uint64
	var unhr []string

	for _, snp := range sl.Records {
		if snp.Healthy {
			msg = fmt.Sprintf("Snapmirror relationship with UUID %s from cluster %s, path %s on SVM %s is healthy", snp.UUID, snp.Source.Cluster.Name, snp.Source.Path, snp.Source.SVM.Name)
			result.OK = append(result.OK, msg)
			ok++
		} else {
			for _, s := range snp.UnhealthyReason {
				unhr = append(unhr, fmt.Sprintf("%s (code %s)", s.Message, s.Code))
			}

			msg = fmt.Sprintf("Snapmirror relationship with UUID %s from cluster %s, path %s on SVM %s is unhealthy (%s)", snp.UUID, snp.Source.Cluster.Name, snp.Source.Path, snp.Source.SVM.Name, strings.Join(unhr, ", "))
			result.Critical = append(result.Critical, msg)
			nok++
		}

		snpCount++
	}
	if snpCount == 0 {
		result.Critical = append(result.Critical, "No snapmirror relations found")
	}

	// add performance data
	result.PerfData = append(result.PerfData, shared.FormatPerfData("disk_health_healthy", float64(ok), "", math.NaN(), math.NaN(), 0.0, float64(snpCount)))
	result.PerfData = append(result.PerfData, shared.FormatPerfData("disk_health_unhealthy", float64(nok), "", math.NaN(), math.NaN(), 0.0, float64(snpCount)))
	return result
}
