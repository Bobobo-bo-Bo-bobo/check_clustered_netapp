package main

import (
	"fmt"
	"math"
	"shared"
)

func processDisks(dl shared.DiskList) shared.NagiosState {
	var dskCount uint64
	var result shared.NagiosState
	var msg string
	var ok uint64
	var brk uint64
	var cpy uint64
	var mnt uint64
	var prt uint64
	var prs uint64
	var rcs uint64
	var rmv uint64
	var spr uint64
	var unf uint64
	var zrn uint64

	for _, dsk := range dl.Records {
		if dsk.State == "" {
			continue
		}

		msg = fmt.Sprintf("%s %s disk %s (serial number: %s) in drawer %d, bay %d is %s", dsk.Vendor, dsk.Model, dsk.Name, dsk.SerialNumber, dsk.Drawer.ID, dsk.Drawer.Slot, dsk.State)

		switch dsk.State {
		case "broken":
			result.Critical = append(result.Critical, msg)
			brk++
		case "copy":
			ok++
			cpy++
		case "maintenance":
			result.Warning = append(result.Warning, msg)
			mnt++
		case "partner":
			ok++
			prt++
		case "present":
			ok++
			prs++
		case "reconstructing":
			ok++
			rcs++
		case "removed":
			result.Warning = append(result.Warning, msg)
			rmv++
		case "spare":
			ok++
			spr++
		case "unfail":
			ok++
			unf++
		case "zeroing":
			result.Warning = append(result.Warning, msg)
			zrn++
		default:
		}

		dskCount++
	}
	if dskCount == 0 {
		result.Critical = append(result.Critical, "No disks found")
	}

	if ok > 0 {
		result.OK = append(result.OK, fmt.Sprintf("%d disks are ok (%d present, %d copying, %d partner, %d reconstructing, %d spare, %d unfail)", ok, prs, cpy, prt, rcs, spr, unf))
	}

	// add performance data
	result.PerfData = append(result.PerfData, shared.FormatPerfData("disk_health_ok", float64(ok), "", math.NaN(), math.NaN(), 0.0, float64(dskCount)))
	result.PerfData = append(result.PerfData, shared.FormatPerfData("disk_health_broken", float64(brk), "", math.NaN(), math.NaN(), 0.0, float64(dskCount)))
	result.PerfData = append(result.PerfData, shared.FormatPerfData("disk_health_copy", float64(cpy), "", math.NaN(), math.NaN(), 0.0, float64(dskCount)))
	result.PerfData = append(result.PerfData, shared.FormatPerfData("disk_health_maintenance", float64(mnt), "", math.NaN(), math.NaN(), 0.0, float64(dskCount)))
	result.PerfData = append(result.PerfData, shared.FormatPerfData("disk_health_partner", float64(prt), "", math.NaN(), math.NaN(), 0.0, float64(dskCount)))
	result.PerfData = append(result.PerfData, shared.FormatPerfData("disk_health_present", float64(prs), "", math.NaN(), math.NaN(), 0.0, float64(dskCount)))
	result.PerfData = append(result.PerfData, shared.FormatPerfData("disk_health_reconstructing", float64(rcs), "", math.NaN(), math.NaN(), 0.0, float64(dskCount)))
	result.PerfData = append(result.PerfData, shared.FormatPerfData("disk_health_removed", float64(rmv), "", math.NaN(), math.NaN(), 0.0, float64(dskCount)))
	result.PerfData = append(result.PerfData, shared.FormatPerfData("disk_health_spare", float64(spr), "", math.NaN(), math.NaN(), 0.0, float64(dskCount)))
	result.PerfData = append(result.PerfData, shared.FormatPerfData("disk_health_unfail", float64(unf), "", math.NaN(), math.NaN(), 0.0, float64(dskCount)))
	result.PerfData = append(result.PerfData, shared.FormatPerfData("disk_health_zeroing", float64(zrn), "", math.NaN(), math.NaN(), 0.0, float64(dskCount)))
	return result
}
