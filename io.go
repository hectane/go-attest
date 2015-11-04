package attest

import (
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
)

func Read(r io.Reader, data []byte) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if !reflect.DeepEqual(b, data) {
		return fmt.Errorf("%v != %v", b, data)
	}
	return nil
}
