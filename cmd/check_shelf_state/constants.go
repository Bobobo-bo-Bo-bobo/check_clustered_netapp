package main

const name = "check_shelf_state"
const helpText = `Usage %s --host=<host> --user=<user> --password=<pwd>|--password-file=<pwdf>
         [--ca-file=<file>] [--help] [--insecure] [--shelf=<uuid>] [--timeout=<sec>] [--version]

  --ca-file=<file>          Use CA certificate from <file> for validation of SSL certificate

  --help                    This text

  --host=<host>             Host to connect to

  --insecure                Skip validation of SSL server certificate

  --password=<pwd>          Password for authentication

  --password-file=<pwdf>    Read password from password file
                            Note: Only the first line of the file is used

  --shelf=<uuid>            Only check shelf with UUID <uuid> instead of all shelves

  --timeout=<sec>           HTTP connection timeout in seconds
                            Default: %d sec.

  --user=<user>             Username for authentication

  --version                 Shows version

  --volume=<vol>            Only check volume <vol>
                            Default: check all volumes

`

const defaultHTTPTimeout = 60
