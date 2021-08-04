package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

//PoW流程:
//从区块取出数据
//创建一个从0开始的计数器（随机数）
//将数据与随机数结合，计算哈希
//验证哈希是否符合要求

//要求：
//哈希的前几个字节必须为0

const Difficulty = 20

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

//创建新的工作量证明对象
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)                  //target初始化为1
	target.Lsh(target, uint(256-Difficulty)) //左移(256-Difficulty)位，得到最终target

	return &ProofOfWork{b, target}
}

//初始化数据，将区块数据拼接成字节数组
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

//穷举计算符合target的nonce值
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	//死循环
	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])
		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash[:]
}

//验证哈希是否满足target
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int
	data := pow.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])
	return intHash.Cmp(pow.Target) == -1
}

//整型数据转换成十六进制的字节数组，用以让Nonce、Difficulty参与区块哈希计算
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)                        //缓冲区
	err := binary.Write(buff, binary.BigEndian, num) //BigEndian，高位字节放在内存的低地址端，低位字节放在内存的高地址端
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
