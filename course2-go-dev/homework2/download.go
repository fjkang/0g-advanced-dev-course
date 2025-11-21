package homework2

import (
	"context"
	"fmt"
	"os"

	"github.com/0gfoundation/0g-storage-client/common/blockchain"
	"github.com/0gfoundation/0g-storage-client/indexer"
	"github.com/joho/godotenv"
)

func Download(roots []string, filename string) {

	// 配置基础参数
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	evmRpc := os.Getenv("evmRpc")
	indRpc := os.Getenv("indRpc")
	privateKey := os.Getenv("privateKey")

	// 创建w3client对象
	w3client := blockchain.MustNewWeb3(evmRpc, privateKey)
	defer w3client.Close()

	// 创建indexerClient
	indexerClient, err := indexer.NewClient(indRpc)
	if err != nil {
		fmt.Println("Failed to create indexer client:", err)
	}
	defer indexerClient.Close()
	ctx := context.Background()

	// Download with optional verification
	indexerClient.DownloadFragments(ctx, roots, filename, false)

}
