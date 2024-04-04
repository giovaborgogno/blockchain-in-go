package core

import (
	"bytes"
	"testing"
	"time"

	"github.com/giovaborgogno/blockchain-in-go/types"
	"github.com/stretchr/testify/assert"
)

func TestHeader_Enconde_Decode(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     989394,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))

	assert.Equal(t, h, hDecode)
}

func TestBlock_Enconde_Decode(t *testing.T) {
	h := Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     989394,
	}

	b := &Block{
		Header:       h,
		Transactions: nil,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buf))

	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBinary(buf))

	assert.Equal(t, b, bDecode)
}

func TestBlockHash(t *testing.T) {
	h := Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     989394,
	}

	b := &Block{
		Header:       h,
		Transactions: []Transaction{},
	}

	ha := b.Hash()
	assert.False(t, ha.IsZero())
}
