# haste

haste is a cross-platform hastebin CLI client.

## installation

```sh
go install github.com/x6r/haste@latest
```

## usage

```sh
$ haste # interactively prompts you for data
$ haste "text here" # uploads "text here" to hastebin and prints the url
$ haste -f file.go # uploads content of file.go to hastebin and prints the url
$ echo "text" | haste -i "https://hastebin.cc" -r # uploads "text" to a custom haste instance and prints the raw url
```

`haste -h` for more info.

## notes

default haste instance is [my own](https://p.x4.pm), i **highly** recommend you use a more reliable instance such as [hastebin.cc](https://hastebin.cc).

## license

[ISC](https://github.com/x6r/haste/blob/master/LICENSE)
