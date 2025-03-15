package main

import (
	"lobmto-echo-example/server"
)

func main() {
	e := server.NewRouter()
	e.Logger.Fatal(e.Start(":8000"))
}
