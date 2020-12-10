package account

import (
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
	_=AddEOAAccount(statedb, testdata.EthAccount2,big.NewInt(200))

	// 智能合约
	contractAccount :=AddSmartContractAccount(statedb, testdata.EthAccount1,[]byte("contracts code bytes"))
	SetSmartContractState(statedb,contractAccount,[][2][]byte{
		{[]byte("owner"),account1.Bytes()},
		{[]byte("name"),[]byte("ysqi")},
		{[]byte("online"),[]byte{1}},
		{[]byte("online"),[]byte{}},
	})

	PrintStateDB(statedb)
}
