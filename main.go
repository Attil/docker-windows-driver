// Implemented according to
// https://github.com/docker/libnetwork/blob/master/docs/remote.md

package main

import "net/http"

type HandlerType func(http.ResponseWriter, *http.Request)

func main() {

	routeHandlers := map[string]HandlerType{
		"/Plugin.Activate":                activateHandler,
		"/NetworkDriver.GetCapabilities":  getCapabilitiesHandler,
		"/NetworkDriver.CreateNetwork":    nullResponseHandler,
		"/NetworkDriver.DeleteNetwork":    nullResponseHandler,
		"/NetworkDriver.CreateEndpoint":   createEndpointHandler,
		"/NetworkDriver.EndpointOperInfo": endpointOperInfoHandler,
		"/NetworkDriver.DeleteEndpoint":   nullResponseHandler,
		"/NetworkDriver.Join":             joinHandler,
		"/NetworkDriver.Leave":            nullResponseHandler,
		"/NetworkDriver.DiscoverNew":      nullResponseHandler,
		"/NetworkDriver.DiscoverDelete":   nullResponseHandler,
	}

	for route, handler := range routeHandlers {
		http.HandleFunc(route, handler)
	}
	http.ListenAndServe(":8080", nil)
}

func activateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"Implements": ["NetworkDriver"]}`))
}

func getCapabilitiesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"Scope": "global"}`))
}

func nullResponseHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{}`))
}

func createEndpointHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		"Interface": {
        "Address": "1.2.3.4",
        "MacAddress": "AA:BB:CC:DD:EE:FF"
    	}
	`))
}

func endpointOperInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		{
			"Value": {}
		}
	`))
}

func joinHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		{
			"InterfaceName": {
				SrcName: "dummy0",
				DstPrefix: "wut"
			},
			"StaticRoutes": [{
				"Destination": "asdf",
				"RouteType": 1,
			}]
		}
	`))
}
