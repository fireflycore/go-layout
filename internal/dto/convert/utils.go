package convert

import (
	"github.com/google/uuid"
	gorm "github.com/lhdhtrc/gorm/pkg"
	"time"
)

func TimeToString(t time.Time) string {
	return t.Format(time.DateTime)
}

func StringToUUID(s string) gorm.BinUUID {
	u, e := uuid.Parse(s)
	if e != nil {
		return gorm.BinUUID(uuid.Nil)
	}
	return gorm.BinUUID(u)
}

func UUIDToString(u gorm.BinUUID) string {
	if uuid.UUID(u) == uuid.Nil {
		return ""
	}
	return u.String()
}

func UUIDPtrToString(u *gorm.BinUUID) string {
	if u != nil {
		return u.String()
	}
	return ""
}
