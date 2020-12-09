package account

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

// 获取账户余额
func GetAccountBalance(client *ethclient.Client, addr common.Address, blockNum ...*big.Int)(*big.Int, error){
	if len(blockNum)==1{
		return client.BalanceAt(context.Background(), addr, blockNum[0])
	}
	return client.BalanceAt(context.Background(), addr, nil)
}

// 获取待处理的账户余额
func GetPendingBalance(client *ethclient.Client, addr common.Address)(*big.Int, error){
	return client.PendingBalanceAt(context.Background(),addr)
}