package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// BlockDevicePartitions implements api.BlockDevicePartitions
type BlockDevicePartitions struct {
	APIClient APIClient
}

func (p *BlockDevicePartitions) client(systemID string, blockDeviceID int) *APIClient {
	return p.APIClient.SubClient(fmt.Sprintf("nodes/%s/blockdevices/%v/partitions", systemID, blockDeviceID))
}

// Get lists the BlockDevicePartition objects for a given system_id and BlockDevice id
func (p *BlockDevicePartitions) Get(ctx context.Context, systemID string, blockDeviceID int) ([]entity.BlockDevicePartition, error) {
	partitions := make([]entity.BlockDevicePartition, 0)
	err := p.client(systemID, blockDeviceID).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &partitions)
	})

	return partitions, err
}

// Create creats a new BlockDevicePartition for a given system_id and BlockDevice id
func (p *BlockDevicePartitions) Create(ctx context.Context, systemID string, blockDeviceID int, params *entity.BlockDevicePartitionParams) (*entity.BlockDevicePartition, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	partition := new(entity.BlockDevicePartition)
	err = p.client(systemID, blockDeviceID).Post(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, partition)
	})

	return partition, err
}
