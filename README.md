# Simple todo rest application in Go

## Running the service

Start the database:

```
$ docker-compose up -d
```

WIP-Create the tables (up.sql) in the docker container.
 
Run service.

```
$ go run main.go
$ go run main.go serve
$ go run main.go version
$ go run main.go migrate
```

```
go mod tidy
```

## User api calls

```
$ http -v POST http://127.0.0.1:9090/api/users/ email=foo@bar.com password=123456
```

## Todo api calls

```
$ http -v -a foo@bar.com:123456 POST http://127.0.0.1:9090/api/todos/ name=foo done:=true
$ http -v -a foo@bar.com:123456 GET http://127.0.0.1:9090/api/todos/
$ http -v -a foo@bar.com:123456 GET http://127.0.0.1:9090/api/todos/1/
$ http -v -a foo@bar.com:123456 DELETE http://127.0.0.1:9090/api/todos/1/
$ http -v -a foo@bar.com:123456 PATCH http://127.0.0.1:9090/api/todos/1/ name=foo done:=false
```

## Todo
* automatically create the default tables
* write tests
* figure out migration service (not sure if it is working)

## Links
Highly inspired from [here](https://github.com/theaaf/todos) but did not want to rely on an ORM.

### Notes
* https://www.calhoun.io/updating-and-deleting-records-using-sql/
* https://github.com/go-pg/pg/wiki/Model-Definition
* https://www.calhoun.io/inserting-records-into-a-postgresql-database-with-gos-database-sql-package/
* https://github.com/go-sql-driver/mysql/wiki/Examples
* http://go-database-sql.org/retrieving.html
* https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html
* https://github.com/tinrab/curly-waddle
