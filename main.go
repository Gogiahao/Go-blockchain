package main

import (
	"fmt"
	blc "github.com/Gogiahao/Go-blockchain/blockchain"
)

func main() {
	chain := blc.InitBlockChain()

	chain.AddBlock("第一个区块")
	chain.AddBlock("第二个区块")
	chain.AddBlock("第三个区块")

	for _, block := range chain.Blocks {
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
