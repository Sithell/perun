# Perun

_Perun_ is a Peer-To-Peer Infrastructure-As-A-Service platform

![](https://upload.wikimedia.org/wikipedia/commons/e/e4/Thunder_mark_%283%29.svg)

Perun allows you to run code on other peoples machines or let other people run their code on your machine

## Architecture

Perun consists of 3 main components: _Client_, _Manager_ and _Provider_

- Clients request Manager to run their code remotely
- Manager assigns execution of client code to one of the Providers
- Providers run code assigned by Manager and return back the results

## Roadmap

### TODO

- [x] Make a provider-api (name in progress) with RegisterProvider method
- [x] Message broker for passing jobs to assigner?
- [x] Make assigner check for updates in db:jobs and assign their execution to providers
- [x] Connector: register providers in DB on connect
- [ ] Connector: method to return active connections
- [ ] Connector: provider parameter on every Api request
- [ ] Assigner: call Connector methods
- [ ] Provider: Call Connector InitConnection on start and listen for commands
- [ ] Add an UpdateRunStatus method to provider-api
- [ ] Make provider push container status and stdout updates to provider-api
- [ ] Ability to kill a job via client-api
- [ ] Integration tests
- [ ] Add docker-compose for all control-plane services
- [ ] Add Job and Run status constants
- [ ] Make job id a string
- [ ] Use OpenAPI 3.0 instead of Swagger 2.0
- [ ] Handle provider host aliases (domain.com, 123.122.33.22, etc.)

### Features possible in the future
- re-run container if interrupted
- get live stdout/stderr via websockets
- ssh into containers
- privacy for providers -> ip address hidden
- shh proxy
- exposed ports limited
- proxy exposed ports
- allow to mount drives (?)
- providers rating
- cli client
- client libraries

## How to contribute

### Code generation

```shell
go generate ./...
```

