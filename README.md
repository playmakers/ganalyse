# Ganalyse
Google Analytics Rocks


## Install
Given `$GOPATH` is set and `$GOPATH/bin` added to `$PATH`. Install dependencies:

```
    go get -u github.com/djimenez/iconv-go
    go get github.com/paulrosania/go-charset/charset
    go get -u github.com/go-martini/martini
    go get -u github.com/martini-contrib/render
    go get -u github.com/onsi/ginkgo/ginkgo
    go get -u github.com/onsi/gomega
    go get -u github.com/tools/godep
```

## Execute
To run script, execute:

```
make run

```

and open browser on http://localhost:3000

## Test
To run specs, execute:

```
make test

```

## Deploy

Before deployment, make sure to update dependencies

`make dependencies`

more see: http://mmcgrana.github.io/2012/09/getting-started-with-go-on-heroku.html

### run single spec
Add a `F` in front of the line to focus on. See http://onsi.github.io/ginkgo/#focused-specs

## TODO

- [x] Resolve Multiple URL
- [x] Forelle, fix when not color
- [ ] Color, Size Mapping
- [ ] better JSON output?
