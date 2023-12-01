package store

type _Entrance struct {
	Config _ConfigStoreEntrance
	DB     _DBStoreEntrance
	Micro  _MicroStoreEntrance
	Logger _LoggerStoreEntrance
}

var Use = new(_Entrance)
