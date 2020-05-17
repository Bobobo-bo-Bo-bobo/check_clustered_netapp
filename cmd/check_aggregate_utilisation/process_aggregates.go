package main

import (
	"fmt"
	"math"
	"shared"
)

func processAggregates(al shared.AggregateList, warn float64, critical float64) shared.NagiosState {
	var aggrCount uint64
	var result shared.NagiosState
	var pct float64
	var uw float64
	var uc float64
	var fw float64
	var fc float64
	var msg string

	for _, aggr := range al.Records {
		if aggr.Space == nil {
			continue
		}

		pct = 100.0 * float64(aggr.Space.BlockStorage.Used) / float64(aggr.Space.BlockStorage.Size)
		uw = warn * float64(aggr.Space.BlockStorage.Used) / 100.0
		uc = critical * float64(aggr.Space.BlockStorage.Used) / 100.0
		fw = float64(aggr.Space.BlockStorage.Size) - uw
		fc = float64(aggr.Space.BlockStorage.Size) - uc

		msg = fmt.Sprintf("Usage of aggregate %s (name: %s) is %.2f%%", aggr.UUID, aggr.Name, pct)

		if pct >= critical {
			result.Critical = append(result.Critical, msg)
		} else if pct >= warn {
			result.Warning = append(result.Warning, msg)
		} else {
			result.OK = append(result.OK, msg)
		}

		// add performance data
		result.PerfData = append(result.PerfData, shared.FormatPerfData(aggr.UUID+"_used", float64(aggr.Space.BlockStorage.Used), "B", uw, uc, 0.0, float64(aggr.Space.BlockStorage.Size)))
		result.PerfData = append(result.PerfData, shared.FormatPerfData(aggr.UUID+"_free", float64(aggr.Space.BlockStorage.Available), "B", fw, fc, 0.0, float64(aggr.Space.BlockStorage.Size)))
		result.PerfData = append(result.PerfData, shared.FormatPerfData(aggr.UUID+"_pct_used", pct, "%", warn, critical, math.NaN(), math.NaN()))

		aggrCount++
	}
	if aggrCount == 0 {
		result.Critical = append(result.Critical, "No aggregates found")
	}

	return result
}
