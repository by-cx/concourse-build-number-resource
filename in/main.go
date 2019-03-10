package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strconv"

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

	buildNumber := storage.Version.BuildNumber

	// In case bump is needed, do a bump :)
	if storage.Params.DoBump {
		err = storage.Bump()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		number, err := storage.Get()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		buildNumber = strconv.Itoa(number)
	} else { // and if not we just get the value from the storage
		number, err := storage.Get()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		buildNumber = strconv.Itoa(number)
	}

	// Writing the build number into a file
	fmt.Fprintln(os.Stderr, "..writing the build number")
	f, err := os.OpenFile(path.Join(directory, "build-number"), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	_, err = f.Write([]byte(buildNumber))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Return saved build number
	fmt.Fprintln(os.Stderr, "..handling output to Concourse")
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
