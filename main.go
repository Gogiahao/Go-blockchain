package main

import (
	"fmt"
	blc "github.com/Gogiahao/Go-blockchain/blockchain"
	"strconv"
)

func main() {
	chain := blc.InitBlockChain()

	chain.AddBlock("第二个区块")
	chain.AddBlock("第三个区块")
	chain.AddBlock("第四个区块")

	for _, block := range chain.Blocks {
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blc.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
