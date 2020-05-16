**_Note:_** Because I'm running my own servers for several years, main development is done at at https://git.ypbind.de/cgit/check_clustered_netapp/

----

# Preface
Starting with ONTAP 9.6 NetApp file server provide a REST API (in addition to the classic XML API used by the NetApp SDK).
The general documentation can be found at the [ONTAP 9 documentation center](http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html) and the endpoint documentation can be found at [ONTAP REST API documentation](https://library.netapp.com/ecmdocs/ECMLP2856304/html/index.html).

# Nagios checks
## Check volume utilisation
To check the utilisation of a single volume or all volumes `check_volume_utilisation` can be used.

| *Option* | *Description* | *Default* | *Note* |
|:---------|:--------------|:---------:|:------:|
| `--ca-file=<file> | Use CA certificate from <file> for validation of SSL certificate | - |
| `--critical=<pct>` | Report CRITICAL if used space is larger then <cpct> percent of the volume | 90.0% | - |
| `--help` | Show help text | - | - |
| `--insecure` | Skip validation of SSL server certificate | - | - |
| `--host=<host>` | Host to connect to | - | **mandatory** |
| `--password=<pass>` | Password for authentication | - | - |
| `--password-file=<passf>` | Read password for authentication from file <passf> | - | Only the first line in <passf> is used as password |
| `--user=<user>` | Username for authentication | - | **mandatory** |
| `--timeout=<sec>` | HTTP connection timeout in seconds | 60 | - |
| `--version` | Show version information | - | - |
| `--volume=<uuid>` | Only check volume with UUID <uuid> | - | If omitted all volumes will be checked |
| `--warning=<pct>` | Report WARNING if used space is larger then <cpct> percent of the volume | 80.0% | - |

## Check volume health
To check the utilisation of a single volume or all volumes `check_volume_state` can be used.

| *Option* | *Description* | *Default* | *Note* |
|:---------|:--------------|:---------:|:------:|
| `--ca-file=<file> | Use CA certificate from <file> for validation of SSL certificate | - |
| `--help` | Show help text | - | - |
| `--insecure` | Skip validation of SSL server certificate | - | - |
| `--host=<host>` | Host to connect to | - | **mandatory** |
| `--password=<pass>` | Password for authentication | - | - |
| `--password-file=<passf>` | Read password for authentication from file <passf> | - | Only the first line in <passf> is used as password |
| `--user=<user>` | Username for authentication | - | **mandatory** |
| `--timeout=<sec>` | HTTP connection timeout in seconds | 60 | - |
| `--version` | Show version information | - | - |
| `--volume=<uuid>` | Only check volume with UUID <uuid> | - | If omitted all volumes will be checked |

