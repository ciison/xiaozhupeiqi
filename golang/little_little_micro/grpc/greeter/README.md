# Greeter Service

This is the Greeter service

Generated with

```
micro new xiaozhupeiqi/golang/little_little_micro/grpc/greeter --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.greeter
- Type: srv
- Alias: greeter

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./greeter-srv
```

Build a docker image
```
make docker
```