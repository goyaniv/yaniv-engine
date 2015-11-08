package main

var s *Server

func main() {
	s = ServerNew()
	LaunchHTTP()
}
