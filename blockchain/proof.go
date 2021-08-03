package blockchain

import (
	"bytes"
	"math/big"
)

//从区块取出数据

//创建一个从0开始的计数器（随机数）

//将数据与随机数结合，计算哈希

//验证哈希是否符合要求

//要求：
//前几个字节必须为0

const Difficulty = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	return &ProofOfWork{b, target}
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
		},
		[]byte{},
	)
	return data
}
