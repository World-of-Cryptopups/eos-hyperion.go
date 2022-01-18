package hyperion

import (
	"fmt"

	"github.com/TheBoringDude/go-urljoin"
)

const HISTORY_API = "history"

func (h *HyperionClient) historyUrl(url string, params interface{}) string {
	return urljoin.UrlJoin(h.ApiUrl, HISTORY_API, url, fmt.Sprintf("?%s", parseQuery(params)))

}

type GetAbiSnapshotParams struct {
	Contract string `url:"contract"`
	Block    int64  `url:"block,omitempty"`
	Fetch    bool   `url:"fetch,omitempty"`
}

// Fetch contract abi at specific block.
func (h *HyperionClient) HistoryGetAbiSnapshot(params GetAbiSnapshotParams) (HistoryABISnapshotResponse, error) {
	var r = HistoryABISnapshotResponse{}

	err := h.Client.Get(h.historyUrl("/get_abi_snapshot", params), &r)

	return r, err

}

type GetActionsParams struct {
	Limit    int64  `url:"limit,omitempty"`
	Skip     int64  `url:"skip,omitempty"`
	Account  string `url:"account,omitempty"`
	Track    string `url:"track,omitempty"`
	Filter   string `url:"filter,omitempty"`
	Sort     string `url:"sort,omitempty"`
	After    string `url:"after,omitempty"`
	Before   string `url:"before,omitempty"`
	Simple   bool   `url:"simple,omitempty"`
	NoBinary bool   `url:"noBinary,omitempty"`
	CheckLib bool   `url:"checkLib,omitempty"`
}

// Get actions based on notified account.
// This endpoint also accepts generic filters based on indexed fields
// (e.g. act.authorization.actor=eosio or act.name=delegatebw), if included they will be combined with a AND operator.
func (h *HyperionClient) HistoryGetActions(params GetActionsParams) (HistoryActionsResponse, error) {
	var r = HistoryActionsResponse{}

	err := h.Client.Get(h.historyUrl("/get_actions", params), &r)

	return r, err
}

type GetDeltasParams struct {
	Limit int64  `url:"limit,omitempty"`
	Skip  int64  `url:"skip,omitempty"`
	Code  string `json:"code,omitempty"`
	Scope string `json:"scope,omitempty"`
	Table string `json:"table,omitempty"`
	Payer string `json:"payer,omitempty"`
}

// Get state deltas.
func (h *HyperionClient) HistoryGetDeltas(params GetDeltasParams) (HistoryDeltasResponse, error) {
	var r = HistoryDeltasResponse{}

	err := h.Client.Get(h.historyUrl("/get_deltas", params), &r)

	return r, err
}

type GetScheduleParams struct {
	Producer string `json:"producer,omitempty"`
	Key      string `json:"key,omitempty"`
	After    string `json:"after,omitempty"`
	Before   string `json:"before,omitempty"`
	Version  int64  `json:"version,omitempty"`
}

// Get producer schedule by version.
func (h *HyperionClient) HistoryGetSchedule(params GetScheduleParams) (HistoryScheduleResponse, error) {
	var r = HistoryScheduleResponse{}

	err := h.Client.Get(h.historyUrl("/get_schedule", params), &r)

	return r, err
}

type GetTransactionParams struct {
	ID int64 `url:"id"`
}

// Get all actions belonging to the same transaction.
func (h *HyperionClient) HistoryGetTransaction(trxId int64) (HistoryTransactionResponse, error) {
	var r = HistoryTransactionResponse{}

	err := h.Client.Get(h.historyUrl("/get_transaction", GetTransactionParams{
		ID: trxId,
	}), &r)

	return r, err
}

type GetCreatedAccountsParams struct {
	Limit   int64  `url:"limit,omitempty"`
	Skip    int64  `url:"skip,omitempty"`
	Account string `url:"account"`
}

// Get all accounts created by one creator.
func (h *HyperionClient) HistoryGetCreatedAccounts(params GetCreatedAccountsParams) (HistoryCreatedAccountsResponse, error) {
	var r = HistoryCreatedAccountsResponse{}

	err := h.Client.Get(h.historyUrl("/get_created_accounts", params), &r)

	return r, err
}

type GetCreatorParams struct {
	Account string `url:"account"`
}

// Get account creator.
func (h *HyperionClient) HistoryGetCreator(account string) (HistoryCreatorResponse, error) {
	var r = HistoryCreatorResponse{}

	err := h.Client.Get(h.historyUrl("/get_creator", GetCreatorParams{
		Account: account,
	}), &r)

	return r, err
}
