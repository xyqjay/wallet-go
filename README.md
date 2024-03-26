# wallet-go
Web3 wallet address generate using golang, The same generation algorithm as the mainstream wallet apps, for example: MetaMask, Phantom, OKX Wallet, UniSat Wallet


# BTC Taproot Example:
```go
package main

import (
	"fmt"
	"wallet-golang/btc/taproot"
	mnemonicgen "wallet-golang/mnemonic"
)

func main() {
	fmt.Println("Hello")

	mnemonic := mnemonicgen.GenerateOne(mnemonicgen.Num12)

	fmt.Println(taproot.GetAddress(mnemonic, false))

	fmt.Println(taproot.GetAddresses(mnemonic, 0, 5, false))
}
```