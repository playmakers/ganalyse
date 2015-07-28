# Ganalyse
Google Analytics Rocks


## Install
Given `$GOPATH` is set and `$GOPATH/bin` added to `$PATH`. Install dependencies:

```
    go get github.com/djimenez/iconv-go
    go get github.com/go-martini/martini
    go get github.com/martini-contrib/render
    go get github.com/onsi/ginkgo/ginkgo
    go get github.com/onsi/gomega
    go get github.com/tools/godep
```

## Execute
To run script, execute:

```
go run main.go

```

and open browser on http://localhost:3000

## Test
To run specs, execute:

```
ginkgo specs/vendors

```

## Deploy

Before deployment, make sure to update dependencies

`godep save`

more see: http://mmcgrana.github.io/2012/09/getting-started-with-go-on-heroku.html

### run single spec
Add a `F` in front of the line to focus on. See http://onsi.github.io/ginkgo/#focused-specs

## TODO

- [x] Resolve Multiple URL
- [ ] Color, Size Mapping
- [ ] better JSON output?
