// Package client contains the implementation of CRUD operations on MAAS resources.
package client

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/canonical/gomaasclient/api"
)

// GetTLSClient creates a Client configured with TLS
func GetTLSClient(apiURL string, apiKey string, apiVersion string, tlsConfig *tls.Config) (*Client, error) {
	var tr http.RoundTripper

	if tlsConfig != nil {
		val, ok := http.DefaultTransport.(*http.Transport)
		if !ok {
			return nil, fmt.Errorf("unexpected error")
		}

		defaultTransportCopy := val.Clone()
		defaultTransportCopy.TLSClientConfig = tlsConfig
		tr = defaultTransportCopy
	}

	apiClient, err := getAPIClient(apiURL, apiKey, apiVersion, tr)

	if err != nil {
		return nil, err
	}

	return constructClient(apiClient), nil
}

// GetClient creates a client
func GetClient(apiURL string, apiKey string, apiVersion string) (*Client, error) {
	apiClient, err := getAPIClient(apiURL, apiKey, apiVersion, nil)
	if err != nil {
		return nil, err
	}

	return constructClient(apiClient), nil
}

// GetClientWithTransport creates a Client configured with the specified http.Transport
func GetClientWithTransport(apiURL string, apiKey string, apiVersion string, tr http.RoundTripper) (*Client, error) {
	apiClient, err := getAPIClient(apiURL, apiKey, apiVersion, tr)
	if err != nil {
		return nil, err
	}

	return constructClient(apiClient), nil
}

func constructClient(apiClient *APIClient) *Client {
	client := Client{
		Device:                &Device{APIClient: *apiClient},
		Devices:               &Devices{APIClient: *apiClient},
		Domain:                &Domain{APIClient: *apiClient},
		Domains:               &Domains{APIClient: *apiClient},
		DNSResource:           &DNSResource{APIClient: *apiClient},
		DNSResources:          &DNSResources{APIClient: *apiClient},
		DNSResourceRecord:     &DNSResourceRecord{APIClient: *apiClient},
		DNSResourceRecords:    &DNSResourceRecords{APIClient: *apiClient},
		Events:                &Events{APIClient: *apiClient},
		Fabric:                &Fabric{APIClient: *apiClient},
		Fabrics:               &Fabrics{APIClient: *apiClient},
		VLAN:                  &VLAN{APIClient: *apiClient},
		VLANs:                 &VLANs{APIClient: *apiClient},
		Space:                 &Space{APIClient: *apiClient},
		Spaces:                &Spaces{APIClient: *apiClient},
		Machine:               &Machine{APIClient: *apiClient},
		Machines:              &Machines{APIClient: *apiClient},
		VMHost:                &VMHost{APIClient: *apiClient},
		VMHosts:               &VMHosts{APIClient: *apiClient},
		NetworkInterface:      &NetworkInterface{APIClient: *apiClient},
		NetworkInterfaces:     &NetworkInterfaces{APIClient: *apiClient},
		NodeDevice:            &NodeDevice{APIClient: *apiClient},
		NodeDevices:           &NodeDevices{APIClient: *apiClient},
		RAID:                  &RAID{APIClient: *apiClient},
		RAIDs:                 &RAIDs{APIClient: *apiClient},
		Subnet:                &Subnet{APIClient: *apiClient},
		Subnets:               &Subnets{APIClient: *apiClient},
		IPRange:               &IPRange{APIClient: *apiClient},
		IPRanges:              &IPRanges{APIClient: *apiClient},
		IPAddresses:           &IPAddresses{APIClient: *apiClient},
		Tag:                   &Tag{APIClient: *apiClient},
		Tags:                  &Tags{APIClient: *apiClient},
		BlockDevice:           &BlockDevice{APIClient: *apiClient},
		BlockDevices:          &BlockDevices{APIClient: *apiClient},
		BlockDevicePartition:  &BlockDevicePartition{APIClient: *apiClient},
		BlockDevicePartitions: &BlockDevicePartitions{APIClient: *apiClient},
		User:                  &User{APIClient: *apiClient},
		Users:                 &Users{APIClient: *apiClient},
		ResourcePool:          &ResourcePool{APIClient: *apiClient},
		ResourcePools:         &ResourcePools{APIClient: *apiClient},
		MAASServer:            &MAASServer{APIClient: *apiClient},
		PackageRepository:     &PackageRepository{APIClient: *apiClient},
		PackageRepositories:   &PackageRepositories{APIClient: *apiClient},
		BootSource:            &BootSource{APIClient: *apiClient},
		BootSources:           &BootSources{APIClient: *apiClient},
		BootSourceSelection:   &BootSourceSelection{APIClient: *apiClient},
		BootSourceSelections:  &BootSourceSelections{APIClient: *apiClient},
		BootResource:          &BootResource{APIClient: *apiClient},
		BootResources:         &BootResources{APIClient: *apiClient},
		NodeResults:           &NodeResults{APIClient: *apiClient},
		Zone:                  &Zone{APIClient: *apiClient},
		Zones:                 &Zones{APIClient: *apiClient},
		BCache:                &BCache{APIClient: *apiClient},
		BCaches:               &BCaches{APIClient: *apiClient},
		BCacheCacheSet:        &BCacheCacheSet{APIClient: *apiClient},
		BCacheCacheSets:       &BCacheCacheSets{APIClient: *apiClient},
		SSHKey:                &SSHKey{APIClient: *apiClient},
		SSHKeys:               &SSHKeys{APIClient: *apiClient},
		SSLKey:                &SSLKey{APIClient: *apiClient},
		SSLKeys:               &SSLKeys{APIClient: *apiClient},
		Account:               &Account{APIClient: *apiClient},
		Version:               &Version{APIClient: *apiClient},
		RackController:        &RackController{APIClient: *apiClient},
		RackControllers:       &RackControllers{APIClient: *apiClient},
		NodeScript:            &NodeScript{APIClient: *apiClient},
		NodeScripts:           &NodeScripts{APIClient: *apiClient},
		VolumeGroup:           &VolumeGroup{APIClient: *apiClient},
		VolumeGroups:          &VolumeGroups{APIClient: *apiClient},
	}

	return &client
}

// Client is an object providing API interactions
// with a configured MAAS installation
type Client struct {
	Device                api.Device
	Devices               api.Devices
	DNSResource           api.DNSResource
	DNSResources          api.DNSResources
	DNSResourceRecord     api.DNSResourceRecord
	DNSResourceRecords    api.DNSResourceRecords
	Domain                api.Domain
	Domains               api.Domains
	Events                api.Events
	Fabric                api.Fabric
	Fabrics               api.Fabrics
	VLAN                  api.VLAN
	VLANs                 api.VLANs
	Space                 api.Space
	Spaces                api.Spaces
	Machine               api.Machine
	Machines              api.Machines
	VMHost                api.VMHost
	VMHosts               api.VMHosts
	NetworkInterface      api.NetworkInterface
	NetworkInterfaces     api.NetworkInterfaces
	NodeDevice            api.NodeDevice
	NodeDevices           api.NodeDevices
	RAID                  api.RAID
	RAIDs                 api.RAIDs
	Subnet                api.Subnet
	Subnets               api.Subnets
	IPRange               api.IPRange
	IPRanges              api.IPRanges
	IPAddresses           api.IPAddresses
	Tag                   api.Tag
	Tags                  api.Tags
	BlockDevice           api.BlockDevice
	BlockDevices          api.BlockDevices
	BlockDevicePartition  api.BlockDevicePartition
	BlockDevicePartitions api.BlockDevicePartitions
	User                  api.User
	Users                 api.Users
	ResourcePool          api.ResourcePool
	ResourcePools         api.ResourcePools
	MAASServer            api.MAASServer
	PackageRepository     api.PackageRepository
	PackageRepositories   api.PackageRepositories
	BootSource            api.BootSource
	BootSources           api.BootSources
	BootSourceSelection   api.BootSourceSelection
	BootSourceSelections  api.BootSourceSelections
	BootResource          api.BootResource
	BootResources         api.BootResources
	NodeResults           api.NodeResults
	Zone                  api.Zone
	Zones                 api.Zones
	BCache                api.BCache
	BCaches               api.BCaches
	BCacheCacheSet        api.BCacheCacheSet
	BCacheCacheSets       api.BCacheCacheSets
	SSHKey                api.SSHKey
	SSHKeys               api.SSHKeys
	SSLKey                api.SSLKey
	SSLKeys               api.SSLKeys
	Account               api.Account
	Version               api.Version
	RackController        api.RackController
	RackControllers       api.RackControllers
	NodeScript            api.NodeScript
	NodeScripts           api.NodeScripts
	VolumeGroup           api.VolumeGroup
	VolumeGroups          api.VolumeGroups
}

func getAPIClient(apiURL string, apiKey string, apiVersion string, tr http.RoundTripper) (*APIClient, error) {
	httpClient := &http.Client{
		Transport: tr,
	}

	if tr == nil {
		httpClient.Transport = http.DefaultTransport
	}

	return NewAPIClient(apiURL, apiVersion, apiKey, httpClient)
}
