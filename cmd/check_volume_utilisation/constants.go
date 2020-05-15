package main

const name = "check_volume_utilisation"
const helpText = `Usage %s --host=<host> --user=<user> --password=<pwd>|--password-file=<pwdf> [--criticcl=<cpct>]
         [--help] [--timeout=<sec>] [--version] [--volume=<vol>] [--warning=<wpct>]

  --critical=<cpct>         Report critical if used space is larger then <cpct> percent of the volume
                            Default: %.1f%%

  --help                    This text

  --host=<host>             Host to connect to

  --password=<pwd>          Password for authentication

  --password-file=<pwdf>    Read password from password file
                            Note: Only the first line of the file is used

  --user=<user>             Username for authentication

  --timeout=<sec>           HTTP connection timeout in seconds
                            Default: %d sec.

  --version                 Shows version

  --volume=<vol>            Only check volume <vol>
                            Default: check all volumes

  --warning=<wpct>          Report warning if used space is larger than <wpct> percent of the volume
                            Default: %.1f%%
`

const defaultHTTPTimeout = 60
const defaultWarningThreshold = 80.0
const defaultCriticalThreshold = 90.0
