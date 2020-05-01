// Package roku provides the ability to query and control roku devices on the network
package roku

import (
	"fmt"
	"github.com/koron/go-ssdp"
	"net/http"
)

// Scan uses ssdp to discover roku devices on the network
func Scan(timeout int) ([]*Device, error) {
	list, err := ssdp.Search("roku:ecp", timeout, "")
	if err != nil {
		return nil, fmt.Errorf("failed to discover roku devices on network: %w", err)
	}

	return ProcessDevices(list, http.DefaultClient), nil
}

// ProcessDevices takes a list of ssdp services and builds an array of
// Device objects.
func ProcessDevices(list []ssdp.Service, client *http.Client) []*Device {
	devices := make([]*Device, 0)
	for _, srv := range list {
		device := Device{}
		device.URL = srv.Location
		device.client = client
		devices = append(devices, &device)
	}
	return devices
}
