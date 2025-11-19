package homework2

import (
	"context"
	"fmt"
	"os"

	"github.com/0gfoundation/0g-storage-client/common/blockchain"
	"github.com/0gfoundation/0g-storage-client/core"
	"github.com/0gfoundation/0g-storage-client/indexer"
	"github.com/0gfoundation/0g-storage-client/transfer"
	"github.com/joho/godotenv"
)

func Upload(filePath string, fragmentSize int64) ([]string, []string, error) {
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
	nodes, err := indexerClient.SelectNodes(ctx, 1, []string{}, "max", true)
	if err != nil {
		fmt.Println("Failed to select nodes:", err)
	}

	// 创建上传器对象
	uploader, err := transfer.NewUploader(ctx, w3client, nodes)
	if err != nil {
		fmt.Println("Failed to create uploader:", err)
	}

	// Upload file
	file, err := core.Open(filePath)
	if err != nil {
		fmt.Println(err, "failed to open file %s", filePath)
	}
	defer file.Close()
	// 设置分片大小为fragmentSize
	txHashs, rootHashs, err := uploader.SplitableUpload(ctx, file, fragmentSize)
	if err != nil {
		fmt.Println(err, "failed to upload file %s", filePath)
	}
	// []common.Hash 转换为 []string
	var txHashStrs []string
	for _, txHash := range txHashs {
		txHashStrs = append(txHashStrs, txHash.Hex())
	}
	var rootHashStrs []string
	for _, rootHash := range rootHashs {
		rootHashStrs = append(rootHashStrs, rootHash.Hex())
	}
	return txHashStrs, rootHashStrs, nil
}
