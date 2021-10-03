package main

import (
	"flag"
	"fmt"
	"log"
)

const (
	MODE_SERVER = "server"
	MODE_STATIC = "static"
)

func main() {
	mode := flag.String("mode", MODE_SERVER, fmt.Sprintf(
		"supports: '%s', '%s'. Default: '%s'",
		MODE_STATIC,
		MODE_SERVER,
		MODE_SERVER,
	))

	flag.Parse()
	initApp()

	switch *mode {
	case MODE_SERVER:
		startHttpServer()
	case MODE_STATIC:
		generateStaticWebsite()
	default:
		log.Fatalf("Unsupported mode: %s", *mode)
	}
}
