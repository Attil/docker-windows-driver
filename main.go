// Implemented according to
// https://github.com/docker/libnetwork/blob/master/docs/remote.md

package main

import (
	"fmt"
	"net/http"

	"github.com/docker/go-plugins-helpers/network"
	"github.com/Microsoft/go-winio"
)

type dummyNetworkDriver struct{}

func main() {
	d := dummyNetworkDriver{}
	h := network.NewHandler(d)

	config := winio.PipeConfig {
		SecurityDescriptor: "S:(ML;;NW;;;LW)D:(A;;0x12019f;;;WD)",	// for everyone?
		MessageMode: true,
		InputBufferSize: 4096,
		OutputBufferSize: 4096}
	listener, err := winio.ListenPipe("\\\\.\\pipe\\driverpipe", config)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	server := http.Server{
		Addr:    addr,
		Handler: h.mux,
	}
	return server.Serve(listener)
}

func (dummyNetworkDriver) GetCapabilities() (*network.CapabilitiesResponse, error) {
	fmt.Println("GetCapabilities")
	r := &network.CapabilitiesResponse{}
	r.Scope = network.LocalScope
	return r, nil
}

func (dummyNetworkDriver) CreateNetwork(*network.CreateNetworkRequest) error {
	fmt.Println("CreateNetwork")
	return nil
}

func (dummyNetworkDriver) AllocateNetwork(*network.AllocateNetworkRequest) (*network.AllocateNetworkResponse, error) {
	fmt.Println("AllocateNetwork")
	r := &network.AllocateNetworkResponse{}
	return r, nil
}

func (dummyNetworkDriver) DeleteNetwork(*network.DeleteNetworkRequest) error {
	fmt.Println("DeleteNetwork")
	return nil
}

func (dummyNetworkDriver) FreeNetwork(*network.FreeNetworkRequest) error {
	fmt.Println("FreeNetwork")
	return nil
}

func (dummyNetworkDriver) CreateEndpoint(*network.CreateEndpointRequest) (*network.CreateEndpointResponse, error) {
	fmt.Println("CreateEndpoint")
	r := &network.CreateEndpointResponse{}
	return r, nil
}

func (dummyNetworkDriver) DeleteEndpoint(*network.DeleteEndpointRequest) error {
	fmt.Println("DeleteEndpoint")
	return nil
}

func (dummyNetworkDriver) EndpointInfo(*network.InfoRequest) (*network.InfoResponse, error) {
	fmt.Println("EndpointInfo")
	r := &network.InfoResponse{}
	return r, nil
}

func (dummyNetworkDriver) Join(*network.JoinRequest) (*network.JoinResponse, error) {
	fmt.Println("Join")
	r := &network.JoinResponse{}
	return r, nil
}

func (dummyNetworkDriver) Leave(*network.LeaveRequest) error {
	fmt.Println("Leave")
	return nil
}

func (dummyNetworkDriver) DiscoverNew(*network.DiscoveryNotification) error {
	fmt.Println("DiscoverNew")
	return nil
}

func (dummyNetworkDriver) DiscoverDelete(*network.DiscoveryNotification) error {
	fmt.Println("DiscoverDelete")
	return nil
}

func (dummyNetworkDriver) ProgramExternalConnectivity(*network.ProgramExternalConnectivityRequest) error {
	fmt.Println("ProgramExternalConnectivity")
	return nil
}

func (dummyNetworkDriver) RevokeExternalConnectivity(*network.RevokeExternalConnectivityRequest) error {
	fmt.Println("RevokeExternalConnectivity")
	return nil
}
