package main

import (
	"encoding/base64"
	"fmt"

	"github.com/RadicalApp/libsignal-protocol-go/ecc"
	"github.com/RadicalApp/libsignal-protocol-go/logger"
	"github.com/RadicalApp/libsignal-protocol-go/util/keyhelper"
)

func generateIdentityKeyPair() interface{} {
	identityKeyPair, err := keyhelper.GenerateIdentityKeyPair()
	if err != nil {
		logger.Error(err)
		return nil
	}
	publicKey := identityKeyPair.PublicKey().Serialize()
	privateKey := identityKeyPair.PrivateKey().Serialize()
	pub := base64.StdEncoding.EncodeToString(publicKey)
	priv := base64.StdEncoding.EncodeToString(privateKey[:])
	return map[string]interface{}{"pub": pub, "priv": priv}
}

func generateKeyPair() interface{} {
	keyPair, err := ecc.GenerateKeyPair()
	if err != nil {
		logger.Error(err)
		return nil
	}
	public := keyPair.PublicKey().Serialize()
	private := keyPair.PrivateKey().Serialize()
	pub := base64.StdEncoding.EncodeToString(public)
	priv := base64.StdEncoding.EncodeToString(private[:])
	return map[string]interface{}{"pub": pub, "priv": priv}
}

func main() {
	k := generateIdentityKeyPair()
	fmt.Println(k)
}
