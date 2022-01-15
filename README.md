# haste

cross-platform cli hastebin client.

## installation

```sh
go install github.com/x6r/haste@latest
```

## usage

```sh
$ haste "text here" # uploads "text here" to hastebin and copies the returned url
$ haste -f file.go # uploads content of file.go to hastebin and copies the returned url
$ haste -i "https://hastebin.cc" -r "text" # uploads "text" to a custom hastebin instance and copies the raw url
```

`haste -h` for more info.

## notes

- default instance is [my own](https://p.x4.pm), if you _for some reason_ are using this client i highly recommend you use a more relibale instance such as [hastebin.cc](https://hastebin.cc).
