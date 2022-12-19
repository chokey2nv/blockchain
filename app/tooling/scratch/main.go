package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"encoding/json"
	"errors"
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
	fmt.Printf("sig: 0x%s\n", hex.EncodeToString(sig))

	// =========================================

	// hack the public key
	v2 := struct {
		Name string
	}{
		Name: "chijio", // WRONG CALCULATION WILL PRODUCE DIFFERENT ADDERSS (0xF01813E4B85e178A83e29B8E7bF26BD830a25f32 | 0xfEef90eE1cfba406a9d671CBEa76Cf666D321b58)
	}

	data2, err := json.Marshal(v2)
	if err != nil {
		return err
	}

	dataHash2, err := stamp(data2)
	if err != nil {
		return fmt.Errorf("stamp error: %w %w", err, dataHash2)
	}

	// hack ends

	signPublicKey, err := crypto.Ecrecover(dataHash, sig)
	if err != nil {
		return fmt.Errorf("recover pub key: %w", err)
	}
	// fmt.Println(signPublicKey)

	//======== check hack =========
	//dataHash generated sig not dataHash2
	rs := sig[:crypto.RecoveryIDOffset]
	if !crypto.VerifySignature(signPublicKey, dataHash2, rs) {
		return errors.New("invalid signature")
	}

	//======

	x, y := elliptic.Unmarshal(crypto.S256(), signPublicKey)
	publicKey := ecdsa.PublicKey{
		Curve: crypto.S256(),
		X:     x,
		Y:     y,
	}
	fmt.Println(crypto.PubkeyToAddress(publicKey))
	return nil
}
func stamp(value any) ([]byte, error) {

	// Marshal the data.
	v, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	// This stamp is used so signatures we produce when signing data
	// are always unique to the Ardan blockchain.
	stamp := []byte(fmt.Sprintf("\x19Ardan Signed Message:\n%d", len(v)))

	// Hash the stamp and txHash together in a final 32 byte array
	// that represents the data.
	data := crypto.Keccak256(stamp, v)

	return data, nil
}
