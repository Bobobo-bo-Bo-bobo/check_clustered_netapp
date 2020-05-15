package main

import (
	"fmt"
	"math"
	"shared"
)

func processVolumes(vl shared.VolumeList, warn float64, critical float64) shared.NagiosState {
	var volCount uint64
	var result shared.NagiosState
	var pct float64
	var uw float64
	var uc float64
	var fw float64
	var fc float64
	var msg string

	for _, vol := range vl.Records {
		if vol.Space == nil {
			continue
		}

		pct = 100.0 * float64(vol.Space.Used) / float64(vol.Space.Size)
		uw = warn * float64(vol.Space.Used) / 100.0
		uc = critical * float64(vol.Space.Used) / 100.0
		fw = float64(vol.Space.Size) - uw
		fc = float64(vol.Space.Size) - uc

		msg = fmt.Sprintf("Usage of volume %s (name: %s) is %.2f%%", vol.UUID, vol.Name, pct)

		if pct >= critical {
			result.Critical = append(result.Critical, msg)
		} else if pct >= warn {
			result.Warning = append(result.Warning, msg)
		} else {
			result.OK = append(result.OK, msg)
		}

		// add performance data
		result.PerfData = append(result.PerfData, shared.FormatPerfData(vol.UUID+"_used", float64(vol.Space.Used), "B", uw, uc, 0.0, float64(vol.Space.Size)))
		result.PerfData = append(result.PerfData, shared.FormatPerfData(vol.UUID+"_free", float64(vol.Space.Available), "B", fw, fc, 0.0, float64(vol.Space.Size)))
		result.PerfData = append(result.PerfData, shared.FormatPerfData(vol.UUID+"_pct_used", pct, "%", warn, critical, math.NaN(), math.NaN()))

		volCount++
	}
	if volCount == 0 {
		result.Critical = append(result.Critical, "No volumes found")
	}

	return result
}
