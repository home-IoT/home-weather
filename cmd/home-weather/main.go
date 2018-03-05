package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/home-IoT/home-weather/internal/cli"
	"github.com/home-IoT/home-weather/internal/config"
	"github.com/home-IoT/home-weather/internal/log"
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

	// list sensors
	listCommand = app.Command("list", "list sensors")

	// read sensors
	readCommand              = app.Command("read", "read sensor(s)")
	readCommandSensorList    = readCommand.Arg("sensor-ids", "comma-separated list of sensors").String()
	readCommandVerbosityFlag = readCommand.Flag("full", "print out details and a summary").Short('f').Bool()
	readCommandAllFlag       = readCommand.Flag("all", "read all sensors").Short('a').Bool()
)

func main() {
	initArgs()

	parse, err := app.Parse(os.Args[1:])
	log.InitLog(*args.DebugFlag)

	log.Debugf("args <%#v>\n", os.Args[1:])

	switch kingpin.MustParse(parse, err) {

	case versionCommand.FullCommand():
		cli.ShowVersion()

	case configSetJupiterURLEntry.FullCommand():
		config.SetJupiterURLString(*configSetJupiterURLEntryArg)

	case configGetJupiterURLEntry.FullCommand():
		fmt.Println(config.GetJupiterURLString())

	case listCommand.FullCommand():
		cli.ListSensors()

	case readCommand.FullCommand():
		cli.ReadSensors(*readCommandSensorList, *readCommandVerbosityFlag, *readCommandAllFlag)

	default:
		fmt.Printf("no matching command found")
		os.Exit(1)
	}
}

func initArgs() {
	args.DebugFlag = app.Flag("debug", "debug information on").Short('D').Bool()
}
