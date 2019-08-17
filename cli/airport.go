package cli

import (
	"fmt"
	"os"

	"github.com/oleewere/cldnt/client"
	"github.com/urfave/cli"
)

func ListAirportsCommand() cli.Command {
	var airportSearchUrl string
	return cli.Command{
		Name:  "airport",
		Usage: "Get closest airports",
		Action: func(c *cli.Context) error {
			if c.Bool("ip") {
				client.CalculateLocationFromIP()
			} else if len(c.String("la")) > 0 && len(c.String("lo")) > 0 {

			} else {
				fmt.Println()
				os.Exit(1)
			}
			return nil
		},
		Flags: []cli.Flag{
			cli.StringFlag{Name: "airport-db-search-url, a", Usage: "Search URL for airport DB", Value: "https://mikerhodes.cloudant.com/airportdb/_design/view1/_search/geo", Destination: &airportSearchUrl},
			cli.StringFlag{Name: "latitude, la", Usage: "Latitude parameter for calculating nearest airport"},
			cli.StringFlag{Name: "longitude, lo", Usage: "Longitude parameter for calculating nearest airport"},
			cli.BoolFlag{Name: "use-ip-address, ip", Usage: "Use IP address to calculate latitude and longitude"},
		},
	}
}
