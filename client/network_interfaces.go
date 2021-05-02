package client

import (
	"encoding/json"
	"net/url"

	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type NetworkInterfaces struct {
	ApiClient ApiClient
}

func (n *NetworkInterfaces) client(systemID string) ApiClient {
	return n.ApiClient.GetSubObject("nodes").GetSubObject(systemID).GetSubObject("interfaces")
}

func (n *NetworkInterfaces) Get(systemID string) (networkInterfaces []entity.NetworkInterface, err error) {
	err = n.client(systemID).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &networkInterfaces)
	})
	return
}

func (n *NetworkInterfaces) CreateBond(systemID string, params *entity.NetworkInterfaceBondParams) (networkInterface *entity.NetworkInterface, err error) {
	networkInterface = new(entity.NetworkInterface)
	err = n.client(systemID).Post("create_bond", ToQSP(params), func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})
	return
}

func (n *NetworkInterfaces) CreateBridge(systemID string, params *entity.NetworkInterfaceBridgeParams) (networkInterface *entity.NetworkInterface, err error) {
	networkInterface = new(entity.NetworkInterface)
	err = n.client(systemID).Post("create_bridge", ToQSP(params), func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})
	return
}

func (n *NetworkInterfaces) CreatePhysical(systemID string, params *entity.NetworkInterfacePhysicalParams) (networkInterface *entity.NetworkInterface, err error) {
	networkInterface = new(entity.NetworkInterface)
	err = n.client(systemID).Post("create_physical", ToQSP(params), func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})
	return
}

func (n *NetworkInterfaces) CreateVLAN(systemID string, params *entity.NetworkInterfaceVLANParams) (networkInterface *entity.NetworkInterface, err error) {
	networkInterface = new(entity.NetworkInterface)
	err = n.client(systemID).Post("create_vlan", ToQSP(params), func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})
	return
}
