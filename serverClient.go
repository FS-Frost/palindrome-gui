package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func getServerVersion() (int64, error) {
	url := fmt.Sprintf("http://localhost:%d/version", port)
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(url)

	if err != nil {
		return 0, fmt.Errorf("error: %s", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return 0, fmt.Errorf("error: %s", err)
	}

	jsonResponse := &ResponseGetVersion{}
	err = json.Unmarshal(body, jsonResponse)

	if err != nil {
		return 0, fmt.Errorf("error: %s", err)
	}

	return jsonResponse.Version, nil
}

func startVersionChecking(ctx app.Context) {
	if clientVersion == 0 {
		fmt.Printf("Invalid client version: %d\n", clientVersion)
		return
	}

	for range time.Tick(time.Millisecond * 1000) {
		serverVersion, err := getServerVersion()

		if err != nil {
			handleError("error fetching server version", err)
			continue
		}

		if serverVersion == 0 || serverVersion == clientVersion {
			// fmt.Println("No need to reload")
			continue
		}

		fmt.Printf("New version found (client: %d, server: %d), reloading...\n", clientVersion, serverVersion)
		ctx.Reload()
	}
}
