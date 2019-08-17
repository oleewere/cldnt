package cli

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func StartApplication(version string, gitRevString string) {
	app := cli.NewApp()
	app.Name = "cldnt"
	app.Usage = "Tool to get the closest airports"
	app.EnableBashCompletion = true
	app.UsageText = "cldnt command [command options] [arguments...]"
	if len(version) > 0 {
		app.Version = version
	} else {
		app.Version = "0.1.0"
	}
	if len(gitRevString) > 0 {
		app.Version = app.Version + fmt.Sprintf(" (git short hash: %v)", gitRevString)
	}
	app.Email = "oleewere@gmail.com"
	app.Author = "Oliver Mihaly Szabo"
	app.Copyright = "Copyright 2019 Oliver Mihaly Szabo"
	app.Commands = []cli.Command{}
	app.Commands = append(app.Commands, ListAirportsCommand())
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
