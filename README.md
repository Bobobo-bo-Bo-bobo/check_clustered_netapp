# Preface
Starting with ONTAP 9.6 NetApp file server provide a REST API (in addition to the classic XML API used by the NetApp SDK).
The general documentation can be found at the [ONTAP 9 documentation center](http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html) and the endpoint documentation can be found at [ONTAP REST API documentation](https://library.netapp.com/ecmdocs/ECMLP2856304/html/index.html).

# Nagios checks
## Aggregates
### Health
**_Command:_** `check_aggregate_state`

**_Options:_**

| *Option* | *Description* | *Default* | *Note* |
|:---------|:--------------|:---------:|:------:|
| `--aggregate=<uuid>` | Only check aggregate with UUID <uuid> | - | If omitted all volumes will be checked |
| `--ca-file=<file>` | Use CA certificate from <file> for validation of SSL certificate | - |
| `--help` | Show help text | - | - |
| `--insecure` | Skip validation of SSL server certificate | - | - |
| `--host=<host>` | Host to connect to | - | **mandatory** |
| `--password=<pass>` | Password for authentication | - | - |
| `--password-file=<passf>` | Read password for authentication from file <passf> | - | Only the first line in <passf> is used as password |
| `--timeout=<sec>` | HTTP connection timeout in seconds | 60 | - |
| `--user=<user>` | Username for authentication | - | **mandatory** |
| `--version` | Show version information | - | - |

### Utilisation
**_Command:_** `check_aggregate_utilisation`

**_Options:_**

| *Option* | *Description* | *Default* | *Note* |
|:---------|:--------------|:---------:|:------:|
| `--aggregate=<uuid>` | Only check aggregate with UUID <uuid> | - | If omitted all volumes will be checked |
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
| `--warning=<pct>` | Report WARNING if used space is larger then <cpct> percent of the volume | 80.0% | - |


## Disks
### Health
**_Command:_** `check_disk_health`

**_Options:_**

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

## Shelves
### Health
**_Command:_** `check_shelf_health`

**_Options:_**

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

## Volumes
### Health
**_Command:_** `check_volume_state`

**_Options:_**

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


### Utilisation
**_Command:_** `check_volume_utilisation`

**_Options:_**

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


