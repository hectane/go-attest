package attest

import (
	"errors"
	"reflect"
)

// Ensure that a value is able to be immediately sent on the specified channel.
func ChanSend(c interface{}, v interface{}) error {
	cases := []reflect.SelectCase{
		{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(c),
			Send: reflect.ValueOf(v),
		},
		{
			Dir: reflect.SelectDefault,
		},
	}
	i, _, _ := reflect.Select(cases)
	if i != 0 {
		return errors.New("sending on channel failed")
	}
	return nil
}

// Ensure that a value is able to be immediately received on the specified
// channel. If a value is received, it is returned.
func ChanRecv(c interface{}) (interface{}, error) {
	cases := []reflect.SelectCase{
		{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(c),
		},
		{
			Dir: reflect.SelectDefault,
		},
	}
	_, v, ok := reflect.Select(cases)
	if !ok {
		return nil, errors.New("receiving on channel failed")
	}
	return v, nil
}

// Ensure that the channel is closed.
func ChanClosed(c interface{}) error {
	cases := []reflect.SelectCase{
		{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(c),
		},
		{
			Dir: reflect.SelectDefault,
		},
	}
	i, _, ok := reflect.Select(cases)
	if i != 0 || ok {
		return errors.New("channel is not closed")
	}
	return nil
}
