package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/oleewere/cldnt/client"
	"github.com/oleewere/cldnt/model"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

// ListAirportsCommand list airports by distance (in table or JSON format)
func ListAirportsCommand() cli.Command {
	var airportSearchUrl string
	var rows int
	var output string
	return cli.Command{
		Name:  "airports",
		Usage: "Get closest airports",
		Action: func(c *cli.Context) error {
			var location model.Location
			if len(c.String("la")) > 0 && len(c.String("lo")) > 0 {
				lat, latErr := strconv.ParseFloat(c.String("la"), 64)
				if latErr != nil {
					fmt.Println(latErr)
					os.Exit(1)
				}
				lon, lonErr := strconv.ParseFloat(c.String("lo"), 64)
				if lonErr != nil {
					fmt.Println(lonErr)
					os.Exit(1)
				}
				location = model.Location{Latitude: lat, Longitude: lon}
			} else {
				lotLatPair, err := client.CalculateLocationFromIP()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				location = *lotLatPair
			}
			airports, errAirp := client.ListAirportsByDistance(location, rows, airportSearchUrl)
			if errAirp != nil {
				fmt.Println(errAirp)
				os.Exit(1)
			}
			if output == "json" {
				jsonAirportResp, jsonAirportRespErr := json.Marshal(airports)
				if jsonAirportRespErr != nil {
					fmt.Println(jsonAirportRespErr)
					os.Exit(1)
				}
				printJson(jsonAirportResp)
			} else {
				var tableData [][]string
				if airports != nil {
					for _, airport := range airports {
						tableData = append(tableData, []string{airport.AirportFields.Name, fmt.Sprintf("%f", airport.AirportFields.Latitude), fmt.Sprintf("%f", airport.AirportFields.Longitude)})
					}
				}
				tableTitle := fmt.Sprintf("NEAREST AIRPORTS: (lat: %f, lon: %f)", location.Latitude, location.Longitude)
				printTable(tableTitle, []string{"NAME", "LATITUDE", "LONGITUDE"}, tableData)
			}
			return nil
		},
		Flags: []cli.Flag{
			cli.StringFlag{Name: "airport-db-search-url, a", Usage: "Search URL for airport DB", Value: "https://mikerhodes.cloudant.com/airportdb/_design/view1/_search/geo", Destination: &airportSearchUrl},
			cli.StringFlag{Name: "latitude, la", Usage: "Latitude parameter for calculating nearest airports"},
			cli.StringFlag{Name: "longitude, lo", Usage: "Longitude parameter for calculating nearest airports"},
			cli.StringFlag{Name: "output, o", Usage: "Output type (json or table)", Value: "table", Destination: &output},
			cli.IntFlag{Name: "rows, r", Usage: "Maximum number of rows to show.", Value: 10, Destination: &rows},
		},
	}
}

func printTable(title string, headers []string, data [][]string) {
	fmt.Println(title)
	if len(data) > 0 {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(headers)
		for _, v := range data {
			table.Append(v)
		}
		table.Render()
	} else {
		for i := 1; i <= len(title); i++ {
			fmt.Print("-")
		}
		fmt.Println()
		fmt.Println("NO ENTRIES FOUND!")
	}
}

func printJson(b []byte) {
	fmt.Println(formatJson(b).String())
}

func formatJson(b []byte) *bytes.Buffer {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &out
}
