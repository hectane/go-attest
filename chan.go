package attest

import (
	"errors"
)

// Ensure that a value is able to be immediately sent on the specified channel.
func ChanSend(c chan<- interface{}, v interface{}) error {
	select {
	case c <- v:
		return nil
	default:
		return errors.New("sending on channel failed")
	}
}

// Ensure that a value is able to be immediately received on the specified
// channel. If a value is received, it is returned.
func ChanRecv(c <-chan interface{}) (interface{}, error) {
	select {
	case v := <-c:
		return v, nil
	default:
		return nil, errors.New("receiving on channel failed")
	}
}

// Ensure that the channel is closed.
func ChanClosed(c <-chan interface{}) error {
	select {
	case _, ok := <-c:
		if !ok {
			return nil
		}
	default:
	}
	return errors.New("channel is not closed")
}
