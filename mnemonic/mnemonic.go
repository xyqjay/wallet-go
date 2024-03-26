package mnemonic

import (
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/tyler-smith/go-bip39"
)

type NUMOFWORDS int

const (
	Num12 NUMOFWORDS = 12
	Num24 NUMOFWORDS = 24
)
const ()

func GenerateOne(numOfWords NUMOFWORDS) string {
	bitSize := 128
	if Num24 == numOfWords {
		bitSize = 256
	}
	entropy, _ := bip39.NewEntropy(bitSize)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	return mnemonic
}

func Generate(numOfWords NUMOFWORDS, count int) []string {
	if count <= 0 {
		return []string{}
	}
	bitSize := 128
	if Num24 == numOfWords {
		bitSize = 256
	}

	res := []string{}

	for i := 0; i < count; i++ {
		entropy, _ := bip39.NewEntropy(bitSize)
		mnemonic, _ := bip39.NewMnemonic(entropy)
		res = append(res, mnemonic)
	}
	return res
}

func ECPubKey(mnemonic string, path string, net *chaincfg.Params) (*btcec.PublicKey, error) {
	seed := bip39.NewSeed(mnemonic, "")

	masterKey, err := hdkeychain.NewMaster(seed, net)
	if err != nil {
		return nil, err
	}

	dpath, err := accounts.ParseDerivationPath(path)
	if err != nil {
		return nil, err
	}

	key := masterKey
	for _, n := range dpath {
		key, err = key.Derive(n)
		if err != nil {
			return nil, err
		}
	}
	publicKey, err := key.ECPubKey()
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}
