package sdk

import (
	"io/ioutil"
	"net"
	"os"
	"path/filepath"

	"github.com/docker/go-connections/sockets"
)


var windowsPluginSpecDir = ([]string{filepath.Join(os.Getenv("programdata"), "docker", "plugins")})[0]

// TODO: groups
func newWindowsListener(address string, pluginName string) (net.Listener, string, error) {
	listener, err := sockets.NewWindowsSocket(address)
	if err != nil {
		return nil, "", err
	}
	spec, err := writeWindowsSpec(pluginName, listener.Addr().String())
	if err != nil {
		return nil, "", err
	}
	return listener, spec, nil
}

func writeWindowsSpec(name string, address string) (string, error) {
	if err := os.MkdirAll(windowsPluginSpecDir, 0755); err != nil {
		return "", err
	}
	spec := filepath.Join(windowsPluginSpecDir, name+".spec")
	url := "npipe://" + address
	if err := ioutil.WriteFile(spec, []byte(url), 0644); err != nil {
		return "", err
	}
	return spec, nil
}
