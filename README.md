# GG-IceCreamShop

GraphQL + gRPC Ice Cream Shop

## Table of Contents
1. [Overview](#overview)
1. [Setup](#setup)
	+ [Setup Protobuf](#protobuf)
	+ [Setup DB Migration Tool](#migration)
1. [Getting Started](#getting-started)
	+ [Running Migrations](#running-migrations)
	+ [Using `docker-compose`](#docker-compose)
1. [Importing Ice Creams JSON](#import)

### <a name="overview">Overview</a>

```
                                                    --------
                                                    | auth |
                                                    --------
                                                   -       -
                                                 -           -
                                              grpc            gprc
                                             -                   -
                                           -                       -
-------------                      ---------------                   --------          -----------
| api-users | < - - graphql - - >  | api-gateway | - - -  grpc - - - | user |  < --- > | user_db |
-------------    GET /graphiql     ---------------                   --------          -----------
                  POST /query              -
                                             -               
                                              grpc
                                                 -
                                                   -
                                                    -------------         ----------------
                                                    | ice_cream | < --- > | ice_cream_db |
                                                    -------------         ----------------
```

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

#### <a name="docker-compose">Using `docker-compose`</a>

To start the entire stack in a single command with `docker-compose`:

```shell
docker-compose up -d
```

To stop:

```shell
docker-compose down
```

To reset the data volumes while stopping the docker services:

```shell
docker-compose down -v
```

To apply changes in microservice:

```shell
docker-compose up --build -d
```

*Note: In the auth section there is the `JWT_SECRET_KEY` value that can be & should be changed to some secure value*

### <a name="import">Importing Ice Creams JSON</a>

The ice icream microservice has the import cli tool included in `cmd/import/main.go`. Presently, it only accepts a `-url` flag and the JSON has to conform to the schema as shown in the <a href="https://gist.githubusercontent.com/penmanglewood/f264e8d926b4c4a9926aa1de8fdb509a/raw/992f3c8a519ecd3d947bc48627ffefcf947f80bd/icecream.json" target="_blank">sample ice creams json</a>.

**Sample invocation**
From the ice cream microservice root dir:
```shell
go run cmd/import/main.go -url http://somehost/somejson.json
```

The import command is an idempotent operation, ice creams that are already exists wil be skipped.

*Note: If the `-url` flag is not supplied, it will default to the sample ice creams json.*