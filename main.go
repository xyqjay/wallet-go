package main

import (
	"fmt"
	"wallet-go/btc/taproot"
	mnemonicgen "wallet-go/mnemonic"
)

func main() {
	fmt.Println("Hello")

	mnemonic := mnemonicgen.GenerateOne(mnemonicgen.Num12)

	fmt.Println(taproot.GetAddress(mnemonic, false))

	fmt.Println(taproot.GetAddresses(mnemonic, 0, 5, false))
}
