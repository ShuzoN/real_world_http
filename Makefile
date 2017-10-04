all: help

help: 
	"install"

install:
	/usr/local/bin/brew install go@1.8
	/usr/local/bin/go get "github.com/k0kubun/pp"

server:
	/usr/local/bin/go run server.go

server_digest:
	/usr/local/bin/go run server_digest.go
