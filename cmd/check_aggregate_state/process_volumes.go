package main

import (
	"fmt"
	"math"
	"shared"
)

func processAggregates(al shared.AggregateList) shared.NagiosState {
	var aggrCount uint64
	var result shared.NagiosState
	var msg string
	var online uint64
	var offline uint64
	var relocate uint64
	var inconsistent uint64
	var failed uint64
	var unmounted uint64
	var unknown uint64

	for _, aggr := range al.Records {
		if aggr.State == "" {
			continue
		}

		msg = fmt.Sprintf("Aggregate %s (name: %s) is in %s state", aggr.UUID, aggr.Name, aggr.State)

		switch aggr.State {
		case "online":
			result.OK = append(result.OK, msg)
			online++
		case "onlining":
			result.OK = append(result.OK, msg)
			online++
		case "restricted":
			result.OK = append(result.OK, msg)
			online++
		case "offline":
			result.Critical = append(result.Critical, msg)
			offline++
		case "offlining":
			result.Critical = append(result.Critical, msg)
			offline++
		case "relocating":
			result.Warning = append(result.Warning, msg)
			relocate++
		case "unmounted":
			result.Warning = append(result.Warning, msg)
			unmounted++
		case "inconsistent":
			result.Critical = append(result.Critical, msg)
			inconsistent++
		case "failed":
			result.Critical = append(result.Critical, msg)
			failed++
		case "unknown":
			result.Unknown = append(result.Unknown, msg)
			unknown++
		default:
			msg = fmt.Sprintf("Aggregate %s (name: %s) is in unknown state (reports: %s)", aggr.UUID, aggr.Name, aggr.State)
			result.Unknown = append(result.Unknown, msg)
			unknown++
		}

		// add performance data
		result.PerfData = append(result.PerfData, shared.FormatPerfData("aggregate_health_online", float64(online), "", math.NaN(), math.NaN(), 0.0, math.NaN()))
		result.PerfData = append(result.PerfData, shared.FormatPerfData("aggregate_health_offline", float64(offline), "", math.NaN(), math.NaN(), 0.0, math.NaN()))
		result.PerfData = append(result.PerfData, shared.FormatPerfData("aggregate_health_relocate", float64(relocate), "", math.NaN(), math.NaN(), 0.0, math.NaN()))
		result.PerfData = append(result.PerfData, shared.FormatPerfData("aggregate_health_unmounted", float64(unmounted), "", math.NaN(), math.NaN(), 0.0, math.NaN()))
		result.PerfData = append(result.PerfData, shared.FormatPerfData("aggregate_health_inconsistent", float64(inconsistent), "", math.NaN(), math.NaN(), 0.0, math.NaN()))
		result.PerfData = append(result.PerfData, shared.FormatPerfData("aggregate_health_unknown", float64(unknown), "", math.NaN(), math.NaN(), 0.0, math.NaN()))

		aggrCount++
	}
	if aggrCount == 0 {
		result.Critical = append(result.Critical, "No aggregates found")
	}

	return result
}
