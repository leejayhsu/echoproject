# echoproject
 simple API built with echo, gorm

### running go app
you'll need to have a postgres instance running (docker container or otherwise), and set these env vars:
```sh
POSTGRES_HOST=<your_postgres_host>
POSTGRES_USER=<your_postgres_user>
POSTGRES_DB=<your_db_name>
POSTGRES_PORT=<your_db_port>
POSTGRES_PASSWORD=<your_postgres_password>
B_ENV=local
```
`B_ENV` is an env var to tell the go app that you are running the app locally, and the only reason for this is to set the echo domain to "localhost" so that OSX doesn't ask you for network permissions everytime you compile the app.

To start the app, do:
```sh
go run main.go
```

### running with docker-compose
no need for any env vars, these are preset by `docker-compose.yml`
```sh
docker-compose build
docker-compose up
```