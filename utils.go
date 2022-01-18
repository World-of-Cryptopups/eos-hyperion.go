package hyperion

import (
	"fmt"

	"github.com/google/go-querystring/query"
)

// parses url query options
func parseQuery(v interface{}) string {
	x, err := query.Values(v)
	if err != nil {
		fmt.Println(err)
	}

	return x.Encode()
}
