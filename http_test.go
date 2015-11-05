package attest

import (
	"net/http"
	"testing"
)

func TestHttpStatusCode(t *testing.T) {
	s, err := NewHttpServer()
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	r, err := http.NewRequest("GET", s.Addr(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := HttpStatusCode(r, http.StatusNotFound); err != nil {
		t.Fatal(err)
	}
}
