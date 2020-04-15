package main

import (
	"flag"
	"fmt"
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
		os.Exit(1)
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

	fmt.Printf("http://%s:%s\n", flagHost, flagPort)
	fmt.Println("__Resources__")
	server.Parse(reader.ReadJSON(filePath[0]))

	server.Serve(flagHost, flagPort)
}
