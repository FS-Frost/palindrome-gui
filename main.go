package main

import (
	"flag"
	"fmt"
	"log"
)

const (
	modeServer = "server"
	modeStatic = "static"
)

var (
	autoupdateEnabled = false
	isDevEnabled      = false
)

func main() {
	mode := flag.String("mode", modeServer, fmt.Sprintf(
		"supports: '%s', '%s'",
		modeStatic,
		modeServer,
	))

	flag.BoolVar(&isDevEnabled, "dev", false, "turns on dev mode")
	flag.BoolVar(&autoupdateEnabled, "autoupdate", false, "turns on auto-update")
	flag.Parse()
	initApp()

	switch *mode {
	case modeServer:
		startHttpServer()
	case modeStatic:
		generateStaticWebsite()
	default:
		log.Fatalf("Unsupported mode: %s", *mode)
	}
}
