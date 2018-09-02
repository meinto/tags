package main

import "tags/server"

func main() {
	s := server.Create()
	s.Start()
}
