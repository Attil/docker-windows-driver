package sockets

import (
	"net"

	"github.com/Microsoft/go-winio"
)

// NewWindowsSocket creates a Windows named pipe with the specified path.
func NewWindowsSocket(addr string) (net.Listener, error) {
	config := winio.PipeConfig {
		SecurityDescriptor: "S:(ML;;NW;;;LW)D:(A;;0x12019f;;;WD)",	// for everyone?
		MessageMode: true,
		InputBufferSize: 4096,
		OutputBufferSize: 4096}
	listener, err := winio.ListenPipe(addr, &config)
	if err != nil {
		return nil, err
	}
	return listener, nil
}