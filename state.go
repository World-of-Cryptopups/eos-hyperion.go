package hyperion

import (
	"fmt"

	"github.com/TheBoringDude/go-urljoin"
)

const STATE_API = "state"

func (h *HyperionClient) stateUrl(url string, params interface{}) string {
	return urljoin.UrlJoin(h.ApiUrl, STATE_API, url, fmt.Sprintf("?%s", parseQuery(params)))
}

type GetAccountParams GetCreatedAccountsParams

// Get account data / summary.
func (h *HyperionClient) StateGetAccount(params GetAccountParams) (StateAccountResponse, error) {
	var r = StateAccountResponse{}

	err := h.Client.Get(h.stateUrl("/get_account", params), &r)

	return r, err
}

type GetKeyAccountsParams struct {
	PublicKey string `url:"public_key"`
	Details   bool   `url:"details,omitempty"`
}

// Get accounts by public key.
func (h *HyperionClient) StateGetKeyAccounts(params GetKeyAccountsParams) (StateKeyAccountsResponse, error) {
	var r = StateKeyAccountsResponse{}

	err := h.Client.Get(h.stateUrl("/get_key_accounts", params), &r)

	return r, err
}

// Get accounts by public key.
func (h *HyperionClient) StateGetKeyAccountsPost(publicKey string) (StateKeyAccountsResponse, error) {
	var r = StateKeyAccountsResponse{}

	err := h.Client.Post(
		urljoin.UrlJoin(h.ApiUrl, STATE_API, "/get_key_accounts"),
		map[string]string{
			"public_key": publicKey,
		}, &r)

	return r, err
}

type GetLinksParams struct {
	Account    string `url:"account"`
	Code       string `url:"code"`
	Action     string `url:"action"`
	Permission string `url:"permission"`
}

// Get permission links.
func (h *HyperionClient) StateGetLinks(params GetLinksParams) (StateLinksResponse, error) {
	var r = StateLinksResponse{}

	err := h.Client.Get(h.stateUrl("/get_links", params), &r)

	return r, err
}

type GetTokenParams GetAccountParams

// Get account tokens data.
func (h *HyperionClient) StateGetTokens(params GetTokenParams) (StateTokensResponse, error) {
	var r = StateTokensResponse{}

	err := h.Client.Get(h.stateUrl("/get_tokens", params), &r)

	return r, err
}
