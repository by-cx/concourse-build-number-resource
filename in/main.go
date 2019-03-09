package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"

	"github.com/by-cx/concourse-build-number-resource/common"
)

func main() {
	var directory string
	if len(os.Args) > 1 {
		directory = os.Args[1]
	} else {
		fmt.Fprintln(os.Stderr, "Directory not found")
		os.Exit(1)
	}

	// Read input from Concourse
	storage, err := common.Load(os.Stdin)
	if err != nil && err != io.EOF {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	f, err := os.OpenFile(path.Join(directory, "build-number"), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	_, err = f.Write([]byte(strconv.Itoa(storage.Version.BuildNumber)))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Return saved build number
	buildNumber, err := storage.Get()

	response := &common.InOut{
		Version: common.Version{buildNumber},
	}

	data, err := json.Marshal(response)
	if err != nil && err != io.EOF {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(string(data))
}
