# Golang 200

This project is meant to help learning go. It provides a basic implementation of a REST microservice exposing a CRUD API.
Data are persisted in a MongoDB NoSQL database and the application is deployed in Docker.

## Technical stack

* [Docker](https://www.docker.com)
* [MongoDB NoSQL database](https://www.mongodb.com)
* [Go is the language](https://golang.org)
* [Gorilla Mux the URL router](https://github.com/gorilla/mux)
* [Urfave negroni Web HTTP middleware](https://github.com/urfave/negroni)
* [Urfave cli the command line client parser](https://gopkg.in/urfave/cli.v1)
* [Sirupsen the logger](https://github.com/Sirupsen/logrus)
* [The database driver](https://gopkg.in/mgo.v2)
* [Glide the dependency manager](https://github.com/Masterminds/glide)
* [Golint the source linter](https://github.com/golang/lint)

## Architecture

![main architecture](doc/img/main_architecture.png "Main architecture")

![web architecture](doc/img/web_architecture.png "Web architecture")

## Build

```shell
make help
```
