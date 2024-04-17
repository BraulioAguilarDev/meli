# Meli

## General aspects
This project follows the hexagonal architecture and implements a `QueryFetcher` mechanism for getting additional information from external services using Goroutines.

### Theoretical

Go to this [link](/theoretical_exercise/README.md) for reading


## Tools project
- [Gin - Web framework ](https://gin-gonic.com/) 
- [Sqlc - SQL query builder](https://sqlc.dev/)
- [Postgresql - SQL Engine](https://https://www.postgresql.org/)
- [Docker](https://docs.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Build project
```sh
# Configure environment
$ cp .env.example .env

# Build docker image required for container
$ make docker
```

## Docker compose
```sh
# Creating container
$ make dc-up

# Deleting container & volume
$ make dc-down
```

## Test code

```sh
$ git clone https://github.com/BraulioAguilarDev/meli.git

# Position in meli folder
$ cp .env.example .env # update MELI_TOKEN

# Building meli image
$ make docker

# Starting api
$ make dc-up
```

### Request by Postman/Insomnia

#### Endpoint
`POST: http://localhost:8080/v1/items/upload-file`

![http request](/assets/image/request.png "Request example")

#### Success
```json
{
    "success": true,
    "message": "File has been processed successfully"
}
```

#### Errors
```json
{
    "success": false,
    "message": "some error message"
}
```

### Example run testing

```sh
$ go test ./internal/adapter/storage/postgres/... -count=1
```

### File examples

You can find [files](/assets/files/) allowed in this project (csv, xml, txt)
