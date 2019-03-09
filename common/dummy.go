package common

// DummyBackend is used for testing. It follows BackendInterface.
type DummyBackend struct {
	Exists      bool
	Err         error
	BuildNumber int
}

// IsExist return true if the object exists
func (d *DummyBackend) IsExist() (bool, error) {
	return d.Exists, d.Err
}

// Delete removes configured object
func (d *DummyBackend) Delete() error {
	return d.Err
}

// Read returns values saved in the configured object
func (d *DummyBackend) Read() (int, error) {
	return d.BuildNumber, d.Err
}

// Write puts given value into configured object
func (d *DummyBackend) Write(buildNumber int) error {
	d.BuildNumber = buildNumber
	return d.Err
}
