### Overview
Gor architecture tries to follow UNIX philosophy: everything made of pipes, various inputs multiplexing data to outputs.

You can [rate limit](../Rate-limiting.md), [filter](../Request-filtering.md), [rewrite](../Request-rewriting.md) requests or even use your own [middleware](../Middleware.md) to implement custom logic. Also, it is possible to replay requests at the higher rate for [load testing](../Saving-and-Replaying-from-file.md).

### Available inputs

* `--input-raw` - used to capture HTTP traffic, you should specify IP address or interface and application port. More about [capturing and replaying traffic](../Capturing-and-replaying-traffic.md). 
* `--input-file` - accepts file which previously was recorded using ` --output-file`. More about [saving and replaying from file](../Saving-and-Replaying-from-file.md)
* `--input-tcp` - used by Gor aggregation instance if you decided forward traffic from multiple forwarder Gor instances to it. Read about using [distributed configuration](../Distributed-configuration.md).

### Available outputs

* `--output-http` - replay HTTP traffic to given endpoint, accepts base url. Read [more about it](../Replaying-HTTP-traffic.md)
* `--output-file` - records incoming traffic to the file. More about [saving and replaying from file](../Saving-and-Replaying-from-file.md)
* `--output-tcp` - forward incoming data to another Gor instance, used in conjunction with `--input-tcp`. Read more about [distributed configuration](../Distributed-configuration.md).
* `--output-stdout` - used for debugging, outputs all data to stdout.
