# f2go converter

Converts any file to Golang variable file.

[![Build Status](https://travis-ci.org/geniusrabbit/f2go.svg?branch=master)](https://travis-ci.org/geniusrabbit/f2go)
[![Go Report Card](https://goreportcard.com/badge/github.com/geniusrabbit/f2go)](https://goreportcard.com/report/github.com/geniusrabbit/f2go)
[![GoDoc](https://godoc.org/github.com/geniusrabbit/f2go?status.svg)](https://godoc.org/github.com/geniusrabbit/f2go)
[![Coverage Status](https://coveralls.io/repos/github/geniusrabbit/f2go/badge.svg)](https://coveralls.io/github/geniusrabbit/f2go)

> License Apache 2.0

## Using

```sh
f2go --help

NAME:
   f2go - converts any file to golang

USAGE:
   main [global options] command [command options] source_file [destination file]

VERSION:
   0.1.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --bytes           convert to the bytes array (default: false)
   --template value  (default: "data.go")
   --package value   Custom name of the package
   --varname value   Custom name of the variable
   -f                Save result to the file (default: false)
   --help, -h        show help (default: false)
   --version, -v     print the version (default: false)
```

```sh
f2go swagger.json > 'swagger.json.go'
```
