package attest

import (
	"testing"
)

func TestChanSend(t *testing.T) {
	c := make(chan bool, 1)
	if err := ChanSend(c, true); err != nil {
		t.Fatal("unexpected error")
	}
	if err := ChanSend(c, true); err == nil {
		t.Fatal("error expected")
	}
}

func TestChanRecv(t *testing.T) {
	var (
		c    = make(chan bool, 1)
		data = true
	)
	c <- data
	v, err := ChanRecv(c)
	if err != nil {
		t.Fatal("unexpected error")
	}
	if v != data {
		t.Fatalf("%v != %v", v, data)
	}
	if _, err := ChanRecv(c); err == nil {
		t.Fatal("error expected")
	}
	close(c)
	if _, err := ChanRecv(c); err == nil {
		t.Fatal("error expected")
	}
}

func TestChanClosed(t *testing.T) {
	c := make(chan bool)
	if err := ChanClosed(c); err == nil {
		t.Fatal("error expected")
	}
	close(c)
	if err := ChanClosed(c); err != nil {
		t.Fatal("unexpected error")
	}
}
