package main

import (
	"twit/servers"
)

func main() {
	router := servers.SetupServer()

	router.Run(":9090")
}
