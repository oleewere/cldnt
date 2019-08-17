package main

import (
	cldntcli "github.com/oleewere/cldnt/cli"
)

// Version that will be generated during the build as a constant
var Version string

// GitRevString that will be generated during the build as a constant - represents git revision value
var GitRevString string

func main() {
	cldntcli.StartApplication(Version, GitRevString)
}
