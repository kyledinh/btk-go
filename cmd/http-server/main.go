package main

import "github.com/kyledinh/btk-go/cmd/http-server/server"

func main() {
	go server.Server()
}
