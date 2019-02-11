package byexample

import (
	"github.com/mitchellh/hashstructure"
)

type Set map[interface{}]interface{}

type PowerMap map[uint64]interface{}

func (setMap *PowerMap) Set(key interface{}, value interface{}) (keyHash uint64, err error) {
	keyHash, err = hashstructure.Hash(key, nil)
	if err != nil {
		(*setMap)[keyHash] = value
	}
	return
}

func (setMap *PowerMap) Get(key interface{}) (interface{}, error) {
	keyHash, err := hashstructure.Hash(key, nil)
	if err != nil {
		return (*setMap)[keyHash], nil
	}
	return nil, err
}

func (setMap *PowerMap) Delete(key interface{}) error {
	keyHash, err := hashstructure.Hash(key, nil)
	if err != nil {
		delete(*setMap, keyHash)
	}
	return err
}

func NewPowerMap() *PowerMap {
	powerMap := make(PowerMap)
	return &powerMap
}
