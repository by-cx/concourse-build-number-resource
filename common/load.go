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
	var buffer = make([]byte, bufferSize)

	var storage = &BuildNumberStorage{}

	reader := bufio.NewReader(input)

	n, err := reader.Read(buffer)
	if err != nil {
		return storage, err
	}
	fmt.Println(string(buffer))

	err = json.Unmarshal(buffer[0:n], storage)
	if err != nil {
		return storage, err
	}

	storage.Backend = &S3Backend{
		Source: storage.Source,
	}

	return storage, nil
}
