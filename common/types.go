// Package common contains common code for all three Concourse commands (check, in, out).
package common

import (
	"runtime/debug"
	"strconv"
)

// ObjectName where build number is saved
const ObjectName = "build-number"

// BackendInterface describes interface for storage backend
type BackendInterface interface {
	Delete() error
	Read() (int, error)
	IsExist() (bool, error)
	Write(buildNumber int) error
}

// InOut is struct used to format output of IN and OUT commands
type InOut struct {
	Version  Version           `json:"version"`
	Metadata map[string]string `json:"metadata"`
}

// Version is used in output and input of the check command. It contains version compatible with Concourse format.
type Version struct {
	BuildNumber string `json:"num"`
}

// Params is coming from Concurs in "in" and "out" commands
type Params struct {
}

// SourceS3 contains configuration to access the build number storage
type SourceS3 struct {
	Endpoint     string `json:"endpoint"`                  // Endpoint of the S3 service
	AccessKey    string `json:"access_key"`                // S3 access key
	SecretKey    string `json:"secret_key"`                // S3 secret access key
	Region       string `json:"region"`                    // AWS region
	UseSSL       bool   `json:"use_ssl" default:"true"`    // True of SSL required, by default enabled
	Bucket       string `json:"bucket"`                    // Name of the bucket
	Project      string `json:"project"`                   // Project name - prefix where the object with build number will be saved
	InitialValue string `json:"initial_value" default:"1"` // Build number used when no object found in S3
	DoBump       bool   `json:"bump"`                      // True if the version should be bumped before it's returned
}

// BuildNumberStorage is representation of a storage where the build number is saved
type BuildNumberStorage struct {
	Source  SourceS3 `json:"source"` // Configuration of the S3 bucket
	Backend BackendInterface
	Version Version `json:"version"` // Version coming from Concourse
	Params  Params  `json:"params"`  // Params is coming from Concourse in "in" and "out" commmands

	buildNumber int // Internal build number
}

// Reset sets the version back to b.InitialValue
func (b *BuildNumberStorage) Reset() error {
	initialValue, err := strconv.Atoi(b.Source.InitialValue)
	if err != nil {
		return err
	}

	b.buildNumber = initialValue
	return b.Backend.Write(b.buildNumber)
}

// Get returns build number saved in the storage
func (b *BuildNumberStorage) Get() (int, error) {
	loadedBuildNumber, err := b.Backend.Read()
	b.buildNumber = loadedBuildNumber

	exists, err := b.Backend.IsExist()
	if err != nil {
		debug.PrintStack()
		return b.buildNumber, err
	}
	if !exists {
		buildNumber, err := strconv.Atoi(b.Source.InitialValue)
		if err != nil {
			debug.PrintStack()
			return b.buildNumber, err
		}
		b.buildNumber = buildNumber
		err = b.Backend.Write(buildNumber)
		if err != nil {
			debug.PrintStack()
		}
		return buildNumber, err
	}

	if err != nil {
		debug.PrintStack()
		return -1, err
	}
	return b.buildNumber, nil
}

// Bump increases build number by one
func (b *BuildNumberStorage) Bump() error {
	_, err := b.Get()
	if err != nil {
		return err
	}
	b.buildNumber++
	return b.Backend.Write(b.buildNumber)
}
