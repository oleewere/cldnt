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
		Name:  "airports",
		Usage: "Get closest airports",
		Action: func(c *cli.Context) error {
			if len(c.String("la")) > 0 && len(c.String("lo")) > 0 {

			} else {
				location, err := client.CalculateLocationFromIP()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Println(location)
			}
			return nil
		},
		Flags: []cli.Flag{
			cli.StringFlag{Name: "airport-db-search-url, a", Usage: "Search URL for airport DB", Value: "https://mikerhodes.cloudant.com/airportdb/_design/view1/_search/geo", Destination: &airportSearchUrl},
			cli.StringFlag{Name: "latitude, la", Usage: "Latitude parameter for calculating nearest airports"},
			cli.StringFlag{Name: "longitude, lo", Usage: "Longitude parameter for calculating nearest airports"},
			cli.IntFlag{Name: "rows, r", Usage: "Maximum number of rows to show."},
		},
	}
}
