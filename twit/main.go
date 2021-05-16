package main

import (
	"twit/servers"
)

func main() {
	router, _ := servers.SetupServer()

	router.Run(":9090")
}
