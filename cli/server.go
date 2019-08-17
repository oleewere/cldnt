package cli

import (
	"github.com/oleewere/cldnt/web"
	"github.com/urfave/cli"
)

// ServerCommand command for starting a web server
func ServerCommand() cli.Command {
	var port int
	return cli.Command{
		Name:  "serve",
		Usage: "Start web server",
		Action: func(c *cli.Context) error {
			web.StartServer(port)
			return nil
		},
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:        "port",
				Value:       7777,
				Destination: &port,
			},
		},
	}
}
