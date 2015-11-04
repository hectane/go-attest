package attest

import (
	"net"
	"net/http"
	"net/url"
	"testing"
)

func TestHttpStatusCode(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	var (
		s http.Server
		r = &http.Request{
			Method: "GET",
			URL: &url.URL{
				Scheme: "http",
				Host:   l.Addr().String(),
			},
		}
	)
	go s.Serve(l)
	defer l.Close()
	if err := HttpStatusCode(r, 404); err != nil {
		t.Fatal(err)
	}
}
