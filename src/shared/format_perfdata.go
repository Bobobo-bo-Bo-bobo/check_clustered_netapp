package shared

import (
	"fmt"
	"math"
	"strings"
)

func floatToString(f float64) string {
	if math.IsNaN(f) {
		return ""
	}
	return fmt.Sprintf("%f", f)
}

// FormatPerfData - Format perfdata string (see https://nagios-plugins.org/doc/guidelines.html#AEN200)
func FormatPerfData(label string, value float64, uom string, warn float64, crit float64, min float64, max float64) string {
	var v = floatToString(value)
	var w = floatToString(warn)
	var c = floatToString(crit)
	var mn = floatToString(min)
	var mx = floatToString(max)

	// label can contain any characters except the equals sign or single quote
	l := strings.Replace(label, "'", "_", -1)
	l = strings.Replace(l, "=", "_", -1)

	arr := [5]string{v + uom, w, c, mn, mx}

	s := strings.Join(arr[:], ";")

	// discard "unfilled" trailing ";"
	return fmt.Sprintf("'%s'=%s", l, strings.TrimRight(s, ";"))
}
