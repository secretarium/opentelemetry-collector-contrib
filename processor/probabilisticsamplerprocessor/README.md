# Probabilistic Sampling Processor

<!-- status autogenerated section -->
| Status        |           |
| ------------- |-----------|
| Stability     | [alpha]: logs   |
|               | [beta]: traces   |
| Distributions | [core], [contrib], [aws], [observiq], [splunk], [sumo] |
| Issues        | ![Open issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aopen%20label%3Aprocessor%2Fprobabilisticsampler%20&label=open&color=orange&logo=opentelemetry) ![Closed issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aclosed%20label%3Aprocessor%2Fprobabilisticsampler%20&label=closed&color=blue&logo=opentelemetry) |

[alpha]: https://github.com/open-telemetry/opentelemetry-collector#alpha
[beta]: https://github.com/open-telemetry/opentelemetry-collector#beta
[core]: https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol
[contrib]: https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-contrib
[aws]: https://github.com/aws-observability/aws-otel-collector
[observiq]: https://github.com/observIQ/observiq-otel-collector
[splunk]: https://github.com/signalfx/splunk-otel-collector
[sumo]: https://github.com/SumoLogic/sumologic-otel-collector
<!-- end autogenerated section -->

The probabilistic sampler supports two types of sampling for traces:

1. `sampling.priority` [semantic
convention](https://github.com/opentracing/specification/blob/master/semantic_conventions.md#span-tags-table)
as defined by OpenTracing
1. Trace ID hashing

The `sampling.priority` semantic convention takes priority over trace ID hashing. As the name
implies, trace ID hashing samples based on hash values determined by trace IDs.  See [Hashing](#hashing) for more information.

The following configuration options can be modified:
- `hash_seed` (no default): An integer used to compute the hash algorithm. Note that all collectors for a given tier (e.g. behind the same load balancer) should have the same hash_seed.
- `sampling_percentage` (default = 0): Percentage at which traces are sampled; >= 100 samples all traces

Examples:

```yaml
processors:
  probabilistic_sampler:
    hash_seed: 22
    sampling_percentage: 15.3
```

The probabilistic sampler supports sampling logs according to their trace ID, or by a specific log record attribute.

The probabilistic sampler optionally may use a `hash_seed` to compute the hash of a log record.
This sampler samples based on hash values determined by log records. See [Hashing](#hashing) for more information.

The following configuration options can be modified:
- `hash_seed` (no default, optional): An integer used to compute the hash algorithm. Note that all collectors for a given tier (e.g. behind the same load balancer) should have the same hash_seed.
- `sampling_percentage` (required): Percentage at which logs are sampled; >= 100 samples all logs, 0 rejects all logs.
- `attribute_source` (default = traceID, optional): defines where to look for the attribute in from_attribute. The allowed values are `traceID` or `record`.
- `from_attribute` (default = null, optional): The optional name of a log record attribute used for sampling purposes, such as a unique log record ID. The value of the attribute is only used if the trace ID is absent or if `attribute_source` is set to `record`.
- `sampling_priority` (default = null, optional): The optional name of a log record attribute used to set a different sampling priority from the `sampling_percentage` setting. 0 means to never sample the log record, and >= 100 means to always sample the log record.

## Hashing

In order for hashing to work, all collectors for a given tier (e.g. behind the same load balancer)
must have the same `hash_seed`. It is also possible to leverage a different `hash_seed` at
different collector tiers to support additional sampling requirements. Please refer to
[config.go](./config.go) for the config spec.

Examples:

Sample 15% of the logs:
```yaml
processors:
  probabilistic_sampler:
    sampling_percentage: 15
```

Sample logs according to their logID attribute:

```yaml
processors:
  probabilistic_sampler:
    sampling_percentage: 15
    attribute_source: record # possible values: one of record or traceID
    from_attribute: logID # value is required if the source is not traceID
```

Sample logs according to the attribute `priority`:

```yaml
processors:
  probabilistic_sampler:
    sampling_percentage: 15
    sampling_priority: priority
```


Refer to [config.yaml](./testdata/config.yaml) for detailed
examples on using the processor.
