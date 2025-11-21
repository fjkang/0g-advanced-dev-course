package homework2

import "os"

func CreateNullFile(filePath string, size int64) error {
	// 创建文件
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 设置文件大小
	if _, err := file.Seek(size-1, 0); err != nil {
		return err
	}

	// 写入一个字节，确保文件大小
	if _, err := file.Write([]byte{0}); err != nil {
		return err
	}

	return nil
}
