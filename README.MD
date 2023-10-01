# Golang POS

API backend rest Admin POS Service using Gin Framework and Gorm ORM Database.

## Command

- ### App Lifecyle

```sh
$ go run main.go || make go run || ./run.sh for start service and generate swagger
```

- ### Docker Lifecycle

```sh
$ docker-compose build
```
```sh
$ docker login
```
```sh
$ docker tag admin-pos-service-app:latest digylabs/admin-pos-service-app:latest
```
```sh
$ docker push digylabs/admin-pos-service-app:latest
```

- ### Swagger

```sh
$ go get -u github.com/swaggo/swag/cmd/swag
$ $HOME/go/bin/swag init
```
## Endpoint
[open swagger](/docs/index.html)

https://www.kaggle.com/datasets/anggagewor/data-wilayah-republic-indonesia?resource=download