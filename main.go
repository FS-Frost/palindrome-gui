package main

import (
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &hello{})
	app.Route("/palindrome", &palindromeCheker{})
	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite("docs", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Resources:   app.GitHubPages("palindrome-gui"),
	})

	if err != nil {
		log.Fatal(err)
	}

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// On the server-side, RunWhenOnBrowser() does nothing, which allows the
	// writing of server logic without needing precompiling instructions.
	// app.RunWhenOnBrowser()

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	// http.Handle("/", &app.Handler{
	// 	Name:        "Hello",
	// 	Description: "An Hello World! example",
	// })

	// address := ":8000"
	// fmt.Printf("Listening on http://localhost%s\n", address)

	// if err := http.ListenAndServe(address, nil); err != nil {
	// 	log.Fatal(err)
	// }
}
