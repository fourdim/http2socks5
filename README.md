# http2socks5
Transfer protocol from http to socks5

## TODOS

- TODO: Resolve chrome DNS requests with random DNS names(prevent DNS hijacking).
- TODO: Make the ports changeable.

This program can be used normally without these featrues.

## Dependencies

- golang.org/x/net

## Build

```
$ go get -u golang.org/x/net
$ cd http2socks5
$ go build
```

## Ports
By default

127.0.0.1:8080 for http in

127.0.0.1:51837 for socks5 out

You can change it in the source code.
