package common

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	// Input
	f, err := ioutil.TempFile("/tmp", "load_test")
	assert.Nil(t, err)
	defer f.Close()

	_, err = f.Write([]byte(`
	{
		"source": {
			"endpoint": "s3.example.com",
			"access_key": "abcdef",
			"secret_key": "ghijkl",
			"use_ssl": true,
			"bucket": "builds.bucket",
			"project": "test",
			"initial_value": "1"
		},
		"version": {
			"num": "1"
		}
	}
	`))
	assert.Nil(t, err)

	_, err = f.Seek(0, 0)
	assert.Nil(t, err)

	// Output
	storage := BuildNumberStorage{
		Source: SourceS3{
			Endpoint:     "s3.example.com",
			AccessKey:    "abcdef",
			SecretKey:    "ghijkl",
			UseSSL:       true,
			Bucket:       "builds.bucket",
			Project:      "test",
			InitialValue: "1",
		},
		Version: Version{
			BuildNumber: "1",
		},
	}

	//Processing
	type args struct {
		input *os.File
	}
	tests := []struct {
		name    string
		args    args
		want    *BuildNumberStorage
		wantErr bool
	}{
		{
			"Load test",
			args{
				input: f,
			},
			&storage,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Load(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want.Backend = got.Backend
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
