package dictionary

import "errors"

type dictionary map[string]string

var (
	errNotFound = errors.New("Not Found")
	errKeyExists = errors.New("Key Exists")
)

func CreateDictionary(key, value string) dictionary {
	return dictionary{key: value}
}

func (d dictionary) Search(key string) (string, error) {
	value, isExists := d[key]

	if isExists {
		return value, nil
	}

	return "", errNotFound
}

func (d dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	if err != nil {
		d[key] = value

		return nil
	} else {
		return errKeyExists
	}
}

func (d dictionary) Update(key, value string) error{
	_, err := d.Search(key)

	if err != nil {
		return errNotFound
	}

	d[key] = value

	return nil
}