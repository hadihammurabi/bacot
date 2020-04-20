package chain

import (
	"bacot/src/core/block"
)

type Chain struct {
	Name string
	Difficulty int
	Blocks []*block.Block
}

func New(name string, difficulty int) *Chain {
	var blocks []*block.Block
	genesis := CreateGenesis()
	blocks = append(blocks, genesis)
	chain := &Chain{
		Name: name,
		Difficulty: difficulty,
		Blocks: blocks,
	}

	return chain
}

func CreateGenesis() *block.Block {
	genesis := block.New(0, "", "genesis")
	return genesis
}

func (chain *Chain) Add(data interface{}) *block.Block {
	length := len(chain.Blocks)
	prevHash := chain.Blocks[length-1].Meta.Hash
	newblock := block.New(length-1, prevHash, data)
	newblock.ProofOfWork(chain.Difficulty)
	chain.Blocks = append(chain.Blocks, newblock)
	return newblock
}

func (chain *Chain) Peek() *block.Block {
	return chain.Blocks[len(chain.Blocks) - 1]
}

func (chain *Chain) IsValid() bool {
	for i, currentblock := range chain.Blocks {
		var before *block.Block
		if i > 0 {
			before = chain.Blocks[i - 1]
		}

		if before != nil && currentblock.Meta.PrevHash != before.Meta.Hash {
			return false
		}

	}

	return true
}
