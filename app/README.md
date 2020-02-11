Dummy application which display hostname and version
====================================================

> GoLang webserver which purpose is to reply with the hostname and if existing,
the environment variable VERSION.

## Getting started

### Docker

#### Build

```
$ docker build -t felipecruz/go-server .
```

#### Run

```shell
$ docker run -d \
    --name app \
    -p 8080:8080 \
    -h host-1 \
    -e VERSION=v1.0.0 \
    felipecruz/go-server
```

#### Test

```
$ curl localhost:8080
Host: host-1, Version: v1.0.0
```

Liveness and readiness probes are replying on `:8080/live` and `:8080/ready`.

#### Cleanup

```
$ docker stop app
```
