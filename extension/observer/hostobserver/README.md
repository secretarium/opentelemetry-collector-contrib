# Host Observer Extension

<!-- status autogenerated section -->
| Status        |           |
| ------------- |-----------|
| Stability     | [beta]  |
| Distributions | [contrib], [splunk], [sumo] |
| Issues        | ![Open issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aopen%20label%3Aextension%2Fhostobserver%20&label=open&color=orange&logo=opentelemetry) ![Closed issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aclosed%20label%3Aextension%2Fhostobserver%20&label=closed&color=blue&logo=opentelemetry) |

[beta]: https://github.com/open-telemetry/opentelemetry-collector#beta
[contrib]: https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-contrib
[splunk]: https://github.com/signalfx/splunk-otel-collector
[sumo]: https://github.com/SumoLogic/sumologic-otel-collector
<!-- end autogenerated section -->

The `host_observer` looks at the current host for listening network endpoints.

It will look for all listening sockets on TCP and UDP over IPv4 and IPv6.

It uses the /proc filesystem and requires the SYS_PTRACE and DAC_READ_SEARCH capabilities so that it can determine what processes own the listening sockets.

### Configuration

#### `refresh_interval`

Determines how often to look for changes in endpoints.

default: `10s`

### Endpoint Variables

Endpoint variables exposed by this observer are as follows.

| Variable  | Description                                                                                |
|-----------|--------------------------------------------------------------------------------------------|
| type      | `"port"`                                                                                     |
| name      | name of the process associated to the port                                                 |
| port      | port number                                                                                |
| command   | full command used to invoke this process, including the executable itself at the beginning |
| is_ipv6   | `true` if the endpoint is IPv6                                                             |
| transport | "TCP" or "UDP"                                                                             |
