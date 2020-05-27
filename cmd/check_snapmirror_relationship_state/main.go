package main

import (
	"flag"
	"fmt"
	"os"
	"shared"
	"time"
)

func main() {
	var help = flag.Bool("help", false, "This text")
	var host = flag.String("host", "", "Host to connect to")
	var password = flag.String("password", "", "Password for authentication")
	var passwordFile = flag.String("password-file", "", "Read password from password file")
	var username = flag.String("user", "", "Username for authentication")
	var version = flag.Bool("version", false, "Show version")
	var timeout = flag.Int64("timeout", defaultHTTPTimeout, "HTTP connection timeout")
	var httpTimeout time.Duration
	var cafile = flag.String("ca-file", "", "Use CA certificate for validation of server SSL from <cafile>")
	var insecure = flag.Bool("insecure", false, "Skip SSL verificattion")
	var snapList shared.SnapmirrorRelationshipList
	var err error

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

	snapList, err = shared.GetNetAppSnapmirrorRelationshipList(*host, "vendor,model,serial_number,name,bay,drawer,state", *username, *password, *cafile, *insecure, httpTimeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't get disk information: %s\n", err)
		os.Exit(shared.Unknown)
	}
	state := processSnapmirrorRelationships(snapList)
	message, ret := shared.ParseNagiosState(state)
	fmt.Println(message)
	os.Exit(ret)
}
