package main

import (
	"flag"
	"fmt"
	"os"
	"shared"
	"time"
)

func main() {
	var critical = flag.Float64("critical", defaultCriticalThreshold, "Report critical if used space is larger then <cpct> percent of the volume")
	var warning = flag.Float64("warning", defaultWarningThreshold, "Report warning if used space is larger than <wpct> percent of the volume")
	var help = flag.Bool("help", false, "This text")
	var host = flag.String("host", "", "Host to connect to")
	var password = flag.String("password", "", "Password for authentication")
	var passwordFile = flag.String("password-file", "", "Read password from password file")
	var username = flag.String("user", "", "Username for authentication")
	var version = flag.Bool("version", false, "Show version")
	var timeout = flag.Int64("timeout", defaultHTTPTimeout, "HTTP connection timeout")
	var httpTimeout time.Duration
	var volume = flag.String("volume", "", "Only check this particular volume")
	var cafile = flag.String("ca-file", "", "Use CA certificate for validation of server SSL from <cafile>")
	var insecure = flag.Bool("insecure", false, "Skip SSL verificattion")

	flag.Usage = showUsage
	flag.Parse()

	if *help {
		showUsage()
		os.Exit(0)
	}

	if *version {
		showVersion()
		os.Exit(0)
	}

	trailing := flag.Args()
	if len(trailing) > 0 {
		fmt.Fprintln(os.Stderr, "Too many arguments")
		fmt.Fprintln(os.Stderr, "")
		showUsage()
		os.Exit(shared.Unknown)
	}

	// check for mandatory options
	if *host == "" {
		fmt.Fprintln(os.Stderr, "Missing host parameter")
		fmt.Fprintln(os.Stderr, "")
		showUsage()
		os.Exit(shared.Unknown)
	}
	if *username == "" {
		fmt.Fprintln(os.Stderr, "Missing username parameter")
		fmt.Fprintln(os.Stderr, "")
		showUsage()
		os.Exit(shared.Unknown)
	}

	// sanity checks
	if *warning < 0.0 || *critical < 0.0 {
		fmt.Fprintln(os.Stderr, "Warning and critical thresholds must be greater than 0")
		fmt.Fprintln(os.Stderr, "")
		showUsage()
		os.Exit(shared.Unknown)
	}
	if *warning > 100.0 || *critical > 100.0 {
		fmt.Fprintln(os.Stderr, "Warning and critical thresholds must be less or equal than 100")
		fmt.Fprintln(os.Stderr, "")
		showUsage()
		os.Exit(shared.Unknown)
	}
	if *warning > *critical {
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Warning threshold must be less or equal than critical threshold")
		showUsage()
		os.Exit(shared.Unknown)
	}
	if *passwordFile != "" && *password != "" {
		fmt.Fprintln(os.Stderr, "Password and password file options are mutually exclusive")
		fmt.Fprintln(os.Stderr, "")
		showUsage()
		os.Exit(shared.Unknown)
	}
	if *timeout <= 0 {
		fmt.Fprintln(os.Stderr, "Connection timeout must be greater than 0")
		fmt.Fprintln(os.Stderr, "")
		showUsage()
		os.Exit(shared.Unknown)
	}

	// process provided options
	httpTimeout = time.Duration(*timeout) * time.Second

	if *passwordFile != "" {
		_pwd, err := shared.ReadSingleLine(*passwordFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't read password from %s: %s\n", *passwordFile, err)
			os.Exit(shared.Unknown)
		}
		password = &_pwd
	}

	if *cafile != "" {
		_, err := shared.ReadFileContent(*cafile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't read data from CA file %s: %s\n", *cafile, err)
			os.Exit(shared.Unknown)
		}
	}

    if *volume != "" {
    } else {
        volumeList, err := shared.GetNetAppVolumeList(*host, "space", *username, *password, *cafile, *insecure, httpTimeout)
        if err != nil {
            os.Exit(shared.Unknown)
        }
        fmt.Printf("%+v\n", volumeList)
    }
}
