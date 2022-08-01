package main

import "sunborn.com/protobuf-demo/server"

func main() {
	s := server.NewServer("tcp4", "127.0.0.1", "8099")
	s.Serve()
}
