// Implemented according to
// https://github.com/docker/libnetwork/blob/master/docs/remote.md

package main

import "github.com/docker/go-plugins-helpers/network"

type dummyNetworkDriver struct{}

func main() {
	d := dummyNetworkDriver{}
	h := network.NewHandler(d)
	h.ServeTCP("test_network", ":8080", nil)
}

func (dummyNetworkDriver) GetCapabilities() (*network.CapabilitiesResponse, error) {
	var r network.CapabilitiesResponse
	r.Scope = network.GlobalScope
	return &r, nil
}

func (dummyNetworkDriver) CreateNetwork(*network.CreateNetworkRequest) error {
	return nil
}

func (dummyNetworkDriver) AllocateNetwork(*network.AllocateNetworkRequest) (*network.AllocateNetworkResponse, error) {
	var r network.AllocateNetworkResponse
	return &r, nil
}

func (dummyNetworkDriver) DeleteNetwork(*network.DeleteNetworkRequest) error {
	return nil
}

func (dummyNetworkDriver) FreeNetwork(*network.FreeNetworkRequest) error {
	return nil
}

func (dummyNetworkDriver) CreateEndpoint(*network.CreateEndpointRequest) (*network.CreateEndpointResponse, error) {
	var r network.CreateEndpointResponse
	return &r, nil
}

func (dummyNetworkDriver) DeleteEndpoint(*network.DeleteEndpointRequest) error {
	return nil
}

func (dummyNetworkDriver) EndpointInfo(*network.InfoRequest) (*network.InfoResponse, error) {
	var r network.InfoResponse
	return &r, nil
}

func (dummyNetworkDriver) Join(*network.JoinRequest) (*network.JoinResponse, error) {
	var r network.JoinResponse
	return &r, nil
}

func (dummyNetworkDriver) Leave(*network.LeaveRequest) error {
	return nil
}

func (dummyNetworkDriver) DiscoverNew(*network.DiscoveryNotification) error {
	return nil
}

func (dummyNetworkDriver) DiscoverDelete(*network.DiscoveryNotification) error {
	return nil
}

func (dummyNetworkDriver) ProgramExternalConnectivity(*network.ProgramExternalConnectivityRequest) error {
	return nil
}

func (dummyNetworkDriver) RevokeExternalConnectivity(*network.RevokeExternalConnectivityRequest) error {
	return nil
}
