package main

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/zengmenglu/ethtransaction/testdata"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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

// 验证签名
func VeryfiSign(){
	// 生成公私钥
	privateKey, err := crypto.HexToECDSA(testdata.Secp256k1PriKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println("data hash:", hash.Hex())

	// 生成签名
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("signature:",hexutil.Encode(signature))

	// 方法1：调用 Ecrecover（椭圆曲线签名恢复) 来检索签名者的公钥
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}
	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println("verify result:",matches)

	// 方法2
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}
	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println("verify result:", matches) // true

	// 方法3
	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println("verify result:",verified) // true
}
