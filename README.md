# eos-hyperion.go

GoLang API Wrapper for the Hyperion state api.

## Install

```sh
go get -u github.com/World-of-Cryptopups/eos-hyperion.go
```

## Usage

```go
package main

import (
    "fmt"
    "log"

    hyperion "github.com/World-of-Cryptopups/eos-hyperion.go"
)

func main() {
    client := hyperion.New("https://testnet.waxsweden.org")

    health, err := client.GetHealth()
    if err != nil {
        log.Fatalln(err)
    }

    fmt.Println(health.Version)
}

```

##

**2022 | World of Cryptopups**
