package hyperion

import (
	"net/http"

	"github.com/TheBoringDude/go-urljoin"
	"github.com/TheBoringDude/scuffed-go/requester"
)

type HyperionClient struct {
	Client *requester.Requester
	ApiUrl string
}

// New creates a new HyperionClient object.
// There is no need to include the `v2` at the end.
func New(api string) *HyperionClient {
	return &HyperionClient{
		ApiUrl: urljoin.UrlJoin(api, "v2"),
		Client: requester.NewRequester(&http.Client{}),
	}
}
