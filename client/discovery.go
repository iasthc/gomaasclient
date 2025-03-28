package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
)

// Discovery implements api.Discovery
type Discovery struct {
	APIClient APIClient
}

func (d *Discovery) client(id string) *APIClient {
	return d.APIClient.SubClient("discovery").
		SubClient(fmt.Sprintf("%v", id))
}

// Get discovery by id
func (d *Discovery) Get(ctx context.Context, id string) (*entity.Discovery, error) {
	deviceDiscovery := new(entity.Discovery)
	err := d.client(id).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &deviceDiscovery)
	})

	return deviceDiscovery, err
}
