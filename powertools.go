package byexample

import (
	"github.com/mitchellh/hashstructure"
)

type Set map[interface{}]interface{}

type PowerMap map[uint64]interface{}

func (setMap PowerMap) Set(key interface{}, value interface{}) (uint64, error) {
	keyHash, error := hashstructure.Hash(key, nil)
	if error != nil {
		setMap[keyHash] = value
	}
	return keyHash, error
}

func (setMap PowerMap) Get(key interface{}) (interface{}, error) {
	keyHash, error := hashstructure.Hash(key, nil)
	if error != nil {
		return setMap[keyHash], nil
	}
	return nil, error
}

func (setMap PowerMap) Delete(key interface{}) error {
	keyHash, error := hashstructure.Hash(key, nil)
	if error != nil {
		delete(setMap, keyHash)
	}
	return error
}

func NewPowerMap() PowerMap {
	return make(PowerMap)
}
