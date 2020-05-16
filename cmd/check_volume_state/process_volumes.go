package main

import (
	"fmt"
	"math"
	"shared"
)

func processVolumes(vl shared.VolumeList) shared.NagiosState {
	var volCount uint64
	var result shared.NagiosState
	var msg string
	var online uint64
	var offline uint64
	var mixed uint64
	var _error uint64
	var unknown uint64

	for _, vol := range vl.Records {
		if vol.State == "" {
			continue
		}

		msg = fmt.Sprintf("Volume %s (name: %s) is in %s state", vol.UUID, vol.Name, vol.State)

		if vol.State == "error" {
			if vol.ErrorState != nil {
				msg += fmt.Sprintf(" (bad blocks: %t, inconsistent: %t)", vol.ErrorState.HasBadBlocks, vol.ErrorState.IsInconsitent)
			} else {
				msg += " (no details found)"
			}
			result.Critical = append(result.Critical, msg)
			_error++
		} else if vol.State == "offline" {
			if vol.ErrorState != nil {
				msg = fmt.Sprintf(" (bad blocks: %t, inconsistent: %t)", vol.ErrorState.HasBadBlocks, vol.ErrorState.IsInconsitent)
			} else {
				msg = " (no details found)"
			}
			result.Warning = append(result.Warning, msg)
			offline++
		} else if vol.State == "online" {
			result.OK = append(result.OK, msg)
			online++
		} else if vol.State == "mixed" {
			result.OK = append(result.OK, msg)
			mixed++
		} else {
			msg = fmt.Sprintf("Volume %s (name: %s) is in unknown state (reports: %s)", vol.UUID, vol.Name, vol.State)
			result.Unknown = append(result.Unknown, msg)
			unknown++
		}

		// add performance data
		result.PerfData = append(result.PerfData, shared.FormatPerfData("volume_health_online", float64(online), "", math.NaN(), math.NaN(), 0.0, math.NaN()))
		result.PerfData = append(result.PerfData, shared.FormatPerfData("volume_health_mixed", float64(mixed), "", math.NaN(), math.NaN(), 0.0, math.NaN()))
		result.PerfData = append(result.PerfData, shared.FormatPerfData("volume_health_error", float64(_error), "", math.NaN(), math.NaN(), 0.0, math.NaN()))
		result.PerfData = append(result.PerfData, shared.FormatPerfData("volume_health_offline", float64(offline), "", math.NaN(), math.NaN(), 0.0, math.NaN()))

		volCount++
	}
	if volCount == 0 {
		result.Critical = append(result.Critical, "No volumes found")
	}

	return result
}
