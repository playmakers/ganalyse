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
To run script, execute:

```
go run main.go

```

## Test
To run specs, execute:

```
ginkgo specs/vendors

```

## Todo

* [x] udpate source to have PM column in first place
* [x] use: http://onsi.github.io/ginkgo/
* [ ] remember column/row and write back
* [ ] store to DB?
* [ ] add server to start job?


## Help

* https://gobyexample.com/channels
* https://golang.org/doc/effective_go.html
* https://github.com/go-sql-driver/mysql/wiki/Examples


```
  bundle install --path vendor/bundle
```



  // DEFAULT_SIZE  = "L"
  // DEFAULT_COLOR = "Schwarz"
  // DEFAULT_POS   = "???"
