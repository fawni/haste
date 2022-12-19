# haste

> Be really pretty. Be really simple.

[![Latest Release](https://img.shields.io/github/release/x6r/haste.svg)](https://github.com/x6r/haste/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/x6r/haste/build.yml?logo=github&branch=master)](https://github.com/x6r/haste/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/x6r/haste)](https://goreportcard.com/report/github.com/x6r/haste)

haste is a simple cross-platform hastebin command-line client.

## Installation

### Binaries

download a binary from the [releases](https://github.com/x6r/haste/releases)
page.

### Build from source

Go 1.16 or higher required. ([install instructions](https://golang.org/doc/install.html))

```sh
go install github.com/x6r/haste@latest
```

## Usage

```sh
$ haste # interactively prompts you for data
$ haste "text here" # uploads "text here" to hastebin and prints the url
$ haste -f file.go # uploads content of file.go to hastebin and prints the url
$ echo "text" | haste -i "https://hastebin.cc" -r # uploads "text" to a custom haste instance and prints the raw url
```

`haste -h` for more info.

## Disclaimer

the client defaults to [my own](https://p.x4.pm) instance. it is recommended to use a more reliable instance such as [hastebin.cc](https://hastebin.cc).

## license

[ISC](https://github.com/x6r/haste/blob/master/LICENSE)
