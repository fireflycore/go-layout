package repo

type LoggerRepo interface {
	CreateAccessLogger(appId string, raw []byte)
	CreateServerLogger(appId string, raw []byte)
	CreateOperationLogger(appId string, raw []byte)
}
