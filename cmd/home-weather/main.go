package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/home-IoT/jupiter-cli/internal/cli"
	"github.com/home-IoT/jupiter-cli/internal/config"
)

var (
	args = cli.Flags{}

	app = kingpin.New("home-weather", "A command line interface for reading Jupiter temperature/humidity sensor data")

	versionCommand = app.Command("version", "show version")

	// config
	configCommand      = app.Command("config", "config command")
	configSetCommand   = configCommand.Command("set", "set subcommand")
	configGetCommand   = configCommand.Command("get", "get subcommand")
	configResetCommand = configCommand.Command("reset", "reset subcommand")

	// config Jupiter URL
	configSetJupiterURLEntry    = configSetCommand.Command("jupiter", "set jupiter url")
	configSetJupiterURLEntryArg = configSetJupiterURLEntry.Arg("url", "url arg").String()

	configGetJupiterURLEntry = configGetCommand.Command("jupiter", "get jupiter url")
)

func main() {
	initArgs()

	parse, err := app.Parse(os.Args[1:])
	cli.InitLog(*args.DebugFlag)

	cli.Debugf("args <%#v>\n", os.Args[1:])

	switch kingpin.MustParse(parse, err) {

	case versionCommand.FullCommand():
		cli.ShowVersion()

	case configSetJupiterURLEntry.FullCommand():
		config.SetJupiterURL(*configSetJupiterURLEntryArg)

	case configGetJupiterURLEntry.FullCommand():
		fmt.Println(config.GetJupiterURL())

	default:
		fmt.Printf("no matching command found")
		os.Exit(1)
	}
}

func initArgs() {
	args.DebugFlag = app.Flag("debug", "debug information on").Short('D').Bool()
}