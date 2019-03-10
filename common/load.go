package common

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

const bufferSize = 1024

// Load processes stdin from Concourse
func Load(input *os.File) (*BuildNumberStorage, error) {
	var buffer string

	var storage = &BuildNumberStorage{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		buffer += scanner.Text()
	}
	err := scanner.Err()
	if err != nil {
		return storage, err
	}

	fmt.Fprint(os.Stderr, string(buffer))

	err = json.Unmarshal([]byte(buffer), storage)
	if err != nil {
		return storage, err
	}

	storage.Backend = &S3Backend{
		Source: storage.Source,
	}

	return storage, nil
}
