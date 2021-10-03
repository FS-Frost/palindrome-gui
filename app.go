package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func initApp() {
	app.Route("/", &palindromeCheker{})
	app.RunWhenOnBrowser()
}

func generateStaticWebsite() {
	dir := "docs"
	err := app.GenerateStaticWebsite(dir, &app.Handler{
		Name:        "Palindrome GUI",
		Description: "A palindrome checker!",
		Resources:   app.GitHubPages("palindrome-gui"),
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Static website generated at: %s", dir)
}

func startHttpServer() {
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})

	address := ":3000"
	log.Printf("Listening on http://localhost%s\n", address)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal(err)
	}
}
