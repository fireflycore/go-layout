package conf

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	microConfig "github.com/fireflycore/go-micro/config"
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

// Encryptor 返回统一配置读取链路使用的加解密实现。
func (ist *Utils) Encryptor() microConfig.Encryptor {
	return ist.crypto
}

// Compressor 返回统一配置读取链路使用的压缩实现。
func (ist *Utils) Compressor() microConfig.Compressor {
	return ist.compress
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
