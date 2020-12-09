package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func sign(nonce,gasLimit uint64, toAddress common.Address, value,gasPrice,chainID *big.Int, data []byte,privateKey *ecdsa.PrivateKey)(string, error){
	// 构造交易
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	// 交易签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil{
		return "", err
	}

	// 放入交易列表数组
	trans := types.Transactions{signedTx}

	// 将签名后的交易序列化, 并进行16进制编码
	rawTxBytes := trans.GetRlp(0)
	rawTxHex := hex.EncodeToString(rawTxBytes)
	return rawTxHex,nil
}
