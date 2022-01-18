package hyperion

type StateAccountResponse struct {
	QueryTimeMS  float64                     `json:"query_time_ms"`
	Account      StateAccountResponseAccount `json:"account"`
	Links        []StateAccountLinks         `json:"links"`
	Tokens       []Token                     `json:"tokens"`
	TotalActions int64                       `json:"total_actions"`
	Actions      []Action                    `json:"actions"`
}

type StateKeyAccountsResponse struct {
	AccountNames []string                       `json:"account_names"`
	Permissions  []KeyAccountsPermissionElement `json:"permissions,omitempty"`
}

type StateLinksResponse struct {
	QueryTimeMS float64 `json:"query_time_ms"`
	Cached      bool    `json:"cached"`
	Total       Total   `json:"total"`
	Links       []Links `json:"links"`
}

type StateTokensResponse struct {
	Account     string  `json:"account"`
	Tokens      []Token `json:"tokens"`
	QueryTimeMS float64 `json:"query_time_ms"`
}

type Links struct {
	BlockNum     int64  `json:"block_num"`
	Timestamp    string `json:"timestamp"`
	Account      string `json:"account"`
	Permission   string `json:"permission"`
	Code         string `json:"code"`
	Action       string `json:"action"`
	Irreversible string `json:"irreversible"`
}

type KeyAccountsPermissionElement struct {
	Owner       string `json:"owner"`
	BlockNum    int64  `json:"block_num"`
	Parent      string `json:"parent"`
	LastUpdated string `json:"last_updated"`
	Auth        Auth   `json:"auth"`
	Name        string `json:"name"`
	Present     bool   `json:"present"`
}

type Auth struct {
	Keys      []Key         `json:"keys"`
	Threshold int64         `json:"threshold"`
	Accounts  []AuthAccount `json:"accounts"`
}

type AuthAccount struct {
	Weight     int64             `json:"weight"`
	Permission AccountPermission `json:"permission"`
}

type AccountPermission struct {
	Actor      string `json:"actor"`
	Permission string `json:"permission"`
}

type StateAccountLinks struct {
	Timestamp  string `json:"timestamp"`
	Permission string `json:"permission"`
	Code       string `json:"code"`
	Action     string `json:"action"`
}

type StateAccountResponseAccount struct {
	AccountName            AccountName            `json:"account_name"`
	HeadBlockNum           int64                  `json:"head_block_num"`
	HeadBlockTime          string                 `json:"head_block_time"`
	Privileged             bool                   `json:"privileged"`
	LastCodeUpdate         string                 `json:"last_code_update"`
	Created                string                 `json:"created"`
	CoreLiquidBalance      string                 `json:"core_liquid_balance"`
	RAMQuota               int64                  `json:"ram_quota"`
	NetWeight              string                 `json:"net_weight"`
	CPUWeight              string                 `json:"cpu_weight"`
	NetLimit               Limit                  `json:"net_limit"`
	CPULimit               Limit                  `json:"cpu_limit"`
	RAMUsage               int64                  `json:"ram_usage"`
	Permissions            []PermissionElement    `json:"permissions"`
	TotalResources         TotalResources         `json:"total_resources"`
	SelfDelegatedBandwidth SelfDelegatedBandwidth `json:"self_delegated_bandwidth"`
	RefundRequest          interface{}            `json:"refund_request"`
	VoterInfo              VoterInfo              `json:"voter_info"`
	RexInfo                interface{}            `json:"rex_info"`
	SubjectiveCPUBillLimit Limit                  `json:"subjective_cpu_bill_limit"`
}

type Limit struct {
	Used      int64 `json:"used"`
	Available int64 `json:"available"`
	Max       int64 `json:"max"`
}

type PermissionElement struct {
	PermName     string       `json:"perm_name"`
	Parent       string       `json:"parent"`
	RequiredAuth RequiredAuth `json:"required_auth"`
}

type RequiredAuth struct {
	Threshold int64            `json:"threshold"`
	Keys      []Key            `json:"keys"`
	Accounts  []AccountElement `json:"accounts"`
	Waits     []interface{}    `json:"waits"`
}

type AccountElement struct {
	Permission Ion   `json:"permission"`
	Weight     int64 `json:"weight"`
}

type Ion struct {
	Actor      AccountName    `json:"actor"`
	Permission PermissionEnum `json:"permission"`
}

type Key struct {
	Key    string `json:"key"`
	Weight int64  `json:"weight"`
}

type SelfDelegatedBandwidth struct {
	From      AccountName `json:"from"`
	To        AccountName `json:"to"`
	NetWeight string      `json:"net_weight"`
	CPUWeight string      `json:"cpu_weight"`
}

type TotalResources struct {
	Owner     AccountName `json:"owner"`
	NetWeight string      `json:"net_weight"`
	CPUWeight string      `json:"cpu_weight"`
	RAMBytes  int64       `json:"ram_bytes"`
}

type VoterInfo struct {
	Owner                      AccountName   `json:"owner"`
	Proxy                      string        `json:"proxy"`
	Producers                  []interface{} `json:"producers"`
	Staked                     string        `json:"staked"`
	UnpaidVoteshare            string        `json:"unpaid_voteshare"`
	UnpaidVoteshareLastUpdated string        `json:"unpaid_voteshare_last_updated"`
	UnpaidVoteshareChangeRate  string        `json:"unpaid_voteshare_change_rate"`
	LastClaimTime              string        `json:"last_claim_time"`
	LastVoteWeight             string        `json:"last_vote_weight"`
	ProxiedVoteWeight          string        `json:"proxied_vote_weight"`
	IsProxy                    int64         `json:"is_proxy"`
	Flags1                     int64         `json:"flags1"`
	Reserved2                  int64         `json:"reserved2"`
	Reserved3                  string        `json:"reserved3"`
}

type Token struct {
	Symbol    string  `json:"symbol"`
	Precision int64   `json:"precision"`
	Amount    float64 `json:"amount"`
	Contract  string  `json:"contract"`
}

type AccountName string

type PermissionEnum string

const (
	Active    PermissionEnum = "active"
	EosioCode PermissionEnum = "eosio.code"
)
