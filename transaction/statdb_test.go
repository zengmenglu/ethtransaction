package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)
var toAddr =common.HexToAddress
var toHash =common.BytesToHash

func TestAddStatedb(t *testing.T){
	// 创建一个内存KV数据库，再包装为 state 数据库实例， 最后利用一个空的DB级的StateRoot，初始化一个以太坊 statedb。
	statedb, _ := state.New(common.Hash{},	state.NewDatabase(rawdb.NewMemoryDatabase()),nil)

	// 创建2个EOA账户，并设置余额
	acct1:=toAddr("0x0bB141C2F7d4d12B1D27E62F86254e6ccEd5FF9a")
	acct2:=toAddr("0x77de172A492C40217e48Ebb7EEFf9b2d7dF8151B")
	statedb.AddBalance(acct1,big.NewInt(100))
	statedb.AddBalance(acct2,big.NewInt(888))

	// 创建智能合约账户
	contract:=crypto.CreateAddress(acct1, statedb.GetNonce(acct1))
	statedb.CreateAccount(contract)
	statedb.SetCode(contract,[]byte("contract code bytes"))
	// 设置智能合约
	statedb.SetNonce(contract,1)
	statedb.SetState(contract,toHash([]byte("owner")),toHash(acct1.Bytes())) //新增owner状态数据字段
	statedb.SetState(contract,toHash([]byte("name")),toHash([]byte("ysqi")))

	statedb.SetState(contract,toHash([]byte("online")),toHash([]byte{1}))
	statedb.SetState(contract,toHash([]byte("online")),toHash([]byte{})) //赋空值，则会删除该数据字段

	//将statedb内存中数据写入数据库文件
	statedb.Commit(true)

	fmt.Println(string(statedb.Dump(false,false,false)))
}
