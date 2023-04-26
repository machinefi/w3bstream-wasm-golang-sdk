package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"errors"

	"github.com/dustinxie/ecc"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
)

// main is required for TinyGo to compile to Wasm.
func main() {}

//export start
func _start(rid uint32) int32 {
	if err := signVerify(); err != nil {
		log.Log("failed")
		return -1
	}

	log.Log("success")
	return 0
}

func signVerify() error {
	// generate secp256k1 private key
	p256k1 := ecc.P256k1()
	privKey, err := ecdsa.GenerateKey(p256k1, rand.Reader)
	if err != nil {
		// handle error
		return err
	}

	msg := []byte("test")

	// sign message
	hash := sha256.Sum256(msg)
	sig, err := ecc.SignBytes(privKey, hash[:], ecc.Normal)
	if err != nil {
		return err
	}

	// verify message
	if !ecc.VerifyBytes(&privKey.PublicKey, hash[:], sig, ecc.Normal) {
		return errors.New("failed to verify secp256k1 signature")
	}
	return nil
}
