# strategy Service

This is the strategy service

Generated with

```
micro new github.com/jason-wj/bitesla/service/service-strategy --namespace=bitesla --alias=strategy --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: bitesla.srv.strategy
- Type: srv
- Alias: strategy

## Dependencies

Micro services depend on service discovery. The default is consul.

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
./strategy-srv
```

Build a docker image
```
make docker
```