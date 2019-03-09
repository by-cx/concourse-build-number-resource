// Package common contains common code for all three Concourse commands (check, in, out).
package common

// ObjectName where build number is saved
const ObjectName = "build-number"

// BackendInterface describes interface for storage backend
type BackendInterface interface {
	Delete() error
	Read() (int, error)
	Write(buildNumber int) error
}

// InOut is struct used to format output of IN and OUT commands
type InOut struct {
	Version Version `json:"version"`
}

// Version is used in output and input of the check command. It contains version compatible with Concourse format.
type Version struct {
	BuildNumber int `json:"num"`
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
	InitialValue int    `json:"initial_value" default:"1"` // Build number used when no object found in S3
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
	b.buildNumber = b.Source.InitialValue
	return b.Backend.Write(b.buildNumber)
}

// Get returns build number saved in the storage
func (b *BuildNumberStorage) Get() (int, error) {
	loadedBuildNumber, err := b.Backend.Read()
	b.buildNumber = loadedBuildNumber

	if err != nil {
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
