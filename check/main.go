package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/by-cx/concourse-build-number-resource/common"
)

func main() {
	// Read input from Concourse
	storage, err := common.Load(os.Stdin)
	if err != nil && err != io.EOF {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Bump build number if it's required
	if storage.Source.DoBump {
		err = storage.Bump()
		if err != nil && err != io.EOF {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	// Return saved build number
	buildNumber, err := storage.Get()

	response := []common.Version{
		{BuildNumber: buildNumber},
	}

	data, err := json.Marshal(response)
	if err != nil && err != io.EOF {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(string(data))
}
