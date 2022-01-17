package hyperion

import (
	"net/http"

	"github.com/TheBoringDude/scuffed-go/requester"
)

type HyperionClient struct {
	Client *requester.Requester
	ApiUrl string
}

// New creates a new HyperionClient object.
func New(api string) *HyperionClient {
	return &HyperionClient{
		ApiUrl: api,
		Client: requester.NewRequester(&http.Client{}),
	}
}
