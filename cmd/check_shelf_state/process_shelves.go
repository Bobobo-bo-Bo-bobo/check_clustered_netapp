package main

import (
	"fmt"
	"shared"
)

func processShelves(sl shared.ShelfList) shared.NagiosState {
	var shlCount uint64
	var result shared.NagiosState
	var msg string

	for _, shl := range sl.Records {
		msg = fmt.Sprintf("%s shelf %s (name: %s, serial number: %s) is in %s state", shl.Model, shl.UUID, shl.Name, shl.SerialNumber, shl.State)

		if shl.State == "ok" {
			result.OK = append(result.OK, msg)
		} else {
			// XXX: Is there a documentation of known states (except "ok") ?
			result.Critical = append(result.Critical, msg)
		}

		shlCount++
	}
	if shlCount == 0 {
		result.Critical = append(result.Critical, "No shelves found")
	}

	return result
}
