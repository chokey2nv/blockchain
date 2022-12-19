package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	if err := sign(); err != nil {
		log.Fatalln(err)
	}
}

func sign() error {
	path := fmt.Sprintf("%s%s.ecdsa", "zblock/accounts/", "kennedy")

	privateKey, err := crypto.LoadECDSA(path)
	if err != nil {
		return fmt.Errorf("load: %w", err)
	}
	// privateKey, err := crypto.GenerateKey()
	// fmt.Println(privateKey)
	// if err != nil {
	// 	return err
	// }
	address := crypto.PubkeyToAddress(privateKey.PublicKey).String()
	fmt.Println(address)

	v := struct {
		Name string
	}{
		Name: "chijioke",
	}

	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	dataHash := crypto.Keccak256(data)

	sig, err := crypto.Sign(dataHash, privateKey)
	if err != nil {
		return fmt.Errorf("sign: %w", err)
	}
	fmt.Println(sig)

	// =========================================

	signPublicKey, err := crypto.Ecrecover(dataHash, sig)
	if err != nil {
		return fmt.Errorf("recover pub key: %w", err)
	}
	fmt.Println(signPublicKey)

	x, y := elliptic.Unmarshal(crypto.S256(), signPublicKey)
	publicKey := ecdsa.PublicKey{
		Curve: crypto.S256(),
		X:     x,
		Y:     y,
	}
	fmt.Println(crypto.PubkeyToAddress(publicKey))
	return nil
}
