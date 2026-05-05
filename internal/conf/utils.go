package conf

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fireflycore/go-utils/compress"
	"github.com/fireflycore/go-utils/crypto"
)

type Utils struct {
	crypto   crypto.Crypto
	compress compress.Compress
}

// NewConfigUtils 组装配置读取和解密阶段共用的工具能力。
func NewConfigUtils(crypto crypto.Crypto, compress compress.Compress) *Utils {
	return &Utils{
		crypto:   crypto,
		compress: compress,
	}
}

// GetConfigFilePath 获取配置路径
func (ist *Utils) GetConfigFilePath(filename string) string {
	cur, err := os.Getwd()
	if err != nil {
		panic("failed to get working directory: " + err.Error())
	}
	return filepath.Join(cur, "conf", filename)
}

// LoadJSONConfig 获取本地配置
func (ist *Utils) LoadJSONConfig(file string, target any) error {
	b, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("failed to read config file %s: %w", file, err)
	}

	if err = json.Unmarshal(b, target); err != nil {
		return fmt.Errorf("failed to parse config file %s: %w", file, err)
	}
	return nil
}

// AnalyzeData 按“整份配置项”解密并反序列化目标对象。
func (ist *Utils) AnalyzeData(content string, key []byte, val any) error {
	decode, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return err
	}

	decrypt, de := ist.crypto.Decrypt(decode, key)
	if de != nil {
		return de
	}

	decompress, err := ist.compress.Decompress(decrypt)
	if err != nil {
		return err
	}

	return json.Unmarshal(decompress, val)
}
