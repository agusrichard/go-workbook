package main

import "db-experiment/server"

func main() {
	s := server.SetupServer()

	s.Run(":9000")
}