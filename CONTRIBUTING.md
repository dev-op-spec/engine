# Implementation details
To streamline dev efforts, the opctl project is maintained as a monorepo containing semi-independent sub-projects.

Typically at a minimum, each sub-project includes it's own `README.md`, `CONTRIBUTING.md`, and`.opspec`. 

The project is configured as a single go module, which allows dependencies to be shared during development, and should allow most IDEs to understand the project configuration without manual configuration.

# How it's laid out

## [api](api)
OpenAPI spec for the opctl ReST API.

## [cli](cli)
CLI for opctl.

## [opspec](opspec)
JSON schema for the opspec language.

used by:
- [sdks/go](sdks/go)

## [sdks/go](sdks/go)
SDK for integrating with opctl from golang.

used by:
- [cli](cli)

## [sdks/js](sdks/js)
SDK (written in typescript) for integrating with opctl from javascript/typescript.

used by:
- [webapp](webapp)
- [sdks/react](sdks/react)

## [sdks/react](sdks/react)
SDK (written in typescript) for integrating with opctl from react.

used by:
- [webapp](webapp)

## [test-suite](test-suite)
End to end tests, maintained as ops, for testing opctl and SDKs.

used by:
- [cli](cli)
- [sdks/go](sdks/go)

## [webapp](webapp)
Webapp for opctl.

used by:
- [cli](cli)

## [website](website)
Website hosted at [https://opctl.io](https://opctl.io).

# Pull requests
Pull requests are subject to:

- approval by one or more [maintainers](https://github.com/orgs/opctl/teams/maintainers/members)
- the [build](.opspec/build) op continuing to run with a successful outcome
