package common

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/minio/minio-go"
)

// S3Backend covers handling of S3 storage
type S3Backend struct {
	Source SourceS3
}

func (s *S3Backend) getClient() (*minio.Client, error) {
	var client *minio.Client
	var err error

	// Initialize minio client object.
	if s.Source.Region == "" {
		client, err = minio.New(s.Source.Endpoint, s.Source.AccessKey, s.Source.SecretKey, s.Source.UseSSL)
	} else {
		client, err = minio.NewWithRegion(s.Source.Endpoint, s.Source.AccessKey, s.Source.SecretKey, s.Source.UseSSL, s.Source.Region)
	}

	return client, err
}

// IsExist return true if the object exists
func (s *S3Backend) IsExist() (bool, error) {
	client, err := s.getClient()
	if err != nil {
		return false, err
	}

	_, err = client.StatObject(s.Source.Bucket, s.Source.Project+"/"+ObjectName, minio.StatObjectOptions{})
	if err != nil {
		if err.Error() == "The specified key does not exist." {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// Delete removes configured object
func (s *S3Backend) Delete() error {
	client, err := s.getClient()
	if err != nil {
		return err
	}

	exists, err := s.IsExist()
	if err != nil {
		return err
	}

	if exists {
		err = client.RemoveObject(s.Source.Bucket, s.Source.Project+"/"+ObjectName)
		if err != nil {
			return err
		}
	}

	return nil
}

// Read returns values saved in the configured object
func (s *S3Backend) Read() (int, error) {
	client, err := s.getClient()
	if err != nil {
		return -1, err
	}

	object, err := client.GetObject(s.Source.Bucket, s.Source.Project+"/"+ObjectName, minio.GetObjectOptions{})
	if err != nil {
		return -1, err
	}

	data, err := ioutil.ReadAll(object)
	if err != nil {
		return -1, err
	}

	buildNumberString := strings.TrimSpace(string(data))
	buildNumber, err := strconv.Atoi(buildNumberString)
	return buildNumber, err
}

// Write puts given value into configured object
func (s *S3Backend) Write(buildNumber int) error {
	client, err := s.getClient()
	if err != nil {
		return err
	}

	buildNumberReader := strings.NewReader(strconv.Itoa(buildNumber))

	_, err = client.PutObject(
		s.Source.Bucket,
		s.Source.Project+"/"+ObjectName,
		buildNumberReader,
		buildNumberReader.Size(),
		minio.PutObjectOptions{ContentType: "application/octet-stream"},
	)

	return err
}
