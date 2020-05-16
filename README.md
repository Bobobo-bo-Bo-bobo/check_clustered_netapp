# Preface
Starting with ONTAP 9.6 NetApp file server provide a REST API (in addition to the classic XML API used by the NetApp SDK).
The general documentation can be found at the [ONTAP 9 documentation center](http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html) and the endpoint documentation can be found at [ONTAP REST API documentation](https://library.netapp.com/ecmdocs/ECMLP2856304/html/index.html).

# Nagios checks
## Check volume utilisation
`check_volume_utilisation`: check the utilisation of a single volume or all volumes

| *Option* | *Description* | *Default* | *Note* |
|:---------|:--------------|:---------:|:------:|
| `--ca-file=<file>` | Use CA certificate from <file> for validation of SSL certificate | - |
| `--critical=<pct>` | Report CRITICAL if used space is larger then <cpct> percent of the volume | 90.0% | - |
| `--help` | Show help text | - | - |
| `--insecure` | Skip validation of SSL server certificate | - | - |
| `--host=<host>` | Host to connect to | - | **mandatory** |
| `--password=<pass>` | Password for authentication | - | - |
| `--password-file=<passf>` | Read password for authentication from file <passf> | - | Only the first line in <passf> is used as password |
| `--timeout=<sec>` | HTTP connection timeout in seconds | 60 | - |
| `--user=<user>` | Username for authentication | - | **mandatory** |
| `--version` | Show version information | - | - |
| `--volume=<uuid>` | Only check volume with UUID <uuid> | - | If omitted all volumes will be checked |
| `--warning=<pct>` | Report WARNING if used space is larger then <cpct> percent of the volume | 80.0% | - |

## Check volume health
`check_volume_state`: check the health of a single volume or all volumes

| *Option* | *Description* | *Default* | *Note* |
|:---------|:--------------|:---------:|:------:|
| `--ca-file=<file>` | Use CA certificate from <file> for validation of SSL certificate | - |
| `--help` | Show help text | - | - |
| `--insecure` | Skip validation of SSL server certificate | - | - |
| `--host=<host>` | Host to connect to | - | **mandatory** |
| `--password=<pass>` | Password for authentication | - | - |
| `--password-file=<passf>` | Read password for authentication from file <passf> | - | Only the first line in <passf> is used as password |
| `--timeout=<sec>` | HTTP connection timeout in seconds | 60 | - |
| `--user=<user>` | Username for authentication | - | **mandatory** |
| `--version` | Show version information | - | - |
| `--volume=<uuid>` | Only check volume with UUID <uuid> | - | If omitted all volumes will be checked |

## Check shelf health
`check_shelf_health`: check health of a shelf or all attached shelves

| *Option* | *Description* | *Default* | *Note* |
|:---------|:--------------|:---------:|:------:|
| `--ca-file=<file>` | Use CA certificate from <file> for validation of SSL certificate | - |
| `--help` | Show help text | - | - |
| `--insecure` | Skip validation of SSL server certificate | - | - |
| `--host=<host>` | Host to connect to | - | **mandatory** |
| `--password=<pass>` | Password for authentication | - | - |
| `--password-file=<passf>` | Read password for authentication from file <passf> | - | Only the first line in <passf> is used as password |
| `--shelf=<uuid>` | Only check shelf with UUID <uuid> | - | If omitted all shelves will be checked |
| `--timeout=<sec>` | HTTP connection timeout in seconds | 60 | - |
| `--user=<user>` | Username for authentication | - | **mandatory** |
| `--version` | Show version information | - | - |

## Check disk health
`check_disk_health`: check health of disks

| *Option* | *Description* | *Default* | *Note* |
|:---------|:--------------|:---------:|:------:|
| `--ca-file=<file>` | Use CA certificate from <file> for validation of SSL certificate | - |
| `--help` | Show help text | - | - |
| `--insecure` | Skip validation of SSL server certificate | - | - |
| `--host=<host>` | Host to connect to | - | **mandatory** |
| `--password=<pass>` | Password for authentication | - | - |
| `--password-file=<passf>` | Read password for authentication from file <passf> | - | Only the first line in <passf> is used as password |
| `--timeout=<sec>` | HTTP connection timeout in seconds | 60 | - |
| `--user=<user>` | Username for authentication | - | **mandatory** |
| `--version` | Show version information | - | - |


