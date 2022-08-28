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

func main() {
	mode := flag.String("mode", modeServer, fmt.Sprintf(
		"supports: '%s', '%s'",
		modeStatic,
		modeServer,
	))

	isDevEnabled := flag.Bool("dev", false, "supports: 'true', 'false'. Default: 'false'")

	flag.Parse()
	initApp(*isDevEnabled)

	switch *mode {
	case modeServer:
		startHttpServer()
	case modeStatic:
		generateStaticWebsite()
	default:
		log.Fatalf("Unsupported mode: %s", *mode)
	}
}
