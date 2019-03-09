// Package common contains common code for all three Concourse commands (check, in, out).
package common

import "testing"

var TestSource = SourceS3{
	Endpoint:     "xxx",
	AccessKey:    "xxx",
	SecretKey:    "xxx",
	UseSSL:       true,
	Bucket:       "xxx",
	Project:      "test",
	InitialValue: "1",
}

func TestBuildNumberStorage_Reset(t *testing.T) {
	type fields struct {
		Source      SourceS3
		Backend     BackendInterface
		buildNumber int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"Reset test",
			fields{
				Source:  TestSource,
				Backend: &DummyBackend{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BuildNumberStorage{
				Source:      tt.fields.Source,
				Backend:     tt.fields.Backend,
				buildNumber: tt.fields.buildNumber,
			}
			if err := b.Reset(); (err != nil) != tt.wantErr {
				t.Errorf("BuildNumberStorage.Reset() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBuildNumberStorage_Get(t *testing.T) {
	type fields struct {
		Source      SourceS3
		Version     string
		Backend     BackendInterface
		buildNumber int
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			"Get test",
			fields{
				Source:  TestSource,
				Backend: &DummyBackend{BuildNumber: 123},
			},
			123,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BuildNumberStorage{
				Source:      tt.fields.Source,
				Backend:     tt.fields.Backend,
				buildNumber: tt.fields.buildNumber,
			}
			got, err := b.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildNumberStorage.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BuildNumberStorage.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildNumberStorage_Bump(t *testing.T) {
	type fields struct {
		Source      SourceS3
		Version     string
		Backend     BackendInterface
		buildNumber int
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			"Bump test",
			fields{
				Source:  TestSource,
				Backend: &DummyBackend{BuildNumber: 2},
			},
			3,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BuildNumberStorage{
				Source:      tt.fields.Source,
				Backend:     tt.fields.Backend,
				buildNumber: tt.fields.buildNumber,
			}
			if err := b.Bump(); (err != nil) != tt.wantErr {
				t.Errorf("BuildNumberStorage.Bump() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := b.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildNumberStorage.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("BuildNumberStorage.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
