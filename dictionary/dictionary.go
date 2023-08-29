package dictionary

import "errors"

type dictionary map[string]string

var errNotFound = errors.New("Not Found")

func CreateDictionary(key, value string) *dictionary {
	return &dictionary{key: value}
}

func (d dictionary) Search(key string) (string, error) {
	value, isExists := d[key]

	if isExists {
		return value, nil
	}

	return "", errNotFound
}