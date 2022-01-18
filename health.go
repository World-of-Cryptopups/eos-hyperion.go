package hyperion

import "github.com/TheBoringDude/go-urljoin"

type HealthProps struct {
	Version     string   `json:"version"`
	VersionHash string   `json:"version_hash"`
	Host        string   `json:"host"`
	Features    Features `json:"features"`
	Health      []Health `json:"health"`
	QueryTimeMS float64  `json:"query_time_ms"`
}

type Features struct {
	Streaming         Streaming `json:"streaming"`
	Tables            Tables    `json:"tables"`
	IndexDeltas       bool      `json:"index_deltas"`
	IndexTransferMemo bool      `json:"index_transfer_memo"`
	IndexAllDeltas    bool      `json:"index_all_deltas"`
	DeferredTrx       bool      `json:"deferred_trx"`
	FailedTrx         bool      `json:"failed_trx"`
	ResourceLimits    bool      `json:"resource_limits"`
	ResourceUsage     bool      `json:"resource_usage"`
}

type Streaming struct {
	Enable bool `json:"enable"`
	Traces bool `json:"traces"`
	Deltas bool `json:"deltas"`
}

type Tables struct {
	Proposals bool `json:"proposals"`
	Accounts  bool `json:"accounts"`
	Voters    bool `json:"voters"`
}

type Health struct {
	Service     string       `json:"service"`
	Status      string       `json:"status"`
	Time        int64        `json:"time"`
	ServiceData *ServiceData `json:"service_data,omitempty"`
}

type ServiceData struct {
	HeadBlockNum          *int64  `json:"head_block_num,omitempty"`
	HeadBlockTime         *string `json:"head_block_time,omitempty"`
	TimeOffset            *int64  `json:"time_offset,omitempty"`
	LastIrreversibleBlock *int64  `json:"last_irreversible_block,omitempty"`
	ChainID               *string `json:"chain_id,omitempty"`
	LastIndexedBlock      *int64  `json:"last_indexed_block,omitempty"`
	TotalIndexedBlocks    *int64  `json:"total_indexed_blocks,omitempty"`
	ActiveShards          *string `json:"active_shards,omitempty"`
}

// GetHealth returns the API service health report.
func (h *HyperionClient) GetHealth() (HealthProps, error) {
	var r = HealthProps{}

	err := h.Client.Get(urljoin.UrlJoin(h.ApiUrl, "health"), &r)

	return r, err
}
