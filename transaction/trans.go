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

func main(){
	rawTxHex:=createTrans()
	sendTrans(rawTxHex)
}

// 交易签名
func createTrans()string{
	// 连接以太坊测试链
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/fb3b49110c2a4509b449923a936dcfc7")
	if err != nil {
		log.Fatal(err)
	}

	// 解析 secp256k1 private key，获取了一个私钥
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//value := big.NewInt(1000000000000000000) // in wei (1 eth)
	value := big.NewInt(	1) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte

	// 构造事务
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	fmt.Printf("trans:%+v\n",tx)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	ts := types.Transactions{signedTx}
	rawTxBytes := ts.GetRlp(0)
	rawTxHex := hex.EncodeToString(rawTxBytes)

	fmt.Printf("rawTxHex:%s\n",rawTxHex) // f86...772
	return rawTxHex
}

func sendTrans(rawTxHex string){
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/fb3b49110c2a4509b449923a936dcfc7")
	if err != nil {
		log.Fatal(err)
	}

	//rawTxHex := "f86d8202b28477359400825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d880de0b6b3a7640000802ca05924bde7ef10aa88db9c66dd4f5fb16b46dff2319b9968be983118b57bb50562a001b24b31010004f13d9a26b320845257a6cfc2bf819a3d55e3fc86263c5f0772"

	rawTxBytes, err := hex.DecodeString(rawTxHex)

	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)
	fmt.Printf("rawTx:%+v\n",tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
}

// 验证签名
func verifyTrans(){
}


