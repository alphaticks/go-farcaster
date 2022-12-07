package utils

import (
	"encoding/json"
	"github.com/tyler-smith/go-bip32"
	"net/http"
)

const Purpose uint32 = 0x8000002C

func PerformJSONRequest(client *http.Client, r *http.Request, target interface{}) error {
	res, err := client.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	/*
		b, err := ioutil.ReadAll(res.Body)
		fmt.Println(string(b), res.StatusCode)
	*/

	if target != nil {
		return json.NewDecoder(res.Body).Decode(target)
	}

	return nil
}

func NewKeyFromMasterKey(masterKey *bip32.Key, coin, account, chain, address uint32) (*bip32.Key, error) {
	child, err := masterKey.NewChildKey(Purpose)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(coin)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(account)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(chain)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(address)
	if err != nil {
		return nil, err
	}

	return child, nil
}
