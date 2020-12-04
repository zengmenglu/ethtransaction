package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

const (
	testNet         = "https://ropsten.infura.io/v3/fb3b49110c2a4509b449923a936dcfc7"
	secp256k1PriKey = "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
	ethAccount      = "0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d" // 标准eth账户地址
)

func main() {
	rawTxHex := createTrans()
	sendTrans(rawTxHex)

}

// 创建交易和签名
func createTrans() string {
	// 连接以太坊测试链
	client, err := ethclient.Dial(testNet)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	// 加载私钥
	privateKey, err := crypto.HexToECDSA(secp256k1PriKey)
	if err != nil {
		return ""
	}

	// 构造交易
	tx, err := formTrans(privateKey,client)
	if err != nil {
		return ""
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
		return ""
	}

	// 签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	fmt.Printf("signedTx:%s\n", signedTx.Hash().Hex())

	ts := types.Transactions{signedTx}
	rawTxBytes := ts.GetRlp(0)
	rawTxHex := hex.EncodeToString(rawTxBytes)

	fmt.Printf("rawTxHex:%s\n", rawTxHex)
	return rawTxHex
}

// 生成庄户随机数
func genNonce(privateKey *ecdsa.PrivateKey, client *ethclient.Client) (uint64, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return client.PendingNonceAt(context.Background(), fromAddress)
}

func formTrans(privateKey *ecdsa.PrivateKey, client *ethclient.Client) (*types.Transaction, error) {
	nonce, err := genNonce(privateKey, client)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// 交易的ETH数量
	value := big.NewInt(1)    // in wei (10^-18 eth)
	gasLimit := uint64(22000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	toAddress := common.HexToAddress(ethAccount)
	bytecode, err := client.CodeAt(context.Background(), toAddress, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("bytecode:%v\n",bytecode)

	var data = []byte("zml trans")

	// 构造事务
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	fmt.Printf("trans:%+v\n", tx)
	return tx, nil
}

// 发送Trans
func sendTrans(rawTxHex string) {
	client, err := ethclient.Dial(testNet)
	if err != nil {
		log.Fatal(err)
	}

	rawTxBytes, err := hex.DecodeString(rawTxHex)

	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)
	fmt.Printf("rawTx:%+v\n", tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

// 验证签名
func verifyTrans() {
}
