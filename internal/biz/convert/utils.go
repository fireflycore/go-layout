package convert

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"time"
)

func TimeToString(t time.Time) string {
	return t.Format(time.DateTime)
}

func StringToUUID(s string) datatypes.UUID {
	u, e := uuid.Parse(s)
	if e != nil {
		return datatypes.UUID(uuid.Nil)
	}
	return datatypes.UUID(u)
}
