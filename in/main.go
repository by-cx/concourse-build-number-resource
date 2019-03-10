package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/by-cx/concourse-build-number-resource/common"
)

func main() {
	fmt.Fprintln(os.Stderr, "IN")
	fmt.Fprintln(os.Stderr, ".. processing output directory")
	var directory string
	if len(os.Args) > 1 {
		directory = os.Args[1]
	} else {
		fmt.Fprintln(os.Stderr, "Directory not found")
		os.Exit(1)
	}

	// Read input from Concourse
	fmt.Fprintln(os.Stderr, "..loading input data from stdin")
	storage, err := common.Load(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, "..writing the build number")
	f, err := os.OpenFile(path.Join(directory, "build-number"), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	_, err = f.Write([]byte(storage.Version.BuildNumber))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Return saved build number
	fmt.Fprintln(os.Stderr, "..handling output to Concourse")
	buildNumber := storage.Version.BuildNumber
	response := &common.InOut{
		Version: common.Version{BuildNumber: buildNumber},
		Metadata: []common.MetadataField{
			{
				Name:  "ver",
				Value: buildNumber,
			},
		},
	}

	json.NewEncoder(os.Stdout).Encode(response)
}
