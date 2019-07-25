# asciinema-player
[![Build Status](https://travis-ci.org/xakep666/asciinema-player.svg?branch=master)](https://travis-ci.org/xakep666/asciinema-player)
[![codecov](https://codecov.io/gh/xakep666/asciinema-player/branch/master/graph/badge.svg)](https://codecov.io/gh/xakep666/asciinema-player)
[![Go Report Card](https://goreportcard.com/badge/github.com/soopsio/asciinema-player)](https://goreportcard.com/report/github.com/soopsio/asciinema-player)
[![GoDoc](https://godoc.org/github.com/soopsio/asciinema-player/pkg/asciicast?status.svg)](https://godoc.org/github.com/soopsio/asciinema-player/pkg/asciicast)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

asciinema-player is a library and cli-app to play terminal sessions recorded by asciinema (http://github.com/asciinema/asciinema)

## Prerequisites
* Golang >= 1.10 or Vgo

## Installation
Library:
```bash
go get -v -u github.com/soopsio/pkg/asciicast
```

App:
```bash
go get -v -u github.com/soopsio/cmd/asciinema-player
```

## Usage
### App
```
$ ./asciinema-player --help
  Usage of ./asciinema-player:
    -f string
          path to asciinema v2 file
    -maxWait duration
          maximum time between frames (default 2s)
    -speed float
          speed adjustment: <1 - increase, >1 - decrease (default 1)
```
For example you can play test session `./asciinema-player -f test.cast`

[![asciicast](https://asciinema.org/a/189343.png)](https://asciinema.org/a/189343)

### Library
```go
parsed, err := parser.Parse(file)
if err != nil {
    return err
}

term, err := terminal.NewPty()
if err != nil {
    return err
}

if err := term.ToRaw(); err != nil {
    return err
}

defer term.Reset()

tp := &asciicast.TerminalPlayer{Terminal: term}

err = tp.Play(parsed, maxWait, speed)
if err != nil {
    return err
}
```
Library usage example is app actually.

## License
Asciinema-player project is licensed under the terms of the MIT license. Please see LICENSE in this repository for more details.