package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"strings"

	"golang.org/x/net/proxy"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	l, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Panic(err)
	}

	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}

		go handleClientRequest(client)
	}
}

func handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()

	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}
	var method, host, address string
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
	fmt.Println(host)
	hostPortURL, err := url.Parse(host)
	if err != nil {
		log.Println(err)
		return
	}

	if hostPortURL.Opaque == "443" { //https
		address = hostPortURL.Scheme + ":443"
	} else { //http
		if strings.Index(hostPortURL.Host, ":") == -1 { //http port 80 by default
			address = hostPortURL.Host + ":80"
		} else {
			address = hostPortURL.Host
		}

	}

	scoks5Dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:51837", nil, proxy.Direct)
	if err != nil {
		log.Println(err)
		return
	}
	server, err := scoks5Dialer.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		return
	}
	if method == "CONNECT" {
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\n\n")
	} else {
		server.Write(b[:n])
	}

	go io.Copy(server, client)
	io.Copy(client, server)
}
