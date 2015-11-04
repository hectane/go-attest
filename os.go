package attest

import (
	"errors"
	"os"
)

// Ensure that the file is in the specified state of existence.
func FileState(name string, exists bool) error {
	if _, err := os.Stat(name); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		if exists {
			return errors.New("file does not exist")
		}
	} else {
		if !exists {
			return errors.New("file exists")
		}
	}
	return nil
}
