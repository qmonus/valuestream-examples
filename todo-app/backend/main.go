package main

import (
	"demo-backend-app/pkg/server"
)

func main() {
	s := server.Server{}
	s.Run()
}
