package web

import (
	"fmt"
	"errors"
	"encoding/json"

	"github.com/gofiber/fiber"
	"github.com/go-resty/resty/v2"
)

type Node struct {
	URL string    `json:"url"`
}

func nodeRoute(app *fiber.App, nodes []*Node) {

	app.Get("/nodes", func(c *fiber.Ctx) {
		c.JSON(map[string][]*Node{
			"nodes": nodes,
		})
	})

	app.Post("/nodes", func(c *fiber.Ctx) {
		data := make(map[string]interface{})
		json.Unmarshal([]byte(c.Body()), &data)

		nodeURL := data["url"].(string)

		err := healthCheck(nodeURL)

		if err != nil {
			c.JSON(err)
		} else {
			nodes = append(nodes, &Node{URL: nodeURL})
			c.JSON(map[string][]*Node{
				"nodes": nodes,
			})
		}
	})

}

func healthCheck(nodeURL string) error {
	http := resty.New()
	res, err := http.R().Get(fmt.Sprintf("%s/health-check", nodeURL))
	if err != nil {
		return err
	}

	var data map[string]interface{}
	err = json.Unmarshal(res.Body(), &data)
	if err != nil {
		return err
	}

	if data["status"] != "good" {
		return errors.New("this node doesn't healthy")
	}

	return nil
}