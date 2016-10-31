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
	var c network.CapabilitiesResponse
	c.Scope = network.GlobalScope
	return &c, nil
}

func (dummyNetworkDriver) CreateNetwork(*network.CreateNetworkRequest) error {
	return nil
}

func (dummyNetworkDriver) AllocateNetwork(*network.AllocateNetworkRequest) (*network.AllocateNetworkResponse, error) {
	return nil, nil
}

func (dummyNetworkDriver) DeleteNetwork(*network.DeleteNetworkRequest) error {
	return nil
}

func (dummyNetworkDriver) FreeNetwork(*network.FreeNetworkRequest) error {
	return nil
}

func (dummyNetworkDriver) CreateEndpoint(*network.CreateEndpointRequest) (*network.CreateEndpointResponse, error) {
	return nil, nil
}

func (dummyNetworkDriver) DeleteEndpoint(*network.DeleteEndpointRequest) error {
	return nil
}

func (dummyNetworkDriver) EndpointInfo(*network.InfoRequest) (*network.InfoResponse, error) {
	return nil, nil
}

func (dummyNetworkDriver) Join(*network.JoinRequest) (*network.JoinResponse, error) {
	return nil, nil
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
