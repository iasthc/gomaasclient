package client

import (
	"context"
	"encoding/json"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// NodeScripts implements api.NodeScripts
type NodeScripts struct {
	APIClient APIClient
}

func (ns *NodeScripts) client() *APIClient {
	return ns.APIClient.SubClient("scripts")
}

// Get fetches a list of NodeScripts
func (ns *NodeScripts) Get(ctx context.Context, nodeScriptParams *entity.NodeScriptReadParams) ([]entity.NodeScript, error) {
	qsp, err := query.Values(nodeScriptParams)
	if err != nil {
		return nil, err
	}

	nodeScripts := make([]entity.NodeScript, 0)
	err = ns.client().Get(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, &nodeScripts)
	})

	return nodeScripts, err
}

// Create creates a new NodeScript
func (ns *NodeScripts) Create(ctx context.Context, nodeScriptParams *entity.NodeScriptParams, script []byte) (*entity.NodeScript, error) {
	qsp, err := query.Values(nodeScriptParams)
	if err != nil {
		return nil, err
	}

	files := map[string][]byte{}
	if script != nil {
		files["script"] = script
	}

	nodeScript := new(entity.NodeScript)
	err = ns.client().PostFiles(ctx, "", qsp, files, func(data []byte) error {
		return json.Unmarshal(data, nodeScript)
	})

	return nodeScript, err
}
