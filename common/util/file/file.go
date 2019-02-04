package file

import (
	"encoding/base64"
	"io/ioutil"
	"os"
)

//读取文件数据
func GetFileData(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

//保存base64数据为文件
func SaveBase64ToFile(content, path string) error {
	data, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return err
	}
	f.Write(data)
	return nil
}

//获取文件数据 转base64
func GetFileToBase64(filePath string) (string, error) {
	data, err := GetFileData(filePath)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}
