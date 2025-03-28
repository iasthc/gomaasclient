//nolint:dupl // disable dupl check on client for now
package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// BCache Contains functionality for manipulating the BCache entity.
type BCache struct {
	APIClient APIClient
}

func (b *BCache) client(systemID string, id int) *APIClient {
	return b.APIClient.
		SubClient("nodes").SubClient(systemID).
		SubClient("bcache").SubClient(fmt.Sprintf("%v", id))
}

// Get BCache details.
func (b *BCache) Get(ctx context.Context, systemID string, id int) (*entity.BCache, error) {
	bCache := new(entity.BCache)
	err := b.client(systemID, id).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, bCache)
	})

	return bCache, err
}

// Update BCache.
func (b *BCache) Update(ctx context.Context, systemID string, id int, params *entity.BCacheParams) (*entity.BCache, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	bCache := new(entity.BCache)
	err = b.client(systemID, id).Put(ctx, qsp, func(data []byte) error {
		return json.Unmarshal(data, bCache)
	})

	return bCache, err
}

// Delete BCache.
func (b *BCache) Delete(ctx context.Context, systemID string, id int) error {
	return b.client(systemID, id).Delete(ctx)
}
