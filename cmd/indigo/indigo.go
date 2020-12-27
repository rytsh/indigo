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

	"indigo/internal/common"
	"indigo/internal/reader"
	"indigo/internal/server"
)

const helpText = `indigo [OPTIONS] <file_or_URL>
Generate RestAPI with JSON file and serve folder

Options:
  --port, -P <3000>
    Set port, tool default is '3000'
  --host, -H <localhost>
    Set host, tool default is 'localhost'

  --location <./data.json>
    Change save location
  --api-path <api_url_path>
    Set API path prefix
  --ui-path <ui_url_path>
    Set UI path default '/indigo'

  --folder <./public>
    Serve folder
  --folder-path <folder_path>
    Set Folder path, works with folder option
  --browsable
    Enable folder browsable
  --spa
    Enable SPA mode
  --no-index
    Stop redirect to index

  --no-api
    Close API server, use just serve folder
  --no-ui
    Close UI server

  --auth-basic <username:password>
    Enable basic authentication with username and password

  --no-color
    Disable color output

  -v, --version
    Show version number
  -h, --help
    Show help

  Examples:
    indigo --api-path api/v1 --auth-basic user:pass test/users.json
`

func usage() {
	fmt.Println(helpText)
	os.Exit(0)
}

var (
	flagVersion, flagNoColor                                                   bool
	flagPort, flagHost, flagAPIPath, flagUIPath, flagFolderPath, flagAuthBasic string
)

func flagParse() []string {
	flag.Usage = usage

	flag.BoolVar(&flagVersion, "v", false, "")
	flag.BoolVar(&flagVersion, "version", false, "")

	flag.StringVar(&flagPort, "port", "3000", "")
	flag.StringVar(&flagPort, "P", "3000", "")
	flag.StringVar(&flagHost, "host", "localhost", "")
	flag.StringVar(&flagHost, "H", "localhost", "")

	flag.StringVar(&reader.FPath, "location", "data.json", "")
	flag.StringVar(&flagAPIPath, "api-path", "", "")
	flag.StringVar(&flagUIPath, "ui-path", "indigo", "")
	flag.StringVar(&flagFolderPath, "folder-path", "", "")

	flag.StringVar(&common.StaticFolder, "folder", "", "")
	flag.StringVar(&flagAuthBasic, "auth-basic", "", "")
	flag.BoolVar(&common.StaticBrowsable, "browsable", false, "")
	flag.BoolVar(&common.StaticSPA, "spa", false, "")
	flag.BoolVar(&common.NoIndex, "no-index", false, "")

	flag.BoolVar(&flagNoColor, "no-color", false, "")
	flag.BoolVar(&common.NoAPI, "no-api", false, "")
	flag.BoolVar(&common.NoUI, "no-ui", false, "")

	flag.Parse()

	// Check Values
	if flagVersion {
		fmt.Println(common.Version)
		os.Exit(0)
	}

	// Color disable
	if flagNoColor == true {
		common.DisableColor()
	}

	// Set paths
	common.SetURL(flagAPIPath, false, &common.APIPath)
	common.SetURL(flagUIPath, false, &common.UIPath)
	common.SetURL(flagFolderPath, false, &common.FolderPath)

	// Check folder exist
	if common.StaticFolder != "" && common.FolderExists(common.StaticFolder) == false {
		common.ErrorPrintExit(common.StaticFolder+" folder not exist", 2)
	}

	// Check basic auth user-pass
	if flagAuthBasic != "" {
		if strings.Contains(flagAuthBasic, ":") == false {
			common.ErrorPrintExit("auth-basic must contain ':'", 3)
		} else {
			tmpSplit := strings.Split(flagAuthBasic, ":")
			common.AuthBasic = append(common.AuthBasic, tmpSplit[0], strings.Join(tmpSplit[1:], ""))
		}
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

	signalCheck()

	// Read file
	if common.NoAPI == false {
		if len(filePath) == 0 {
			common.Color["Red"].Println("Not given a json file, using empty", reader.FPath)
			reader.All = nil
		} else {
			var err error
			common.Color["Magenta"].Println("Loading ", filePath[0])
			if reader.IsURL(filePath[0]) {
				err = reader.GetFile(filePath[0])
			} else {
				err = reader.ReadJSON(filePath[0])
			}

			if err != nil {
				common.ErrorPrintExit(err.Error(), 4)
			} else {
				common.Color["Magenta"].Println("Done")
			}
		}
	}

	// Print information
	IPs := common.GetIPs(flagHost)
	if common.NoUI == false {
		common.Color["Bold"].Println("UI_Path: ", common.UIPath)
	}
	if common.NoAPI == false {
		common.Color["Bold"].Println("API_Path: ", common.APIPath)
		common.Color["Bold"].Println("FILE_Path: ", reader.FPath)
	}
	if common.StaticFolder != "" {
		common.Color["Bold"].Println("FOLDER_Path: ", common.FolderPath)
		common.Color["Bold"].Println("Static Folder: ", common.StaticFolder)
		if common.StaticBrowsable {
			common.Color["Bold"].Println("Browsable: ", common.StaticBrowsable)
		}
		if common.StaticSPA {
			common.Color["Bold"].Println("SPA: ", common.StaticSPA)
		}
		if common.NoIndex {
			common.Color["Bold"].Println("No index redirect: ", common.NoIndex)
		}
	}

	for _, IP := range IPs {
		common.Color["Yellow"].Printf(" - http://%s:%s\n", IP, flagPort)
	}

	if len(common.AuthBasic) > 1 {
		common.Color["Blue"].Println("Basic auth activated")
	}

	// listen key input
	go reader.ReadKey()

	// Start Serve
	server.Serve(flagHost, flagPort)
}
