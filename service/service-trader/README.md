# Trader Service

This is the Trader service

Generated with

```
micro new github.com/jason-wj/bitesla/service/service-trader --namespace=bitesla --alias=trader --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: bitesla.srv.trader
- Type: srv
- Alias: trader

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
./trader-srv
```

Build a docker image
```
make docker
```