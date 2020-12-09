package wallet

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"log"
)

// 创建并存储密钥对
func CreateKs(passphrase string) *keystore.KeyStore{
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(passphrase)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())
	return ks
}

// 导出地址
func importKs(ks *keystore.KeyStore, passphrase string) {
	file:=ks.Accounts()[0].URL.Path
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	account, err := ks.Import(jsonBytes, passphrase, passphrase)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(account.Address.Hex())
}
