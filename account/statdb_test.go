package account

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zengmenglu/ethtransaction/testdata"
	"math/big"
	"testing"
)
var toAddr =common.HexToAddress
var toHash =common.BytesToHash

func TestAddStatedb(t *testing.T){
	statedb:=NewStateDB()

	// EOA账户
	account1:=AddEOAAccount(statedb, testdata.EthAccount1,big.NewInt(100))
	fmt.Printf("EOA addr:%v\n", account1.Hex())
	_=AddEOAAccount(statedb, testdata.EthAccount2,big.NewInt(200))

	// 智能合约
	contractAccount :=AddSmartContractAccount(statedb, testdata.EthAccount1,[]byte("contracts code bytes"))
	fmt.Printf("contract account addr:%v\n", contractAccount.Hex())
	SetSmartContractState(statedb,contractAccount,[][2][]byte{
		{[]byte("owner"),account1.Bytes()},
		{[]byte("name"),[]byte("zml")},
		{[]byte("online"),[]byte{1}},
		//{[]byte("online"),[]byte{}},
	})

	PrintStateDB(statedb)
}
