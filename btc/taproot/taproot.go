package taproot

import (
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/tyler-smith/go-bip39"
)

var (
	xPath_TAPROOT       = "m/86'/0'/0'/0/0"
	xPath_TAPROOT_INDEX = "m/86'/0'/0'/0/__INDEX__"
	xIndex              = "__INDEX__"
)

func GetAddress(mnemonic string, isTestNet bool) (string, error) {
	seed := bip39.NewSeed(mnemonic, "")

	BTCParams := chaincfg.MainNetParams
	if isTestNet {
		BTCParams = chaincfg.TestNet3Params
	}

	masterKey, err := hdkeychain.NewMaster(seed, &BTCParams)
	if err != nil {
		return "", err
	}

	dpath, err := accounts.ParseDerivationPath(xPath_TAPROOT)
	if err != nil {
		return "", err
	}

	key := masterKey
	for _, n := range dpath {
		key, err = key.Derive(n)
		if err != nil {
			return "", err
		}
	}
	publicKey, err := key.ECPubKey()
	if err != nil {
		return "", err
	}

	p2tr, err := btcutil.NewAddressTaproot(txscript.ComputeTaprootKeyNoScript(publicKey).SerializeCompressed()[1:], &BTCParams)
	if err != nil {
		return "", err
	}
	return p2tr.EncodeAddress(), nil
}

func GetAddresses(mnemonic string, start int, count int, isTestNet bool) ([]string, error) {
	seed := bip39.NewSeed(mnemonic, "")
	BTCParams := chaincfg.MainNetParams
	if isTestNet {
		BTCParams = chaincfg.TestNet3Params
	}

	masterKey, err := hdkeychain.NewMaster(seed, &BTCParams)
	if err != nil {
		return []string{}, err
	}

	res := []string{}

	for i := start; i < count; i++ {
		dpath, err := accounts.ParseDerivationPath(strings.ReplaceAll(xPath_TAPROOT_INDEX, xIndex, strconv.Itoa(i)))
		if err != nil {
			return []string{}, err
		}
		key := masterKey
		for _, n := range dpath {
			key, err = key.Derive(n)
			if err != nil {
				return []string{}, err
			}
		}
		publicKey, err := key.ECPubKey()
		if err != nil {
			return []string{}, err
		}

		p2tr, err := btcutil.NewAddressTaproot(txscript.ComputeTaprootKeyNoScript(publicKey).SerializeCompressed()[1:], &BTCParams)
		if err != nil {
			return []string{}, err
		}
		res = append(res, p2tr.EncodeAddress())
	}

	return res, nil
}
