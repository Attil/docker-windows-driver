// +build !linux,!freebsd

package sdk

import (
	"errors"
	"net"
)

var (
	errOnlySupportedOnLinuxAndFreeBSD = errors.New("unix socket creation is only supported on Linux and FreeBSD")
)

func newUnixListener(pluginName string, group string) (net.Listener, string, error) {
	return nil, "", errOnlySupportedOnLinuxAndFreeBSD
}
