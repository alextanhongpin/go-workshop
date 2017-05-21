# Go Workshop

Structuring a golang's web api


## Useful commands

Go imports
```bash
$ go get golang.org/x/tools/cmd/goimports
$ goimport
```


## Benchmarking HTTP

```bash

$ brew install wrk
$ wrk -c100 -d10 -t10 "http://localhost:8080/jobs"
```
