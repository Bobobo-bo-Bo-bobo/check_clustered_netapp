package shared

import (
	"strings"
)

// ParseNagiosState - parse Nagios state and return message and exit code
func ParseNagiosState(n NagiosState) (string, int) {
	var allMsg []string
	var rc = 99

	allMsg = append(allMsg, n.Unknown...)
	allMsg = append(allMsg, n.Critical...)
	allMsg = append(allMsg, n.Warning...)
	allMsg = append(allMsg, n.OK...)

	if len(n.Unknown) > 0 {
		rc = Unknown
	} else if len(n.Critical) > 0 {
		rc = Critical
	} else if len(n.Warning) > 0 {
		rc = Warning
	} else {
		rc = Ok
	}

	return strings.Join(allMsg, "; "), rc
}
