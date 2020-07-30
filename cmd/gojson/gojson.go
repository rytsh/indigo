package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"gojson/internal/common"
	"gojson/internal/reader"
	"gojson/internal/server"
)

const helpText = `gojson [OPTIONS] <source>
Generate RestAPI with JSON file and serve folder

Options:
  --port, -P <3000>
    Set port, tool default is '3000'
  --host, -H <localhost>
    Set host, tool default is 'localhost'

  --empty <data.json>
    If not given any json file, this is the save location
    Tool default is 'data.json'
  --api <api>
    Start API url with this string
  --folder <./public>
    Serve folder
    if api option is empty, auto set 'api'

  --auth-basic <username:password>
    Enable basic authentication with username and password

  --no-color
    Disable color output

  -v, --version
    Show version number
  -h, --help
    Show help

  Examples:
    gojson --api api/v1 --folder /server/public --auth-basic user:pass db.json
`

func usage() {
	fmt.Println(helpText)
	os.Exit(0)
}

var (
	flagVersion, flagNoColor bool
	flagPort, flagHost       string
)

func flagParse() []string {
	flag.Usage = usage

	flag.BoolVar(&flagVersion, "v", false, "")
	flag.BoolVar(&flagVersion, "version", false, "")

	flag.StringVar(&flagPort, "port", "3000", "")
	flag.StringVar(&flagPort, "P", "3000", "")
	flag.StringVar(&flagHost, "host", "localhost", "")
	flag.StringVar(&flagHost, "H", "localhost", "")

	flag.StringVar(&reader.FPath, "empty", "data.json", "")
	flag.StringVar(&common.API, "api", "", "")
	flag.StringVar(&common.StaticFolder, "folder", "", "")
	flag.StringVar(&common.AuthBasic, "auth-basic", "", "")

	flag.BoolVar(&flagNoColor, "no-color", false, "")

	flag.Parse()

	// Check Values
	if flagVersion {
		fmt.Println(common.Version)
		os.Exit(0)
	}

	// color disable
	if flagNoColor == true {
		common.DisableColor()
	}

	// API Trim
	common.API = strings.Trim(common.API, "/ ")
	if common.StaticFolder != "" && common.API == "" {
		common.API = "api"
	}

	return flag.Args()
}

func exit(out int) {
	if err := server.SRV.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		log.Printf("HTTP server Shutdown: %v", err)
	}
	server.Close()
	if out >= 0 {
		common.Color["Reset"].Print("\n")
		os.Exit(out)
	}
}

func signalCheck() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs,
		syscall.SIGTERM,
		syscall.SIGINT,
	)

	go func() {
		<-sigs
		exit(-1)
		fmt.Println(`Byeee..`)
	}()
}

func main() {
	common.SetCustomLog()

	filePath := flagParse()
	common.PrintIntro()

	if common.StaticFolder != "" && common.FolderExists(common.StaticFolder) == false {
		common.ErrorPrintExit(common.StaticFolder+" folder not exist", 2)
	}
	if common.AuthBasic != "" && strings.Contains(common.AuthBasic, ":") == false {
		common.ErrorPrintExit("auth-basic must contain ':'", 3)
	}

	signalCheck()

	if len(filePath) == 0 {
		common.Color["Red"].Println("Not given a json file, using empty", reader.FPath)
	} else {
		common.Color["Magenta"].Println("Loading ", filePath[0])
		if err := reader.ReadJSON(filePath[0]); err != nil {
			common.ErrorPrintExit(err.Error(), 4)
		}
		common.Color["Magenta"].Println("Done")
	}

	// Start Serve
	server.SetHandle()

	common.Color["Bold"].Println("API: ", reader.FPath)
	common.Color["Yellow"].Printf("http://%s:%s/%s\n", flagHost, flagPort, common.API)
	if common.StaticFolder != "" {
		common.Color["Bold"].Println("Static Folder: ", common.StaticFolder)
		common.Color["Yellow"].Printf("http://%s:%s\n", flagHost, flagPort)
	}

	if common.AuthBasic != "" {
		common.Color["Blue"].Println("Basic auth activated")
	}

	go reader.ReadKey()

	server.Serve(flagHost, flagPort)
}
