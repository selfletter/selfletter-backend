package main

import (
	"selfletter-backend/initialization"
)

func main() {
	err := initialization.InitializeConfigAndDatabase()
	if err != nil {
		panic(err.Error())
	}

	// todo: figure out a way to restart router on config change
	err = initialization.InitializeRouter()

	if err != nil {
		panic("api: failed to start: " + err.Error())
	}
}
