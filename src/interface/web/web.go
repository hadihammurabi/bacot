package web

import (
	"fmt"
	"encoding/json"

	"github.com/gofiber/fiber"

	"bacot/src/core/chain"
)

var nodes []*Node

func New(blockchain *chain.Chain, port string) {
	chain := chain.New("bacot", 1)

	app := fiber.New()
	nodes = append(nodes, &Node{ URL: fmt.Sprintf("http://localhost:%s", port) })

	app.Get("/health-check", func(c *fiber.Ctx) {
		data := map[string]string{
			"status": "good",
		}
		c.JSON(data)
	})

	app.Get("/chain", func(c *fiber.Ctx) {
    c.JSON(chain)
	})
	
	app.Post("/chain", func(c *fiber.Ctx) {
		data := make(map[string]interface{})
		json.Unmarshal([]byte(c.Body()), &data)

		block := chain.Add(data)
	
    c.JSON(block)
	})

	nodeRoute(app, nodes)
	
	app.Listen(port)
}
