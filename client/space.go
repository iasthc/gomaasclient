package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
)

// Space implements api.Space
type Space struct {
	APIClient APIClient
}

func (s *Space) client(id int) *APIClient {
	return s.APIClient.SubClient("spaces").SubClient(fmt.Sprintf("%v", id))
}

// Get fetches a given Space
func (s *Space) Get(ctx context.Context, id int) (*entity.Space, error) {
	space := new(entity.Space)
	err := s.client(id).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, space)
	})

	return space, err
}

// Update updates a given Space
func (s *Space) Update(ctx context.Context, id int, name string) (*entity.Space, error) {
	space := new(entity.Space)
	qsp := url.Values{}
	qsp.Set("name", name)
	err := s.client(id).Put(ctx, qsp, func(data []byte) error {
		return json.Unmarshal(data, space)
	})

	return space, err
}

// Delete deletes a given Space
func (s *Space) Delete(ctx context.Context, id int) error {
	return s.client(id).Delete(ctx)
}
