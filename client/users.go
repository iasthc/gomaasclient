//nolint:dupl // disable dupl check on client for now
package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// Users implements api.Users
type Users struct {
	APIClient APIClient
}

func (u *Users) client() *APIClient {
	return u.APIClient.SubClient("users")
}

// Get fetches a list of User objects
func (u *Users) Get(ctx context.Context) ([]entity.User, error) {
	users := make([]entity.User, 0)
	err := u.client().Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &users)
	})

	return users, err
}

// Create creates a new User
func (u *Users) Create(ctx context.Context, params *entity.UserParams) (*entity.User, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	user := new(entity.User)
	err = u.client().Post(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, user)
	})

	return user, err
}
