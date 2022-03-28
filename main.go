package main

import "shorty/server"

func main() {
	s := server.New()
	s.Logger.Fatal(s.Start(":80"))
}
