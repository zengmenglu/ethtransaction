package wallet

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"golang.org/x/crypto/sha3"
)

// 生成EOA账户
func NewWallet()(string,string, error){
	// 私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil{
		return "","", err
	}
	privateKeyByte := crypto.FromECDSA(privateKey)
	privateKeyString := hexutil.Encode(privateKeyByte)[2:]//删除了前导0x
	log.Printf("privateKeyByte:%v\n",privateKeyString)

	//公钥
	publicKey:=privateKey.Public()
	publicKeyECDAS, ok := publicKey.(*ecdsa.PublicKey)
	if !ok{
		return "", "", fmt.Errorf("publickey error")
	}
	publicKeyByte:=crypto.FromECDSAPub(publicKeyECDAS)
	publicKeyString:=hexutil.Encode(publicKeyByte)[4:]//删除前导0x04
	log.Printf("publicKeyByte:%v\n",publicKeyString)

	//钱包地址
	addr := crypto.PubkeyToAddress(*publicKeyECDAS).Hex()
	log.Printf("addr:%v\n",addr)

	return publicKeyString, addr, nil
}

// 根据公钥获取钱包地址
// 以下函数等效于
// addr:=crypto.PubkeyToAddress(*publicKeyECDAS).Hex()
func GetWalletAddr(publicKey []byte)string{
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKey[1:])
	addr:=hexutil.Encode(hash.Sum(nil)[12:])
	return addr
}