package pidfile

import (
	"fmt"
	"io/ioutil"
	"os"
)

// PIDFile holds information about created pidfile.
type PIDFile struct {
	Path string
}

// New creates new pidfile at provided path and returns a PIDFile.
func New(path string) (*PIDFile, error) {
	err := ioutil.WriteFile(path, []byte(fmt.Sprintf("%d", os.Getpid())), 0644)
	if err != nil {
		return nil, err
	}

	return &PIDFile{path}, nil
}

// Remove deletes pidfile.
func (p PIDFile) Remove() error {
	err := os.Remove(p.Path)
	if err != nil {
		return err
	}

	return nil
}
