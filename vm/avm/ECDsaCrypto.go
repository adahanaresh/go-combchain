package vm

import (
	"errors"
	"github.com/combchain/combchain/crypto"
	"github.com/combchain/combchain/log"
	. "github.com/combchain/go-combchain/vm/avm/errors"
)

type ECDsaCrypto struct {
}

func (c *ECDsaCrypto) Hash160(message []byte) []byte {
	return []byte{}
}

func (c *ECDsaCrypto) Hash256(message []byte) []byte {
	return []byte{}
}

func (c *ECDsaCrypto) VerifySignature(message []byte, signature []byte, pubkey []byte) (bool, error) {

	log.Debug("message: %x \n", message)
	log.Debug("signature: %x \n", signature)
	log.Debug("pubkey: %x \n", pubkey)

	pk, err := crypto.DecodePoint(pubkey)
	if err != nil {
		return false, NewDetailErr(errors.New("[ECDsaCrypto], crypto.DecodePoint failed."), ErrNoCode, "")
	}

	err = crypto.Verify(*pk, message, signature)
	if err != nil {
		return false, NewDetailErr(errors.New("[ECDsaCrypto], VerifySignature failed."), ErrNoCode, "")
	}
	return true, nil
}
