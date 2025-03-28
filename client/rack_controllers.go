package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// RackControllers contains functionality for manipulating the RackControllers entity.
type RackControllers struct {
	APIClient APIClient
}

func (r *RackControllers) client() *APIClient {
	return r.APIClient.SubClient("rackcontrollers")
}

// Get fetches a list rackControllers.
func (r *RackControllers) Get(ctx context.Context, rackControllersParams *entity.RackControllersGetParams) ([]entity.RackController, error) {
	qsp, err := query.Values(rackControllersParams)
	if err != nil {
		return nil, err
	}

	rackControllers := make([]entity.RackController, 0)
	err = r.client().Get(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, &rackControllers)
	})

	return rackControllers, err
}

// DescribePowerTypes returns the supported power types of the rack controller.
func (r *RackControllers) DescribePowerTypes(ctx context.Context) ([]entity.PowerType, error) {
	powerTypes := make([]entity.PowerType, 0)
	err := r.client().Get(ctx, "describe_power_types", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &powerTypes)
	})

	return powerTypes, err
}

// IsRegistered confirms that a MAC Address is registered to MAAS.
func (r *RackControllers) IsRegistered(ctx context.Context, macAddress string) (bool, error) {
	qsp := url.Values{}
	qsp.Set("mac_address", macAddress)

	isRegistered := new(bool)
	err := r.client().Get(ctx, "is_registered", qsp, func(data []byte) error {
		return json.Unmarshal(data, isRegistered)
	})

	return *isRegistered, err
}

// GetPowerParameters of a given node.
func (r *RackControllers) GetPowerParameters(ctx context.Context, systemIDs []string) (map[string]interface{}, error) {
	qsp, err := query.Values(map[string][]string{"id": systemIDs})
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{}
	err = r.client().Get(ctx, "power_parameters", qsp, func(data []byte) error {
		return json.Unmarshal(data, &params)
	})

	return params, err
}

// SetZone to given nodes.
func (r *RackControllers) SetZone(ctx context.Context, setZoneParams *entity.RackControllerSetZoneParams) error {
	qsp, err := query.Values(setZoneParams)
	if err != nil {
		return err
	}

	err = r.client().Post(ctx, "set_zone", qsp, func(data []byte) error { return nil })

	return err
}
