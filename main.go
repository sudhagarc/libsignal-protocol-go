package main

import "C"

import (
	"encoding/base64"
	"encoding/json"

	"github.com/RadicalApp/libsignal-protocol-go/ecc"
	"github.com/RadicalApp/libsignal-protocol-go/logger"
	"github.com/RadicalApp/libsignal-protocol-go/util/keyhelper"
)

type Key struct {
	Pub  []byte `json:"pub"`
	Priv []byte `json:"priv"`
}

//export GenerateIdentityKeyPair
func GenerateIdentityKeyPair() interface{} {
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

//export GenerateKeyPair
func GenerateKeyPair(keyStr *string) int {
	keyPair, err := ecc.GenerateKeyPair()
	if err != nil {
		*keyStr = ""
		return 0
	}
	public := keyPair.PublicKey().Serialize()
	private := keyPair.PrivateKey().Serialize()
	key := Key{
		Pub:  public,
		Priv: private[:],
	}
	k, _ := json.Marshal(key)
	*keyStr = string(k)
	return 1
}

func main() {
}
