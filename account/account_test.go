package account

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zengmenglu/ethtransaction/testdata"
	"testing"
)

func TestGetAccountBalance(t *testing.T) {
	client, err := ethclient.Dial(testdata.TestNet)
	if err != nil {
		t.Errorf("connect to eth fail, err:%s\n", err)
		return
	}
	addr := common.HexToAddress(testdata.EthAccount)
	balance, err:= GetAccountBalance(client,addr)
	if err != nil{
		t.Errorf("get balance fail, err:%s\n", err)
		return
	}

	t.Logf("balance:%v\n", balance)
}

func TestGetPendingBalance(t *testing.T) {
	client, err := ethclient.Dial(testdata.TestNet)
	if err != nil {
		t.Errorf("connect to eth fail, err:%s\n", err)
		return
	}
	addr := common.HexToAddress(testdata.EthAccount)
	balance, err:= GetPendingBalance(client,addr)
	if err != nil{
		t.Errorf("get balance fail, err:%s\n", err)
		return
	}

	t.Logf("pending balance:%v\n", balance)
}