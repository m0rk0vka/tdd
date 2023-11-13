package maps

import "errors"

var (
	ErrNotFound = errors.New("could not find the key you are looking for")
	ErrKeyExist = errors.New("cannot add key because it already exist")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	val, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return val, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrKeyExist
	default:
		return err
	}
	return nil
}
