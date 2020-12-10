package account

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

func NewStateDB()*state.StateDB{
	// 创建一个内存KV数据库，再包装为 state 数据库实例， 最后利用一个空的DB级的StateRoot，初始化一个以太坊 statedb。
	statedb, _ := state.New(common.Hash{},	state.NewDatabase(rawdb.NewMemoryDatabase()),nil)
	return statedb
}

func AddEOAAccount(db *state.StateDB, addr string,balance *big.Int)common.Address{
	account:= common.HexToAddress(addr)
	db.AddBalance(account,balance)
	db.Commit(true) //将statedb内存中数据写入数据库文件
	return account
}

func AddSmartContractAccount(db *state.StateDB, eoaAddr string, code []byte)common.Address{
	account:=common.HexToAddress(eoaAddr)
	contractAddr :=crypto.CreateAddress(account, db.GetNonce(account))
	db.CreateAccount(contractAddr)
	db.SetCode(contractAddr,code)
	db.Commit(true)

	return contractAddr
}

func SetSmartContractState(db *state.StateDB,addr common.Address,stateDatas [][2][]byte){
	for _,stateData:=range stateDatas{
		db.SetState(addr, common.BytesToHash(stateData[0]),common.BytesToHash(stateData[1]))
	}
	db.Commit(true)
}

func PrintStateDB(db *state.StateDB){
	fmt.Println(string(db.Dump(false,false,false)))
}
