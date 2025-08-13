package conf

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	compress "github.com/lhdhtrc/compress-go/pkg"
	crypto "github.com/lhdhtrc/crypto-go/pkg"
	"github.com/lhdhtrc/func-go/file"
	micro "github.com/lhdhtrc/micro-go/pkg/core"
	config "go-layout/dep/protobuf/gen/acme/config/v1"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

// getConfigFilePath 获取配置路径
func getConfigFilePath(filename string) string {
	cur, err := os.Getwd()
	if err != nil {
		panic("failed to get working directory: " + err.Error())
	}
	return filepath.Join(cur, "conf", filename)
}

// loadJSONConfig 获取本地配置
func loadJSONConfig(file string, target interface{}) error {
	b, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("failed to read config file %s: %w", file, err)
	}

	if err = json.Unmarshal(b, target); err != nil {
		return fmt.Errorf("failed to parse config file %s: %w", file, err)
	}
	return nil
}

// fetchRemoteConfig 获取远程配置
func fetchRemoteConfig(bc *BootstrapConf, cc config.ConfigServiceClient, group, key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := micro.WithRemoteInvoke[*config.Config, *config.GetResponse](func() (*config.GetResponse, error) {
		return cc.Get(ctx, &config.GetRequest{
			Env:   bc.Env,
			AppId: bc.AppId,
			Group: group,
			Key:   key,
		})
	})
	if err != nil {
		return "", err
	}

	return data.Content, nil
}

// analyzeData 解析数据
func analyzeData(content string, key []byte, val any) error {
	decode, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return err
	}

	decrypt, de := crypto.UseAES.Decrypt(decode, key)
	if de != nil {
		return de
	}

	decompress, err := compress.UseGZIP.Decompress(decrypt)
	if err != nil {
		return err
	}

	return json.Unmarshal(decompress, &val)
}

// analyzeTlsData 解析tls数据
func analyzeTlsData(dir string, tls interface{}) error {
	if tls == nil {
		return nil
	}

	dirPath := filepath.Join("dep", "cert", dir)

	// 遍历config的字段
	valueOfConfig := reflect.ValueOf(tls)
	if valueOfConfig.Kind() == reflect.Ptr {
		valueOfConfig = valueOfConfig.Elem()
	}
	typeOfConfig := valueOfConfig.Type()

	for i := 0; i < valueOfConfig.NumField(); i++ {
		fieldValue := valueOfConfig.Field(i)
		fieldType := typeOfConfig.Field(i)
		if fieldValue.IsValid() && !fieldValue.IsZero() && fieldType.Type.Kind() == reflect.String {
			val := fieldValue.String()

			fd := strings.ReplaceAll(val, `\n`, "\n")
			fp := filepath.Join(dirPath, fmt.Sprintf("%s.pem", fieldType.Tag.Get("json")))

			if err := file.WriteLocal(fp, []byte(fd)); err != nil {
				return err
			}

			fieldValue.SetString(fp)
		}
	}
	return nil
}
