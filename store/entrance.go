package store

type _Entrance struct {
	Config _ConfigStoreEntrance
	DB     _DBStoreEntrance
	Micro  _MicroStoreEntrance
}

var Use = new(_Entrance)
