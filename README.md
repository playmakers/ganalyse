# Ganalyse
Google Analytics Rocks

## Algo

### Prepare:

- google sheet with product/variant id, status


### Process

- read file/google doc
- get data from given urls
- send to ga with custom dimensions/value
- update google doc with values & status


## Execute

```
go run main.go

```

## Todo

* udpate source to have PM column in first place
* remember column/row and write back
* store to DB?
* use: http://onsi.github.io/ginkgo/

## Help

* https://gobyexample.com/channels
* https://golang.org/doc/effective_go.html
* https://github.com/go-sql-driver/mysql/wiki/Examples
