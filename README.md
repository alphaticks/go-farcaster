go-farcaster is a library to interact with merkle API V2
    
# Installation
```bash
go get github.com/alphaticks/go-farcaster
```
# Usage

You can either use the client implementation, by supplying it the `ecdsa.PrivateKey` of the custody address of your fid,
which is the `m/44'/60'/0'/0/0` derivation path from the seed of the Merkle client, or pass an application token.
```go
import "github.com/alphaticks/go-farcaster"
client := farcaster.NewClient(privateKey, "")

casts, err := client.GetCasts()
```

You can also build and execute the `http.Request` yourself using the `github.com/alphaticks/go-farcaster/api` package
