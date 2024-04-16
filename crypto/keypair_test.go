package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeypairSignVerifySuccess(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	msg := []byte("hello world")

	sig, err := privKey.Sign(msg)

	assert.Nil(t, err)
	assert.True(t, sig.Verify(pubKey, msg))
}

func TestKeypairSignVerifyFail(t *testing.T) {
	privKey1 := GeneratePrivateKey()
	privKey2 := GeneratePrivateKey()
	pubKey1 := privKey1.PublicKey()
	pubKey2 := privKey2.PublicKey()
	msg := []byte("hello world")

	sig, err := privKey1.Sign(msg)

	assert.Nil(t, err)
	assert.False(t, sig.Verify(pubKey2, msg))
	assert.False(t, sig.Verify(pubKey1, []byte("ïncorrect msg")))
}
