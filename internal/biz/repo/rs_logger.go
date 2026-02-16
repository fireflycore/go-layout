package repo

type LoggerRepo interface {
	CreateAccessLog(appId string, raw []byte)
	CreateServerLog(appId string, raw []byte)
	CreateOperationLog(appId string, raw []byte)
}
