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

func UUIDToString(u datatypes.UUID) string {
	if uuid.UUID(u) == uuid.Nil {
		return ""
	}
	return u.String()
}

func UUIDPtrToString(u *datatypes.UUID) string {
	if u != nil {
		return u.String()
	}
	return ""
}
