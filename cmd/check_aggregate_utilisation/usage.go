package main

import (
	"fmt"
	"runtime"
	"shared"
)

func showVersion() {
	fmt.Printf("%s version %s\n"+
		"Copyright (C) by Andreas Maus <maus@ypbind.de>\n"+
		"This program comes with ABSOLUTELY NO WARRANTY.\n"+
		"\n"+
		"%s is distributed under the Terms of the GNU General\n"+
		"Public License Version 3. (http://www.gnu.org/copyleft/gpl.html)\n"+
		"\n"+
		"Build with go version: %s\n\n", name, shared.Version, name, runtime.Version())
}

func showUsage() {
	showVersion()
	fmt.Printf(helpText, name, defaultCriticalThreshold, defaultHTTPTimeout, defaultWarningThreshold)
}
