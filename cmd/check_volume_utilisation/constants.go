package main

const name = "check_volume_utilisation"
const helpText = `Usage %s --host=<host> --user=<user> --password=<pwd>|--password-file=<pwdf> [--ca-file=<file>]
         [--criticcl=<cpct>] [--help] [--insecure] [--timeout=<sec>] [--version] [--volume=<vol>] [--warning=<wpct>]

  --ca-file=<file>          Use CA certificate from <file> for validation of SSL certificate

  --critical=<cpct>         Report critical if used space is larger then <cpct> percent of the volume
                            Default: %.1f%%

  --help                    This text

  --host=<host>             Host to connect to

  --insecure                Skip validation of SSL server certificate

  --password=<pwd>          Password for authentication

  --password-file=<pwdf>    Read password from password file
                            Note: Only the first line of the file is used

  --timeout=<sec>           HTTP connection timeout in seconds
                            Default: %d sec.

  --user=<user>             Username for authentication

  --version                 Shows version

  --volume=<vol>            Only check volume <vol>
                            Default: check all volumes

  --warning=<wpct>          Report warning if used space is larger than <wpct> percent of the volume
                            Default: %.1f%%
`

const defaultHTTPTimeout = 60
const defaultWarningThreshold = 80.0
const defaultCriticalThreshold = 90.0
