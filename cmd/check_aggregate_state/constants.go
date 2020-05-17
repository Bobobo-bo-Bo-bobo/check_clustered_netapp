package main

const name = "check_aggregate_state"
const helpText = `Usage %s --host=<host> --user=<user> --password=<pwd>|--password-file=<pwdf>
         [--aggregate=<aggr>] [--ca-file=<file>] [--help] [--insecure] [--timeout=<sec>] [--version]

  --aggregate=<aggr>        Only check aggregate with UUID <aggr>
                            Default: check all aggregates

  --ca-file=<file>          Use CA certificate from <file> for validation of SSL certificate

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

`

const defaultHTTPTimeout = 60
