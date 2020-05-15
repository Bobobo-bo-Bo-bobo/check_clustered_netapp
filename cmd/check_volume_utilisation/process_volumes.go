package main

import (
	"fmt"
	"shared"
)

func processVolumes(vl shared.VolumeList, warn float64, critical float64) shared.NagiosState {
	var volCount uint64
	var result shared.NagiosState
	var pct float64
	var msg string

	for _, vol := range vl.Records {
		if vol.Space == nil {
			continue
		}

		pct = 100.0 * float64(vol.Space.Used) / float64(vol.Space.Size)
		msg = fmt.Sprintf("Usage of volume %s (name: %s) is %.2f%%", vol.UUID, vol.Name, pct)

		if pct >= critical {
			result.Critical = append(result.Critical, msg)
		} else if pct >= warn {
			result.Warning = append(result.Warning, msg)
		} else {
			result.OK = append(result.OK, msg)
		}

		volCount++
	}
	if volCount == 0 {
		result.Critical = append(result.Critical, "No volumes found")
	}

	return result
}
