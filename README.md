# go-wc

[![GoDoc](https://godoc.org/github.com/irenicaa/go-wc?status.svg)](https://godoc.org/github.com/irenicaa/go-wc)
[![Go Report Card](https://goreportcard.com/badge/github.com/irenicaa/go-wc)](https://goreportcard.com/report/github.com/irenicaa/go-wc)
[![Build Status](https://app.travis-ci.com/irenicaa/go-wc.svg?branch=master)](https://app.travis-ci.com/irenicaa/go-wc)
[![codecov](https://codecov.io/gh/irenicaa/go-wc/branch/master/graph/badge.svg)](https://codecov.io/gh/irenicaa/go-wc)

The clone of the Unix wc tool.

## Installation

```
$ go get github.com/irenicaa/go-wc/...
```

## Usage

```
$ go-wc -h | -help | --help
$ go-wc [options]
```

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit;
- `-l` &mdash; show line count;
- `-w` &mdash; show word count;
- `-c` &mdash; show byte count;
- `-m` &mdash; show symbol count.

## License

The MIT License (MIT)

Copyright &copy; 2020 irenica
