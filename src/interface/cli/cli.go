package cli

import (
	"github.com/urfave/cli/v2"

	"bacot/src/core/chain"
	"bacot/src/interface/web"
)

func New(blockchain *chain.Chain) *cli.App {
	app := &cli.App{
		Name: "BACOT",
		Usage: "Blocks And Chain Of Transcript (blockchain engine)",
		Commands: []*cli.Command{
      {
        Name: "server",
        Aliases: []string{"s"},
				Usage: "run BACOT's server",
				Flags: []cli.Flag {
					&cli.StringFlag{
						Name: "port",
						Aliases: []string{"p"},
						Value: "8080",
						Usage: "set port of BACOT's server",
					},
				},
        Action:  func(c *cli.Context) error {
					web.New(blockchain, c.String("port"))
					return nil
        },
      },
    },
	}

	return app
}
