package attest

import (
	"fmt"
	"net"
	"net/http"
)

// Basic HTTP server.
type HttpServer struct {
	Handler  *http.ServeMux
	listener net.Listener
	server   http.Server
	stopped  chan bool
}

// Create a new HTTP server for testing purposes. To add a handler for a path,
// use the Handler field. The server will continue to run until the Close()
// method is invoked.
func NewHttpServer() (*HttpServer, error) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, err
	}
	var (
		s = http.NewServeMux()
		c = make(chan bool)
		h = &HttpServer{
			Handler:  s,
			listener: l,
			server: http.Server{
				Handler: s,
			},
			stopped: c,
		}
	)
	go func() {
		h.server.Serve(l)
		close(c)
	}()
	return h, nil
}

// Retrieve the address of the server. The string will be in the form
// "http://host:port".
func (h *HttpServer) Addr() string {
	return fmt.Sprintf("http://%s", h.listener.Addr())
}

// Close the server.
func (h *HttpServer) Close() {
	h.listener.Close()
	<-h.stopped
}
