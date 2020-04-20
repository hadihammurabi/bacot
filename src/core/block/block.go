package block

import (
	"fmt"
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	Meta *Meta
	Data interface{}
}

func New(index int, timestamp int, prev_hash string, data interface{}) *Block {
	meta := &Meta{
		index: index,
		timestamp: timestamp,
		prev_hash: prev_hash,
	}
	block := &Block{
		Meta: meta,
		Data: data,
	}

	block.Meta.hash = block.CreateHash(block.Meta.nonce)

	return block
}

func (b Block) CreateHash(nonce int) string {
	if nonce < 0 {
		nonce = 0
	}

	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%d%d%s%d", b.Meta.index, b.Meta.timestamp, b.Meta.prev_hash, nonce)))

	return hex.EncodeToString(h.Sum(nil))
}

func (b *Block) ProofOfWork(difficulty int) {
	if difficulty <= 0 {
		difficulty = 1
	}

	for {
		break
	}
}
