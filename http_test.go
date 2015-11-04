package attest

import (
	"net/http"
	"testing"
)

func TestHttpServer(t *testing.T) {
	h, err := NewHttpServer()
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("GET", h.Addr(), nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Fatal("%d != %d", resp.StatusCode, http.StatusNotFound)
	}
	h.Close()
	if _, err := http.DefaultClient.Do(req); err == nil {
		t.Fatal("error expected")
	}
}
