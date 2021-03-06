package sdk

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
)

const activatePath = "/Plugin.Activate"

// Handler is the base to create plugin handlers.
// It initializes connections and sockets to listen to.
type Handler struct {
	mux *http.ServeMux
}

// NewHandler creates a new Handler with an http mux.
func NewHandler(manifest string) Handler {
	mux := http.NewServeMux()

	mux.HandleFunc(activatePath, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", DefaultContentTypeV1_1)
		fmt.Fprintln(w, manifest)
	})

	return Handler{mux: mux}
}

// Serve sets up the handler to serve requests on the passed in listener
func (h Handler) Serve(l net.Listener) error {
	server := http.Server{
		Addr:    l.Addr().String(),
		Handler: h.mux,
	}
	return server.Serve(l)
}

// ServeTCP makes the handler to listen for request in a given TCP address.
// It also writes the spec file in the right directory for docker to read.
func (h Handler) ServeTCP(pluginName, addr string, tlsConfig *tls.Config) error {
	return h.listenAndServe(newTCPListener(addr, pluginName, tlsConfig))
}

// ServeUnix makes the handler to listen for requests in a unix socket.
// It also creates the socket file in the right directory for docker to read.
func (h Handler) ServeUnix(systemGroup, addr string) error {
	return h.listenAndServe(newUnixListener(addr, systemGroup))
}

// ServeWindows makes the handler to listen for request in a windows named pipe.
// It also creates the spec file in the right directory for docker to read.
func (h Handler) ServeWindows(addr, pluginName string, pipeConfig *WindowsPipeConfig) error {
	return h.listenAndServe(newWindowsListener(addr, pluginName, pipeConfig))
}

// HandleFunc registers a function to handle a request path with.
func (h Handler) HandleFunc(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	h.mux.HandleFunc(path, fn)
}

type newListenerFunc func() (net.Listener, string, string, error)

func (h Handler) listenAndServe(fn newListenerFunc) error {

	listener, addr, spec, err := fn()

	server := http.Server{
		Addr:    addr,
		Handler: h.mux,
	}

	if spec != "" {
		defer os.Remove(spec)
	}
	if err != nil {
		return err
	}

	return server.Serve(listener)
}
