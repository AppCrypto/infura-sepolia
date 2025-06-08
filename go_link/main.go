// used to interact with the smart contract
package main

import (
	"context"
	"fmt"
	hello "go_link_sol/gen"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//Wallet下是用keystore创建的加密钱包文件，需要用户的password解密得到钱包私钥
	b, err := os.ReadFile("./wallet/UTC--2025-05-29T13-51-04.984946200Z--9dda53414c9a26b1054427718cd991ec14bd5fd4")
	if err != nil {
		log.Fatal(err)
	}
	//解密
	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal(err)
	}
	//创建client连接测试网
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/de28dcce3b8f4d319a904bfab58d1e1a")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	//获取chainID,以太网主网ID为0
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//获取建议的gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//contract address, not sender wallet address
	//下面的地址更改为在运行deploy文件后输出的合约地址，或者可以在etherscan查看
	cAdd := common.HexToAddress("0xA677259156718de6B1F9575Bf498d3663B3629a4")
	t, err := hello.NewHello(cAdd, client)
	if err != nil {
		log.Fatal(err)
	}
	tx, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	tx.GasLimit = 3000000
	tx.GasPrice = gasPrice

	add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	t.SetMsg(&bind.TransactOpts{From: add}, "hello")
	Msg, err := t.GetMsg(&bind.CallOpts{From: add})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Msg)
}
