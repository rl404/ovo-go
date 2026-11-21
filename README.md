# ovo-go

![License: MIT](https://img.shields.io/github/license/rl404/ovo-go.svg)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/rl404/ovo-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/rl404/ovo-go.svg)](https://pkg.go.dev/github.com/rl404/ovo-go)

_ovo-go_ is unofficial golang API wrapper for [OVO](https://www.ovo.id/).

## Features

- Inquiry customer phone number
- Create push-to-pay transaction
- Create reversal push-to-pay transaction
- Void transaction
- Get transaction

## Installation

```sh
go get github.com/rl404/ovo-go
```

## Quick Start

```go
package main

import (
    "log"

    "github.com/rl404/ovo-go"
)

func main() {
    // Prepare credentials.
    appID := "appID"
    key := "key123"
    tid := "123"
    mid := "123"
    merchantID := "123"
    storeCode := "ABC123"

    // Create ovo client.
    o := ovo.NewDefault(appID, key, tid, mid, merchantID, storeCode, ovo.Sandbox)

    // Create transaction.
    tx, code, err := o.GetStatus(ovo.GetStatusRequest{
        Amount:          10000,
        Phone:           "081234567890",
        MerchantInvoice: "invoice123",
        ReferenceNumber: 1,
        BatchNo:         1,
    })
    if err != nil {
        log.Println(code, err)
        return
    }

    log.Println(code, tx)
}
```

_For more detail config and usage, please go to the [documentation](https://pkg.go.dev/github.com/rl404/ovo-go)._

## License

MIT License

Copyright (c) 2021 Axel
