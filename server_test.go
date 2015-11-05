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
	h.Handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "", http.StatusOK)
	})
	req, err := http.NewRequest("GET", h.Addr(), nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("%d != %d", resp.StatusCode, http.StatusOK)
	}
	h.Close()
	if _, err := http.DefaultClient.Do(req); err == nil {
		t.Fatal("error expected")
	}
}
