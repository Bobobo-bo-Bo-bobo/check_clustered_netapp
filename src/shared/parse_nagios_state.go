package shared

import (
	"strings"
)

// ParseNagiosState - parse Nagios state and return message and exit code
func ParseNagiosState(n NagiosState) (string, int) {
	var allMsg []string
	var rc = Unknown
	var msg string

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

	if len(n.PerfData) > 0 {
		msg = NagiosMessagePrefixList[rc] + strings.Join(allMsg, "; ") + "|" + strings.Join(n.PerfData, " ")
	} else {
		msg = NagiosMessagePrefixList[rc] + strings.Join(allMsg, "; ")
	}

	return msg, rc
}
