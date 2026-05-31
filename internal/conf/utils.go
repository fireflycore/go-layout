package conf

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fireflycore/go-micro/config"
	"github.com/fireflycore/go-utils/compress"
	"github.com/fireflycore/go-utils/crypto"
)

// Utils 提供启动期配置文件读取相关辅助能力。
type Utils struct {
	crypto   crypto.Crypto
	compress compress.Compress
}

// NewConfigUtils 创建配置辅助工具。
func NewConfigUtils(crypto crypto.Crypto, compress compress.Compress) *Utils {
	return &Utils{
		crypto:   crypto,
		compress: compress,
	}
}

// Encryptor 返回统一配置读取链路使用的加解密实现。
func (u *Utils) Encryptor() config.Encryptor {
	return u.crypto
}

// Compressor 返回统一配置读取链路使用的压缩实现。
func (u *Utils) Compressor() config.Compressor {
	return u.compress
}

// DecodePayload 按 go-micro/config 的统一规则还原远程配置内容。
func (u *Utils) DecodePayload(content string, encrypted bool, secret []byte, target any) error {
	return config.UnmarshalPayload(content, encrypted, secret, target, u.Compressor(), u.Encryptor(), nil)
}

// GetConfigFilePath 返回 conf 目录下配置文件的绝对路径。
func (u *Utils) GetConfigFilePath(filename string) string {
	cur, err := os.Getwd()
	if err != nil {
		panic("failed to get working directory: " + err.Error())
	}
	return filepath.Join(cur, "conf", filename)
}

// LoadJSONConfig 读取并解析 JSON 配置文件。
func (u *Utils) LoadJSONConfig(file string, target any) error {
	// 先完整读取文件内容，保持启动阶段错误尽早暴露。
	b, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("failed to read config file %s: %w", file, err)
	}
	// 解析失败时保留原始文件路径，方便启动期快速定位问题。
	if err = json.Unmarshal(b, target); err != nil {
		return fmt.Errorf("failed to parse config file %s: %w", file, err)
	}
	return nil
}
