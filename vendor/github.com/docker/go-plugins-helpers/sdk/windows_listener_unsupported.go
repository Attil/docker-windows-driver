// +build !windows

package sdk

import (
	"errors"
	"net"
)

var (
	errOnlySupportedOnWindows = errors.New("named pipe creation is only supported on Windows")
)

func newWindowsListener(_, _ string, _ *WindowsPipeConfig) (net.Listener, string, error) {
	return nil, "", errOnlySupportedOnWindows
}
