# GG-IceCreamShop

GraphQL + gRPC Ice Cream Shop

## Table of Contents

1. [Setup](#setup)
	+ [Setup Protobuf](#protobuf)
	+ [Setup DB Migration Tool](#migration)
2. [Getting Started](#getting-started)
	+ [Running Migrations](#running-migrations)

### <a name="setup">Setup</a>

#### <a name="protobuf">Setting up the `google.protobuf` package & `protoc` binary</a>

This project requires the `google.protobuf` proto package and expects to import from the `/usr/local/include` directory.

Download the applicable release on <a href="https://github.com/protocolbuffers/protobuf/releases" target="_blank">github</a>, and the downloaded zip archive contains the `bin` and `include` directory.

Extract the archive and copy the `bin/protoc` to your `GOBIN` path and copy the contents of the `include` directory to your `/usr/local/include` directory.

#### <a name="migration"></a>Setting up migration tool

This project uses the [`golang-migrate/migrate`](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) package to manage the database migration.

### <a name="getting-started">Getting Started</a>

#### <a name="running-migrations">Running Migrations</a>

To run db migration for each services, one has to go the service root folder and have the `DB_URL` environment set, then run the following command:

```shell
migrate -database "$DB_URL" -path migrations up
```

To reset the db, one can use the following command:

```shell
migrate -database "$DB_URL" drop
```
