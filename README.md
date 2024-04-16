# Meli

## General aspects
This project follows the hexagonal architecture and implements a `QueryFetcher` mechanism for getting additional information from external services using Goroutines.

## Tools
- [Gin - Web framework ](https://gin-gonic.com/) 
- [Sqlc - SQL query builder](https://sqlc.dev/)
- [Postgresql - SQL Engine](https://https://www.postgresql.org/)
- [Docker](https://docs.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [DBeaver - SQL client](https://dbeaver.io/)

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
$ git clone git@github.com:BraulioAguilarDev/meli.git

# Position in meli folder
$ cp .env.example .env # update MELI_TOKEN

# Mandatory commands
$ make build && make dc-up

# Call upload file endpoint
$ curl --location 'http://localhost:8080/v1/items/upload-file' \
--form 'file=@"$(pwd)/meli/technical_challenge_data.csv"'
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