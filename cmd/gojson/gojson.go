package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"gojson/internal/common"
	"gojson/internal/reader"
	"gojson/internal/server"
)

const helpText = `gojson [OPTIONS] <source>
Serve give json file with first child URL seperation.

Options:
  --port <3000>
    Set port
  --host <localhost>
    Set host

  -v, --version
    Show version number
  -h, --help
	Show help

  Examples:
	gojson db.json
`

func usage() {
	fmt.Println(helpText)
	os.Exit(0)
}

var (
	flagVersion        bool
	flagPort, flagHost string
)

func flagParse() []string {
	flag.Usage = usage

	flag.BoolVar(&flagVersion, "v", false, "")
	flag.BoolVar(&flagVersion, "version", false, "")

	flag.StringVar(&flagPort, "port", "3000", "")
	flag.StringVar(&flagHost, "host", "localhost", "")

	flag.Parse()

	return flag.Args()
}

func signalCheck() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs,
		syscall.SIGTERM,
		syscall.SIGINT,
	)

	go func() {
		<-sigs
		fmt.Println(`Byeee..`)
		if err := server.SRV.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		server.Close()
	}()
}

func main() {
	signalCheck()
	filePath := flagParse()

	if flagVersion {
		fmt.Println(common.Version)
		os.Exit(0)
	}

	if len(filePath) == 0 {
		fmt.Println("Give a json file!")
		os.Exit(2)
	}

	fmt.Println(common.Intro)
	fmt.Printf("\033[90mLoading %s\033[0m\n", filePath[0])
	reader.ReadJSON(filePath[0])
	resList := server.SetHandle()
	fmt.Println("\033[90mDone\033[0m")

	home := fmt.Sprintf("http://%s:%s", flagHost, flagPort)
	fmt.Println("\n\033[1mResources\033[0m")
	for _, val := range resList {
		fmt.Printf("%s/%s\n", home, val)
	}

	fmt.Println("\n\033[1mHome\033[0m")
	fmt.Println(home)

	go reader.ReadKey()

	server.Serve(flagHost, flagPort)
}
