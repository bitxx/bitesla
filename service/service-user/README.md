# User Service

This is the User service

Generated with

```
micro new github.com/jason-wj/bitesla/service/service-user --namespace=bitesla --alias=user --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: bitesla.srv.user
- Type: srv
- Alias: user

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
./user-srv
```

Build a docker image
```
make docker
```