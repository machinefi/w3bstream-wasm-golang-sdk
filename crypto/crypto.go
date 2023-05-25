package secp256k1

import (
	"encoding/hex"

	gp "github.com/iotexproject/go-pkgs/crypto"
)

func Verify(pubkeyHex string, dataBytes []byte, sigHex string) (bool, error) {
	pbk, err := gp.HexStringToPublicKey(pubkeyHex)
	if err != nil {
		return false, err
	}
	sig, err := hex.DecodeString(sigHex)
	if err != nil {
		return false, err
	}
	ret := pbk.Verify(dataBytes, sig)
	return ret, nil
}
