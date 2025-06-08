// used to interact with the smart contract
package main

import (
	"context"
	"fmt"
	hello "go_link_sol/gen" // 替换为你的合约生成的Go绑定包名
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"

	//"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 1. 加载加密的钱包文件
	walletPath := "./wallet/UTC--2025-05-29T13-51-04.984946200Z--9dda53414c9a26b1054427718cd991ec14bd5fd4"
	walletData, err := os.ReadFile(walletPath)
	if err != nil {
		log.Fatalf("读取钱包文件失败: %v", err)
	}

	// 2. 解密钱包 (使用你的钱包密码)
	key, err := keystore.DecryptKey(walletData, "password")
	if err != nil {
		log.Fatalf("解密钱包失败: %v", err)
	}

	// 3. 连接到以太坊节点 (这里使用Sepolia测试网)
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/de28dcce3b8f4d319a904bfab58d1e1a")
	if err != nil {
		log.Fatalf("连接以太坊节点失败: %v", err)
	}
	defer client.Close()

	// 4. 获取链ID和gas价格
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("获取网络ID失败: %v", err)
	}

	// 5. 创建合约实例
	contractAddress := common.HexToAddress("0x0b498951dE984dEa3CA5cE3d292ae4c8B5112786")
	contract, err := hello.NewHello(contractAddress, client)
	if err != nil {
		log.Fatalf("创建合约实例失败: %v", err)
	}

	// 6. 准备交易选项
	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatalf("创建交易签名者失败: %v", err)
	}

	// 设置合理的Gas参数
	auth.GasLimit = 300000
	auth.Context = context.Background()

	// 7. 调用setMsg函数
	fmt.Println("正在设置消息...")
	tx, err := contract.SetMsg(auth, "hello")
	if err != nil {
		log.Fatalf("调用SetMsg失败: %v", err)
	}

	fmt.Printf("交易已发送，交易哈希: %s\n", tx.Hash().Hex())

	// 8. 等待交易被挖出
	fmt.Println("等待交易确认...")
	start := time.Now()
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("等待交易确认失败: %v", err)
	}

	if receipt.Status == 0 {
		log.Fatal("交易执行失败")
	}

	fmt.Printf("交易确认成功，耗时: %v\n", time.Since(start))

	// 9. 读取消息
	fmt.Println("正在读取消息...")
	message, err := contract.GetMsg(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("调用GetMsg失败: %v", err)
	}

	fmt.Printf("消息内容: %s\n", message)
}
