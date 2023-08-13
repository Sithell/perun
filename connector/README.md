# Connector

Connector is a microservice that provides an API to run commands on providers

Providers call connector to establish a connection through which they recieve commands

## Getting started
```shell
go run cmd/server.go
```

## Usage
Check `pb/connector.proto` for up-to-date spec

## TODO

- [x] register providers in DB on connect
- [x] method to return active connections
- [x] provider parameter on every Api request
- [ ] refactor: wait for channels via for expression
- [ ] How to identify a provider after reconnect (by host? generated id?)
