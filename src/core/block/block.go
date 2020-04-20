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

func New(index int, timestamp int, prevHash string, data interface{}) *Block {
	meta := &Meta{
		Index: index,
		Timestamp: timestamp,
		PrevHash: prevHash,
	}
	block := &Block{
		Meta: meta,
		Data: data,
	}

	block.Meta.Hash = block.CreateHash(block.Meta.Nonce)

	return block
}

func (b Block) CreateHash(nonce int) string {
	if nonce < 0 {
		nonce = 0
	}

	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%d%d%s%d", b.Meta.Index, b.Meta.Timestamp, b.Meta.PrevHash, nonce)))

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
