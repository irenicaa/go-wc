# go-wc

[![GoDoc](https://godoc.org/github.com/irenicaa/go-wc?status.svg)](https://godoc.org/github.com/irenicaa/go-wc)
[![Go Report Card](https://goreportcard.com/badge/github.com/irenicaa/go-wc)](https://goreportcard.com/report/github.com/irenicaa/go-wc)
[![Build Status](https://app.travis-ci.com/irenicaa/go-wc.svg?branch=master)](https://app.travis-ci.com/irenicaa/go-wc)
[![codecov](https://codecov.io/gh/irenicaa/go-wc/branch/master/graph/badge.svg)](https://codecov.io/gh/irenicaa/go-wc)

The clone of the [Unix wc tool](<https://en.wikipedia.org/wiki/Wc_(Unix)>).

## Installation

```
$ go get github.com/irenicaa/go-wc/...
```

## Usage

```
$ go-wc -h | -help | --help
$ go-wc [options]
```

Stdin: text in the [UTF-8 encoding](https://en.wikipedia.org/wiki/UTF-8).

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit;
- `-l` &mdash; show line count;
- `-w` &mdash; show word count;
- `-c` &mdash; show byte count;
- `-m` &mdash; show symbol count.

## Output Example

```
     10     10     21     21
```

## License

The MIT License (MIT)

Copyright &copy; 2021 irenica
