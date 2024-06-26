# Friend Management API

A simple API for managing friends with features like "Friend", "Unfriend", "Block", "Receive Updates".

## Technologies

* Go language
* go-chi
* volatiletech/sqlboiler
* PostgreSQL
* Docker

## Endpoints

* `POST /users`: creates a new user
* `POST /friends`: creates a new friend connection between two email addresses
* `POST /friends/list`: returns the friends list for an email address
* `POST /friends/common`: returns the common friends list between two email addresses
* `POST /friends/subscription`: subscribes to updates from an email address
* `POST /friends/block`: blocks updates from an email address
* `POST /friends/recipients`: returns all email addresses that can receive updates from an email address

## Run

Run the following command from the root directory of the project to build the docker image

```
make build
```

Then run docker-compose

```
make run
```

Create a database connection with the following attributes:
* host: localhost
* port: 5432
* database: fm-pg
* username: postgres
* password: postgres

Migrate up database
```
make migrate-up
```

## Project structure

```
.friend-management
├── cmd/
│   └── serverd/
│       ├── router/
│       │   └── router.go
│       └── main.go
├── data/
│   └── migration/
├── internal/
│   ├── handler/
│   ├── model/
│   ├── repository/
│   └── service/
├── pkg/
│   ├── db/
│   └── util/
│       └── test/
├── vendor/
├── docker-compose.yaml
├── Dockerfile
├── sqlboiler.toml
├── .env
├── go.mod
├── go.sum
├── Makefile
└── README.md
```
