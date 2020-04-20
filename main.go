package main

import (
	"os"
	"log"

	"bacot/src/core/chain"
	"bacot/src/cli"
)

func main() {
	blockchain := chain.New("bacot", 1)
	app := cli.New(blockchain)
	
	err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
