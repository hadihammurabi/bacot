package main

import (
	"bacot/src/core/chain"
	"bacot/src/web"
)

func main() {
	blockchain := chain.New("bacot", 1)

	app := web.New(blockchain)
	app.Listen(8080)
}
