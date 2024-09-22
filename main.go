package main

import "github-bridge/internal/app"

func main() {
	h := app.NewApp()
	err := h.Run()
	if err != nil {
		panic(err)
	}
}
