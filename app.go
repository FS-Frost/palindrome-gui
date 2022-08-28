package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	port = 3000
)

var (
	clientVersion int64
)

func initApp() {
	app.Route("/", &palindromeCheker{})

	if isRunningOnBrowser() && _isDevEnabled() {
		var err error
		clientVersion, err = getServerVersion()

		if err != nil {
			handleError("error fetching server version", err)
			return
		}
	}

	app.RunWhenOnBrowser()
}

func isRunningOnBrowser() bool {
	return app.Getenv("GOAPP_VERSION") != ""
}

func _isDevEnabled() bool {
	return strings.ToLower(app.Getenv("DEV_ENABLED")) == "true"
}

// func assetUrl(url string) string {
// 	return url
// 	// if _isDevEnabled() {
// 	// }

// 	// return "https://fs-frost.github.io/palindrome-gui/" + url
// }

func appHandler() *app.Handler {
	return &app.Handler{
		Name:        "Palindrome GUI",
		Description: "A palindrome checker!",
		Styles: []string{
			"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css",
			"web/main.css",
		},
		Env: app.Environment{
			"AUTO_UPDATE_ENABLED": fmt.Sprintf("%t", autoupdateEnabled),
			"DEV_ENABLED":         fmt.Sprintf("%t", isDevEnabled),
		},
	}
}

func generateStaticWebsite() {
	dir := "docs"
	handler := appHandler()
	handler.Resources = app.GitHubPages("palindrome-gui")
	err := app.GenerateStaticWebsite(dir, handler)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Static website generated at: %s", dir)
}

func startHttpServer() {
	handler := appHandler()
	http.Handle("/", handler)
	http.Handle("/version", http.HandlerFunc(handleGetVersion))

	address := fmt.Sprintf(":%d", port)
	log.Printf("Listening on http://localhost%s\n", address)
	clientVersion = time.Now().Unix()

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal(err)
	}
}

func handleError(msg string, err error) {
	fmt.Printf("%s: %s\n", msg, err)
}

type ResponseGetVersion struct {
	Version int64
}

func handleGetVersion(w http.ResponseWriter, req *http.Request) {
	response := &ResponseGetVersion{
		Version: clientVersion,
	}

	jsonResponse, err := json.Marshal(&response)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Fprintf(w, "%s\n", string(jsonResponse))
}
