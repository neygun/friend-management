# Friend Management API

A simple API for managing friends

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
$ docker build -t fm-app .
```

Then run docker-compose

```
$ docker-compose --env-file ./.env up
```
