package web

import (
	"encoding/json"

	"github.com/gofiber/fiber"

	"bacot/src/core/chain"
)

func New(blockchain *chain.Chain) {
	chain := chain.New("bacot", 1)

	app := fiber.New()

	app.Get("/chain", func(c *fiber.Ctx) {
    c.JSON(chain)
	})
	
	app.Post("/chain", func(c *fiber.Ctx) {
		data := make(map[string]interface{})
		json.Unmarshal([]byte(c.Body()), &data)

		block := chain.Add(data)
	
    c.JSON(block)
  })

	app.Listen(8080)
}
