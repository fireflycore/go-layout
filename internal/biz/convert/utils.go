package convert

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// ProtoIgnoreFields 定义 Proto 运行时生成字段，做结构差异比较时需要跳过。
var ProtoIgnoreFields = []string{"state", "sizeCache", "unknownFields"}

// TimeToString 统一把时间序列化成字符串。
func TimeToString(t time.Time) string {
	return t.Format(time.DateTime)
}

// StringToUUID 把字符串安全转换成 UUID；解析失败时返回零值 UUID。
func StringToUUID(s string) datatypes.UUID {
	u, e := uuid.Parse(s)
	if e != nil {
		return datatypes.UUID(uuid.Nil)
	}
	return datatypes.UUID(u)
}
